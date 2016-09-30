package client

import (
	"fmt"
	"net/url"

	"github.com/calebamiles/github-client/client/internal"
)

func (c *defaultclient) FetchIssuesPages() ([][]byte, error) {
	f := internal.NewFetcher(c.accessToken)

	urlString := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", c.repoOwner, c.repoName)
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	params := &url.Values{}
	params.Add("state", "all")
	u.RawQuery = params.Encode()

	return f.Fetch(u.String())
}
