package prs

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/calebamiles/github-client/prs/pr"
)

// NewBasicPullRequests returns a slice of BasicPullRequest from raw JSON
func NewBasicPullRequests(msgs []json.RawMessage) ([]pr.BasicPullRequest, error) {
	accumulator := newPRAccumulator(len(msgs))

	for i := range msgs {
		go func(msg json.RawMessage) {
			pull, err := pr.NewBasicPullRequest(msg)
			accumulator.Add(pull, err)
		}(msgs[i])
	}

	pullRequests, err := accumulator.Wait()
	return pullRequests, err
}

func newPRAccumulator(waitForNumber int) *prAccumulator {
	accumulator := &prAccumulator{}
	accumulator.wg.Add(waitForNumber)

	return accumulator
}

type prAccumulator struct {
	pullRequests []pr.BasicPullRequest
	errors       []error

	wg sync.WaitGroup
	sync.Mutex
}

func (a *prAccumulator) Add(p pr.BasicPullRequest, err error) {
	a.Lock()
	defer a.Unlock()

	if p != nil {
		a.pullRequests = append(a.pullRequests, p)
	}

	if err != nil {
		a.errors = append(a.errors, err)
	}

	a.wg.Done()
}

func (a *prAccumulator) Wait() ([]pr.BasicPullRequest, error) {
	a.wg.Wait()

	return a.pullRequests, combineError(a.errors)
}

func combineError(errs []error) error {
	if len(errs) == 0 {
		return nil
	}

	var errStrings []string
	for i := range errs {
		errStrings = append(errStrings, errs[i].Error())
	}

	joinedErrStrings := strings.Join(errStrings, "\n")
	return fmt.Errorf("errors were returned while processing pull requests: \n %s", joinedErrStrings)
}
