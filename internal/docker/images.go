package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/image"
)

func GetNumImages() int {
	images, err := cli.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return -1
	}

	return len(images)
}
