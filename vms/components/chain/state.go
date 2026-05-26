// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package chain

import (
	"context"
	"errors"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

func cachedBlockSize(_ ids.ID, bw *BlockWrapper) int { _ = "STUB: not implemented"; return 0 }

func cachedBlockBytesSize(blockBytes string, _ ids.ID) int { _ = "STUB: not implemented"; return 0 }

// State implements an efficient caching layer used to wrap a VM
// implementation.
type State struct {
	// getBlock retrieves a block from the VM's storage. If getBlock returns
	// a nil error, then the returned block must not have the status Unknown
	getBlock func(context.Context, ids.ID) (snowman.Block, error)
	// unmarshals [b] into a block
	unmarshalBlock        func(context.Context, []byte) (snowman.Block, error)
	batchedUnmarshalBlock func(context.Context, [][]byte) ([]snowman.Block, error)
	// buildBlock attempts to build a block on top of the currently preferred block
	// buildBlock should always return a block with status Processing since it should never
	// create an unknown block, and building on top of the preferred block should never yield
	// a block that has already been decided.
	buildBlock func(context.Context) (snowman.Block, error)

	// If nil, [BuildBlockWithContext] returns [BuildBlock].
	buildBlockWithContext func(context.Context, *block.Context) (snowman.Block, error)

	// verifiedBlocks is a map of blocks that have been verified and are
	// therefore currently in consensus.
	verifiedBlocks map[ids.ID]*BlockWrapper
	// decidedBlocks is an LRU cache of decided blocks.
	decidedBlocks cache.Cacher[ids.ID, *BlockWrapper]
	// unverifiedBlocks is an LRU cache of blocks with status processing
	// that have not yet passed verification.
	unverifiedBlocks cache.Cacher[ids.ID, *BlockWrapper]
	// missingBlocks is an LRU cache of missing blocks
	missingBlocks cache.Cacher[ids.ID, struct{}]
	// string([byte repr. of block]) --> the block's ID
	bytesToIDCache    cache.Cacher[string, ids.ID]
	lastAcceptedBlock *BlockWrapper
}

// Config defines all of the parameters necessary to initialize State
type Config struct {
	// Cache configuration:
	DecidedCacheSize, MissingCacheSize, UnverifiedCacheSize, BytesToIDCacheSize int

	LastAcceptedBlock     snowman.Block
	GetBlock              func(context.Context, ids.ID) (snowman.Block, error)
	UnmarshalBlock        func(context.Context, []byte) (snowman.Block, error)
	BatchedUnmarshalBlock func(context.Context, [][]byte) ([]snowman.Block, error)
	BuildBlock            func(context.Context) (snowman.Block, error)
	BuildBlockWithContext func(context.Context, *block.Context) (snowman.Block, error)
}

func (s *State) initialize(config *Config) { _ = "STUB: not implemented"; return }

func NewState(config *Config) *State { _ = "STUB: not implemented"; return nil }

func NewMeteredState(
	registerer prometheus.Registerer,
	config *Config,
) (*State, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var errSetAcceptedWithProcessing = errors.New("cannot set last accepted block with blocks processing")

// SetLastAcceptedBlock sets the last accepted block to [lastAcceptedBlock].
// This should be called with an internal block - not a wrapped block returned
// from state.
//
// This also flushes [lastAcceptedBlock] from missingBlocks and unverifiedBlocks
// to ensure that their contents stay valid.
func (s *State) SetLastAcceptedBlock(lastAcceptedBlock snowman.Block) error {
	_ = "STUB: not implemented"
	return nil
}

// [lastAcceptedBlock] is no longer missing or unverified, so we evict it from the corresponding
// caches.
//
// Note: there's no need to evict from the decided blocks cache or bytesToIDCache since their
// contents will still be valid.

// Flush each block cache
func (s *State) Flush() { _ = "STUB: not implemented"; return }

// GetBlock returns the BlockWrapper as snowman.Block corresponding to [blkID]
func (s *State) GetBlock(ctx context.Context, blkID ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

// If getBlock returns [database.ErrNotFound], State considers
// this a cacheable miss.

// Since this block is not in consensus, addBlockOutsideConsensus
// is called to add [blk] to the correct cache.

// getCachedBlock checks the caches for [blkID] by priority. Returning
// true if [blkID] is found in one of the caches.
func (s *State) getCachedBlock(blkID ids.ID) (snowman.Block, bool) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), false
}

// GetBlockInternal returns the internal representation of [blkID]
func (s *State) GetBlockInternal(ctx context.Context, blkID ids.ID) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

// ParseBlock attempts to parse [b] into an internal Block and adds it to the
// appropriate caching layer if successful.
func (s *State) ParseBlock(ctx context.Context, b []byte) (snowman.Block, error) {
	_ = "STUB: not implemented"
	// See if we've cached this block's ID by its byte repr.
	return *new(snowman.Block), nil
}

// See if we have this block cached

// We don't have this block cached by its byte repr.
// Parse the block from bytes

// Only check the caches if we didn't do so above

// Check for an existing block, so we can return a unique block
// if processing or simply allow this block to be immediately
// garbage collected if it is already cached.

// Since this block is not in consensus, addBlockOutsideConsensus
// is called to add [blk] to the correct cache.

// BatchedParseBlock implements part of the block.BatchedChainVM interface. In
// addition to performing all the caching as the ParseBlock function, it
// performs at most one call to the underlying VM if [batchedUnmarshalBlock] was
// provided.
func (s *State) BatchedParseBlock(ctx context.Context, blksBytes [][]byte) ([]snowman.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// See if we've cached this block's ID by its byte repr.

// See if we have this block cached

// Check for an existing block, so we can return a unique block
// if processing or simply allow this block to be immediately
// garbage collected if it is already cached.

// BuildBlockWithContext attempts to build a new internal Block, wraps it, and
// adds it to the appropriate caching layer if successful.
// If [s.buildBlockWithContext] is nil, returns [BuildBlock].
func (s *State) BuildBlockWithContext(ctx context.Context, blockCtx *block.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

// BuildBlock attempts to build a new internal Block, wraps it, and adds it
// to the appropriate caching layer if successful.
func (s *State) BuildBlock(ctx context.Context) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func (s *State) deduplicate(blk snowman.Block) snowman.Block {
	_ = "STUB: not implemented"

	// Defensive: buildBlock should not return a block that has already been verified.
	// If it does, make sure to return the existing reference to the block.
	return *new(snowman.Block)
}

// Evict the produced block from missing blocks in case it was previously
// marked as missing.

// wrap the returned block and add it to the correct cache

// addBlockOutsideConsensus adds [blk] to the correct cache and returns
// a wrapped version of [blk]
// assumes [blk] is a known, non-wrapped block that is not currently
// in consensus. [blk] could be either decided or a block that has not yet
// been verified and added to consensus.
func (s *State) addBlockOutsideConsensus(blk snowman.Block) snowman.Block {
	_ = "STUB: not implemented"
	return *new(snowman.Block)
}

func (s *State) LastAccepted(context.Context) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

// LastAcceptedBlock returns the last accepted wrapped block
func (s *State) LastAcceptedBlock() *BlockWrapper { _ = "STUB: not implemented"; return nil }

// LastAcceptedBlockInternal returns the internal snowman.Block that was last accepted
func (s *State) LastAcceptedBlockInternal() snowman.Block {
	_ = "STUB: not implemented"
	return *new(snowman.Block)
}

// IsProcessing returns whether [blkID] is processing in consensus
func (s *State) IsProcessing(blkID ids.ID) bool { _ = "STUB: not implemented"; return false }
