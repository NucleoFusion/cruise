// Package dockerruntime is the implementation of the Runtime interface for Docker
package dockerruntime

import (
	"github.com/docker/docker/client"
)

type DockerRuntime struct {
	Client *client.Client
}

func NewDockerClient() (*DockerRuntime, error) {
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	cli := c

	return &DockerRuntime{
		Client: cli,
	}, nil
}

func (s *DockerRuntime) Name() string { return "docker" }
