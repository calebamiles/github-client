package serializable

import "time"

type serializableilestone struct {
	SerializedOpen            bool
	SerializedDescription     string
	SerializedNumIssuesOpen   int
	SerializedNumIssuesClosed int
	SerializedTitle           string
	SerializedDeadline        time.Time
}

func (m *serializableilestone) Open() bool           { return m.SerializedOpen }
func (m *serializableilestone) Description() string  { return m.SerializedDescription }
func (m *serializableilestone) NumIssuesOpen() int   { return m.SerializedNumIssuesOpen }
func (m *serializableilestone) NumIssuesClosed() int { return m.SerializedNumIssuesClosed }
func (m *serializableilestone) Title() string        { return m.SerializedTitle }
func (m *serializableilestone) Deadline() time.Time  { return m.SerializedDeadline }
func (m *serializableilestone) String() string       { return m.SerializedTitle }
