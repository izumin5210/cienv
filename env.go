package cienv

// Env provides build environment getter functions
type Env interface {
	Service() Service
	Commit() *Commit
	Branch() string
	Tag() string
	PullRequest() *PullRequest
}
