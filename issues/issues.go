package issues

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/calebamiles/github-client/comments"
	"github.com/calebamiles/github-client/handle"
	"github.com/calebamiles/github-client/labels"
	"github.com/calebamiles/github-client/milestone"
	"github.com/calebamiles/github-client/state"
)

// An Issue is a composition of a basic issue with comments added
type Issue interface {
	IssueWithoutComments
	Comments() []comments.Comment
}

// An IssueWithoutComments represents an issue as returned by the GitHub API
type IssueWithoutComments interface {
	ID() string
	OpenedBy() string
	Body() string
	CommentsURL() string
	Title() string
	Open() bool
	Labels() []labels.Label
	Milestone() milestone.Milestone
	Assignees() []string
	CreatedAt() time.Time
	UpdatedAt() time.Time
}

// New returns a slice of issues without comments from raw JSON
func New(rawJSON []byte) ([]IssueWithoutComments, error) {
	var issues []IssueWithoutComments
	is := []*issue{}

	err := json.Unmarshal(rawJSON, &is)
	if err != nil {
		log.Printf("failed unmarshalling: \n %s\n", string(rawJSON))
		return nil, err
	}

	for i := range is {
		// get rid of any pull requests returned in the JSON
		if len(is[i].PullRequest) != 0 {
			continue
		}

		assigneesSet := make(map[string]bool)
		if is[i].Assignee.ID != "" {
			assigneesSet[is[i].Assignee.ID] = true
		}

		for j := range is[i].AssigneeList {
			assigneesSet[is[i].AssigneeList[j].ID] = true
		}

		for assignee := range assigneesSet {
			is[i].assignees = append(is[i].assignees, assignee)
		}

		// add the unmarshalled labels
		lbls, loopErr := labels.New(is[i].LabelsList)
		if loopErr != nil {
			return nil, loopErr
		}

		is[i].labels = lbls

		// add the milestone
		m, loopErr := milestone.New(is[i].MilestoneJSON)
		if loopErr != nil {
			return nil, loopErr
		}

		is[i].milestone = m

		issues = append(issues, is[i])
	}

	return issues, nil
}

type issue struct {
	IssueNumber       int             `json:"number"`
	IssueTitle        string          `json:"title"`
	IssueBody         string          `json:"body"`
	StateString       string          `json:"state"`
	Assignee          handle.GitHub   `json:"assignee"`
	AssigneeList      []handle.GitHub `json:"assignees"`
	CreatedTime       time.Time       `json:"created_at"`
	UpdatedTime       time.Time       `json:"updated_at"`
	CommentsURLString string          `json:"comments_url"`
	LabelsList        json.RawMessage `json:"labels"`
	MilestoneJSON     json.RawMessage `json:"milestone"`
	PullRequest       json.RawMessage `json:"pull_request"` // allows us to filter out PR issues which don't belong here
	handle.GitHub     `json:"user"`

	labels    []labels.Label
	milestone milestone.Milestone
	assignees []string
}

func (i *issue) ID() string                     { return fmt.Sprintf("%d", i.IssueNumber) }
func (i *issue) OpenedBy() string               { return i.GitHub.ID }
func (i *issue) Title() string                  { return i.IssueTitle }
func (i *issue) Open() bool                     { return strings.EqualFold(i.StateString, state.Open) }
func (i *issue) Labels() []labels.Label         { return i.labels }
func (i *issue) Milestone() milestone.Milestone { return i.milestone }
func (i *issue) Assignees() []string            { return i.assignees }
func (i *issue) CreatedAt() time.Time           { return i.CreatedTime }
func (i *issue) UpdatedAt() time.Time           { return i.UpdatedTime }
func (i *issue) Body() string                   { return i.IssueBody }
func (i *issue) CommentsURL() string            { return i.CommentsURLString }
