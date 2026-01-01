package runtimes

import (
	"context"

	"github.com/NucleoFusion/cruise/pkg/types"
)

type Runtime interface {
	Name() string // Container Runtime name

	Containers(ctx context.Context) (*[]types.Container, error)
	Images(ctx context.Context) (*[]types.Image, error)
	Networks(ctx context.Context) (*[]types.Network, error)
	Volumes(ctx context.Context) (*[]types.Volume, error)

	// TODO: Add all relevant function definitions

	// Container Specific
	ContainerDetails(ctx context.Context, id string) (*types.Container, error) // TODO: Type creation
}
