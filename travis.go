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
		commit: Commit{
			Hash:    os.Getenv("TRAVIS_COMMIT"),
			Message: os.Getenv("TRAVIS_COMMIT_MESSAGE"),
		},
		branch: os.Getenv("TRAVIS_BRANCH"),
		tag:    os.Getenv("TRAVIS_TAG"),
	}
}
