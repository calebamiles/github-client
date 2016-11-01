package client

import (
	"fmt"
	"net/url"

	"github.com/calebamiles/github-client/commits"
)

// FetchCommits returns commits without comments to a path, an empty string for path will return all commits
func (c *DefaultClient) FetchCommits() ([]commits.Commit, error) {
	urlString := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", c.RepoOwner, c.RepoName)
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	params := &url.Values{}
	params.Add("per_page", NumberOfPagesToRequest)
	u.RawQuery = params.Encode()

	commitPages, _, err := c.Fetcher.Fetch(u.String())
	if err != nil {
		return nil, err
	}

	commitsWithoutComments, err := commits.New(commitPages)
	if err != nil {
		return nil, err
	}

	return commitsWithoutComments, nil
}
