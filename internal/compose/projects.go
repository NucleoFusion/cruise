package compose

import (
	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/docker/docker/api/types/container"
)

type Project struct {
	Name     string
	Services map[string]ServiceSummary
	// Should be functions
	TotalNetworks      int
	TotalVolumes       int
	TotalContainers    int    // Should be func
	Status             string // Should be func
	LastUpdated        string
	RegistryConfigured bool // Should be func
}

type ProjectSummary struct {
	Name               string
	Containers         int
	Services           map[string]bool
	Volumes            int
	Networks           int
	RegistryConfigured int
}

func GetProjects() ([]Project, error) {
	projects := map[string]bool{}
	res := make([]Project, 0)

	containers, err := docker.GetContainers()
	if err != nil {
		return res, err
	}

	for _, v := range containers {
		if proj, ok := v.Labels["com.docker.compose.project"]; ok {
			projects[proj] = true
		}
	}

	for k := range projects {
		prj, err := GetSummary(k)
		if err != nil {
			return res, err
		}

		res = append(res, prj)
	}

	return res, nil
}

func GetSummary(proj string) (Project, error) {
	summary := Project{Name: proj}
	containers := make([]container.Summary, 0)

	// All services
	serviceMap := map[string]ServiceSummary{}
	for _, v := range containers {
		if srv, ok := v.Labels["com.docker.compose.service"]; ok {
			if _, ok := serviceMap[srv]; ok {
				*serviceMap[srv].Containers = append(*serviceMap[srv].Containers, v)
			} else {
				serviceMap[srv] = ServiceSummary{
					Name:       srv,
					Containers: &[]container.Summary{v},
				}
			}
		}
	}

	summary.Services = serviceMap

	return summary, nil
}
