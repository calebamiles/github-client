package client

import (
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/calebamiles/github-client/commits"
)

// FetchCommitsToPathSince returns commits with comments to a path, an empty string for path will return all commits
func (c *DefaultClient) FetchCommitsToPathSince(path string, since time.Time) ([]commits.Commit, error) {
	urlString := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", c.RepoOwner, c.RepoName)
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	params := &url.Values{}
	if path != "" {
		params.Add("path", path)
	}

	if !since.IsZero() {
		params.Add("since", since.Format(DateFormat))
	}

	params.Add("per_page", NumberOfPagesToRequest)
	u.RawQuery = params.Encode()

	commitPages, err := c.Fetcher.Fetch(u.String())
	if err != nil {
		return nil, err
	}

	var commitsWithoutComments []commits.CommitWithoutComments
	for i := range commitPages {
		commitsOnPage, loopErr := commits.New(commitPages[i])
		if loopErr != nil {
			return nil, loopErr
		}

		commitsWithoutComments = append(commitsWithoutComments, commitsOnPage...)
	}

	errs := &errorAccumulator{}
	wg := &sync.WaitGroup{}
	cs := &commitAccumulator{}
	ready := make(chan struct{}, 50)

	for i := range commitsWithoutComments {
		wg.Add(1)
		go c.processCommits(commitsWithoutComments[i], cs, ready, wg, errs)
	}

	wg.Wait()
	if !errs.IsNil() {
		return nil, errs
	}

	return cs.GetAll(), nil
}
