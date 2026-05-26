// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/vms/avm/txs"
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

func (m *txMetrics) BaseTx(*txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

func (m *txMetrics) CreateAssetTx(*txs.CreateAssetTx) error { _ = "STUB: not implemented"; return nil }

func (m *txMetrics) OperationTx(*txs.OperationTx) error { _ = "STUB: not implemented"; return nil }

func (m *txMetrics) ImportTx(*txs.ImportTx) error { _ = "STUB: not implemented"; return nil }

func (m *txMetrics) ExportTx(*txs.ExportTx) error { _ = "STUB: not implemented"; return nil }
