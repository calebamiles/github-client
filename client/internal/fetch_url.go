package internal

func (c *DefaultClient) FetchURL(urlString string) ([][]byte, error) {
	return c.Fetcher.Fetch(urlString)
}
