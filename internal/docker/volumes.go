package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/volume"
)

func GetNumVolumes() int {
	vols, err := cli.VolumeList(context.Background(), volume.ListOptions{})
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return -1
	}

	return len(vols.Volumes)
}
