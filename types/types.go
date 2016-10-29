package types

import "time"

/*
	TODO:
	- add BasicPullRequest back
	- add New(BasicPullRequest, []comments.Comment, []comments.ReviewComment, []commits.Commit, []file.Change) PullRequest ? (maybe create subpackage for proper pluralization)
	- update client to fetch additional fields
	- check if richer model should be returned for issues
	-
*/

/*
  Scratch space for working out data models, delete this when done
*/

type Commit interface {
	SHA() string
	AuthorID() string
	AuthorEmail() string
	CommitterID() string
	Message() string
	Date() time.Time
	ParentSHAs() []string

	Comments() []Comment
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
