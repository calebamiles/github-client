package types

import "time"

/*
  Scratch space for working out data models, delete this when done
*/

type FileChange interface {
	FileSHA() string
	Filename() string
	NumDeletions() int
	NumAdditions() int
	NumChanges() int
}

type Commit interface {
	SHA() string
	AuthorID() string
	AuthorEmail() string
	Committer() string
	Message() string
	Date() time.Time
	ParentSHAs() []string
}

type Comment interface {
	ID() int
	Body() string
	Author() string
	CreatedAt() time.Time
	UpodatedAt() time.Time
}

type ReviewComment interface {
	Comment
	Filename() string
	CommitSHA() string
	OriginalCommitSHA() string
}

type Milestone interface {
	Title() string
	Description() string
	Id() int
	Number() int

	NumOpenIssues() int
	NumClosedIssues() int
	Open() bool

	CreatedAt() time.Time
	UpdatedAt() time.Time
	ClosedAt() time.Time
	DueAt() time.Time
}
