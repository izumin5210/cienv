package cienv

// Env provides build environment getter functions
type Env struct {
	Commit      Commit
	Branch      string
	Tag         string
	PullRequest PullRequest
}
