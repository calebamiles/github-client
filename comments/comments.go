package comments

import (
	"encoding/json"
	"log"
	"time"
)

// A Comment represents a comment made on an Issue or Pull Request
type Comment interface {
	Author() string
	Body() string
}

type ReviewComment interface{}

// New returns a list of Comment from raw JSON
func New(rawJSON []byte) ([]Comment, error) {
	var comments []Comment
	cs := []*comment{}

	err := json.Unmarshal(rawJSON, &cs)
	if err != nil {
		log.Printf("failed unmarshalling: \n %s\n", string(rawJSON))
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
}

func (c *comment) Author() string { return c.commentAuthor.GithubID }
func (c *comment) Body() string   { return c.BodyString }

type commentAuthor struct {
	GithubID string `json:"login"`
}
