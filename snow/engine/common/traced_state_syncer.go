// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"context"

	"github.com/ava-labs/avalanchego/trace"
)

var _ StateSyncer = (*tracedStateSyncer)(nil)

type tracedStateSyncer struct {
	Engine
	stateSyncer StateSyncer
	tracer      trace.Tracer
}

func TraceStateSyncer(stateSyncer StateSyncer, tracer trace.Tracer) StateSyncer {
	_ = "STUB: not implemented"
	return *new(StateSyncer)
}

func (e *tracedStateSyncer) IsEnabled(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
