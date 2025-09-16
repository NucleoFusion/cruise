package compose

import (
	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/types"
	"github.com/docker/docker/api/types/container"
)

func GetProjects() ([]types.Project, error) {
	projects := map[string]bool{}
	res := make([]types.Project, 0)

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

func GetSummary(proj string) (types.Project, error) {
	summary := types.Project{Name: proj}
	containers := make([]container.Summary, 0)

	// All services
	serviceMap := map[string]types.ServiceSummary{}
	for _, v := range containers {
		if srv, ok := v.Labels["com.docker.compose.service"]; ok {
			if _, ok := serviceMap[srv]; ok {
				*serviceMap[srv].Containers = append(*serviceMap[srv].Containers, v)
			} else {
				serviceMap[srv] = types.ServiceSummary{
					Name:       srv,
					Containers: &[]container.Summary{v},
				}
			}
		}
	}

	summary.Services = serviceMap

	return summary, nil
}
