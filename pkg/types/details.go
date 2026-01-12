package types

import (
	"context"

	"github.com/docker/docker/client"
)

type StatCard interface {
	Title() string
	Stats(ctx context.Context, cli *client.Client) (*map[string]string, error)
}
