package client

import (
	"fmt"
	"net/url"
	"time"
)

const (
	numberPagesToRequest = "1000"
)

func (c *defaultclient) FetchCommitsToPathSince(path string, since time.Time) ([][]byte, error) {
	f := NewFetcher(c.accessToken)
	urlString := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", c.repoOwner, c.repoName)
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	params := &url.Values{}
	//TODO add test for this
	if path != "" {
		params.Add("path", path)
	}

	params.Add("since", since.Format(DateFormat))
	params.Add("per_page", numberPagesToRequest)
	u.RawQuery = params.Encode()

	commitPages, err := f.Fetch(PaginateGitHubResponse, u.String())
	if err != nil {
		return nil, err
	}

	return commitPages, nil
}
