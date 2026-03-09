package dockerruntime

import (
	"context"

	"github.com/cruise-org/cruise/pkg/types"
	"github.com/docker/docker/client"
)

func (s *DockerRuntime) VolumeDetails(ctx context.Context, id string) ([]types.StatCard, *types.StatMeta) {
	stats := []types.StatCard{
		NewVolumeDetails(ctx, id, s.Client),
		NewVolumeLabelDetails(ctx, id, s.Client),
		NewVolumeOptionsDetails(ctx, id, s.Client),
	}

	return stats, &types.StatMeta{
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

type VolumeDetails struct {
	ID   string
	data *map[string]string
	err  error
}

func NewVolumeDetails(ctx context.Context, id string, cli *client.Client) *VolumeDetails {
	s := VolumeDetails{ID: id}

	res, err := cli.VolumeInspect(ctx, s.ID)
	if err != nil {
		s.err = err
		return &s
	}

	s.data = &map[string]string{
		"Volume":     res.Name,
		"Mountpoint": res.Mountpoint,
		"Driver":     res.Driver,
		"Scope":      res.Scope,
		"Created":    res.CreatedAt,
	}

	return &s
}

func (s *VolumeDetails) Title() string { return "Volume Details" }

func (s *VolumeDetails) Stats(ctx context.Context) (*map[string]string, error) {
	return s.data, s.err
}

type VolumeLabels struct {
	ID   string
	data *map[string]string
	err  error
}

func NewVolumeLabelDetails(ctx context.Context, id string, cli *client.Client) *VolumeLabels {
	s := VolumeLabels{ID: id}

	res, err := cli.VolumeInspect(ctx, s.ID)
	if err != nil {
		s.err = err
		return &s
	}

	s.data = &res.Labels
	return &s
}

func (s *VolumeLabels) Title() string { return "Labels" }

func (s *VolumeLabels) Stats(ctx context.Context) (*map[string]string, error) {
	return s.data, s.err
}

type VolumeOptions struct {
	ID   string
	data *map[string]string
	err  error
}

func NewVolumeOptionsDetails(ctx context.Context, id string, cli *client.Client) *VolumeOptions {
	s := VolumeOptions{ID: id}

	res, err := cli.VolumeInspect(ctx, s.ID)
	if err != nil {
		s.err = err
		return &s
	}

	s.data = &res.Options

	return &s
}

func (s *VolumeOptions) Title() string { return "Options" }

func (s *VolumeOptions) Stats(ctx context.Context) (*map[string]string, error) {
	return s.data, s.err
}
