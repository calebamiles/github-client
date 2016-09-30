package commits

import (
	"encoding/json"
	"time"
)

// A Commit provides interesting information relating to a commit
type Commit interface {
	Author() Author
	Date() time.Time
	SHA() string
	ParentSHAs() []string
	String() string
}

// New returns a list of Commits from raw JSON
func New(rawJSON json.RawMessage) ([]Commit, error) {
	var commits []Commit
	cs := []*githubCommit{}

	err := json.Unmarshal(rawJSON, &cs)
	if err != nil {
		return nil, err
	}

	for i := range cs {
		for j := range cs[i].Parents {
			cs[i].parentSHAs = append(cs[i].parentSHAs, cs[i].Parents[j].CommitSHA)
		}

		commits = append(commits, cs[i])

		cs[i].author = &author{
			githubID: cs[i].githubAuthor.GithubID,
			name:     cs[i].gitCommit.CommitterName,
			email:    cs[i].gitCommit.CommitterEmail,
		}
	}

	return commits, nil
}

type githubCommit struct {
	CommitSHA    string `json:"sha"`
	gitCommit    `json:"commit"`
	githubAuthor `json:"author"`
	Parents      []commitParent `json:"parents"`
	author       Author
	parentSHAs   []string
}

func (c *githubCommit) String() string       { return c.CommitSHA }
func (c *githubCommit) Author() Author       { return c.author }
func (c *githubCommit) Date() time.Time      { return c.CommitDate }
func (c *githubCommit) SHA() string          { return c.CommitSHA }
func (c *githubCommit) ParentSHAs() []string { return c.parentSHAs }

type githubAuthor struct {
	GithubID string `json:"login"`
}

type commitAuthor struct {
	CommitterName  string    `json:"name"`
	CommitterEmail string    `json:"email"`
	CommitDate     time.Time `json:"date"`
}

type gitCommit struct {
	commitAuthor `json:"author"`
}

type commitParent struct {
	CommitSHA string `json:"sha"`
}
