package comments

import (
	"encoding/json"
	"time"

	"github.com/calebamiles/github-client/author"
)

var _ Comment = &comment{}

// A Comment represents a comment made on an Issue or Pull Request
type Comment interface {
	Author() string
	Body() string
}

// New returns a list of comments from raw JSON
func New(rawJSON json.RawMessage) ([]Comment, error) {
	var comments []Comment
	cs := []*comment{}

	err := json.Unmarshal(rawJSON, &cs)
	if err != nil {
		return nil, err
	}

	for i := range cs {
		comments = append(comments, cs[i])
	}

	return comments, nil
}

type comment struct {
	BodyString    string    `json:"body"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	commentAuthor `json:"user"`
	author        author.Author
}

func (c *comment) Author() string { return c.commentAuthor.GithubID }
func (c *comment) Body() string   { return c.BodyString }

type commentAuthor struct {
	GithubID string `json:"login"`
}
