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

	// Container Specific
	StartContainer(ctx context.Context, id string) error
	StopContainer(ctx context.Context, id string) error
	PauseContainer(ctx context.Context, id string) error
	UnpauseContainer(ctx context.Context, id string) error
	RestartContainer(ctx context.Context, id string) error
	RemoveContainer(ctx context.Context, id string) error
	ExecContainer(ctx context.Context, id string) *exec.Cmd
	// TODO: In containers these are left:-
	// 1. Port map maker (probably define a concrete type)
	// 2. Details Object (concrete type as well, or maybe a custom one for each????? concrete maybe)
}
