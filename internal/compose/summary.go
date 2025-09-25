package compose

import (
	"context"
	"fmt"
	"strings"

	"github.com/NucleoFusion/cruise/internal/docker"
	"github.com/NucleoFusion/cruise/internal/utils"
	"github.com/compose-spec/compose-go/v2/types"
)

func GetProjects() ([]*types.Project, error) {
	containers, err := docker.GetContainers()
	if err != nil {
		return nil, err
	}

	projects := map[string]*types.Project{}
	for _, v := range containers {
		if cfg, ok := v.Labels["com.docker.compose.project.config_files"]; ok {
			proj := v.Labels["com.docker.compose.project"] // Project Name

			if _, ok := projects[proj]; ok {
				continue
			}

			project, err := LoadCompose(context.Background(), proj, strings.Split(cfg, ",")) // config_files are in format `./path/to/one.yml,./two.yml`
			if err != nil {
				return nil, err
			}

			projects[proj] = project
		}
	}

	var results []*types.Project
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

func ProjectFormatted(proj *types.Project, width int) string {
	w := width / 7

	return fmt.Sprintf("%-*s %-*s %-*d %-*d %-*d %-*d",
		2*w, proj.Name,
		1*w, utils.Shorten(proj.WorkingDir, w-5),
		1*w, len(proj.Services),
		1*w, len(proj.Volumes),
		1*w, len(proj.Networks),
		1*w, len(proj.ConfigNames()),
	)
}
