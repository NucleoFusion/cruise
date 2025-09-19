package compose

import (
	"time"

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
			cntrs := make([]container.InspectResponse, 0)
			project.Services[srv] = &types.ServiceSummary{
				Name:       srv,
				Containers: &cntrs,
			}

			service = project.Services[srv] // Updating Variable
		}

		insp, err := docker.InspectContainer(v.ID)
		if err != nil {
			return nil, err
		}

		*service.Containers = append(*service.Containers, insp)
	}

	return &project, nil
}

// If all "running", running
// if some running "partially running"
// if none running "exited"
func Status(s *types.Project) string {
	total := 0
	running := 0

	for _, srv := range s.Services {
		for _, v := range *srv.Containers {
			if v.State.Running {
				running++
			}

			total++
		}
	}

	var status string
	if total == running {
		status = "Running"
	} else if running != 0 {
		status = "Partially Running"
	} else if running == 0 {
		status = "Exited"
	} else {
		status = "Uknown"
	}

	return status
}

func StartedAt(s *types.Project) string {
	startedAt := time.Unix(1<<63-1, 0) // Max Time
	for _, srv := range s.Services {
		for _, v := range *srv.Containers {
			t, err := time.Parse(time.RFC3339Nano, v.State.StartedAt)
			if err != nil {
				continue
			}

			if t.Before(startedAt) {
				startedAt = t
			}
		}
	}

	return startedAt.Format("15:04 02 January")
}
