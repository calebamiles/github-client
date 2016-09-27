package client

import (
	"fmt"
	"net/url"
	"time"
)

const (
	// DateFormat is an ISO 8601 formatish format
	DateFormat = "2006-01-02T15:04:05-0700"
)

type defaultclient struct {
	accessToken string
}

func (c *defaultclient) FetchCommitsToPathSince(repoOwner string, repoName string, path string, since time.Time) ([][]byte, error) {
	f := NewFetcher()
	urlString := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", repoOwner, repoName)
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	params := &url.Values{}
	//TODO add test for this
	if path != "." && path != "" {
		params.Add("path", path)
	}

	params.Add("since", since.Format(DateFormat))
	params.Add("access_token", c.accessToken)
	params.Add("per_page", fmt.Sprintf("%d", 1000)) //TODO get this tuning variable out of here

	u.RawQuery = params.Encode()

	commitPages, err := f.Fetch(PaginateGitHubResponse, u.String())
	if err != nil {
		return nil, err
	}

	return commitPages, nil
}

// having a second user of a GitHub client begs some interesting
// questions for reuse. It would seem that in both the findowners work
// and in the Tracker annotation work that we're really only interested
// in actions against the repository so maybe it's best to remove
// the raw "client" interface and expose more interesting actions
// that can be narrowly consumed downstream e.g.
/*

type Repository interface {
	// some interesting getters
}

type Comment interface {
	Body() string
}

type Issue interface {
	Description() string
	Comments() []Comment
}

type IssueFetcher interface {
	FetchIssues() ([]Issue, error)
}

*/

// creating a client will likely involve specifiying the repo and
// some secret for connecting but that messiness could probably be
// hidden in someone's main() :)

// A Client knows how to support interesting interactions with the GitHub API
// type Client interface {
// 	// TODO flesh this out around time of extraction
// }
//
// // NewClient returns a new github/client.Client that is ready for use
// func NewClient(accessToken string) Client {
// 	return &defaultclient{accessToken: accessToken}
// }
