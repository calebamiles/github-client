package fetcher

import (
	"fmt"
	"net/http"
	"time"

	"github.com/calebamiles/github-client/cache"
	"github.com/calebamiles/github-client/client/fetcher"
)

var _ fetcher.Fetcher = &DefaultCachingFetcher{}

const (
	cacheRequestHeader  = "If-None-Match"
	cacheResponseHeader = "ETag"
)

// NewCachingFetcher returns the default caching fetcher which understands how to use
// an authorization header to authenticate HTTP requests
func NewCachingFetcher(accessToken string, pageCache cache.PageCache) fetcher.Fetcher {
	return &DefaultCachingFetcher{
		cache:       pageCache,
		fetcher:     NewFetcher(accessToken),
		accessToken: accessToken,
	}
}

// DefaultCachingFetcher uses a paginator to fetch all reachable pages from a base URL
// and uses a PageCache to store previously fetched pages
type DefaultCachingFetcher struct {
	cache       cache.PageCache
	fetcher     fetcher.Fetcher
	accessToken string
}

// Fetch fetches all pages reachable from url, according to the PaginationFunc
func (f *DefaultCachingFetcher) Fetch(url string) ([]byte, error) {
	err := f.cache.Open() // make sure the cache is open before attempting to use it, relies on PageCache.Open() being idempotent
	if err != nil {
		return nil, err
	}

	cacheKey, err := f.cache.KeyForPage(url)
	if err != nil {
		return nil, err
	}

	updatedCacheKey, contentUnchanged, err := f.pageContentUnchanged(url, cacheKey)
	if err != nil {
		return nil, err
	}

	if contentUnchanged {
		cachedPage, cacheFetchErr := f.cache.FetchPageByKey(cacheKey)
		if cacheFetchErr != nil {
			return nil, cacheFetchErr
		}

		return cachedPage, nil
	}

	updatedPage, err := f.fetcher.Fetch(url)
	if err != nil {
		return nil, err
	}

	err = f.cache.AddPage(url, updatedCacheKey, updatedPage)
	if err != nil {
		return nil, err
	}

	return updatedPage, nil
}

func (f *DefaultCachingFetcher) pageContentUnchanged(url string, etag string) (string, bool, error) {
	if etag == "" {
		return "", false, nil
	}

	c := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", false, err
	}

	req.Header.Set(authorizationHeader, fmt.Sprintf("token %s", f.accessToken))
	req.Header.Set(cacheRequestHeader, etag)
	resp, err := c.Do(req)
	if err != nil {
		return "", false, err
	}

	defer resp.Body.Close()
	unchanged := (resp.StatusCode == http.StatusNotModified)
	currentEtag := resp.Header.Get(cacheResponseHeader)

	return currentEtag, unchanged, nil
}
