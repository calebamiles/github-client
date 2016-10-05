package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/calebamiles/github-client/client/fetcher"
	"github.com/calebamiles/github-client/client/paginator"
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
		Paginate:    PaginateGitHubResponse,
	}
}

// DefaultFetcher uses a paginator to fetch all reachable pages from a base URL
type DefaultFetcher struct {
	AccessToken string
	Paginate    paginator.PaginationFunc
}

// Fetch fetches all pages reachable from url, according to the PaginationFunc
func (f *DefaultFetcher) Fetch(url string) ([][]byte, error) {
	var allPageBodies [][]byte
	var c *http.Client

	c = &http.Client{Timeout: 10 * time.Second}

	var nextPageLink string
	var loopReq *http.Request
	var loopResp *http.Response
	var loopBody []byte
	var loopErr error

	for nextPageLink = url; nextPageLink != ""; {
		loopReq, loopErr = http.NewRequest(http.MethodGet, nextPageLink, nil)
		if loopErr != nil {
			return nil, loopErr
		}

		loopReq.Header.Set(authorizationHeader, fmt.Sprintf("token %s", f.AccessToken))
		loopResp, loopErr = c.Do(loopReq)
		if loopErr != nil {
			return nil, loopErr
		}

		loopBody, nextPageLink, loopErr = f.Paginate(loopResp)
		if loopErr != nil {
			return nil, loopErr
		}

		allPageBodies = append(allPageBodies, loopBody)
	}

	return allPageBodies, nil
}
