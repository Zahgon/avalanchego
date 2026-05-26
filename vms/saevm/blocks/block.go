// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package blocks defines [Streaming Asynchronous Execution] (SAE) blocks.
//
// [Streaming Asynchronous Execution]: https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/194-streaming-asynchronous-execution
package blocks

import (
	"errors"
	"sync/atomic"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/params"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/saevm/proxytime"

	saetypes "github.com/ava-labs/avalanchego/vms/saevm/types"
)

// A Block extends a [types.Block] to track SAE-defined concepts of async
// execution and settlement. It MUST be constructed with [New].
type Block struct {
	b *types.Block
	// Invariant: ancestry is non-nil and contains non-nil pointers i.f.f. the
	// block hasn't itself been settled. A synchronous block (e.g. SAE genesis
	// or the last pre-SAE block) is always considered settled. See [New] for
	// caveats during construction.
	//
	// Rationale: the ancestral pointers form a linked list that would prevent
	// garbage collection if not severed. Once a block is settled there is no
	// need to inspect its history so we sacrifice the ancestors to the GC
	// Overlord as a sign of our unwavering fealty. See [InMemoryBlockCount] for
	// observability.
	ancestry atomic.Pointer[ancestry]
	// Only the genesis block or the last pre-SAE block is synchronous. These
	// are self-settling by definition so their `ancestry` MUST be nil.
	synchronous bool
	// Determined during block building and SHOULD be set before execution as
	// an early warning system in case of near-miss incorrect predictions.
	bounds *WorstCaseBounds
	// Non-nil i.f.f. [Block.MarkExecuted] has returned without error.
	execution atomic.Pointer[executionResults]

	// Allows this block to be ruled out as able to be settled at a particular
	// time (i.e. if this field is >= said time). The pointer MAY be nil if
	// execution is yet to commence. For more details, see
	// [Block.SetInterimExecutionTime for setting and [LastToSettleAt] for
	// usage.
	interimExecutionTime atomic.Pointer[proxytime.Time[gas.Gas]]

	executed chan struct{} // closed after `execution` is set
	settled  chan struct{} // closed after `ancestry` is cleared

	log logging.Logger
}

var inMemoryBlockCount atomic.Int64

// InMemoryBlockCount returns the number of blocks created with [New] that are
// yet to have their GC finalizers run.
func InMemoryBlockCount() int64 { _ = "STUB: not implemented"; return 0 }

// New constructs a new Block.
//
// While both the `parent` and `lastSettled` arguments MAY be nil, this will
// result in an invalid Block as it breaks important invariants. In such
// situations, [Block.CopyAncestorsFrom] MUST then be called before further use
// of the Block. In practice, this SHOULD only be done when parsing an encoded
// Block.
func New(eth *types.Block, parent, lastSettled *Block, log logging.Logger) (*Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RestoreSettledBlock constructs a new block with [New] and restores it to an
// settled state before returning it. By definition of being settled, the
// returned block also includes post-execution artefacts.
func RestoreSettledBlock(eth *types.Block, log logging.Logger, db ethdb.Database, xdb saetypes.ExecutionResults, config *params.ChainConfig) (*Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	errParentHashMismatch         = errors.New("block-parent hash mismatch")
	errBlockHeightNotIncrementing = errors.New("block height not incrementing")
	errHashMismatch               = errors.New("block hash mismatch")
)

// SetAncestors sets the block's ancestry while enforcing invariants.
func (b *Block) SetAncestors(parent, lastSettled *Block) error {
	_ = "STUB: not implemented"
	return nil
}

// CopyAncestorsFrom populates the [Block.ParentBlock] and [Block.LastSettled]
// values, typically only required during database recovery or block
// verification. The source block MUST have the same hash as b.
//
// Although the individual ancestral blocks are shallow copied, calling
// [Block.MarkSettled] on either the source or destination will NOT clear the
// pointers of the other.
func (b *Block) CopyAncestorsFrom(c *Block) error { _ = "STUB: not implemented"; return nil }

// Signer returns the transaction signer for the block.
func (b *Block) Signer(c *params.ChainConfig) types.Signer {
	_ = "STUB: not implemented"
	return *new(types.Signer)
}

// Signer returns the transaction signer for the block.
func Signer(b *types.Block, c *params.ChainConfig) types.Signer {
	_ = "STUB: not implemented"
	return *new(types.Signer)
}

// A Source returns a [Block] that matches both a hash and number, and a
// boolean indicating if such a block was found.
type Source func(hash common.Hash, number uint64) (*Block, bool)

// AsEthBlockSource returns a [saetypes.BlockSource] backed by the original
// [Source].
func (s Source) AsEthBlockSource() saetypes.BlockSource {
	_ = "STUB: not implemented"
	return *new(saetypes.BlockSource)
}

// AsHeaderSource returns a [saetypes.HeaderSource] backed by the original
// [Source].
func (s Source) AsHeaderSource() saetypes.HeaderSource {
	_ = "STUB: not implemented"
	return *new(saetypes.HeaderSource)
}
