package client

import (
	"fmt"
	"net/url"
	"time"

	"github.com/calebamiles/github-client/client/internal/pages"
	"github.com/calebamiles/github-client/prs"
)

// FetchPullRequestsSince returns all pull requests against the repository, optionally since a non zero time
func (c *DefaultClient) FetchPullRequestsSince(since time.Time) ([]prs.PullRequest, error) {
	urlString := fmt.Sprintf("https://api.github.com/repos/%s/%s/pullRequests", c.RepoOwner, c.RepoName)
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	params := &url.Values{}
	if !since.IsZero() {
		params.Add("since", since.Format(DateFormat))
	}

	params.Add("per_page", NumberOfPagesToRequest)
	u.RawQuery = params.Encode()

	pullRequestPages, err := c.Fetcher.Fetch(u.String())
	if err != nil {
		return nil, err
	}

	joinedPullRequestPages := pages.Join(pullRequestPages)
	pullRequestsWithoutComments, err := prs.New(joinedPullRequestPages)
	if err != nil {
		return nil, err
	}

	return pullRequestsWithoutComments, nil
}
