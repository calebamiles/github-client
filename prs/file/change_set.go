package file

import "encoding/json"

// A FileChange represents a diff to a file associated with a Pull Request
type ChangeSet interface {
	FileSHA() string
	Filename() string
	NumDeletions() int
	NumAdditions() int
	NumChanges() int
}

func NewChangeSets(rawJSON []byte) ([]ChangeSet, error) {
	var changeSets []ChangeSet
	var cs []*changeSet
	err := json.Unmarshal(rawJSON, cs)
	if err != nil {
		return nil, err
	}

	for i := range cs {
		changeSets = append(changeSets, cs[i])
	}

	return changeSets, nil
}

type changeSet struct {
	SerializableFileSHA      string `json:"sha"`
	SerializableFilename     string `json:"filename"`
	SerializableNumDeletions int    `json:"deletions"`
	SerializableNumAdditions int    `json:"additions"`
	SerializableNumChanges   int    `json:"changes"`
}

func (c *changeSet) FileSHA() string   { return c.SerializableFileSHA }
func (c *changeSet) Filename() string  { return c.SerializableFilename }
func (c *changeSet) NumDeletions() int { return c.SerializableNumDeletions }
func (c *changeSet) NumAdditions() int { return c.SerializableNumAdditions }
func (c *changeSet) NumChanges() int   { return c.SerializableNumChanges }
