// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
)

func (b *backend) GetPoolNonce(ctx context.Context, addr common.Address) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *backend) Stats() (pending int, queued int) { _ = "STUB: not implemented"; return 0, 0 }

func (b *backend) TxPoolContent() (map[common.Address][]*types.Transaction, map[common.Address][]*types.Transaction) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *backend) TxPoolContentFrom(addr common.Address) ([]*types.Transaction, []*types.Transaction) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetPoolTransactions returns only pending transactions from the mempool.
func (b *backend) GetPoolTransactions() (types.Transactions, error) {
	_ = "STUB: not implemented"
	return *new(types.Transactions), nil
}
