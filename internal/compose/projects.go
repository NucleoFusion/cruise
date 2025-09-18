package compose

import (
	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/types"
	"github.com/docker/docker/api/types/container"
)

func Inspect(s *types.ProjectSummary) (*types.Project, error) {
	containers, err := docker.GetContainers()
	if err != nil {
		return nil, err
	}

	project := types.Project{Name: s.Name, Services: make(map[string]*types.ServiceSummary)}

	for _, v := range containers {
		proj, ok := v.Labels["com.docker.compose.project"]
		if !ok || proj != s.Name {
			continue
		}

		srv, ok := v.Labels["com.docker.compose.service"]
		if !ok {
			continue
		}

		service, ok := project.Services[srv]
		if !ok {
			cntrs := make([]container.Summary, 0)
			project.Services[srv] = &types.ServiceSummary{
				Name:       srv,
				Containers: &cntrs,
			}

			service = project.Services[srv] // Updating Variable
		}

		*service.Containers = append(*service.Containers, v)
	}

	return &project, nil
}
