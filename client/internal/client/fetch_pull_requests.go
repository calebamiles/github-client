package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"sync"

	"github.com/calebamiles/github-client/client/fetcher"
	"github.com/calebamiles/github-client/comments"
	"github.com/calebamiles/github-client/commits"
	"github.com/calebamiles/github-client/prs"
	"github.com/calebamiles/github-client/prs/file"
	"github.com/calebamiles/github-client/prs/pr"
)

// FetchPullRequestsSince returns all pull requests against the repository, optionally since a non zero time
func (c *DefaultClient) FetchPullRequests() ([]pr.PullRequest, error) {
	urlString := fmt.Sprintf("https://api.github.com/repos/%s/%s/pullRequests", c.RepoOwner, c.RepoName)
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	params := &url.Values{}
	params.Add("per_page", NumberOfPagesToRequest)
	u.RawQuery = params.Encode()

	pullRequestPages, _, err := c.Fetcher.Fetch(u.String())
	if err != nil {
		return nil, err
	}

	// make processing these PRs a little bit easier to do
	var jsonPRs []json.RawMessage
	err = json.Unmarshal(pullRequestPages, &jsonPRs)
	if err != nil {
		return nil, err
	}

	basicPullRequests, err := prs.NewBasicPullRequests(jsonPRs)
	if err != nil {
		return nil, err
	}

	accumulator := newPRAccumulator(c.Fetcher, len(basicPullRequests))
	for i := range basicPullRequests {
		go func(basicPR pr.BasicPullRequest) {
			commitsPages, _, err := c.Fetcher.Fetch(basicPR.CommitsURL())
			if err != nil {
				accumulator.Add(nil, err)
				return
			}

			commits, err := commits.New(commitsPages)
			if err != nil {
				accumulator.Add(nil, err)
				return
			}

			commentsPages, _, err := c.Fetcher.Fetch(basicPR.CommentsURL())
			if err != nil {
				accumulator.Add(nil, err)
				return
			}

			comments, err := comments.New(commentsPages)
			if err != nil {
				accumulator.Add(nil, err)
				return
			}

			filesChangedPages, _, err := c.Fetcher.Fetch(basicPR.FileChangeURL())
			if err != nil {
				accumulator.Add(nil, err)
				return
			}

			filesChanged, err := file.NewChangeSets(filesChangedPages)
			if err != nil {
				accumulator.Add(nil, err)
				return
			}

			accumulator.Add(&pullRequest{
				BasicPullRequest: basicPR,
				commits:          commits,
				comments:         comments,
				filesChanged:     filesChanged,
			}, nil)

			return
		}(basicPullRequests[i])
	}

	return nil, nil
}

type pullRequest struct {
	pr.BasicPullRequest
	commits      []commits.Commit
	comments     []comments.Comment
	filesChanged []file.ChangeSet
}

func (p *pullRequest) Commits() []commits.Commit      { return p.commits }
func (p *pullRequest) Comments() []comments.Comment   { return p.comments }
func (p *pullRequest) FilesChanged() []file.ChangeSet { return p.filesChanged }

func newPRAccumulator(fetcher fetcher.Fetcher, waitForNumber int) *prAccumulator {
	accumulator := &prAccumulator{}
	accumulator.wg.Add(waitForNumber)

	return accumulator
}

type prAccumulator struct {
	pullRequests []pr.PullRequest
	errors       []error

	wg sync.WaitGroup
	sync.Mutex
}

func (a *prAccumulator) Add(p pr.PullRequest, err error) {
	a.Lock()
	defer a.Unlock()

	if p != nil {
		a.pullRequests = append(a.pullRequests, p)
	}

	if err != nil {
		a.errors = append(a.errors, err)
	}

	a.wg.Done()
}

func (a *prAccumulator) Wait() ([]pr.PullRequest, error) {
	a.wg.Wait()

	return a.pullRequests, combineError(a.errors)
}

func combineError(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	var errStrings []string
	for i := range errs {
		errStrings = append(errStrings, errs[i].Error())
	}

	joinedErrStrings := strings.Join(errStrings, "\n")
	return fmt.Errorf("errors were returned while processing pull requests: \n %s", joinedErrStrings)
}
