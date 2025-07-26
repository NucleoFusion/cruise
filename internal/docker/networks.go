package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/network"
)

func GetNumNetworks() int {
	networks, err := cli.NetworkList(context.Background(), network.ListOptions{})
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return -1
	}

	return len(networks)
}
