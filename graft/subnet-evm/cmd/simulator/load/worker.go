// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package load

import (
	"context"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/ethclient"
)

type ethereumTxWorker struct {
	client ethclient.Client

	acceptedNonce uint64
	address       common.Address

	newHeads chan *types.Header
}

// NewSingleAddressTxWorker creates and returns a new ethereumTxWorker that confirms transactions by checking the latest
// nonce of [address] and assuming any transaction with a lower nonce was already accepted.
func NewSingleAddressTxWorker(client ethclient.Client, address common.Address) *ethereumTxWorker {
	_ = "STUB: not implemented"
	return nil
}

// NewTxReceiptWorker creates and returns a new ethereumTxWorker that confirms transactions by checking for the
// corresponding transaction receipt.
func NewTxReceiptWorker(client ethclient.Client) *ethereumTxWorker {
	_ = "STUB: not implemented"
	return nil
}

func (tw *ethereumTxWorker) IssueTx(ctx context.Context, tx *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

func (tw *ethereumTxWorker) ConfirmTx(ctx context.Context, tx *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

func (tw *ethereumTxWorker) confirmTxByNonce(ctx context.Context, tx *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

// If the is less than what has already been accepted, the transaction is confirmed

func (tw *ethereumTxWorker) confirmTxByReceipt(ctx context.Context, tx *types.Transaction) error {
	_ = "STUB: not implemented"
	return nil
}

func (tw *ethereumTxWorker) LatestHeight(ctx context.Context) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
