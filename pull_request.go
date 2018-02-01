package cienv

// PullRequest contains GitHub pull request metadata
type PullRequest struct {
	Number     int
	HeadBranch string
	BaseBranch string
	Slug       string
}
