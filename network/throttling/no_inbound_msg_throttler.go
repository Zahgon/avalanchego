// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package throttling

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
)

var _ InboundMsgThrottler = (*noInboundMsgThrottler)(nil)

// Returns an InboundMsgThrottler where Acquire() always returns immediately.
func NewNoInboundThrottler() InboundMsgThrottler {
	_ = "STUB: not implemented"
	return *new(InboundMsgThrottler)
}

// [Acquire] always returns immediately.
type noInboundMsgThrottler struct{}

func (*noInboundMsgThrottler) Acquire(context.Context, uint64, ids.NodeID) ReleaseFunc {
	_ = "STUB: not implemented"
	return *new(ReleaseFunc)
}

func (*noInboundMsgThrottler) AddNode(ids.NodeID) { _ = "STUB: not implemented"; return }

func (*noInboundMsgThrottler) RemoveNode(ids.NodeID) { _ = "STUB: not implemented"; return }
