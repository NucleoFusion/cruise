package compose

import (
	"fmt"

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
				summary = m[proj]
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
			summary, ok := m[proj]
			if !ok {
				continue
			}

			summary.Volumes += 1
		}
	}

	nets, err := docker.GetNetworks()
	if err != nil {
		return nil, err
	}

	for _, v := range nets {
		if proj, ok := v.Labels["com.docker.compose.project"]; ok {
			summary, ok := m[proj]
			if !ok {
				continue
			}

			summary.Networks += 1
		}
	}

	results := make([]*types.ProjectSummary, 0, len(m))
	for _, v := range m {
		results = append(results, v)
	}

	return results, nil
}

func ProjectHeaders(width int) string {
	w := width / 7
	return fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s %-*s",
		2*w, "Name",
		1*w, "Containers",
		1*w, "Services",
		1*w, "Volumes",
		1*w, "Networks",
		1*w, "Configured",
	)
}

func ProjectSummaryFormatted(proj *types.ProjectSummary, width int) string {
	configured := "\u2714" // Tick
	if proj.RegistryConfigured {
		configured = "\u2716" // Cross
	}
	w := width / 7

	return fmt.Sprintf("%-*s %-*d %-*d %-*d %-*d %-*s",
		2*w, proj.Name,
		1*w, proj.Containers,
		1*w, proj.NumServices(),
		1*w, proj.Volumes,
		1*w, proj.Networks,
		1*w, configured,
	)
}
