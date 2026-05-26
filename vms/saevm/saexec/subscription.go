// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package saexec

import (
	"github.com/ava-labs/libevm/core"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
)

func (e *Executor) sendPostExecutionEvents(b *types.Block, receipts types.Receipts) {
	_ = "STUB: not implemented"
	return
}

// SubscribeChainHeadEvent returns a new subscription for each
// [core.ChainHeadEvent] emitted after execution of a [blocks.Block].
func (e *Executor) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// SubscribeChainEvent returns a new subscription for each [core.ChainEvent]
// emitted after execution of a [blocks.Block].
func (e *Executor) SubscribeChainEvent(ch chan<- core.ChainEvent) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// SubscribeLogsEvent returns a new subscription for logs emitted after
// execution of a [blocks.Block].
func (e *Executor) SubscribeLogsEvent(ch chan<- []*types.Log) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}
