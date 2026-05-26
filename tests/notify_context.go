// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"time"
)

// DefaultNotifyContext returns a context that is marked done when signals indicating
// process termination are received. If a non-zero duration is provided, the parent to the
// notify context will be a context with a timeout for that duration.
func DefaultNotifyContext(duration time.Duration, cleanup func(func())) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
