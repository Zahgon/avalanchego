// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sae

import (
	"context"
	"errors"
	"time"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/vms/saevm/blocks"

	saetypes "github.com/ava-labs/avalanchego/vms/saevm/types"
)

// maxFutureBlockDuration is the maximum time from the current time allowed for
// blocks before they're considered future blocks and fail parsing or
// verification.
const (
	maxFutureBlockSeconds  uint64 = 10
	maxFutureBlockDuration        = time.Duration(maxFutureBlockSeconds) * time.Second
)

var (
	errBlockHeightNotUint64 = errors.New("block height not uint64")
	errBlockTooFarInFuture  = errors.New("block too far in the future")

	errTxHashMismatch         = errors.New("transaction hash mismatch")
	errUncleHashMismatch      = errors.New("uncle hash mismatch")
	errWithdrawalHashMismatch = errors.New("withdrawals hash mismatch")
)

// ParseBlock parses the buffer as [rlp] encoding of a [types.Block]. It does
// NOT populate the block ancestry, which is done by [VM.VerifyBlock] i.f.f.
// verification passes.
func (vm *VM) ParseBlock(ctx context.Context, buf []byte) (*blocks.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The uint64 timestamp can't underflow [time.Time] but it can overflow so
// make this some future engineer's problem in a few millennia.

// Block body must match what is declared by the header.

// The withdrawals hash being set depends on the Ethereum hard fork.

func compareHashPtrs(a, b *common.Hash) bool { _ = "STUB: not implemented"; return false }

// BuildBlock builds a new block, using the last block passed to
// [VM.SetPreference] as the parent. The block context MAY be nil.
func (vm *VM) BuildBlock(ctx context.Context, bCtx *block.Context) (*blocks.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var (
	errUnknownParent     = errors.New("unknown parent")
	errBlockHeightTooLow = errors.New("block height too low")
	errHashMismatch      = errors.New("hash mismatch")
)

// VerifyBlock validates the block and, if successful, populates its ancestry.
// The block context MAY be nil.
func (vm *VM) VerifyBlock(ctx context.Context, bCtx *block.Context, b *blocks.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// Sanity check that we aren't verifying an accepted block.

// Although this is also checked in [blocks.Block.CopyAncestorsFrom], it is
// key to the purpose of this method so included here to be defensive. It
// also provides a clearer failure message.

var (
	errSettledRootMismatch   = errors.New("settled root mismatch")
	errSettledHeightMismatch = errors.New("settled height mismatch")
)

// verifyWhenBootstrapping skips verification in its entirety. It is expected
// for blocks to be verified by hash in the bootstrapping engine. This supports
// hooks, such as Coreth and Subnet-EVM, that are unable to fully verify blocks
// during bootstrapping.
func (vm *VM) verifyWhenBootstrapping(b, parent *blocks.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// Sanity checks to ensure the in-memory settled block matches the expected
// settled block.

func canonicalBlock(db ethdb.Database, num uint64) (*types.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *VM) settledBlockFromDB(db ethdb.Reader, hash common.Hash, num uint64) (*blocks.Block, error) {
	_ = "STUB: not implemented"
	// Before doing any disk IO, we sanity check that num is for a settled
	// block.
	//
	// If using this function with [readByHash] this check is required.
	// Otherwise, there is a possible (read: near impossible but non-zero)
	// chance that [VM.VerifyBlock] and [VM.AcceptBlock] were *both* called
	// between checking the in-memory block store and loading the canonical
	// number from the database. That could result in attempting to restore an
	// unexecuted block, which would report an error.
	//
	// TODO(arr4n) I think [readHash] should be providing this guarantee
	// as it has access to the [syncMap] and its lock.
	return nil, nil
}

// Excess is only used for executing the next block, which can never
// be the case if `b` isn't actually the last synchronous block, so
// passing the same value for all is OK.

// GetBlock returns the block with the given ID, or [database.ErrNotFound].
//
// It is expected that blocks that have been successfully verified should be
// returned correctly. It is also expected that blocks that have been
// accepted by the consensus engine should be able to be fetched. It is not
// required for blocks that have been rejected by the consensus engine to be
// able to be fetched.
func (vm *VM) GetBlock(ctx context.Context, id ids.ID) (*blocks.Block, error) {
	_ = "STUB: not implemented"
	// protect the input to allow comment linking
	return nil, nil
}

// consensus MAY request verified-but-not-accepted blocks

// GetBlockIDAtHeight returns the accepted block at the given height, or
// [database.ErrNotFound].
func (vm *VM) GetBlockIDAtHeight(ctx context.Context, height uint64) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

var (
	_ saetypes.BlockSource  = (*VM)(nil).ethBlockSource
	_ saetypes.HeaderSource = (*VM)(nil).headerSource
)

func (vm *VM) ethBlockSource(hash common.Hash, num uint64) (*types.Block, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (vm *VM) headerSource(hash common.Hash, num uint64) (*types.Header, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func source[T any](vm *VM, hash common.Hash, num uint64, fromMem blocks.Extractor[T], fromDB blocks.DBReader[T]) (*T, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
