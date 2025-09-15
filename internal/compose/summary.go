package compose

import "github.com/NucleoFusion/cruise/internal/docker"

type ProjectSummary struct {
	Name               string
	Containers         int
	Services           map[string]bool
	Volumes            int
	Networks           int
	RegistryConfigured bool
}

// TODO: Add RegistryConfigured
func GetProjectSummaries() ([]*ProjectSummary, error) {
	m := map[string]*ProjectSummary{}

	containers, err := docker.GetContainers()
	if err != nil {
		return nil, err
	}

	for _, v := range containers {
		if proj, ok := v.Labels["com.docker.compose.project"]; ok {
			summary, ok := m[proj]
			if !ok {
				m[proj] = &ProjectSummary{
					Name:     proj,
					Services: make(map[string]bool),
				}
			}

			summary.Containers += 1

			if srv, ok := v.Labels["com.docker.compose.service"]; ok {
				summary.Services[srv] = true
			}
		}
	}

	vols, err := docker.GetVolumes()
	if err != nil {
		return nil, err
	}

	for _, v := range vols.Volumes {
		if proj, ok := v.Labels["com.docker.compose.project"]; ok {
			m[proj].Volumes += 1
		}
	}

	nets, err := docker.GetNetworks()
	if err != nil {
		return nil, err
	}

	for _, v := range nets {
		if proj, ok := v.Labels["com.docker.compose.project"]; ok {
			m[proj].Networks += 1
		}
	}

	results := make([]*ProjectSummary, 0, len(m))
	for _, v := range m {
		results = append(results, v)
	}

	return results, nil
}
