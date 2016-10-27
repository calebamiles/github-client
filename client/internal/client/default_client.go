package client

import (
	"os"
	"sync"

	"github.com/calebamiles/github-client/client/fetcher"
)

const (
	// DateFormat is an ISO 8601 formatish format
	DateFormat = "2006-01-02T15:04:05-0700"

	// NumberOfPagesToRequest in an API request
	NumberOfPagesToRequest = "1000"
)

// DefaultClient is returned by client.New()
type DefaultClient struct {
	Fetcher           fetcher.Fetcher
	RepoOwner         string
	RepoName          string
	TempCacheFilePath string
	doneOnce          sync.Once
}

// Done closes the cache and removes the temporary datastore if it exists
func (c *DefaultClient) Done() error {
	var err error

	c.doneOnce.Do(func() {
		err = c.Fetcher.Done()
		if err != nil {
			return
		}

		if c.TempCacheFilePath != "" {
			err = os.Remove(c.TempCacheFilePath)
			if err != nil {
				return
			}
		}
	})

	return err
}
