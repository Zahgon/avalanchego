// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"errors"

	"github.com/ava-labs/avalanchego/snow/uptime"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm/block"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var (
	_ block.Visitor = (*options)(nil)

	errUnexpectedProposalTxType           = errors.New("unexpected proposal transaction type")
	errFailedFetchingStakerTx             = errors.New("failed fetching staker transaction")
	errUnexpectedStakerTxType             = errors.New("unexpected staker transaction type")
	errFailedFetchingPrimaryStaker        = errors.New("failed fetching primary staker")
	errFailedFetchingSubnetTransformation = errors.New("failed fetching subnet transformation")
	errFailedCalculatingUptime            = errors.New("failed calculating uptime")
)

// options supports build new option blocks
type options struct {
	// inputs populated before calling this struct's methods:
	log                     logging.Logger
	primaryUptimePercentage float64
	uptimes                 uptime.Calculator
	state                   state.Chain

	// outputs populated by this struct's methods:
	preferredBlock block.Block
	alternateBlock block.Block
}

func (*options) BanffAbortBlock(*block.BanffAbortBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (*options) BanffCommitBlock(*block.BanffCommitBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *options) BanffProposalBlock(b *block.BanffProposalBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// We fall back to commit here to err on the side of over-rewarding
// rather than under-rewarding.
//
// Invariant: We must not return the error here, because the error would
// be treated as fatal. Errors can occur here due to a malicious block
// proposer or even in unusual virtuous cases.

func (*options) BanffStandardBlock(*block.BanffStandardBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (*options) ApricotAbortBlock(*block.ApricotAbortBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (*options) ApricotCommitBlock(*block.ApricotCommitBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *options) ApricotProposalBlock(b *block.ApricotProposalBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (*options) ApricotStandardBlock(*block.ApricotStandardBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (*options) ApricotAtomicBlock(*block.ApricotAtomicBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *options) prefersCommit(tx *txs.Tx) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
