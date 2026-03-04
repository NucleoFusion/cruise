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
	stats := []types.StatCard{
		NewNetworkDetails(ctx, id, s.Client),
		NewNetworkLabelDetails(ctx, id, s.Client),
		NewNetworkOptionsDetails(ctx, id, s.Client),
		NewNetworkIPAMDetails(ctx, id, s.Client),
	}
	return stats, &types.StatMeta{
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

type NetworkDetails struct {
	ID   string
	data *map[string]string
	err  error
}

func NewNetworkDetails(ctx context.Context, id string, cli *client.Client) *NetworkDetails {
	s := NetworkDetails{ID: id}

	res, err := cli.NetworkInspect(ctx, s.ID, network.InspectOptions{})
	if err != nil {
		s.err = err
		return &s
	}

	s.data = &map[string]string{
		"Network":  res.Name,
		"ID":       res.ID,
		"Created":  res.Created.Format(time.DateOnly) + " " + res.Created.Format(time.Kitchen),
		"Driver":   res.Driver,
		"Scope":    res.Scope,
		"Internal": strconv.FormatBool(res.Internal),
		"Ingress":  strconv.FormatBool(res.Ingress),
	}

	return &s
}

func (s *NetworkDetails) Title() string { return "Network Details" }

func (s *NetworkDetails) Stats(ctx context.Context) (*map[string]string, error) {
	return s.data, s.err
}

type NetworkLabels struct {
	ID   string
	data *map[string]string
	err  error
}

func NewNetworkLabelDetails(ctx context.Context, id string, cli *client.Client) *NetworkLabels {
	s := NetworkLabels{ID: id}

	res, err := cli.NetworkInspect(ctx, s.ID, network.InspectOptions{})
	if err != nil {
		s.err = err
		return &s
	}

	s.data = &res.Labels

	return &s
}

func (s *NetworkLabels) Title() string { return "Labels" }

func (s *NetworkLabels) Stats(ctx context.Context) (*map[string]string, error) {
	return s.data, s.err
}

type NetworkOptions struct {
	ID   string
	data *map[string]string
	err  error
}

func NewNetworkOptionsDetails(ctx context.Context, id string, cli *client.Client) *NetworkOptions {
	s := NetworkOptions{ID: id}

	res, err := cli.NetworkInspect(ctx, s.ID, network.InspectOptions{})
	if err != nil {
		s.err = err
		return &s
	}

	s.data = &res.Options

	return &s
}
func (s *NetworkOptions) Title() string { return "Options" }

func (s *NetworkOptions) Stats(ctx context.Context) (*map[string]string, error) {
	return s.data, s.err
}

type NetworkIPAM struct {
	ID   string
	data *map[string]string
	err  error
}

func NewNetworkIPAMDetails(ctx context.Context, id string, cli *client.Client) *NetworkIPAM {
	s := NetworkIPAM{ID: id}

	res, err := cli.NetworkInspect(ctx, s.ID, network.InspectOptions{})
	if err != nil {
		s.err = err
		return &s
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

	s.data = &stat

	return &s
}
func (s *NetworkIPAM) Title() string { return "IPAM" }

func (s *NetworkIPAM) Stats(ctx context.Context) (*map[string]string, error) {
	return s.data, s.err
}
