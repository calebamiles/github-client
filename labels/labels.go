package labels

import "encoding/json"

// A Label provides basic information about a GitHub label
type Label interface {
	Name() string
}

// New returns a slice of Label from raw JSON
func New(rawJSON []byte) ([]Label, error) {
	ls := []Label{}

	// don't blow up if the JSON is empty
	if len(rawJSON) == 0 {
		return ls, nil
	}

	var labelsFromJSON []*label

	err := json.Unmarshal(rawJSON, &labelsFromJSON)
	if err != nil {
		return nil, err
	}

	for i := range labelsFromJSON {
		ls = append(ls, labelsFromJSON[i])
	}

	return ls, nil
}

type label struct {
	NameString string `json:"name"`
}

func (l *label) Name() string {
	return l.NameString
}
