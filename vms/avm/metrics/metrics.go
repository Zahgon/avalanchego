// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/utils/metric"
	"github.com/ava-labs/avalanchego/vms/avm/block"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
)

var _ Metrics = (*metrics)(nil)

type Metrics interface {
	metric.APIInterceptor

	IncTxRefreshes()
	IncTxRefreshHits()
	IncTxRefreshMisses()

	// MarkBlockAccepted updates all metrics relating to the acceptance of a
	// block, including the underlying acceptance of the contained transactions.
	MarkBlockAccepted(b block.Block) error
	// MarkTxAccepted updates all metrics relating to the acceptance of a
	// transaction.
	//
	// Note: This is not intended to be called during the acceptance of a block,
	// as MarkBlockAccepted already handles updating transaction related
	// metrics.
	MarkTxAccepted(tx *txs.Tx) error
}

type metrics struct {
	txMetrics *txMetrics

	numTxRefreshes, numTxRefreshHits, numTxRefreshMisses prometheus.Counter

	metric.APIInterceptor
}

func (m *metrics) IncTxRefreshes() { _ = "STUB: not implemented"; return }

func (m *metrics) IncTxRefreshHits() { _ = "STUB: not implemented"; return }

func (m *metrics) IncTxRefreshMisses() { _ = "STUB: not implemented"; return }

func (m *metrics) MarkBlockAccepted(b block.Block) error { _ = "STUB: not implemented"; return nil }

func (m *metrics) MarkTxAccepted(tx *txs.Tx) error { _ = "STUB: not implemented"; return nil }

func New(registerer prometheus.Registerer) (Metrics, error) {
	_ = "STUB: not implemented"
	return *new(Metrics), nil
}
