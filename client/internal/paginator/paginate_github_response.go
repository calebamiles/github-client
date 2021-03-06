package paginator

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/calebamiles/github-client/client/paginator"
)

var _ paginator.PaginationFunc = PaginateGitHubResponse

const (
	gitHubPaginationHeader = "Link"
	gitHubNextPageTag      = `rel="next"`
)

// PaginateGitHubResponse knows how to walk a HTTP response from GitHub to find multiple pages
func PaginateGitHubResponse(resp *http.Response) ([]byte, string, error) {
	var nextPageLink string
	var linkHeader string
	var headerLinks []string

	linkHeader = resp.Header.Get(gitHubPaginationHeader)
	headerLinks = strings.Split(linkHeader, ",")
	for i := range headerLinks {
		if strings.Contains(headerLinks[i], gitHubNextPageTag) {
			nonTrimLink := strings.Split(headerLinks[i], ";")[0] //TODO check whether this is a actually URL or not rather than assuming
			nextPageLink = strings.Trim(nonTrimLink, "<>")
		}
	}

	bodyBytes, err := ioutil.ReadAll(ioutil.NopCloser(resp.Body))
	if err != nil {
		return nil, "", err
	}

	log.Printf("processed a page %s, next page at: %s", resp.Request.URL.String(), nextPageLink)
	return bodyBytes, nextPageLink, nil
}
