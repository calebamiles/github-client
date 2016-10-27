package prs

import (
	"encoding/json"
	"log"
	"time"

	"github.com/calebamiles/github-client/commits"
	"github.com/calebamiles/github-client/labels"
	"github.com/calebamiles/github-client/milestone"
)

// A FileChange represents a diff to a file associated with a Pull Request
type FileChange interface {
	FileSHA() string
	Filename() string
	NumDeletions() int
	NumAdditions() int
	NumChanges() int
}

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

	String() string

	Commits() []commits.Commit
	Comments() []comments.Comment
	ReviewComments() []comments.ReviewComment
	Labels() []labels.Label
	Milestone() milestone.Milestone
	FilesChanged() []FileChange
}

type PullFoo struct {
	ID        string
	Author    string
	Body      string
	Title     string
	Open      bool
	Merged    bool
	Reviewers []string

	CreatedAt time.Time
	UpdatedAt time.Time
	ClosedAt  time.Time
	MergedAt  time.Time
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

/*
	TODO: decide how to build up a PR from here
*/

// type pullRequest struct {
// 	Number                   int             `json:"number"`
// 	State                    string          `json:"state"`
// 	TitleString              string          `json:"title"`
// 	BodyString               string          `json:"body"`
// 	NumberOfCommentsJSON     int             `json:"comments"`
// 	NumberOfCommitsJSON      int             `json:"commits"`
// 	NumberOfAdditionsJSON    int             `json:"additions"`
// 	NumberOfDeletionsJSON    int             `json:"deletions"`
// 	NumberOfChangedFilesJSON int             `json:"changed_files"`
// 	CreatedTime              time.Time       `json:"created_at"`
// 	UpdatedTime              time.Time       `json:"updated_at"`
// 	MergeTime                *time.Time      `json:"merged_at"` // `null` unless merged so make this a pointer to a time
// 	Assignee                 handle.GitHub   `json:"assignee"`
// 	AssigneeList             []handle.GitHub `json:"reviewers"`
// 	MilestoneJSON            json.RawMessage `json:"milestone"`
// 	handle.GitHub            `json:"user"`
//
// 	labels    []labels.Label
// 	milestone milestone.Milestone
// 	reviewers []string
// }
