// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package builder

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/txs/mempool"

	blockexecutor "github.com/ava-labs/avalanchego/vms/avm/block/executor"
	txexecutor "github.com/ava-labs/avalanchego/vms/avm/txs/executor"
)

// targetBlockSize is the max block size we aim to produce
const targetBlockSize = 128 * units.KiB

var (
	_ Builder = (*builder)(nil)

	ErrNoTransactions = errors.New("no transactions")
)

type Builder interface {
	// WaitForEvent waits until there is at least one tx available to the
	// builder.
	WaitForEvent(ctx context.Context) (common.Message, error)
	// BuildBlock can be called to attempt to create a new block
	BuildBlock(context.Context) (snowman.Block, error)
}

// builder implements a simple builder to convert txs into valid blocks
type builder struct {
	backend *txexecutor.Backend
	manager blockexecutor.Manager
	clk     *mockable.Clock

	// Pool of all txs that may be able to be added
	mempool mempool.Mempool[*txs.Tx]
}

func New(
	backend *txexecutor.Backend,
	manager blockexecutor.Manager,
	clk *mockable.Clock,
	mempool mempool.Mempool[*txs.Tx],
) Builder {
	_ = "STUB: not implemented"
	return *new(Builder)
}

func (b *builder) WaitForEvent(ctx context.Context) (common.Message, error) {
	_ = "STUB: not implemented"
	return *new(common.Message), nil
}

// BuildBlock builds a block to be added to consensus.
func (b *builder) BuildBlock(context.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

// Get the block to build on top of and retrieve the new block's context.

// [timestamp] = max(now, parentTime)

// Invariant: [mempool.MaxTxSize] < [targetBlockSize]. This guarantees
// that we will only stop building a block once there are no
// transactions in the mempool or the block is at least
// [targetBlockSize - mempool.MaxTxSize] bytes full.

// Invariant: [tx] has already been syntactically verified.
