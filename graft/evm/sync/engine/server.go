// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package engine

import (
	"context"
	"errors"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/triedb"

	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var errProviderNotSet = errors.New("provider not set")

// BlockChain provides the blockchain interface needed by the state sync server.
// This interface abstracts the blockchain operations required for serving state summaries.
type BlockChain interface {
	// LastAcceptedBlock returns the last accepted block.
	LastAcceptedBlock() *types.Block

	// GetBlockByNumber returns the block at the given height, or nil if not found.
	GetBlockByNumber(number uint64) *types.Block

	// HasState returns true if the state for the given root hash is available.
	HasState(root common.Hash) bool

	// ResetToStateSyncedBlock resets the blockchain to the given synced block.
	ResetToStateSyncedBlock(block *types.Block) error

	// TrieDB returns the database used for storing the state trie.
	TrieDB() *triedb.Database
}

// SummaryProvider provides state summaries for blocks.
type SummaryProvider interface {
	StateSummaryAtBlock(ethBlock *types.Block) (block.StateSummary, error)
}

type server struct {
	chain BlockChain

	provider         SummaryProvider
	syncableInterval uint64
}

type Server interface {
	GetLastStateSummary(context.Context) (block.StateSummary, error)
	GetStateSummary(context.Context, uint64) (block.StateSummary, error)
}

func NewServer(chain BlockChain, provider SummaryProvider, syncableInterval uint64) Server {
	_ = "STUB: not implemented"
	return *new(Server)
}

// GetLastStateSummary returns the latest state summary.
// State summary is calculated by the block nearest to last accepted
// that is divisible by [syncableInterval]
// If no summary is available, [database.ErrNotFound] must be returned.
func (s *server) GetLastStateSummary(context.Context) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

// GetStateSummary implements StateSyncableVM and returns a summary corresponding
// to the provided [height] if the node can serve state sync data for that key.
// If not, [database.ErrNotFound] must be returned.
func (s *server) GetStateSummary(_ context.Context, height uint64) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

func (s *server) stateSummaryAtHeight(height uint64) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}
