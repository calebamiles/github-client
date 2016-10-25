package client

import (
	"fmt"
	"net/url"
	"time"

	"github.com/calebamiles/github-client/client/internal/pages"
	"github.com/calebamiles/github-client/comments"
	"github.com/calebamiles/github-client/issues"
)

// FetchIssuesSince returns all issues against the repository, optionally since a non zero time
func (c *DefaultClient) FetchIssuesSince(since time.Time) ([]issues.Issue, error) {
	urlString := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", c.RepoOwner, c.RepoName)
	u, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	params := &url.Values{}
	if !since.IsZero() {
		params.Add("since", since.Format(DateFormat))
	}

	params.Add("per_page", NumberOfPagesToRequest)
	u.RawQuery = params.Encode()

	issuePages, err := c.Fetcher.Fetch(u.String())
	if err != nil {
		return nil, err
	}

	joinedIssuePages := pages.Join(issuePages)
	issuesWithoutComments, err := issues.New(joinedIssuePages)
	if err != nil {
		return nil, err
	}

	var allIssues []issues.Issue
	for i := range issuesWithoutComments {
		commentsURL := issuesWithoutComments[i].CommentsURL()
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
