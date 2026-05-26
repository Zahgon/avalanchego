// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
)

func (b *backend) GetTransaction(ctx context.Context, txHash common.Hash) (exists bool, tx *types.Transaction, blockHash common.Hash, blockNumber uint64, index uint64, err error) {
	_ = "STUB: not implemented"
	return false, nil, *new(common.Hash), 0, 0, nil
}

// GetPoolTransaction returns a transaction from the mempool regardless of its
// status (pending or queued).
func (b *backend) GetPoolTransaction(txHash common.Hash) *types.Transaction {
	_ = "STUB: not implemented"
	return nil
}
