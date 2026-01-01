package runtimes

import (
	"context"
	"log"

	"github.com/NucleoFusion/cruise/pkg/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerRuntime struct {
	Client *client.Client
}

func NewDockerClient() *DockerRuntime {
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal("Docker Error: " + err.Error())
	}

	cli := c

	return &DockerRuntime{
		Client: cli,
	}
}

func (s *DockerRuntime) Name() string { return "docker" }

func (s *DockerRuntime) Containers(ctx context.Context) (*[]types.Container, error) {
	dockerCtr, err := s.Client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	// Type Assert
	cnt := make([]types.Container, 0, len(dockerCtr))
	for _, v := range dockerCtr {
		state := types.ContainerState(v.State)

		// asserting ports
		ports := make([]types.ContainerPort, 0, len(v.Ports))
		for _, p := range v.Ports {
			ports = append(ports, types.ContainerPort{
				ContainerPort: p.PrivatePort,
				HostPort:      p.PublicPort,
				Protocol:      p.Type,
			})
		}

		// asserting mounts
		mounts := make([]types.ContainerMount, 0, len(v.Ports))
		for _, m := range v.Mounts {
			mounts = append(mounts, types.ContainerMount{
				Source:      m.Source,
				Destination: m.Destination,
				ReadOnly:    !m.RW,
			})
		}

		cnt = append(cnt, types.Container{
			Name:    v.Names[0],
			ID:      v.ID,
			Image:   v.Image,
			Created: v.Created,
			Ports:   ports,
			State:   state,
			Mounts:  mounts,
			Labels:  v.Labels,
		})
	}

	return &cnt, nil
}
