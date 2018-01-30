package cienv

// PullRequest contains GitHub pull request metadata
type PullRequest struct {
	Number int
	Branch string
	Slug   string
}
