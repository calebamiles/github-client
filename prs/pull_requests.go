package prs

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

// A PullRequest contains basic information about a GitHub pull request, including comments
type PullRequest interface {
	PullRequestWithoutComments
	Comments() []comments.Comment
}

// A PullRequestWithoutComments contains basic information about a GitHub pull request
type PullRequestWithoutComments interface {
	ID() string
	Author() string
	Body() string
	Title() string
	CommentsURL() string
	Open() bool
	Merged() bool
	Labels() []labels.Label
	Milestone() milestone.Milestone
	Reviewers() []string
	CreatedAt() time.Time
	UpdatedAt() time.Time
	String() string
}

// New returns a slice of PullRequestWithoutComments from raw JSON
func New(rawJSON []byte) ([]PullRequestWithoutComments, error) {
	var pulls []PullRequestWithoutComments
	ps := []*pullrequest{}

	err := json.Unmarshal(rawJSON, &ps)
	if err != nil {
		log.Printf("failed unmarshalling: \n %s\n", string(rawJSON))
		return nil, err
	}

	for i := range ps {
		// in the case that the PR hasn't been merged set the pointer
		if ps[i].MergeTime == nil {
			ps[i].MergeTime = &time.Time{}
		}

		reviewersSet := make(map[string]bool)
		if ps[i].Assignee.ID != "" {
			reviewersSet[ps[i].Assignee.ID] = true

		}

		for j := range ps[i].AssigneeList {
			reviewersSet[ps[i].AssigneeList[j].ID] = true
		}

		for reviewer := range reviewersSet {
			ps[i].reviewers = append(ps[i].reviewers, reviewer)
		}

		// add the milestone
		m, loopErr := milestone.New(ps[i].MilestoneJSON)
		if loopErr != nil {
			return nil, loopErr
		}

		ps[i].milestone = m

		pulls = append(pulls, ps[i])
	}

	return pulls, nil
}

type pullrequest struct {
	Number            int             `json:"number"`
	State             string          `json:"state"`
	TitleString       string          `json:"title"`
	BodyString        string          `json:"body"`
	CreatedTime       time.Time       `json:"created_at"`
	UpdatedTime       time.Time       `json:"updated_at"`
	MergeTime         *time.Time      `json:"merged_at"` // `null` unless merged so make this a pointer to a time
	Assignee          handle.GitHub   `json:"assignee"`
	AssigneeList      []handle.GitHub `json:"reviewers"`
	CommentsURLString string          `json:"comments_url"`
	MilestoneJSON     json.RawMessage `json:"milestone"`
	handle.GitHub     `json:"user"`

	labels    []labels.Label
	milestone milestone.Milestone
	reviewers []string
}

func (p *pullrequest) ID() string                     { return fmt.Sprintf("%d", p.Number) }
func (p *pullrequest) Author() string                 { return p.GitHub.ID }
func (p *pullrequest) Body() string                   { return p.BodyString }
func (p *pullrequest) Title() string                  { return p.TitleString }
func (p *pullrequest) CommentsURL() string            { return p.CommentsURLString }
func (p *pullrequest) Open() bool                     { return strings.EqualFold(p.State, state.Open) }
func (p *pullrequest) Merged() bool                   { return !p.MergeTime.IsZero() }
func (p *pullrequest) Labels() []labels.Label         { return p.labels }
func (p *pullrequest) Milestone() milestone.Milestone { return p.milestone }
func (p *pullrequest) Reviewers() []string            { return p.reviewers }
func (p *pullrequest) CreatedAt() time.Time           { return p.CreatedTime }
func (p *pullrequest) UpdatedAt() time.Time           { return p.UpdatedTime }
func (p *pullrequest) String() string                 { return fmt.Sprintf("%d", p.Number) }
