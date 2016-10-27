package fetcher

import (
	"fmt"
	"net/http"
	"time"

	"github.com/calebamiles/github-client/client/fetcher"
	"github.com/calebamiles/github-client/client/internal/pages"
	"github.com/calebamiles/github-client/client/internal/paginator"
	abstract "github.com/calebamiles/github-client/client/paginator"
)

var _ fetcher.Fetcher = &DefaultFetcher{}

const (
	authorizationHeader = "Authorization"
)

// NewFetcher returns the default fetcher which understands how to use
// an authorization header to authenticate HTTP requests
func NewFetcher(accessToken string) fetcher.Fetcher {
	return &DefaultFetcher{
		AccessToken: accessToken,
		Paginate:    paginator.PaginateGitHubResponse,
	}
}

// DefaultFetcher uses a paginator to fetch all reachable pages from a base URL
type DefaultFetcher struct {
	AccessToken string
	Paginate    abstract.PaginationFunc
}

// Fetch fetches all pages reachable from url, according to the PaginationFunc
func (f *DefaultFetcher) Fetch(url string) ([]byte, string, error) {
	var allPageBodies [][]byte
	var c *http.Client
	var firstPageEtag string

	c = &http.Client{Timeout: 10 * time.Second}

	var nextPageLink string
	var loopReq *http.Request
	var loopResp *http.Response
	var loopBody []byte
	var loopErr error

	for currentLink := url; currentLink != ""; currentLink = nextPageLink {
		loopReq, loopErr = http.NewRequest(http.MethodGet, currentLink, nil)
		if loopErr != nil {
			return nil, "", loopErr
		}

		if f.AccessToken != "" {
			loopReq.Header.Set(authorizationHeader, fmt.Sprintf("token %s", f.AccessToken))
		}

		loopResp, loopErr = c.Do(loopReq)
		if loopErr != nil {
			return nil, "", loopErr
		}

		// ugly hack to special case returning the etag for the
		if currentLink == url {
			firstPageEtag = loopResp.Header.Get(CacheResponseHeader)
		}

		if loopResp.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("got unexpected status code: %d, desired 200", loopResp.StatusCode)
		}

		loopBody, nextPageLink, loopErr = f.Paginate(loopResp)
		if loopErr != nil {
			return nil, "", loopErr
		}

		loopErr = loopResp.Body.Close()
		if loopErr != nil {
			return nil, "", loopErr
		}

		allPageBodies = append(allPageBodies, loopBody)
	}

	joinedPages := pages.Join(allPageBodies)

	return joinedPages, firstPageEtag, nil
}

// Done has nothing to close here, so NoOp
func (f *DefaultFetcher) Done() error { return nil }
