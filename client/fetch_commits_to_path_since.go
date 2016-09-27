package client

import (
	"fmt"
	"net/url"
	"time"
)

func (c *defaultclient) FetchCommitsToPathSince(path string, since time.Time) ([][]byte, error) {
	f := NewFetcher()
	urlString := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", c.repoOwner, c.repoName)
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	params := &url.Values{}
	//TODO add test for this
	if path != "." && path != "" && path != "./" {
		params.Add("path", path)
	}

	params.Add("since", since.Format(DateFormat))
	params.Add("access_token", c.accessToken)
	params.Add("per_page", fmt.Sprintf("%d", 1000)) //TODO get this tuning variable out of here

	u.RawQuery = params.Encode()

	commitPages, err := f.Fetch(PaginateGitHubResponse, u.String())
	if err != nil {
		return nil, err
	}

	return commitPages, nil
}
