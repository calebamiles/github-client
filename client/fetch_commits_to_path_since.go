package client

import (
	"fmt"
	"net/url"
	"time"

	"github.com/calebamiles/github-client/client/internal"
)

func (c *defaultclient) FetchCommitsToPathSincePages(path string, since time.Time) ([][]byte, error) {
	f := internal.NewFetcher(c.accessToken)
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

	return f.Fetch(u.String())
}
