package client

import (
	"sync"

	"github.com/calebamiles/github-client/comments"
	"github.com/calebamiles/github-client/commits"
)

func (c *DefaultClient) processCommits(commitWithoutComments commits.CommitWithoutComments, cs *commitAccumulator, ready chan struct{}, wg *sync.WaitGroup, errs *errorAccumulator) {
	defer wg.Done()
	ready <- struct{}{}
	defer func(readyChan chan struct{}) { <-readyChan }(ready)

	var allComments []comments.Comment

	commentsPages, err := c.Fetcher.Fetch(commitWithoutComments.CommentsURL())
	if err != nil {
		errs.Add(err)
		return
	}

	for j := range commentsPages {
		commentsOnPage, commentsLoopErr := comments.New(commentsPages[j])
		if commentsLoopErr != nil {
			errs.Add(commentsLoopErr)
			return
		}

		allComments = append(allComments, commentsOnPage...)
	}

	commitToAdd := &commit{
		CommitWithoutComments: commitWithoutComments,
		comments:              allComments,
	}

	cs.Add(commitToAdd)
	return
}

type commit struct {
	commits.CommitWithoutComments
	comments []comments.Comment
}

func (c *commit) Comments() []comments.Comment { return c.comments }
