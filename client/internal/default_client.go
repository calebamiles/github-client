package internal

import "github.com/calebamiles/github-client/client/fetcher"

const (
	// DateFormat is an ISO 8601 formatish format
	DateFormat             = "2006-01-02T15:04:05-0700"
	NumberOfPagesToRequest = "1000"
)

type DefaultClient struct {
	Fetcher     fetcher.Fetcher
	AccessToken string
	RepoOwner   string
	RepoName    string
}
