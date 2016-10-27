package fetcher

import (
	"fmt"
	"net/http"
	"time"

	"github.com/calebamiles/github-client/cache"
	"github.com/calebamiles/github-client/client/fetcher"
)

var _ fetcher.Fetcher = &DefaultCachingFetcher{}

// NewCachingFetcher returns the default caching fetcher which understands how to use
// an authorization header to authenticate HTTP requests
func NewCachingFetcher(accessToken string, pageCache cache.Page) fetcher.Fetcher {
	return &DefaultCachingFetcher{
		Cache:       pageCache,
		Fetcher:     NewFetcher(accessToken),
		accessToken: accessToken,
	}
}

// DefaultCachingFetcher uses a paginator to fetch all reachable pages from a base URL
// and uses a PageCache to store previously fetched pages
type DefaultCachingFetcher struct {
	Cache       cache.Page
	Fetcher     fetcher.Fetcher
	accessToken string
}

// Fetch fetches all pages reachable from url, according to the PaginationFunc
func (f *DefaultCachingFetcher) Fetch(url string) ([]byte, string, error) {
	err := f.Cache.Open() // make sure the cache is open before attempting to use it, relies on PageCache.Open() being idempotent
	if err != nil {
		return nil, "", err
	}

	cacheKey, err := f.Cache.KeyForPage(url)
	if err != nil {
		return nil, "", err
	}

	contentUnchanged, err := f.pageContentUnchanged(url, cacheKey)
	if err != nil {
		return nil, "", err
	}

	if contentUnchanged {
		cachedPage, cacheFetchErr := f.Cache.FetchPageByKey(cacheKey)
		if cacheFetchErr != nil {
			return nil, "", cacheFetchErr
		}

		return cachedPage, cacheKey, nil
	}

	updatedPage, updatedCacheKey, err := f.Fetcher.Fetch(url)
	if err != nil {
		return nil, "", err
	}

	err = f.Cache.AddPage(url, updatedCacheKey, updatedPage)
	if err != nil {
		return nil, "", err
	}

	return updatedPage, updatedCacheKey, nil
}

func (f *DefaultCachingFetcher) pageContentUnchanged(url string, etag string) (bool, error) {
	if etag == "" {
		return false, nil
	}

	c := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return false, err
	}

	if f.accessToken != "" {
		req.Header.Set(authorizationHeader, fmt.Sprintf("token %s", f.accessToken))
	}

	if etag != "" {
		req.Header.Set(CacheRequestHeader, etag)
	}

	resp, err := c.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()
	unchanged := (resp.StatusCode == http.StatusNotModified)

	return unchanged, nil
}

// Done will close the cache and make the fetcher unusable
func (f *DefaultCachingFetcher) Done() error {
	return f.Cache.Close()
}
