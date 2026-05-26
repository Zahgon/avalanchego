// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metrics

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/vms/platformvm/block"
)

const blkLabel = "blk"

var (
	_ block.Visitor = (*blockMetrics)(nil)

	blkLabels = []string{blkLabel}
)

type blockMetrics struct {
	txMetrics *txMetrics
	numBlocks *prometheus.CounterVec
}

func newBlockMetrics(registerer prometheus.Registerer) (*blockMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *blockMetrics) BanffAbortBlock(*block.BanffAbortBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *blockMetrics) BanffCommitBlock(*block.BanffCommitBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *blockMetrics) BanffProposalBlock(b *block.BanffProposalBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *blockMetrics) BanffStandardBlock(b *block.BanffStandardBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *blockMetrics) ApricotAbortBlock(*block.ApricotAbortBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *blockMetrics) ApricotCommitBlock(*block.ApricotCommitBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *blockMetrics) ApricotProposalBlock(b *block.ApricotProposalBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *blockMetrics) ApricotStandardBlock(b *block.ApricotStandardBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *blockMetrics) ApricotAtomicBlock(b *block.ApricotAtomicBlock) error {
	_ = "STUB: not implemented"
	return nil
}
