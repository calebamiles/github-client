package prs

import (
	"time"

	"github.com/calebamiles/github-client/author"
	"github.com/calebamiles/github-client/comment"
	"github.com/calebamiles/github-client/label"
	"github.com/calebamiles/github-client/milestone"
)

type PullRequest interface {
	ID() string
	Author() author.Author
	Body() string
	Comments() []comment.Comment
	Open() bool
	Merged() bool
	Labels() []label.Label
	Milestone() milestone.Milestone
	Reviewers() []string
	CreatedAt() time.Time
	UpdatedAt() time.Time
	String() string
}
