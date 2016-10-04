package client

import (
	"net/url"
	"time"

	"github.com/calebamiles/github-client/commits"
	"github.com/calebamiles/github-client/issues"
	"github.com/calebamiles/github-client/prs"
)

// A Client handles basic read only operations against the GitHub API
type Client interface {
	FetchCommitsToPathSince(string, time.Time) ([]commits.Commit, error)
	FetchIssuesSince(time.Time) ([]issues.Issue, error)
	FetchPullRequestsSince(time.Time) ([]prs.PullRequest, error)
	FetchPages(url.URL) ([][]byte, error)
}

// New returns a new github/client.Client that is ready for use
// func New(repoOwner string, repoName string, accessToken string) Client {
// 	return &defaultclient{
// 		accessToken: accessToken,
// 		repoOwner:   repoOwner,
// 		repoName:    repoName,
// 	}
// }
//
