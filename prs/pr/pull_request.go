package pr

import (
	"time"

	"github.com/calebamiles/github-client/comments"
	"github.com/calebamiles/github-client/commits"
	"github.com/calebamiles/github-client/milestone"
	"github.com/calebamiles/github-client/prs/file"
)

// A PullRequest contains information about a GitHub pull request
type PullRequest interface {
	ID() string
	Author() string
	Body() string
	Title() string
	Open() bool
	Merged() bool
	Reviewers() []string

	CreatedAt() time.Time
	UpdatedAt() time.Time
	ClosedAt() time.Time
	MergedAt() time.Time

	Milestone() milestone.Milestone

	// expensive methods requiring an additional API call
	Commits() []commits.Commit
	Comments() []comments.Comment
	FilesChanged() []file.ChangeSet

	String() string
}
