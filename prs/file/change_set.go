package file

// A FileChange represents a diff to a file associated with a Pull Request
type Change interface {
	FileSHA() string
	Filename() string
	NumDeletions() int
	NumAdditions() int
	NumChanges() int
}
