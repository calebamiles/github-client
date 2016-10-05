package client

import (
	"sync"

	"github.com/calebamiles/github-client/commits"
)

type commitAccumulator struct {
	sync.Mutex

	commits []commits.Commit
}

func (c *commitAccumulator) Add(commit commits.Commit) {
	c.Lock()
	defer c.Unlock()

	c.commits = append(c.commits, commit)
}

func (c *commitAccumulator) GetAll() []commits.Commit {
	c.Lock()
	defer c.Unlock()

	return c.commits
}
