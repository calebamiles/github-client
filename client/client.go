package client

import (
	"io/ioutil"

	"github.com/calebamiles/github-client/cache"
	"github.com/calebamiles/github-client/client/internal/client"
	"github.com/calebamiles/github-client/client/internal/fetcher"
	"github.com/calebamiles/github-client/prs/pr"
)

// A Client handles basic read only operations against the GitHub API
type Client interface {
	FetchPullRequests() (pullRequests []pr.PullRequest, err error)
	// FetchCommits() (commits []commits.Commit, err error)
	// FetchIssues() (issues []issues.Issue, err error)
	FetchPage(url string) (pageContent []byte, err error)
	Done() error
}

// New returns a new github/client.Client that is ready for use
// callers can control the location of the internal cache file by specifing a
// non empty cacheFile pointing to the desired cache file which will be
//  created if it doesn't exist
func New(repoOwner string, repoName string, accessToken string, cacheFile string) (Client, error) {
	c := &client.DefaultClient{}

	c.RepoOwner = repoOwner
	c.RepoName = repoName

	pathToCache := cacheFile
	if pathToCache == "" {
		f, err := ioutil.TempFile("", "github-client-cache")
		if err != nil {
			return nil, err
		}

		pathToCache = f.Name()
		err = f.Close()
		if err != nil {
			return nil, err
		}

		c.TempCacheFilePath = pathToCache
	}

	pageCache := cache.New(pathToCache)
	cachingFetcher := fetcher.NewCachingFetcher(accessToken, pageCache)
	c.Fetcher = cachingFetcher

	return c, nil
}

// NewNonCachingClient returns a new github/client.Client that is ready for use but does
// not cache page responses. The client returned is not recommended for use with large
// repositories due to rate limiting
func NewNonCachingClient(repoOwner string, repoName string, accessToken string) Client {
	return &client.DefaultClient{
		RepoOwner: repoOwner,
		RepoName:  repoName,
		Fetcher:   fetcher.NewFetcher(accessToken),
	}
}
