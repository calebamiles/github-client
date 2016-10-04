package labels

import "encoding/json"

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

type Label interface {
	Name() string
}

type label struct {
	NameString string `json:"name"`
}

func (l *label) Name() string {
	return l.NameString
}
