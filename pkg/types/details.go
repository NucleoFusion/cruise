package types

import (
	"context"

	"github.com/docker/docker/client"
)

type StatCard interface {
	Title() string
	Stats(ctx context.Context, cli *client.Client) (*map[string]string, error)
}

type StatMeta struct {
	TotalRows    int
	TotalColumns int
	SpanMap      *map[string]struct {
		Rows    int
		Columns int
		Index   int
	}
}
