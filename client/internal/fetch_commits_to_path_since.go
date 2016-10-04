package internal

import (
	"fmt"
	"net/url"
	"time"

	"github.com/calebamiles/github-client/comments"
	"github.com/calebamiles/github-client/commits"
)

// FetchCommitsToPathSince returns commits with comments to a path, an empty string for path will return all commits
func (c *DefaultClient) FetchCommitsToPathSince(path string, since time.Time) ([]commits.Commit, error) {
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

	var allCommits []commits.Commit
	var commitsWithoutComments []commits.CommitWithoutComments
	for i := range commitPages {
		commitsOnPage, loopErr := commits.New(commitPages[i])
		if loopErr != nil {
			return nil, loopErr
		}

		commitsWithoutComments = append(commitsWithoutComments, commitsOnPage...)
	}

	for i := range commitsWithoutComments {
		commentsURL := commitsWithoutComments[i].CommentsURL()
		var allComments []comments.Comment

		commentsPages, loopErr := c.Fetcher.Fetch(commentsURL)
		if loopErr != nil {
			return nil, loopErr
		}

		for j := range commentsPages {
			commentsOnPage, commentsLoopErr := comments.New(commentsPages[j])
			if commentsLoopErr != nil {
				return nil, commentsLoopErr
			}

			allComments = append(allComments, commentsOnPage...)
		}

		c := &commit{
			CommitWithoutComments: commitsWithoutComments[i],
			comments:              allComments,
		}

		allCommits = append(allCommits, c)
	}

	return allCommits, nil
}

type commit struct {
	commits.CommitWithoutComments
	comments []comments.Comment
}

func (c *commit) Comments() []comments.Comment { return c.comments }
