package commits

import (
	"encoding/json"
	"log"
	"time"

	"github.com/calebamiles/github-client/comments"
	"github.com/calebamiles/github-client/handle"
)

// A Commit represents basic information about a commit including comments
type Commit interface {
	SHA() string
	Author() Author
	Date() time.Time
	ParentSHAs() []string
	Comments() []comments.Comment

	String() string
}

// A CommitWithoutComments provides basic information relating to a commit
type CommitWithoutComments interface {
	SHA() string
	Author() Author
	Date() time.Time
	ParentSHAs() []string
	CommentsURL() string
	String() string
}

// New returns a list of CommitWithoutComments from raw JSON
func New(rawJSON []byte) ([]CommitWithoutComments, error) {
	var commits []CommitWithoutComments
	cs := []*githubCommit{}

	err := json.Unmarshal(rawJSON, &cs)
	if err != nil {
		log.Printf("failed unmarshalling: \n %s\n", string(rawJSON))
		return nil, err
	}

	for i := range cs {
		for j := range cs[i].Parents {
			cs[i].parentSHAs = append(cs[i].parentSHAs, cs[i].Parents[j].CommitSHA)
		}

		commits = append(commits, cs[i])

		cs[i].author = &author{
			githubID: cs[i].GitHub.ID,
			name:     cs[i].gitCommit.CommitterName,
			email:    cs[i].gitCommit.CommitterEmail,
		}
	}

	return commits, nil
}

type githubCommit struct {
	CommitSHA         string `json:"sha"`
	gitCommit         `json:"commit"`
	handle.GitHub     `json:"author"`
	FetchCommentsFrom string         `json:"comments_url"`
	Parents           []commitParent `json:"parents"`
	author            Author
	parentSHAs        []string
}

func (c *githubCommit) String() string       { return c.CommitSHA }
func (c *githubCommit) Author() Author       { return c.author }
func (c *githubCommit) Date() time.Time      { return c.CommitDate }
func (c *githubCommit) SHA() string          { return c.CommitSHA }
func (c *githubCommit) CommentsURL() string  { return c.FetchCommentsFrom }
func (c *githubCommit) ParentSHAs() []string { return c.parentSHAs }

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
