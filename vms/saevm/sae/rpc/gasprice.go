// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/libevm/rpc"

	"github.com/ava-labs/avalanchego/vms/saevm/blocks"
	"github.com/ava-labs/avalanchego/vms/saevm/gasprice"
)

type estimatorBackend struct {
	chain Chain
}

var _ gasprice.Backend = (*estimatorBackend)(nil)

func (e *estimatorBackend) BlockByNumber(n rpc.BlockNumber) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *estimatorBackend) LastAcceptedBlock() *blocks.Block { _ = "STUB: not implemented"; return nil }

func (e *estimatorBackend) ResolveBlockNumber(bn rpc.BlockNumber) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e *estimatorBackend) SubscribeAcceptedBlocks(ch chan<- *blocks.Block) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}
