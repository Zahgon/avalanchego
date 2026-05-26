// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blocktest

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

var (
	errGetAncestor       = errors.New("unexpectedly called GetAncestor")
	errBatchedParseBlock = errors.New("unexpectedly called BatchedParseBlock")

	_ block.BatchedChainVM = (*BatchedVM)(nil)
)

// BatchedVM is a BatchedVM that is useful for testing.
type BatchedVM struct {
	T *testing.T

	CantGetAncestors    bool
	CantBatchParseBlock bool

	GetAncestorsF func(
		ctx context.Context,
		blkID ids.ID,
		maxBlocksNum int,
		maxBlocksSize int,
		maxBlocksRetrivalTime time.Duration,
	) ([][]byte, error)

	BatchedParseBlockF func(
		ctx context.Context,
		blks [][]byte,
	) ([]snowman.Block, error)
}

func (vm *BatchedVM) Default(cant bool) { _ = "STUB: not implemented"; return }

func (vm *BatchedVM) GetAncestors(
	ctx context.Context,
	blkID ids.ID,
	maxBlocksNum int,
	maxBlocksSize int,
	maxBlocksRetrivalTime time.Duration,
) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (vm *BatchedVM) BatchedParseBlock(
	ctx context.Context,
	blks [][]byte,
) ([]snowman.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
