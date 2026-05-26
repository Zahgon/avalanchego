// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"github.com/ava-labs/avalanchego/vms/platformvm/block"
)

var _ block.Visitor = (*rejector)(nil)

// rejector handles the logic for rejecting a block.
// All errors returned by this struct are fatal and should result in the chain
// being shutdown.
type rejector struct {
	*backend
	addTxsToMempool bool
}

func (r *rejector) BanffAbortBlock(b *block.BanffAbortBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rejector) BanffCommitBlock(b *block.BanffCommitBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rejector) BanffProposalBlock(b *block.BanffProposalBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rejector) BanffStandardBlock(b *block.BanffStandardBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rejector) ApricotAbortBlock(b *block.ApricotAbortBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rejector) ApricotCommitBlock(b *block.ApricotCommitBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rejector) ApricotProposalBlock(b *block.ApricotProposalBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rejector) ApricotStandardBlock(b *block.ApricotStandardBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rejector) ApricotAtomicBlock(b *block.ApricotAtomicBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *rejector) rejectBlock(b block.Block, blockType string) error {
	_ = "STUB: not implemented"
	return nil
}
