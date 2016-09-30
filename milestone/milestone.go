package milestone

import (
	"encoding/json"
	"time"
)

type Milestone interface {
	Open() bool
	Description() string
	IssuesOpen() int
	IssuesClosed() int
	Deadline() time.Time
	Title() string
	String() string
}

func New(rawJSON json.RawMessage) (Milestone, error) {
	m := &milestone{}

	err := json.Unmarshal(rawJSON, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

var foo = `
"url": "https://api.github.com/repos/octocat/Hello-World/milestones/1",
		 "html_url": "https://github.com/octocat/Hello-World/milestones/v1.0",
		 "labels_url": "https://api.github.com/repos/octocat/Hello-World/milestones/1/labels",
		 "id": 1002604,
		 "number": 1,
		 "state": "open",
		 "title": "v1.0",
		 "description": "Tracking milestone for version 1.0",
		 "open_issues": 4,
		 "closed_issues": 8,
		 "created_at": "2011-04-10T20:09:31Z",
		 "updated_at": "2014-03-03T18:58:10Z",
		 "closed_at": "2013-02-12T13:22:01Z",
		 "due_on": "2012-10-09T23:39:01Z"


`

type milestone struct {
	IDString           string    `json:"id"`
	TitleString        string    `json:"title"`
	StateString        string    `json:"state"`
	DescriptionString  string    `json:"description"`
	OpenIssuesString   string    `json:"open_issues"`
	ClosedIssuesString string    `json:"closed_issues"`
	DueOn              time.Time `json:"due_on"`
	deadline           time.Time
	openIssues         int
	closedIssues       int
}
