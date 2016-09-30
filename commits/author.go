package commits

// An Author contains information relating to a commit as seen on GitHub
type Author interface {
	GitHubID() string
	Email() string
	Name() string
	String() string
}

type author struct {
	name     string
	email    string
	githubID string
}

func (a *author) GitHubID() string { return a.githubID }
func (a *author) Email() string    { return a.email }
func (a *author) Name() string     { return a.name }
func (a *author) String() string   { return a.githubID }
