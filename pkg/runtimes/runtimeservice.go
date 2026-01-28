package runtimes

import (
	"context"
	"errors"

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
