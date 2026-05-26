// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"net/http"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/linked"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
)

type WalletService struct {
	vm         *VM
	pendingTxs *linked.Hashmap[ids.ID, *txs.Tx]
}

func (w *WalletService) decided(txID ids.ID) { _ = "STUB: not implemented"; return }

func (w *WalletService) issue(tx *txs.Tx) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

// IssueTx attempts to issue a transaction into consensus
func (w *WalletService) IssueTx(_ *http.Request, args *api.FormattedTx, reply *api.JSONTxID) error {
	_ = "STUB: not implemented"
	return nil
}
