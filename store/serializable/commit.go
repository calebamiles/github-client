package serializable

import (
	"time"

	"github.com/calebamiles/github-client/comments"
	"github.com/calebamiles/github-client/commits"
)

type serializableCommit struct {
	SerializedSHA        string
	SerializedAuthor     *serializableAuthor
	SerializedDate       time.Time
	SerializedParentSHAs []string
	SerializedComments   []*serializableComment
}

func (c *serializableCommit) SHA() string            { return c.SerializedSHA }
func (c *serializableCommit) Author() commits.Author { return c.SerializedAuthor }
func (c *serializableCommit) Date() time.Time        { return c.SerializedDate }
func (c *serializableCommit) ParentSHAs() []string   { return c.SerializedParentSHAs }

func (c *serializableCommit) Comments() []comments.Comment {
	var cs []comments.Comment
	for i := range c.SerializedComments {
		cs = append(cs, c.SerializedComments[i])
	}

	return cs
}

func (c *serializableCommit) String() string

type serializableAuthor struct {
	SerializedGitHubID string
	SerializedEmail    string
	SerializedName     string
}

func (a *serializableAuthor) GitHubID() string { return a.SerializedGitHubID }
func (a *serializableAuthor) Email() string    { return a.SerializedEmail }
func (a *serializableAuthor) Name() string     { return a.SerializedName }
func (a *serializableAuthor) String() string   { return a.SerializedGitHubID }
