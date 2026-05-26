// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/libevm/ethapi"
	"github.com/ava-labs/libevm/rpc"

	"github.com/ava-labs/avalanchego/vms/saevm/saexec"
)

func (b *backend) GetReceipts(ctx context.Context, hash common.Hash) (types.Receipts, error) {
	_ = "STUB: not implemented"
	return *new(types.Receipts), nil
}

//nolint:nilerr // This follows geth behavior for [ethapi.Backend.GetReceipts]

// getReceipts resolves receipts and the underlying [types.Block] by number or
// hash, checking in-memory blocks first then falling back to the database.
// Returns nils for blocks that are not yet executed.
func (b *backend) getReceipts(numOrHash rpc.BlockNumberOrHash) (types.Receipts, *types.Block, error) {
	_ = "STUB: not implemented"
	return *new(types.Receipts), nil, nil
}

// The use of [notFoundIsNil] in [readByNumberOrHash] means that we know
// this is a "real" error, not just [blocks.ErrNotFound].

type blockChainAPI struct {
	*ethapi.BlockChainAPI
	b *backend
}

// GetBlockReceipts overrides [ethapi.BlockChainAPI.GetBlockReceipts] to avoid
// returning an error when a user queries a known, but not yet executed, block.
func (b *blockChainAPI) GetBlockReceipts(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:nilerr // This follows geth behavior for [ethapi.BlockChainAPI.GetBlockReceipts]

// PendingBlockAndReceipts returns a nil block and receipts. Returning nil tells
// geth that this backend does not support pending blocks. In SAE, the pending
// block is defined as the most recently accepted block, but receipts are only
// available after execution. Returning a non-nil block with incorrect or empty
// receipts could cause geth to encounter errors.
func (*backend) PendingBlockAndReceipts() (*types.Block, types.Receipts) {
	_ = "STUB: not implemented"
	return nil, *new(types.Receipts)
}

func (b *backend) GetLogs(ctx context.Context, blockHash common.Hash, number uint64) ([][]*types.Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type immediateReceipts struct {
	recent func(context.Context, common.Hash) (*saexec.Receipt, bool, error)
	*ethapi.TransactionAPI
}

func (ir immediateReceipts) GetTransactionReceipt(ctx context.Context, h common.Hash) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The transaction has either not been included yet, or it was cleared
// from the [saexec.Executor] cache but is on disk. The standard
// mechanism already differentiates between these scenarios.

//#nosec G115 -- Known to not overflow
