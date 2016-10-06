package commit

import (
	"sync"

	"github.com/calebamiles/github-client/commits"
)

// An Accumulator is a thread safe storage structure for commits.Commit
type Accumulator struct {
	sync.Mutex

	commits []commits.Commit
}

// Add a commit to the accumulator
func (c *Accumulator) Add(commit commits.Commit) {
	c.Lock()
	defer c.Unlock()

	c.commits = append(c.commits, commit)
}

// GetAll returns all commits from the accumulator
func (c *Accumulator) GetAll() []commits.Commit {
	c.Lock()
	defer c.Unlock()

	return c.commits
}
