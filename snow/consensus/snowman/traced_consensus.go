// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/trace"
	"github.com/ava-labs/avalanchego/utils/bag"
)

var _ Consensus = (*tracedConsensus)(nil)

type tracedConsensus struct {
	Consensus
	tracer trace.Tracer
}

func Trace(consensus Consensus, tracer trace.Tracer) Consensus {
	_ = "STUB: not implemented"
	return *new(Consensus)
}

func (c *tracedConsensus) RecordPoll(ctx context.Context, votes bag.Bag[ids.ID]) error {
	_ = "STUB: not implemented"
	return nil
}
