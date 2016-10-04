package fetcher

// A Fetcher uses a paginator to return all pages reachable from a base URL
type Fetcher interface {
	Fetch(string) ([][]byte, error)
}
