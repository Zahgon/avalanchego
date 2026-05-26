// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/fee"
)

var _ txs.Visitor = (*atomicTxExecutor)(nil)

// AtomicTx executes the atomic transaction [tx] and returns the resulting state
// modifications.
//
// This is only used to execute atomic transactions pre-AP5. After AP5 the
// execution was moved to [StandardTx].
func AtomicTx(
	backend *Backend,
	feeCalculator fee.Calculator,
	parentID ids.ID,
	stateVersions state.Versions,
	tx *txs.Tx,
) (*state.Diff, set.Set[ids.ID], map[ids.ID]*atomic.Requests, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

type atomicTxExecutor struct {
	// inputs, to be filled before visitor methods are called
	backend       *Backend
	feeCalculator fee.Calculator
	parentID      ids.ID
	stateVersions state.Versions
	tx            *txs.Tx

	// outputs of visitor execution
	onAccept       *state.Diff
	inputs         set.Set[ids.ID]
	atomicRequests map[ids.ID]*atomic.Requests
}

func (*atomicTxExecutor) AddValidatorTx(*txs.AddValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) AddSubnetValidatorTx(*txs.AddSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) AddDelegatorTx(*txs.AddDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) CreateChainTx(*txs.CreateChainTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) CreateSubnetTx(*txs.CreateSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) AdvanceTimeTx(*txs.AdvanceTimeTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) RewardValidatorTx(*txs.RewardValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) RemoveSubnetValidatorTx(*txs.RemoveSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) TransformSubnetTx(*txs.TransformSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) AddPermissionlessValidatorTx(*txs.AddPermissionlessValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) AddPermissionlessDelegatorTx(*txs.AddPermissionlessDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) TransferSubnetOwnershipTx(*txs.TransferSubnetOwnershipTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) BaseTx(*txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

func (*atomicTxExecutor) ConvertSubnetToL1Tx(*txs.ConvertSubnetToL1Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) RegisterL1ValidatorTx(*txs.RegisterL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) SetL1ValidatorWeightTx(*txs.SetL1ValidatorWeightTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) IncreaseL1ValidatorBalanceTx(*txs.IncreaseL1ValidatorBalanceTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*atomicTxExecutor) DisableL1ValidatorTx(*txs.DisableL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *atomicTxExecutor) ImportTx(*txs.ImportTx) error { _ = "STUB: not implemented"; return nil }

func (e *atomicTxExecutor) ExportTx(*txs.ExportTx) error { _ = "STUB: not implemented"; return nil }

func (e *atomicTxExecutor) atomicTx() error { _ = "STUB: not implemented"; return nil }
