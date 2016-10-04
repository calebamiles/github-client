package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/calebamiles/github-client/client/fetcher"
	"github.com/calebamiles/github-client/client/paginator"
)

const (
	httpsScheme         = "https"
	authorizationHeader = "Authorization"
)

// NewFetcher returns the default fetcher which understands how to use
// an authorization header to authenticate HTTP requests
func NewFetcher(accessToken string) fetcher.Fetcher {
	return &defaultFetcher{
		accessToken: accessToken,
		paginate:    PaginateGitHubResponse,
	}
}

type defaultFetcher struct {
	accessToken string
	paginate    paginator.PaginationFunc
}

func (f *defaultFetcher) Fetch(url string) ([][]byte, error) {
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

		loopBody, nextPageLink, loopErr = f.paginate(loopResp)
		if loopErr != nil {
			return nil, loopErr
		}

		allPageBodies = append(allPageBodies, loopBody)
	}

	return allPageBodies, nil
}
