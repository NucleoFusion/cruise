package compose

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/NucleoFusion/cruise/internal/docker"
	internaltypes "github.com/NucleoFusion/cruise/internal/types"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/docker/docker/api/types/container"
)

func GetProjects() ([]*internaltypes.Project, error) {
	log.Println("Getting Projects ")
	defer log.Println("Returning Projects ")
	containers, err := docker.GetContainers()
	if err != nil {
		return nil, err
	}

	projects := map[string]*internaltypes.Project{}
	for _, v := range containers {
		if cfg, ok := v.Labels["com.docker.compose.project.config_files"]; ok {
			name := v.Labels["com.docker.compose.project"] // Project Name

			log.Println("Found Project ", name)
			if proj, ok := projects[name]; ok {
				service := *proj.Services

				if srvName, ok := v.Labels["com.docker.compose.service"]; ok { // If belonging to a service
					log.Println("Found Existing Service", srvName)
					log.Println("Found Existing Service 2", srvName)

					insp, err := docker.InspectContainer(v.ID)
					if err != nil {
						log.Println("Error Inspecting ", srvName)
						return nil, err
					}

					if currentService, ok := service[srvName]; ok { // If existing in the Service map
						log.Println("Error Appending ", srvName)
						*currentService.Containers = append(*currentService.Containers, insp)
						log.Println("Appended to ", srvName)

						continue
					}

					temp := make([]container.InspectResponse, 0)
					service[srvName] = internaltypes.Service{Containers: &temp}
				}

				continue
			}

			project, err := LoadCompose(context.Background(), name, strings.Split(cfg, ",")) // config_files are in format `./path/to/one.yml,./two.yml`
			if err != nil {
				log.Println("Failed to load project", name)
				return nil, err
			}

			srvcs := make(map[string]internaltypes.Service)
			projects[name] = &internaltypes.Project{Inspect: project, Services: &srvcs}
		}
	}

	var results []*internaltypes.Project
	for _, v := range projects {
		results = append(results, v)
	}

	return results, nil
}

func ProjectHeaders(width int) string {
	w := width / 7
	return fmt.Sprintf("%-*s %-*s %-*s %-*s %-*s %-*s",
		2*w, "Name",
		1*w, "Directory",
		1*w, "Services",
		1*w, "Volumes",
		1*w, "Networks",
		1*w, "No. of Configs",
	)
}

func ProjectFormatted(proj *internaltypes.Project, width int) string {
	w := width / 7

	return fmt.Sprintf("%-*s %-*s %-*d %-*d %-*d %-*d",
		2*w, proj.Inspect.Name,
		1*w, utils.Shorten(proj.Inspect.WorkingDir, w-5),
		1*w, len(proj.Inspect.Services),
		1*w, len(proj.Inspect.Volumes),
		1*w, len(proj.Inspect.Networks),
		1*w, len(proj.Inspect.ConfigNames()),
	)
}
