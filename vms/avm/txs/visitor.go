// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import "github.com/ava-labs/avalanchego/vms/components/avax"

var _ Visitor = (*utxoGetter)(nil)

// Allow vm to execute custom logic against the underlying transaction types.
type Visitor interface {
	BaseTx(*BaseTx) error
	CreateAssetTx(*CreateAssetTx) error
	OperationTx(*OperationTx) error
	ImportTx(*ImportTx) error
	ExportTx(*ExportTx) error
}

// utxoGetter returns the UTXOs transaction is producing.
type utxoGetter struct {
	tx    *Tx
	utxos []*avax.UTXO
}

func (u *utxoGetter) BaseTx(tx *BaseTx) error { _ = "STUB: not implemented"; return nil }

func (u *utxoGetter) ImportTx(tx *ImportTx) error { _ = "STUB: not implemented"; return nil }

func (u *utxoGetter) ExportTx(tx *ExportTx) error { _ = "STUB: not implemented"; return nil }

func (u *utxoGetter) CreateAssetTx(t *CreateAssetTx) error { _ = "STUB: not implemented"; return nil }

func (u *utxoGetter) OperationTx(t *OperationTx) error {
	_ = "STUB: not implemented"
	// The error is explicitly dropped here because no error is ever returned
	// from the utxoGetter.
	return nil
}
