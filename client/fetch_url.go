package client

import "github.com/calebamiles/github-client/client/internal"

func (c *defaultclient) FetchURL(urlString string) ([][]byte, error) {
	f := internal.NewFetcher(c.accessToken)

	return f.Fetch(urlString)
}
