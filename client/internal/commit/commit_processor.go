package commit

import (
	"sync"

	"github.com/calebamiles/github-client/client/fetcher"
	"github.com/calebamiles/github-client/comments"
	"github.com/calebamiles/github-client/commits"
)

// A Processor converts a slice of commits.CommitWithoutComments into a slice of commits.Commit
type Processor struct {
	errs    *errorAccumulator
	cs      *Accumulator
	waiter  *sync.WaitGroup
	ready   chan struct{}
	fetcher fetcher.Fetcher
}

// NewProcessor returns a processor ready to process commits.CommitWithoutComments
func NewProcessor(f fetcher.Fetcher) *Processor {
	return &Processor{
		errs:    &errorAccumulator{},
		cs:      &Accumulator{},
		waiter:  &sync.WaitGroup{},
		ready:   make(chan struct{}, 100),
		fetcher: f,
	}
}

// FetchAndAddComments fetches all comments for each element in a slice of CommitWithoutComments
func (p *Processor) FetchAndAddComments(in []commits.CommitWithoutComments) ([]commits.Commit, error) {
	for i := range in {
		p.waiter.Add(1)
		go processCommits(in[i], p.fetcher, p.cs, p.ready, p.waiter, p.errs)
	}

	p.waiter.Wait()
	if !p.errs.IsNil() {
		return nil, p.errs
	}

	return p.cs.GetAll(), nil
}

// AddEmptyComments adds an empty slice of comments for each element in a slice of CommitWithoutComments
func (p *Processor) AddEmptyComments(in []commits.CommitWithoutComments) ([]commits.Commit, error) {
	var allCommits []commits.Commit
	for i := range in {
		allCommits = append(allCommits, &commit{
			CommitWithoutComments: in[i],
			comments:              []comments.Comment{},
		})
	}

	return allCommits, nil
}

func processCommits(commitWithoutComments commits.CommitWithoutComments, f fetcher.Fetcher, cs *Accumulator, ready chan struct{}, wg *sync.WaitGroup, errs *errorAccumulator) {
	ready <- struct{}{}
	defer wg.Done()
	defer func(readyChan chan struct{}) { <-readyChan }(ready)

	if !errs.IsNil() {
		return
	}

	commentsPage, err := f.Fetch(commitWithoutComments.CommentsURL())
	if err != nil {
		errs.Add(err)
		return
	}

	allComments, err := comments.New(commentsPage)
	if err != nil {
		errs.Add(err)
		return
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
