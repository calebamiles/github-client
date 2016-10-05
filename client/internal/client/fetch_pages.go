package client

// FetchPages returns all pages reachable from urlString by using a Paginator
func (c *DefaultClient) FetchPages(urlString string) ([][]byte, error) {
	return c.Fetcher.Fetch(urlString)
}
