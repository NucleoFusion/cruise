package dockerruntime

import (
	"context"
	"log"

	"github.com/NucleoFusion/cruise/pkg/types"
	"github.com/docker/docker/api/types/volume"
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

func (s *DockerRuntime) Volume(ctx context.Context) (*[]types.Volume, error) {
	dockerVol, err := s.Client.VolumeList(context.Background(), volume.ListOptions{})
	if err != nil {
		return nil, err
	}

	// Type Assert
	vol := make([]types.Volume, 0, len(dockerVol.Volumes))
	for _, v := range dockerVol.Volumes {
		vol = append(vol, types.Volume{
			Name:       v.Name,
			Scope:      v.Scope,
			Driver:     v.Driver,
			Mountpoint: v.Mountpoint,
			CreatedAt:  v.CreatedAt,
		})
	}

	return &vol, nil
}
