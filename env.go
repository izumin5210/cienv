package cienv

// Env provides build environment getter functions
type Env struct {
	Commit      Commit
	Branch      string
	Tag         string
	PullRequest PullRequest
	BuildResult bool
}

// IsPullRequest returns true if this build is triggered by a pull request.
func (e Env) IsPullRequest() bool {
	return e.PullRequest != PullRequest{}
}
