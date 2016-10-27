package fetcher

const (
	// CacheRequestHeader is the header to supply page caching information when making a request
	CacheRequestHeader = "If-None-Match"

	// CacheResponseHeader is the expected header when receiving an HTTP response
	CacheResponseHeader = "ETag"
)
