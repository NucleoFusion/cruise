// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package types

import (
	"context"
)

type StatCard interface {
	Title() string
	Stats(ctx context.Context) (*map[string]string, error)
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
