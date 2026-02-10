package runtimes

import (
	"context"
	"errors"
	"os/exec"
	"sync"

	"github.com/cruise-org/cruise/pkg/config"
	"github.com/cruise-org/cruise/pkg/types"
)

var RuntimeSrv *RuntimeService

func InitializeService() error {
	rts, err := NewRuntimeService()
	if err != nil {
		return err
	}

	RuntimeSrv = rts

	return nil
}

// RuntimeService is an aggregating service for multiple runtimes
type RuntimeService struct {
	Runtimes map[string]Runtime
}

func NewRuntimeService() (*RuntimeService, error) {
	runtimeNames := config.Cfg.Global.Runtimes
	errs := make([]error, 0)

	runtimes := map[string]Runtime{}
	for _, v := range runtimeNames {
		rt, err := runtimeMap[v]()
		if err != nil {
			errs = append(errs, err)
			continue
		}

		runtimes[v] = rt
	}

	return &RuntimeService{
		Runtimes: runtimes,
	}, errors.Join(errs...)
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

func (s *RuntimeService) PortsMap(ctx context.Context, runtime string, id string) ([]string, error) {
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

func (s *RuntimeService) PruneNetworks(ctx context.Context, runtime string) error {
	var errs []error
	for _, v := range s.Runtimes {
		err := v.PruneNetworks(ctx)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (s *RuntimeService) RemoveNetwork(ctx context.Context, runtime string, id string) error {
	return s.Runtimes[runtime].RemoveNetwork(ctx, id)
}

func (s *RuntimeService) NetworkDetails(ctx context.Context, runtime string, id string) ([]types.StatCard, *types.StatMeta) {
	return s.Runtimes[runtime].NetworkDetails(ctx, id)
}

func (s *RuntimeService) PruneVolumes(ctx context.Context, runtime string) error {
	var errs []error
	for _, v := range s.Runtimes {
		err := v.PruneVolumes(ctx)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (s *RuntimeService) RemoveVolume(ctx context.Context, runtime string, id string) error {
	return s.Runtimes[runtime].RemoveVolume(ctx, id)
}

func (s *RuntimeService) VolumeDetails(ctx context.Context, runtime string, id string) ([]types.StatCard, *types.StatMeta) {
	return s.Runtimes[runtime].VolumeDetails(ctx, id)
}

func (s *RuntimeService) ContainerLogs(ctx context.Context, runtime string, id string) (*types.Monitor, error) {
	return s.Runtimes[runtime].ContainerLogs(ctx, id)
}

func (s *RuntimeService) RuntimeLogs(ctx context.Context, runtime string, id string) (*types.Monitor, error) {
	ch := make(chan types.Log)
	errs := make([]error, 0)
	var wg sync.WaitGroup

	for _, v := range s.Runtimes {
		m, err := v.RuntimeLogs(ctx)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		wg.Add(1)
		go func(m *types.Monitor) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return

				case log, ok := <-m.Incoming:
					if !ok {
						return
					}

					// forward into aggregated channel
					select {
					case ch <- log:
					case <-ctx.Done():
						return
					}
				}
			}
		}(m)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return &types.Monitor{
		Runtime:  "root",
		Incoming: ch,
		Ctx:      ctx,
	}, errors.Join(errs...)
}
