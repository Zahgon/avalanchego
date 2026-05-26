// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txgossip

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/state"
	"github.com/ava-labs/libevm/core/txpool"
	"github.com/ava-labs/libevm/core/txpool/legacypool"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/params"

	// Imported for [core.ChainHeadEvent] comment resolution. Already a
	// downstream dependency.
	_ "github.com/ava-labs/libevm/core"

	"github.com/ava-labs/avalanchego/vms/saevm/saexec"

	saetypes "github.com/ava-labs/avalanchego/vms/saevm/types"
)

// A BlockChain is the union of [txpool.BlockChain] and [legacypool.BlockChain].
type BlockChain interface {
	txpool.BlockChain
	legacypool.BlockChain
}

// NewBlockChain wraps an [saexec.Executor] to be compatible with a
// non-blob-transaction mempool.
//
// The wrappers's `CurrentBlock()` method returns the last executed, while the
// `StateAt()` method ignores its argument and always opens the latest
// post-execution state root. The [core.ChainHeadEvent] subscription therefore
// acts only to inform the mempool of some new state, but not which specific
// root as the event contains a [types.Header] carrying the (ignored)
// last-settled state root.
func NewBlockChain(exec *saexec.Executor, blocks saetypes.BlockSource) BlockChain {
	_ = "STUB: not implemented"
	return *new(BlockChain)
}

type blockchain struct {
	*saexec.Executor // exposes SubscribeChainHeadEvent()
	blocks           saetypes.BlockSource
}

func (bc *blockchain) Config() *params.ChainConfig { _ = "STUB: not implemented"; return nil }

func (bc *blockchain) CurrentBlock() *types.Header { _ = "STUB: not implemented"; return nil }

func (bc *blockchain) GetBlock(hash common.Hash, number uint64) *types.Block {
	_ = "STUB: not implemented"
	return nil
}

func (bc *blockchain) StateAt(common.Hash) (*state.StateDB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
