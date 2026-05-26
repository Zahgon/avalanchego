// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package x

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
)

var _ txs.Visitor = (*backendVisitor)(nil)

// backendVisitor handles accepting of transactions for the backend
type backendVisitor struct {
	b    *backend
	ctx  context.Context
	txID ids.ID
}

func (*backendVisitor) BaseTx(*txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

func (*backendVisitor) CreateAssetTx(*txs.CreateAssetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*backendVisitor) OperationTx(*txs.OperationTx) error { _ = "STUB: not implemented"; return nil }

func (b *backendVisitor) ImportTx(tx *txs.ImportTx) error { _ = "STUB: not implemented"; return nil }

func (b *backendVisitor) ExportTx(tx *txs.ExportTx) error { _ = "STUB: not implemented"; return nil }
