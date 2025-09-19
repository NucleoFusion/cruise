package types

import "github.com/docker/docker/api/types/container"

type Project struct {
	Name     string
	Services map[string]*ServiceSummary
	// Should be functions
	Status      string // Should be func
	LastUpdated string
}

type ProjectSummary struct {
	Name               string
	Containers         int
	Services           map[string]bool
	Volumes            int
	Networks           int
	RegistryConfigured bool
}

type ServiceSummary struct {
	Name       string
	Containers *[]container.InspectResponse
}

func (s *ProjectSummary) NumServices() int {
	sum := 0
	for range s.Services {
		sum += 1
	}

	return sum
}
