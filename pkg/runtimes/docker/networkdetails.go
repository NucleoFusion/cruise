package dockerruntime

import (
	"context"
	"strconv"
	"time"

	"github.com/cruise-org/cruise/pkg/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// TODO: THe other Network Details?

func (s *DockerRuntime) NetworkDetails(ctx context.Context, id string) ([]types.StatCard, *types.StatMeta) {
	return []types.StatCard{
			&NetworkDetails{ID: id},
			&NetworkLabels{ID: id},
			&NetworkOptions{ID: id},
			&NetworkIPAM{ID: id},
		}, &types.StatMeta{
			TotalRows:    2,
			TotalColumns: 3,
			SpanMap: &map[string]struct {
				Rows    int
				Columns int
				Index   int
			}{
				"Network Details": {Rows: 1, Columns: 1, Index: 0},
				"Labels":          {Rows: 2, Columns: 1, Index: 1},
				"Options":         {Rows: 2, Columns: 1, Index: 2},
				"IPAM":            {Rows: 1, Columns: 1, Index: 3},
			},
		}
}

type NetworkDetails struct{ ID string }

func (s *NetworkDetails) Title() string { return "Network Details" }

func (s *NetworkDetails) Stats(ctx context.Context, cli *client.Client) (*map[string]string, error) {
	res, err := cli.NetworkInspect(ctx, s.ID, network.InspectOptions{})
	if err != nil {
		return nil, err
	}

	return &map[string]string{
		"Network":  res.Name,
		"ID":       res.ID,
		"Created":  res.Created.Format(time.DateOnly) + " " + res.Created.Format(time.Kitchen),
		"Driver":   res.Driver,
		"Scope":    res.Scope,
		"Internal": strconv.FormatBool(res.Internal),
		"Ingress":  strconv.FormatBool(res.Ingress),
	}, nil
}

type NetworkLabels struct{ ID string }

func (s *NetworkLabels) Title() string { return "Labels" }

func (s *NetworkLabels) Stats(ctx context.Context, cli *client.Client) (*map[string]string, error) {
	res, err := cli.NetworkInspect(ctx, s.ID, network.InspectOptions{})
	if err != nil {
		return nil, err
	}

	return &res.Labels, nil
}

type NetworkOptions struct{ ID string }

func (s *NetworkOptions) Title() string { return "Options" }

func (s *NetworkOptions) Stats(ctx context.Context, cli *client.Client) (*map[string]string, error) {
	res, err := cli.NetworkInspect(ctx, s.ID, network.InspectOptions{})
	if err != nil {
		return nil, err
	}

	return &res.Options, nil
}

type NetworkIPAM struct{ ID string }

func (s *NetworkIPAM) Title() string { return "Options" }

func (s *NetworkIPAM) Stats(ctx context.Context, cli *client.Client) (*map[string]string, error) {
	res, err := cli.NetworkInspect(ctx, s.ID, network.InspectOptions{})
	if err != nil {
		return nil, err
	}

	stat := map[string]string{
		"Driver": res.IPAM.Driver,
	}

	for k, v := range res.IPAM.Options {
		if k == "Driver" {
			continue
		}
		stat[k] = v
	}

	return &stat, nil
}
