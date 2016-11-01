package serializable

type serializableLabel struct {
	SerializedName string
}

func (l *serializableLabel) Name() string { return l.SerializedName }
