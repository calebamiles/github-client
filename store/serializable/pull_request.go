package serializable

import (
	"time"

	"github.com/calebamiles/github-client/comments"
	"github.com/calebamiles/github-client/commits"
	"github.com/calebamiles/github-client/labels"
	"github.com/calebamiles/github-client/milestone"
	"github.com/calebamiles/github-client/prs/file"
)

type serializablePullRequest struct {
	SerializedID        string
	SerializedAuthor    string
	SerializedBody      string
	SerializedTitle     string
	SerializedOpen      bool
	SerializedMerged    bool
	SerializedReviewers []string

	SerializedCreatedAt time.Time
	SerializedUpdatedAt time.Time
	SerializedClosedAt  time.Time
	SerializedMergedAt  time.Time

	SerializedCommits      []*serializableCommit
	SerializedComments     []*serializableComment
	SerializedLabels       []*serializableLabel
	SerializedMilestone    *serializableilestone
	SerializedFilesChanged []*serializableFileChange
}

func (p *serializablePullRequest) ID() string          { return p.SerializedID }
func (p *serializablePullRequest) Author() string      { return p.SerializedAuthor }
func (p *serializablePullRequest) Body() string        { return p.SerializedBody }
func (p *serializablePullRequest) Title() string       { return p.SerializedTitle }
func (p *serializablePullRequest) Open() bool          { return p.SerializedOpen }
func (p *serializablePullRequest) Merged() bool        { return p.SerializedMerged }
func (p *serializablePullRequest) Reviewers() []string { return p.SerializedReviewers }

func (p *serializablePullRequest) CreatedAt() time.Time { return p.SerializedCreatedAt }
func (p *serializablePullRequest) UpdatedAt() time.Time { return p.SerializedUpdatedAt }
func (p *serializablePullRequest) ClosedAt() time.Time  { return p.SerializedClosedAt }
func (p *serializablePullRequest) MergedAt() time.Time  { return p.SerializedMergedAt }

func (p *serializablePullRequest) Commits() []commits.Commit {
	var cs []commits.Commit
	for i := range p.SerializedCommits {
		cs = append(cs, p.SerializedCommits[i])
	}

	return cs

}

func (p *serializablePullRequest) Comments() []comments.Comment {
	var cs []comments.Comment
	for i := range p.SerializedComments {
		cs = append(cs, p.SerializedComments[i])
	}

	return cs
}

func (p *serializablePullRequest) Labels() []labels.Label {
	var ls []labels.Label
	for i := range p.SerializedLabels {
		ls = append(ls, p.SerializedLabels[i])
	}

	return ls
}

func (p *serializablePullRequest) Milestone() milestone.Milestone { return p.SerializedMilestone }

func (p *serializablePullRequest) FilesChanged() []file.Change {
	var fs []file.Change
	for i := range p.SerializedFilesChanged {
		fs = append(fs, p.SerializedFilesChanged[i])
	}

	return fs
}

func (p *serializablePullRequest) String() string { return p.SerializedID }

type serializableFileChange struct {
	SerializedFileSHA      string
	SerializedFilename     string
	SerializedNumDeletions int
	SerializedNumAdditions int
	SerializedNumChanges   int
}

func (c *serializableFileChange) FileSHA() string   { return c.SerializedFileSHA }
func (c *serializableFileChange) Filename() string  { return c.SerializedFilename }
func (c *serializableFileChange) NumDeletions() int { return c.SerializedNumDeletions }
func (c *serializableFileChange) NumAdditions() int { return c.SerializedNumAdditions }
func (c *serializableFileChange) NumChanges() int   { return c.SerializedNumChanges }
