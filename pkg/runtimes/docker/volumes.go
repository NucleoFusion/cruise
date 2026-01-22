package dockerruntime

import (
	"context"

	"github.com/NucleoFusion/cruise/pkg/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/volume"
)

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

func (s *DockerRuntime) PruneVolumes(ctx context.Context) error {
	_, err := s.Client.VolumesPrune(context.Background(), filters.NewArgs())
	return err
}

func (s *DockerRuntime) RemoveVolumes(ctx context.Context, id string) error {
	return s.Client.VolumeRemove(context.Background(), id, false)
}
