package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
)

func GetNumContainers() int {
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return -1
	}

	return len(containers)
}
