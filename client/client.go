package client

import "time"

const (
	// DateFormat is an ISO 8601 formatish format
	DateFormat = "2006-01-02T15:04:05-0700"
)

// A Client returns all pages associated with a GitHub API request
type Client interface {
	FetchCommitsToPathSince(string, time.Time) ([][]byte, error)
}

// New returns a new github/client.Client that is ready for use
func New(repoOwner string, repoName string, accessToken string) Client {
	return &defaultclient{
		accessToken: accessToken,
		repoOwner:   repoOwner,
		repoName:    repoName,
	}
}

type defaultclient struct {
	accessToken string
	repoOwner   string
	repoName    string
}
