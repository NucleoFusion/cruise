package runtimes

import (
	"context"

	"github.com/NucleoFusion/cruise/pkg/types"
)

type Runtime interface {
	Name() string // Container Runtime name

	Containers(ctx context.Context) []types.Container
	Images(ctx context.Context) []types.Image
	Networks(ctx context.Context) []types.Network
	Volumes(ctx context.Context) []types.Volume
}
