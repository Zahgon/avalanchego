// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"context"

	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

func (vm *VM) StateSyncEnabled(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (vm *VM) GetOngoingSyncStateSummary(ctx context.Context) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

// includes database.ErrNotFound case

func (vm *VM) GetLastStateSummary(ctx context.Context) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

// Extract inner vm's last state summary

// including database.ErrNotFound case

// Note: it's important that ParseStateSummary do not use any index or state
// to allow summaries being parsed also by freshly started node with no previous state.
func (vm *VM) ParseStateSummary(ctx context.Context, summaryBytes []byte) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

// it may be a preFork summary

func (vm *VM) GetStateSummary(ctx context.Context, height uint64) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

// including database.ErrNotFound case

// Note: building state summary requires a well formed height index.
func (vm *VM) buildStateSummary(ctx context.Context, innerSummary block.StateSummary) (block.StateSummary, error) {
	_ = "STUB: not implemented"
	return *new(block.StateSummary), nil
}

// fork has not been reached since there is not fork height
// just return innerSummary
