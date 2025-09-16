package compose

import (
	"fmt"
	"strings"

	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/types"
)

// TODO: Add RegistryConfigured
func GetProjectSummaries() ([]*types.ProjectSummary, error) {
	m := map[string]*types.ProjectSummary{}

	containers, err := docker.GetContainers()
	if err != nil {
		return nil, err
	}

	for _, v := range containers {
		if proj, ok := v.Labels["com.docker.compose.project"]; ok {
			summary, ok := m[proj]
			if !ok {
				m[proj] = &types.ProjectSummary{
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

	results := make([]*types.ProjectSummary, 0, len(m))
	for _, v := range m {
		results = append(results, v)
	}

	return results, nil
}

func ProjectHeaders(width int) string {
	format := strings.Repeat(fmt.Sprintf("%%-%ds ", width), 6)

	return fmt.Sprintf(format,
		"Name",
		"Services",
		"Containers",
		"Volumes",
		"Networks",
		"Configured",
	)
}

func ProjectSummaryFormatted(proj *types.ProjectSummary, width int) string {
	return ""
}
