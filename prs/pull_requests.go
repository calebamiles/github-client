package prs

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/calebamiles/github-client/handle"
	"github.com/calebamiles/github-client/labels"
	"github.com/calebamiles/github-client/milestone"
	"github.com/calebamiles/github-client/state"
)

// A PullRequestWithoutComments contains basic information about a GitHub pull request
type PullRequest interface {
	ID() string
	Author() string
	Body() string
	Title() string
	Open() bool
	Merged() bool
	Labels() []labels.Label
	Milestone() milestone.Milestone
	Reviewers() []string
	CreatedAt() time.Time
	UpdatedAt() time.Time
	String() string

	//TODO add expensive methods for fetching richer objects here
	NumberOfCommits() int
	NumberOfAdditions() int
	NumberOfDeletions() int
	NumberOfChangedFiles() int
	NumberOfComments() int
}

// New returns a slice of PullRequestWithoutComments from raw JSON
func New(rawJSON []byte) ([]PullRequest, error) {
	var pulls []PullRequest
	ps := []*pullRequest{}

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

type pullRequest struct {
	Number                   int             `json:"number"`
	State                    string          `json:"state"`
	TitleString              string          `json:"title"`
	BodyString               string          `json:"body"`
	NumberOfCommentsJSON     int             `json:"comments"`
	NumberOfCommitsJSON      int             `json:"commits"`
	NumberOfAdditionsJSON    int             `json:"additions"`
	NumberOfDeletionsJSON    int             `json:"deletions"`
	NumberOfChangedFilesJSON int             `json:"changed_files"`
	CreatedTime              time.Time       `json:"created_at"`
	UpdatedTime              time.Time       `json:"updated_at"`
	MergeTime                *time.Time      `json:"merged_at"` // `null` unless merged so make this a pointer to a time
	Assignee                 handle.GitHub   `json:"assignee"`
	AssigneeList             []handle.GitHub `json:"reviewers"`
	MilestoneJSON            json.RawMessage `json:"milestone"`
	handle.GitHub            `json:"user"`

	labels    []labels.Label
	milestone milestone.Milestone
	reviewers []string
}

func (p *pullRequest) ID() string                     { return fmt.Sprintf("%d", p.Number) }
func (p *pullRequest) Author() string                 { return p.GitHub.ID }
func (p *pullRequest) Body() string                   { return p.BodyString }
func (p *pullRequest) Title() string                  { return p.TitleString }
func (p *pullRequest) Open() bool                     { return strings.EqualFold(p.State, state.Open) }
func (p *pullRequest) Merged() bool                   { return !p.MergeTime.IsZero() }
func (p *pullRequest) Labels() []labels.Label         { return p.labels }
func (p *pullRequest) Milestone() milestone.Milestone { return p.milestone }
func (p *pullRequest) Reviewers() []string            { return p.reviewers }
func (p *pullRequest) CreatedAt() time.Time           { return p.CreatedTime }
func (p *pullRequest) UpdatedAt() time.Time           { return p.UpdatedTime }

func (p *pullRequest) NumberOfComments() int     { return p.NumberOfCommentsJSON }
func (p *pullRequest) NumberOfCommits() int      { return p.NumberOfCommitsJSON }
func (p *pullRequest) NumberOfAdditions() int    { return p.NumberOfAdditionsJSON }
func (p *pullRequest) NumberOfDeletions() int    { return p.NumberOfDeletionsJSON }
func (p *pullRequest) NumberOfChangedFiles() int { return p.NumberOfChangedFilesJSON }

func (p *pullRequest) String() string { return fmt.Sprintf("%d", p.Number) }
