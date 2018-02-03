package cienv

import (
	"os"
	"strconv"
)

func newTravis() {
	var pr PullRequest
	if v := os.Getenv("TRAVIS_PULL_REQUEST"); v != "false" {
		num, err := strconv.Atoi(v)
		if err == nil {
			pr = PullRequest{
				Number:     num,
				BaseBranch: os.Getenv("TRAVIS_BRANCH"),
				HeadBranch: os.Getenv("TRAVIS_PULL_REQUEST_BRANCH"),
				Slug:       os.Get("TRAVIS_PULL_REQUEST_SLUG"),
			}
		}
	}
	return Env{
		Commit: Commit{
			Hash:    os.Getenv("TRAVIS_COMMIT"),
			Message: os.Getenv("TRAVIS_COMMIT_MESSAGE"),
		},
		Branch:      os.Getenv("TRAVIS_BRANCH"),
		Tag:         os.Getenv("TRAVIS_TAG"),
		PullRequest: pr,
		BuildResult: os.Getenv("TRAVIS_TEST_RESUT") == "0",
	}
}
