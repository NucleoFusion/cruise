// SPDX-License-Identifier: Apache-2.0
// Copyright The cruise-org Authors

package types

import (
	"context"
	"time"
)

type Monitor struct {
	Runtime  string
	Incoming chan Log
	Ctx      context.Context
}

type Log struct {
	Timestamp time.Time
	Message   string
}
