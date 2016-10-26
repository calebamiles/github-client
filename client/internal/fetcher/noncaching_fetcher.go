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
func NewFetcher(AccessToken string) fetcher.Fetcher {
	return &DefaultFetcher{
		AccessToken: AccessToken,
		Paginate:    paginator.PaginateGitHubResponse,
	}
}

// DefaultFetcher uses a paginator to fetch all reachable pages from a base URL
type DefaultFetcher struct {
	AccessToken string
	Paginate    abstract.PaginationFunc
}

// Fetch fetches all pages reachable from url, according to the PaginationFunc
func (f *DefaultFetcher) Fetch(url string) ([]byte, error) {
	var allPageBodies [][]byte
	var c *http.Client

	c = &http.Client{Timeout: 10 * time.Second}

	var nextPageLink string
	var loopReq *http.Request
	var loopResp *http.Response
	var loopBody []byte
	var loopErr error

	for currentLink := url; len(currentLink) > 0; currentLink = nextPageLink {
		loopReq, loopErr = http.NewRequest(http.MethodGet, currentLink, nil)
		if loopErr != nil {
			return nil, loopErr
		}

		loopReq.Header.Set(authorizationHeader, fmt.Sprintf("token %s", f.AccessToken))
		loopResp, loopErr = c.Do(loopReq)
		if loopErr != nil {
			return nil, loopErr
		}

		if loopResp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("got unexpected status code: %d, desired 200", loopResp.StatusCode)
		}

		loopBody, nextPageLink, loopErr = f.Paginate(loopResp)
		if loopErr != nil {
			return nil, loopErr
		}

		loopErr = loopResp.Body.Close()
		if loopErr != nil {
			return nil, loopErr
		}

		allPageBodies = append(allPageBodies, loopBody)
	}

	joinedPages := pages.Join(allPageBodies)

	return joinedPages, nil
}
