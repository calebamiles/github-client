package issues

import (
	"encoding/json"
	"time"

	"github.com/calebamiles/github-client/label"
	"github.com/calebamiles/github-client/milestone"
)

type Issue interface {
	ID() string
	Author() string
	Body() string
	Comments() []comment.Comment
	Title() string
	Open() bool
	Labels() []label.Label
	Milestone() milestone.Milestone
	Assignees() []string
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

type issue struct {
	IssueNumber        string          `json:"number"`
	IssueTitle         string          `json:"title"`
	StateString        string          `json:"state"`
	AssigneeString     string          `json:"assignee"`
	AssigneeStringList []string        `json:"assignees"`
	CreatedAtString    time.Time       `json:"created_at"`
	UpdatedAtString    time.Time       `json:"updated_at"`
	CommentsURL        string          `json:"comments_url"`
	LabelsList         json.RawMessage `json:"labels"`
	Milestone          json.RawMessage `json:"milestone"`
	PullRequest        json.RawMessage `json:"pull_request"` // allows us to filter out PR issues which don't belong here
}

type issueAuthor struct {
	GithubID string `json:"login"`
}
