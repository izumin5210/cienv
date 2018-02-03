package cienv

import "os"

const (
	valueTrue = "true"
)

// Get retrieves current git repository status from environment variables.
func Get() Env {
	if os.Getenv("CI") == valueTrue && os.Getenv("TRAVIS") == valueTrue && os.Getenv("SHIPPABLE") != valueTrue {
		return newTravis()
	}
	return Env{Service: Unknown}
}
