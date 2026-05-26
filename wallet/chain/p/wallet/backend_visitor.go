// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package wallet

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var (
	_ txs.Visitor = (*backendVisitor)(nil)

	ErrUnsupportedTxType = errors.New("unsupported tx type")
)

// backendVisitor handles accepting of transactions for the backend
type backendVisitor struct {
	b    *backend
	ctx  context.Context
	txID ids.ID
}

func (*backendVisitor) AdvanceTimeTx(*txs.AdvanceTimeTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*backendVisitor) RewardValidatorTx(*txs.RewardValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) AddValidatorTx(tx *txs.AddValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) AddSubnetValidatorTx(tx *txs.AddSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) AddDelegatorTx(tx *txs.AddDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) CreateChainTx(tx *txs.CreateChainTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) CreateSubnetTx(tx *txs.CreateSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) ImportTx(tx *txs.ImportTx) error { _ = "STUB: not implemented"; return nil }

func (b *backendVisitor) ExportTx(tx *txs.ExportTx) error { _ = "STUB: not implemented"; return nil }

func (b *backendVisitor) RemoveSubnetValidatorTx(tx *txs.RemoveSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) TransformSubnetTx(tx *txs.TransformSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) AddPermissionlessValidatorTx(tx *txs.AddPermissionlessValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) AddPermissionlessDelegatorTx(tx *txs.AddPermissionlessDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) TransferSubnetOwnershipTx(tx *txs.TransferSubnetOwnershipTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) BaseTx(tx *txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

func (b *backendVisitor) ConvertSubnetToL1Tx(tx *txs.ConvertSubnetToL1Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) RegisterL1ValidatorTx(tx *txs.RegisterL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) SetL1ValidatorWeightTx(tx *txs.SetL1ValidatorWeightTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) IncreaseL1ValidatorBalanceTx(tx *txs.IncreaseL1ValidatorBalanceTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) DisableL1ValidatorTx(tx *txs.DisableL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backendVisitor) baseTx(tx *txs.BaseTx) error { _ = "STUB: not implemented"; return nil }
