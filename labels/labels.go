package labels

type Label interface {
}

type LabelStruct struct {
	NameString string `json:"name"`
}
