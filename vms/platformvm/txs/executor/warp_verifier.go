// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"context"

	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

const (
	WarpQuorumNumerator   = 67
	WarpQuorumDenominator = 100
)

var _ txs.Visitor = (*warpVerifier)(nil)

// VerifyWarpMessages verifies all warp messages in the tx. If any of the warp
// messages are invalid, an error is returned.
func VerifyWarpMessages(
	ctx context.Context,
	networkID uint32,
	validatorState validators.State,
	pChainHeight uint64,
	tx txs.UnsignedTx,
) error {
	_ = "STUB: not implemented"
	return nil
}

type warpVerifier struct {
	context        context.Context
	networkID      uint32
	validatorState validators.State
	pChainHeight   uint64
}

func (*warpVerifier) AddValidatorTx(*txs.AddValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*warpVerifier) AddSubnetValidatorTx(*txs.AddSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*warpVerifier) AddDelegatorTx(*txs.AddDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*warpVerifier) CreateChainTx(*txs.CreateChainTx) error { _ = "STUB: not implemented"; return nil }

func (*warpVerifier) CreateSubnetTx(*txs.CreateSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*warpVerifier) ImportTx(*txs.ImportTx) error { _ = "STUB: not implemented"; return nil }

func (*warpVerifier) ExportTx(*txs.ExportTx) error { _ = "STUB: not implemented"; return nil }

func (*warpVerifier) AdvanceTimeTx(*txs.AdvanceTimeTx) error { _ = "STUB: not implemented"; return nil }

func (*warpVerifier) RewardValidatorTx(*txs.RewardValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*warpVerifier) RemoveSubnetValidatorTx(*txs.RemoveSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*warpVerifier) TransformSubnetTx(*txs.TransformSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*warpVerifier) AddPermissionlessValidatorTx(*txs.AddPermissionlessValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*warpVerifier) AddPermissionlessDelegatorTx(*txs.AddPermissionlessDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*warpVerifier) TransferSubnetOwnershipTx(*txs.TransferSubnetOwnershipTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*warpVerifier) BaseTx(*txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

func (*warpVerifier) ConvertSubnetToL1Tx(*txs.ConvertSubnetToL1Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*warpVerifier) IncreaseL1ValidatorBalanceTx(*txs.IncreaseL1ValidatorBalanceTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*warpVerifier) DisableL1ValidatorTx(*txs.DisableL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *warpVerifier) RegisterL1ValidatorTx(tx *txs.RegisterL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *warpVerifier) SetL1ValidatorWeightTx(tx *txs.SetL1ValidatorWeightTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *warpVerifier) verify(message []byte) error { _ = "STUB: not implemented"; return nil }
