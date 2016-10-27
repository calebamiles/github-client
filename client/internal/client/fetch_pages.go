package client

// FetchPage returns all pages reachable from urlString by using a Paginator
func (c *DefaultClient) FetchPage(urlString string) ([]byte, error) {
	pages, _, err := c.Fetcher.Fetch(urlString)
	return pages, err
}
