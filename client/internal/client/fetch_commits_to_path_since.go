package client

import (
	"fmt"
	"net/url"
	"time"

	"github.com/calebamiles/github-client/client/internal/commit"
	"github.com/calebamiles/github-client/client/internal/pages"
	"github.com/calebamiles/github-client/commits"
)

// FetchCommitsToPathSince returns commits without comments to a path, an empty string for path will return all commits
func (c *DefaultClient) FetchCommitsToPathSince(path string, since time.Time) ([]commits.Commit, error) {
	return c.fetchCommitsToPathSince(path, since, false)
}

// FetchCommitsWithCommentsToPathSince returns commits with comments to a path, an empty string for path will return all commits
func (c *DefaultClient) FetchCommitsWithCommentsToPathSince(path string, since time.Time) ([]commits.Commit, error) {
	return c.fetchCommitsToPathSince(path, since, true)
}

func (c *DefaultClient) fetchCommitsToPathSince(path string, since time.Time, fetchComments bool) ([]commits.Commit, error) {
	urlString := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", c.RepoOwner, c.RepoName)
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	params := &url.Values{}
	if path != "" {
		params.Add("path", path)
	}

	if !since.IsZero() {
		params.Add("since", since.Format(DateFormat))
	}

	params.Add("per_page", NumberOfPagesToRequest)
	u.RawQuery = params.Encode()

	commitPages, err := c.Fetcher.Fetch(u.String())
	if err != nil {
		return nil, err
	}

	joinedCommitPages := pages.Join(commitPages)
	commitsWithoutComments, err := commits.New(joinedCommitPages)
	if err != nil {
		return nil, err
	}

	processor := commit.NewProcessor(c.Fetcher)
	if fetchComments {
		return processor.FetchAndAddComments(commitsWithoutComments)
	}

	return processor.AddEmptyComments(commitsWithoutComments)
}
