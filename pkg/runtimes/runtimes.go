package runtimes

import (
	"context"
	"os/exec"

	"github.com/NucleoFusion/cruise/pkg/types"
)

type Runtime interface {
	Name() string // Container Runtime name

	Containers(ctx context.Context) (*[]types.Container, error)
	Images(ctx context.Context) (*[]types.Image, error)
	Networks(ctx context.Context) (*[]types.Network, error)
	Volumes(ctx context.Context) (*[]types.Volume, error)

	// TODO: Add all relevant function definitions

	// Containers
	StartContainer(ctx context.Context, id string) error
	StopContainer(ctx context.Context, id string) error
	PauseContainer(ctx context.Context, id string) error
	UnpauseContainer(ctx context.Context, id string) error
	RestartContainer(ctx context.Context, id string) error
	RemoveContainer(ctx context.Context, id string) error
	ExecContainer(ctx context.Context, id string) *exec.Cmd
	PortsMap(ctx context.Context, id string) map[string][]string
	ContainerDetails(ctx context.Context, id string) []types.StatCard

	// Images
	PruneImages(ctx context.Context) error
	RemoveImage(ctx context.Context, id string) error
	PushImage(ctx context.Context, id string) error
	PullImage(ctx context.Context, id string) error
	BuildImage(ctx context.Context, id string) error
	SyncImage(ctx context.Context, id string) error
	ImageLayers(ctx context.Context, id string) error
}
