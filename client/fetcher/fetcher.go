package fetcher

// A Fetcher uses a paginator to return all pages reachable from a base URL
type Fetcher interface {
	Fetch(url string) (pageContent []byte, err error)
}
