package dockerruntime

import (
	"context"

	"github.com/NucleoFusion/cruise/pkg/types"
	"github.com/docker/docker/client"
)

// TODO: New Volume Details format

func (s *DockerRuntime) VolumeDetails(ctx context.Context, id string) ([]types.StatCard, *types.StatMeta) {
	return []types.StatCard{
			&VolumeDetails{ID: id},
			&VolumeLabels{ID: id},
			&VolumeOptions{ID: id},
		}, &types.StatMeta{
			TotalRows:    1,
			TotalColumns: 3,
			SpanMap: &map[string]struct {
				Rows    int
				Columns int
				Index   int
			}{
				"Volume Details": {Rows: 1, Columns: 1, Index: 0},
				"Labels":         {Rows: 1, Columns: 1, Index: 1},
				"Options":        {Rows: 1, Columns: 1, Index: 2},
			},
		}
}

type VolumeDetails struct{ ID string }

func (s *VolumeDetails) Title() string { return "Volume Details" }

func (s *VolumeDetails) Stats(ctx context.Context, cli *client.Client) (*map[string]string, error) {
	res, err := cli.VolumeInspect(ctx, s.ID)
	if err != nil {
		return nil, err
	}

	return &map[string]string{
		"Volume":     res.Name,
		"Mountpoint": res.Mountpoint,
		"Driver":     res.Driver,
		"Scope":      res.Scope,
		"Created":    res.CreatedAt,
	}, nil
}

type VolumeLabels struct{ ID string }

func (s *VolumeLabels) Title() string { return "Labels" }

func (s *VolumeLabels) Stats(ctx context.Context, cli *client.Client) (*map[string]string, error) {
	res, err := cli.VolumeInspect(ctx, s.ID)
	if err != nil {
		return nil, err
	}

	return &res.Labels, nil
}

type VolumeOptions struct{ ID string }

func (s *VolumeOptions) Title() string { return "Options" }

func (s *VolumeOptions) Stats(ctx context.Context, cli *client.Client) (*map[string]string, error) {
	res, err := cli.VolumeInspect(ctx, s.ID)
	if err != nil {
		return nil, err
	}

	return &res.Options, nil
}
