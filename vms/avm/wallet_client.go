// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/rpc"
)

// WalletClient for interacting with avm managed wallet.
//
// Deprecated: Transactions should be issued using the
// `avalanchego/wallet/chain/x.Wallet` utility.
type WalletClient struct {
	Requester rpc.EndpointRequester
}

// NewWalletClient returns an AVM wallet client for interacting with avm managed
// wallet
//
// Deprecated: Transactions should be issued using the
// `avalanchego/wallet/chain/x.Wallet` utility.
func NewWalletClient(uri, chain string) *WalletClient { _ = "STUB: not implemented"; return nil }

// IssueTx issues a transaction to a node and returns the TxID
func (c *WalletClient) IssueTx(ctx context.Context, txBytes []byte, options ...rpc.Option) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}
