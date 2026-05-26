// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package throttling

import (
	"context"

	"golang.org/x/time/rate"
)

var (
	_ DialThrottler = (*dialThrottler)(nil)
	_ DialThrottler = (*noDialThrottler)(nil)
)

type DialThrottler interface {
	// Block until the event associated with this Acquire can happen.
	// If [ctx] is canceled, gives up and returns an error.
	Acquire(ctx context.Context) error
}

type dialThrottler struct {
	limiter *rate.Limiter
}

type noDialThrottler struct{}

func (t dialThrottler) Acquire(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func NewDialThrottler(throttleLimit int) DialThrottler {
	_ = "STUB: not implemented"
	return *new(DialThrottler)
}

func NewNoDialThrottler() DialThrottler { _ = "STUB: not implemented"; return *new(DialThrottler) }

func (noDialThrottler) Acquire(context.Context) error { _ = "STUB: not implemented"; return nil }
