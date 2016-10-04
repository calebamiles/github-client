package internal

import (
	"fmt"
	"net/url"
	"time"

	"github.com/calebamiles/github-client/comments"
	"github.com/calebamiles/github-client/prs"
)

// FetchPullRequestsSince returns all pull requests against the repository, optionally since a non zero time
func (c *DefaultClient) FetchPullRequestsSince(since time.Time) ([]prs.PullRequest, error) {
	urlString := fmt.Sprintf("https://api.github.com/repos/%s/%s/pullRequests", c.RepoOwner, c.RepoName)
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

	pullRequestPages, err := c.Fetcher.Fetch(u.String())
	if err != nil {
		return nil, err
	}

	var allpullRequests []prs.PullRequest
	var pullRequestsWithoutComments []prs.PullRequestWithoutComments
	for i := range pullRequestPages {
		pullRequestsOnPage, loopErr := prs.New(pullRequestPages[i])
		if loopErr != nil {
			return nil, loopErr
		}

		pullRequestsWithoutComments = append(pullRequestsWithoutComments, pullRequestsOnPage...)
	}

	for i := range pullRequestsWithoutComments {
		commentsURL := pullRequestsWithoutComments[i].CommentsURL()
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

		apullRequest := &pullRequest{
			PullRequestWithoutComments: pullRequestsWithoutComments[i],
			comments:                   allComments,
		}

		allpullRequests = append(allpullRequests, apullRequest)
	}

	return allpullRequests, nil
}

type pullRequest struct {
	prs.PullRequestWithoutComments
	comments []comments.Comment
}

func (i *pullRequest) Comments() []comments.Comment {
	return i.comments

}
