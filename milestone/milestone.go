package milestone

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/calebamiles/github-client/state"
)

// A Milestone provides basic information about a GitHub milestone
type Milestone interface {
	Open() bool
	Description() string
	NumIssuesOpen() int
	NumIssuesClosed() int
	Title() string
	Deadline() time.Time

	String() string
}

// New returns a Milestone from raw JSON
func New(rawJSON []byte) (Milestone, error) {
	m := milestone{}

	if len(rawJSON) == 0 {
		return &m, nil
	}

	err := json.Unmarshal(rawJSON, &m)
	if err != nil {
		log.Printf("failed unmarshalling: \n %s\n", string(rawJSON))
		return nil, err
	}

	return &m, nil
}

type milestone struct {
	IDInt             int       `json:"id"`
	TitleString       string    `json:"title"`
	StateString       string    `json:"state"`
	DescriptionString string    `json:"description"`
	OpenIssuesInt     int       `json:"open_issues"`
	ClosedIssuesInt   int       `json:"closed_issues"`
	DueOn             time.Time `json:"due_on"`
	deadline          time.Time
	openIssues        int
	closedIssues      int
}

func (m *milestone) Open() bool           { return strings.EqualFold(m.StateString, state.Open) }
func (m *milestone) Description() string  { return m.DescriptionString }
func (m *milestone) NumIssuesOpen() int   { return m.OpenIssuesInt }
func (m *milestone) NumIssuesClosed() int { return m.ClosedIssuesInt }
func (m *milestone) Deadline() time.Time  { return m.DueOn }
func (m *milestone) Title() string        { return m.TitleString }
func (m *milestone) String() string       { return m.TitleString }
