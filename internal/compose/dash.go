package compose

import "github.com/NucleoFusion/cruise/internal/docker"

func GetNumProjects() (int, error) {
	projects := map[string]bool{}

	containers, err := docker.GetContainers()
	if err != nil {
		return -1, err
	}

	for _, v := range containers {
		if proj, ok := v.Labels["com.docker.compose.project"]; ok {
			projects[proj] = true
		}
	}

	n := 0
	for range projects {
		n++
	}

	return n, nil
}

func GetNumServices() (int, error) {
	projects := map[string]bool{}

	containers, err := docker.GetContainers()
	if err != nil {
		return -1, err
	}

	for _, v := range containers {
		if proj, ok := v.Labels["com.docker.compose.service"]; ok {
			projects[proj] = true
		}
	}

	n := 0
	for range projects {
		n++
	}

	return n, nil
}

// func GetNumServiceHealth() (int, int, error) {
// }
