package client

import (
	"fmt"
	"net/http"
	"time"
)

const (
	httpsScheme         = "https"
	authorizationHeader = "Authorization"
)

// A PaginationFunc processes an http response and returns the current
// page and a link to the next page if it exists
type PaginationFunc func(*http.Response) (currentPage []byte, nextPageURL string, dontShadowThis error)

// A Fetcher uses a paginator to return all pages reachable from a base URL
type Fetcher interface {
	Fetch(PaginationFunc, string) ([][]byte, error)
}

// NewFetcher returns the default fetcher which understands how to use
// an authorization header to authenticate HTTP requests
func NewFetcher(accessToken string) Fetcher {
	return &defaultFetcher{
		accessToken: accessToken,
	}
}

type defaultFetcher struct {
	accessToken string
}

func (f *defaultFetcher) Fetch(paginate PaginationFunc, url string) ([][]byte, error) {
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

		loopReq.Header.Set("Authorization", fmt.Sprintf("token %s", f.accessToken))
		loopResp, loopErr = c.Do(loopReq)
		if loopErr != nil {
			return nil, loopErr
		}

		loopBody, nextPageLink, loopErr = paginate(loopResp)
		if loopErr != nil {
			return nil, loopErr
		}

		allPageBodies = append(allPageBodies, loopBody)
	}

	return allPageBodies, nil
}
