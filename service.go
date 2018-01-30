package cienv

// Service represents a CI SaaS
type Service int

// Enum values for Service
const (
	Unknown Service = iota
	TravisCI
)

var (
	nameByService = map[Service]string{
		Unknown:  "unknown",
		TravisCI: "Travis CI",
	}
)

func (s Service) String() string {
	if n, ok := nameByService[s]; ok {
		return n
	}
	return nameByService[Unknown]
}
