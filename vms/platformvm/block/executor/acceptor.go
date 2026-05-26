// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"errors"

	"github.com/ava-labs/avalanchego/vms/platformvm/block"
	"github.com/ava-labs/avalanchego/vms/platformvm/metrics"
	"github.com/ava-labs/avalanchego/vms/platformvm/validators"
)

var (
	_ block.Visitor = (*acceptor)(nil)

	errMissingBlockState = errors.New("missing state of block")
)

// acceptor handles the logic for accepting a block.
// All errors returned by this struct are fatal and should result in the chain
// being shutdown.
type acceptor struct {
	*backend
	metrics    metrics.Metrics
	validators *validators.Manager
}

func (a *acceptor) BanffAbortBlock(b *block.BanffAbortBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *acceptor) BanffCommitBlock(b *block.BanffCommitBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *acceptor) BanffProposalBlock(b *block.BanffProposalBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *acceptor) BanffStandardBlock(b *block.BanffStandardBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *acceptor) ApricotAbortBlock(b *block.ApricotAbortBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *acceptor) ApricotCommitBlock(b *block.ApricotCommitBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *acceptor) ApricotProposalBlock(b *block.ApricotProposalBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *acceptor) ApricotStandardBlock(b *block.ApricotStandardBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *acceptor) ApricotAtomicBlock(b *block.ApricotAtomicBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// Update the state to reflect the changes made in [onAcceptState].

// Note that this method writes [batch] to the database.

func (a *acceptor) optionBlock(b block.Block, blockType string) error {
	_ = "STUB: not implemented"
	return nil
}

// Note: we assume this block's sibling doesn't
// need the parent's state when it's rejected.

// Note that the parent must be accepted first.

// Note that this method writes [batch] to the database.

func (a *acceptor) proposalBlock(b block.Block, blockType string) {
	_ = "STUB: not implemented"
	// Note that:
	//
	// * We don't free the proposal block in this method.
	//   It is freed when its child is accepted.
	//   We need to keep this block's state in memory for its child to use.
	//
	// * We only update the metrics to reflect this block's
	//   acceptance when its child is accepted.
	//
	// * We don't write this block to state here.
	//   That is done when this block's child (a CommitBlock or AbortBlock) is accepted.
	//   We do this so that in the event that the node shuts down, the proposal block
	//   is not written to disk unless its child is.
	//   (The VM's Shutdown method commits the database.)
	//   The snowman.Engine requires that the last committed block is a decision block
	return
}

func (a *acceptor) standardBlock(b block.Block, blockType string) error {
	_ = "STUB: not implemented"
	return nil
}

// Update the state to reflect the changes made in [onAcceptState].

// Note that this method writes [batch] to the database.

func (a *acceptor) commonAccept(b *blockState) error { _ = "STUB: not implemented"; return nil }
