package runtimes

import (
	"context"
	"os/exec"

	dockerruntime "github.com/cruise-org/cruise/pkg/runtimes/docker"
	"github.com/cruise-org/cruise/pkg/types"
)

type Runtime interface {
	Name() string // Container Runtime name

	Containers(ctx context.Context) (*[]types.Container, error)
	Images(ctx context.Context) (*[]types.Image, error)
	Networks(ctx context.Context) (*[]types.Network, error)
	Volumes(ctx context.Context) (*[]types.Volume, error)

	// Containers
	StartContainer(ctx context.Context, id string) error
	StopContainer(ctx context.Context, id string) error
	PauseContainer(ctx context.Context, id string) error
	UnpauseContainer(ctx context.Context, id string) error
	RestartContainer(ctx context.Context, id string) error
	RemoveContainer(ctx context.Context, id string) error
	ExecContainer(ctx context.Context, id string) *exec.Cmd
	PortsMap(ctx context.Context, id string) (map[string][]string, error)
	ContainerDetails(ctx context.Context, id string) ([]types.StatCard, *types.StatMeta)

	// Images
	PruneImages(ctx context.Context) error
	RemoveImage(ctx context.Context, id string) error
	PushImage(ctx context.Context, id string) error
	PullImage(ctx context.Context, id string) error
	ImageLayers(ctx context.Context, id string) (string, error)

	// Networks
	PruneNetworks(ctx context.Context) error
	RemoveNetwork(ctx context.Context, id string) error
	NetworkDetails(ctx context.Context, id string) ([]types.StatCard, *types.StatMeta)

	// Volumes
	PruneVolumes(ctx context.Context) error
	RemoveVolume(ctx context.Context, id string) error
	VolumeDetails(ctx context.Context, id string) ([]types.StatCard, *types.StatMeta)

	// Events/Logs
	ContainerLogs(ctx context.Context, id string) (*types.Monitor, error)
	RuntimeLogs(ctx context.Context) (*types.Monitor, error)
}

var runtimeMap = map[string]func() Runtime{
	"docker": func() Runtime {
		return dockerruntime.NewDockerClient()
	},
}
