package compose

type ProjectSummary struct {
	Name               string
	TotalContainers    int
	TotalServices      int
	TotalNetworks      int
	TotalVolumes       int
	Status             string
	LastUpdated        string
	RegistryConfigured bool
}

// func GetProjects() ([]ProjectSummary, error) {
// 	projects := map[string]bool{}
//
// 	containers, err := docker.GetContainers()
// 	if err != nil {
// 		return "", err
// 	}
//
// 	for _, v := range containers {
// 		if proj, ok := v.Labels["com.docker.compose.project"]; ok {
// 			projects[proj] = true
// 		}
// 	}
// }
