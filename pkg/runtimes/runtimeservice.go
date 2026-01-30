package runtimes

import (
	"context"
	"errors"
	"os/exec"

	"github.com/cruise-org/cruise/pkg/config"
	"github.com/cruise-org/cruise/pkg/types"
)

// RuntimeService is an aggregating service for multiple runtimes
type RuntimeService struct {
	Runtimes map[string]Runtime
}

func NewRuntimeService() {
	runtimeNames := config.Cfg.Global.Runtimes

	runtimes := map[string]Runtime{}
	for _, v := range runtimeNames {
		runtimes[v] = runtimeMap[v]()
	}
}

func (s *RuntimeService) Containers(ctx context.Context) (*[]types.Container, error) {
	cnts := make([]types.Container, 0)
	errs := make([]error, 0)

	for _, v := range s.Runtimes {
		res, err := v.Containers(ctx)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		cnts = append(cnts, *res...)
	}

	return &cnts, errors.Join(errs...)
}

func (s *RuntimeService) Volumes(ctx context.Context) (*[]types.Volume, error) {
	cnts := make([]types.Volume, 0)
	errs := make([]error, 0)

	for _, v := range s.Runtimes {
		res, err := v.Volumes(ctx)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		cnts = append(cnts, *res...)
	}

	return &cnts, errors.Join(errs...)
}

func (s *RuntimeService) Networks(ctx context.Context) (*[]types.Network, error) {
	cnts := make([]types.Network, 0)
	errs := make([]error, 0)

	for _, v := range s.Runtimes {
		res, err := v.Networks(ctx)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		cnts = append(cnts, *res...)
	}

	return &cnts, errors.Join(errs...)
}

func (s *RuntimeService) Images(ctx context.Context) (*[]types.Image, error) {
	cnts := make([]types.Image, 0)
	errs := make([]error, 0)

	for _, v := range s.Runtimes {
		res, err := v.Images(ctx)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		cnts = append(cnts, *res...)
	}

	return &cnts, errors.Join(errs...)
}

func (s *RuntimeService) StartContainer(ctx context.Context, runtime string, id string) error {
	return s.Runtimes[runtime].StartContainer(ctx, id)
}

func (s *RuntimeService) StopContainer(ctx context.Context, runtime string, id string) error {
	return s.Runtimes[runtime].StopContainer(ctx, id)
}

func (s *RuntimeService) PauseContainer(ctx context.Context, runtime string, id string) error {
	return s.Runtimes[runtime].PauseContainer(ctx, id)
}

func (s *RuntimeService) UnpauseContainer(ctx context.Context, runtime string, id string) error {
	return s.Runtimes[runtime].UnpauseContainer(ctx, id)
}

func (s *RuntimeService) RestartContainer(ctx context.Context, runtime string, id string) error {
	return s.Runtimes[runtime].RestartContainer(ctx, id)
}

func (s *RuntimeService) RemoveContainer(ctx context.Context, runtime string, id string) error {
	return s.Runtimes[runtime].RemoveContainer(ctx, id)
}

func (s *RuntimeService) ExecContainer(ctx context.Context, runtime string, id string) *exec.Cmd {
	return s.Runtimes[runtime].ExecContainer(ctx, id)
}

func (s *RuntimeService) PortsMap(ctx context.Context, runtime string, id string) (map[string][]string, error) {
	return s.Runtimes[runtime].PortsMap(ctx, id)
}

func (s *RuntimeService) ContainerDetails(ctx context.Context, runtime string, id string) ([]types.StatCard, *types.StatMeta) {
	return s.Runtimes[runtime].ContainerDetails(ctx, id)
}

func (s *RuntimeService) PruneImages(ctx context.Context, runtime string) error {
	var errs []error
	for _, v := range s.Runtimes {
		err := v.PruneImages(ctx)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (s *RuntimeService) RemoveImage(ctx context.Context, runtime string, id string) error {
	return s.Runtimes[runtime].RemoveImage(ctx, id)
}

func (s *RuntimeService) PullImage(ctx context.Context, runtime string, id string) error {
	return s.Runtimes[runtime].PullImage(ctx, id)
}

func (s *RuntimeService) ImageLayers(ctx context.Context, runtime string, id string) (string, error) {
	return s.Runtimes[runtime].ImageLayers(ctx, id)
}

func (s *RuntimeService) PushImage(ctx context.Context, runtime string, id string) error {
	return s.Runtimes[runtime].PushImage(ctx, id)
}
