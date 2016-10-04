package paginator

import "net/http"

// A PaginationFunc processes an http response and returns the current
// page and a link to the next page if it exists
type PaginationFunc func(*http.Response) (currentPage []byte, nextPageURL string, dontShadowThis error)
