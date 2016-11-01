package client

import (
	"fmt"
	"net/url"

	"github.com/calebamiles/github-client/comments"
	"github.com/calebamiles/github-client/issues"
)

// FetchIssuesSince returns all issues against the repository
func (c *DefaultClient) FetchIssues() ([]issues.Issue, error) {
	urlString := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", c.RepoOwner, c.RepoName)
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	params := &url.Values{}
	params.Add("per_page", NumberOfPagesToRequest)
	u.RawQuery = params.Encode()

	issuePages, _, err := c.Fetcher.Fetch(u.String())
	if err != nil {
		return nil, err
	}

	issuesWithoutComments, err := issues.New(issuePages)
	if err != nil {
		return nil, err
	}

	var allIssues []issues.Issue
	for i := range issuesWithoutComments {
		commentsURL := issuesWithoutComments[i].CommentsURL()

		commentsPages, _, loopErr := c.Fetcher.Fetch(commentsURL)
		if loopErr != nil {
			return nil, loopErr
		}

		allComments, loopErr := comments.New(commentsPages)
		if loopErr != nil {
			return nil, loopErr
		}

		anIssue := &issue{
			IssueWithoutComments: issuesWithoutComments[i],
			comments:             allComments,
		}

		allIssues = append(allIssues, anIssue)
	}

	return allIssues, nil
}

type issue struct {
	issues.IssueWithoutComments
	comments []comments.Comment
}

func (i *issue) Comments() []comments.Comment {
	return i.comments
}
