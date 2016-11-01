package pr

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/calebamiles/github-client/handle"
	"github.com/calebamiles/github-client/milestone"
	"github.com/calebamiles/github-client/state"
)

type BasicPullRequest interface {
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

	CommitsURL() string
	CommentsURL() string
	ReviewCommentsURL() string
	FileChangeURL() string
	URL() string

	Milestone() milestone.Milestone

	String() string
}

func NewBasicPullRequest(rawJSON json.RawMessage) (BasicPullRequest, error) {
	basicPR := &basicPullRequest{}

	err := json.Unmarshal(rawJSON, &basicPR)
	if err != nil {
		log.Printf("failed unmarshalling: \n %s\n", string(rawJSON))
		return nil, err
	}

	basicPR.numberString = fmt.Sprintf("%d", basicPR.Number)

	if basicPR.MergeTime == nil {
		basicPR.MergeTime = &time.Time{}
	}

	reviewersSet := make(map[string]bool)
	if basicPR.Assignee.ID != "" {
		reviewersSet[basicPR.Assignee.ID] = true

	}

	for j := range basicPR.AssigneeList {
		reviewersSet[basicPR.AssigneeList[j].ID] = true
	}

	for reviewer := range reviewersSet {
		basicPR.reviewers = append(basicPR.reviewers, reviewer)
	}

	// add the milestone
	m, loopErr := milestone.New(basicPR.MilestoneJSON)
	if loopErr != nil {
		return nil, loopErr
	}

	basicPR.milestone = m
	basicPR.fileChangesURL = fmt.Sprintf("%s/files", basicPR.BaseAPIURL)

	return basicPR, nil
}

type basicPullRequest struct {
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
	ClosedTime               time.Time       `json:"closed_at"`
	MergeTime                *time.Time      `json:"merged_at"` // `null` unless merged so make this a pointer to a time
	Assignee                 handle.GitHub   `json:"assignee"`
	AssigneeList             []handle.GitHub `json:"reviewers"`
	MilestoneJSON            json.RawMessage `json:"milestone"`
	GitHubAuthor             handle.GitHub   `json:"user"`

	BaseAPIURL              string `json:"url"`
	CommitsURLString        string `json:"commits_url"`
	ReviewCommentsURLString string `json:"review_comments_url"`
	CommentsURLString       string `json:"comments_url"`

	fileChangesURL string
	numberString   string
	reviewers      []string
	milestone      milestone.Milestone
}

func (p *basicPullRequest) ID() string                     { return p.numberString }
func (p *basicPullRequest) Author() string                 { return p.GitHubAuthor.ID }
func (p *basicPullRequest) Body() string                   { return p.BodyString }
func (p *basicPullRequest) Title() string                  { return p.TitleString }
func (p *basicPullRequest) Open() bool                     { return p.State == state.Open }
func (p *basicPullRequest) Merged() bool                   { return p.MergeTime.IsZero() }
func (p *basicPullRequest) Reviewers() []string            { return p.reviewers }
func (p *basicPullRequest) CreatedAt() time.Time           { return p.CreatedTime }
func (p *basicPullRequest) UpdatedAt() time.Time           { return p.UpdatedTime }
func (p *basicPullRequest) ClosedAt() time.Time            { return p.ClosedTime }
func (p *basicPullRequest) MergedAt() time.Time            { return p.ClosedTime }
func (p *basicPullRequest) URL() string                    { return p.BaseAPIURL }
func (p *basicPullRequest) CommitsURL() string             { return p.CommitsURLString }
func (p *basicPullRequest) CommentsURL() string            { return p.CommentsURLString }
func (p *basicPullRequest) ReviewCommentsURL() string      { return p.ReviewCommentsURLString }
func (p *basicPullRequest) FileChangeURL() string          { return p.fileChangesURL }
func (p *basicPullRequest) Milestone() milestone.Milestone { return p.milestone }
func (p *basicPullRequest) String() string                 { return p.numberString }
