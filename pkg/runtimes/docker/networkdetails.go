package dockerruntime

import (
	"context"
	"strconv"
	"time"

	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

// TODO: THe other Network Details?

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
