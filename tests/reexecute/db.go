// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package reexecute

import (
	"github.com/ava-labs/avalanchego/tests"
)

// BlockResult represents the result of reading a block from the database.
// It contains either the block data and height, or an error if the read failed.
type BlockResult struct {
	BlockBytes []byte
	Height     uint64
	Err        error
}

// CreateBlockChanFromLevelDB creates a channel that streams blocks from a LevelDB database.
// It opens the database at sourceDir and iterates through blocks from startBlock to endBlock (inclusive).
// Blocks are read sequentially and sent to the returned channel as BlockResult values.
//
// Any validation errors or iteration errors are sent as BlockResult with Err set, then the channel is closed.
func CreateBlockChanFromLevelDB(tc tests.TestContext, sourceDir string, startBlock, endBlock uint64, chanSize int) (<-chan BlockResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BlockKey converts a block height to its corresponding database key.
func BlockKey(height uint64) []byte { _ = "STUB: not implemented"; return nil }
