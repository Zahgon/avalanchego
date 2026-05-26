// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

const txLabel = "tx"

var (
	_ txs.Visitor = (*txMetrics)(nil)

	txLabels = []string{txLabel}
)

type txMetrics struct {
	numTxs *prometheus.CounterVec
}

func newTxMetrics(registerer prometheus.Registerer) (*txMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *txMetrics) AddValidatorTx(*txs.AddValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *txMetrics) AddSubnetValidatorTx(*txs.AddSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *txMetrics) AddDelegatorTx(*txs.AddDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *txMetrics) CreateChainTx(*txs.CreateChainTx) error { _ = "STUB: not implemented"; return nil }

func (m *txMetrics) CreateSubnetTx(*txs.CreateSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *txMetrics) ImportTx(*txs.ImportTx) error { _ = "STUB: not implemented"; return nil }

func (m *txMetrics) ExportTx(*txs.ExportTx) error { _ = "STUB: not implemented"; return nil }

func (m *txMetrics) AdvanceTimeTx(*txs.AdvanceTimeTx) error { _ = "STUB: not implemented"; return nil }

func (m *txMetrics) RewardValidatorTx(*txs.RewardValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *txMetrics) RemoveSubnetValidatorTx(*txs.RemoveSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *txMetrics) TransformSubnetTx(*txs.TransformSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *txMetrics) AddPermissionlessValidatorTx(*txs.AddPermissionlessValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *txMetrics) AddPermissionlessDelegatorTx(*txs.AddPermissionlessDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *txMetrics) TransferSubnetOwnershipTx(*txs.TransferSubnetOwnershipTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *txMetrics) BaseTx(*txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

func (m *txMetrics) ConvertSubnetToL1Tx(*txs.ConvertSubnetToL1Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *txMetrics) RegisterL1ValidatorTx(*txs.RegisterL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *txMetrics) SetL1ValidatorWeightTx(*txs.SetL1ValidatorWeightTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *txMetrics) IncreaseL1ValidatorBalanceTx(*txs.IncreaseL1ValidatorBalanceTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *txMetrics) DisableL1ValidatorTx(*txs.DisableL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}
