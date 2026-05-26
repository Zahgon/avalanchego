// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blockstest

import (
	"errors"
	"sync"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/libevm/libevm/options"
	"github.com/ava-labs/libevm/rpc"

	"github.com/ava-labs/avalanchego/vms/saevm/blocks"
)

// A ChainBuilder builds a chain of blocks, maintaining necessary invariants.
//
// It is not safe for concurrent use.
type ChainBuilder struct {
	chain          []*blocks.Block
	blocksByHash   sync.Map
	acceptedBlocks event.FeedOf[*blocks.Block]

	defaultOpts []ChainOption
}

// NewChainBuilder returns a new ChainBuilder starting from the provided block,
// which MUST NOT be nil.
func NewChainBuilder(genesis *blocks.Block, defaultOpts ...ChainOption) *ChainBuilder {
	_ = "STUB: not implemented"
	return nil
}

// A ChainOption configures [ChainBuilder.NewBlock].
type ChainOption = options.Option[chainOptions]

// SetDefaultOptions sets the default options upon which all
// additional options passed to [ChainBuilder.NewBlock] are appended.
func (cb *ChainBuilder) SetDefaultOptions(opts ...ChainOption) { _ = "STUB: not implemented"; return }

type chainOptions struct {
	eth []EthBlockOption
	sae []BlockOption
}

// WithEthBlockOptions wraps the options that [ChainBuilder.NewBlock] propagates
// to [NewEthBlock].
func WithEthBlockOptions(opts ...EthBlockOption) ChainOption {
	_ = "STUB: not implemented"
	return *new(ChainOption)
}

// WithBlockOptions wraps the options that [ChainBuilder.NewBlock] propagates to
// [NewBlock].
func WithBlockOptions(opts ...BlockOption) ChainOption {
	_ = "STUB: not implemented"
	return *new(ChainOption)
}

// NewBlock constructs a new block and appends it to the chain.
func (cb *ChainBuilder) NewBlock(tb testing.TB, txs []*types.Transaction, opts ...ChainOption) *blocks.Block {
	_ = "STUB: not implemented"
	return nil
}

// Last returns the last block to be built by the builder, which MAY be the
// genesis block passed to the constructor.
func (cb *ChainBuilder) Last() *blocks.Block { _ = "STUB: not implemented"; return nil }

// AllBlocks returns all blocks, including the genesis passed to
// [NewChainBuilder].
func (cb *ChainBuilder) AllBlocks() []*blocks.Block { _ = "STUB: not implemented"; return nil }

// AllExceptGenesis returns all blocks created with [ChainBuilder.NewBlock].
func (cb *ChainBuilder) AllExceptGenesis() []*blocks.Block { _ = "STUB: not implemented"; return nil }

var _ blocks.Source = (*ChainBuilder)(nil).GetBlock

// GetBlock returns the block with specified hash and height, and a flag
// indicating if it was found. If either argument does not match, it returns
// `nil, false`.
func (cb *ChainBuilder) GetBlock(h common.Hash, num uint64) (*blocks.Block, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// ErrBlockNotFound is returned by [ChainBuilder.ResolveBlockNumber] and
// [ChainBuilder.BlockByNumber] when the requested block number exceeds the
// chain height.
var ErrBlockNotFound = errors.New("block not found")

// SubscribeAcceptedBlocks subscribes to accepted block events fired by
// [ChainBuilder.NewBlock].
func (cb *ChainBuilder) SubscribeAcceptedBlocks(ch chan<- *blocks.Block) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}

// LastAcceptedBlock returns the last block in the chain.
func (cb *ChainBuilder) LastAcceptedBlock() *blocks.Block {
	_ = "STUB: not implemented"

	// ResolveBlockNumber resolves special block number aliases to concrete numbers.
	return nil
}

func (cb *ChainBuilder) ResolveBlockNumber(bn rpc.BlockNumber) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//#nosec G115 -- Non-negative checked above

// BlockByNumber returns the accepted block at the specified height.
func (cb *ChainBuilder) BlockByNumber(bn rpc.BlockNumber) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
