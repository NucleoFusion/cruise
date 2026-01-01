package runtimes

import (
	"context"
	"log"

	"github.com/NucleoFusion/cruise/pkg/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
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

func (s *DockerRuntime) Images(ctx context.Context) (*[]types.Image, error) {
	dockerImg, err := s.Client.ImageList(context.Background(), image.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	// Type Assert
	img := make([]types.Image, 0, len(dockerImg))
	for _, v := range dockerImg {
		img = append(img, types.Image{
			ID:            v.ID,
			Tags:          v.RepoTags,
			Size:          v.Size,
			CreatedAt:     v.Created,
			NumContainers: v.Containers,
		})
	}

	return &img, nil
}

func (s *DockerRuntime) Network(ctx context.Context) (*[]types.Network, error) {
	dockerNet, err := s.Client.NetworkList(context.Background(), network.ListOptions{})
	if err != nil {
		return nil, err
	}

	// Type Assert
	net := make([]types.Network, 0, len(dockerNet))
	for _, v := range dockerNet {
		net = append(net, types.Network{
			ID:            v.ID,
			Name:          v.Name,
			Scope:         v.Scope,
			Driver:        v.Driver,
			IPv4:          v.EnableIPv4,
			NumContainers: len(v.Containers),
		})
	}

	return &net, nil
}

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
