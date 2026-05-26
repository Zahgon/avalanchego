// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"errors"

	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/platformvm/block"
	"github.com/ava-labs/avalanchego/vms/platformvm/config"
	"github.com/ava-labs/avalanchego/vms/platformvm/metrics"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/executor"

	txfee "github.com/ava-labs/avalanchego/vms/platformvm/txs/fee"
	validatorfee "github.com/ava-labs/avalanchego/vms/platformvm/validators/fee"
)

var (
	_ block.Visitor = (*verifier)(nil)

	ErrConflictingBlockTxs         = errors.New("block contains conflicting transactions")
	ErrStandardBlockWithoutChanges = errors.New("BanffStandardBlock performs no state changes")

	errApricotBlockIssuedAfterFork           = errors.New("apricot block issued after fork")
	errIncorrectBlockHeight                  = errors.New("incorrect block height")
	errOptionBlockTimestampNotMatchingParent = errors.New("option block proposed timestamp not matching parent block one")
)

// verifier handles the logic for verifying a block.
type verifier struct {
	*backend
	txExecutorBackend *executor.Backend
	pChainHeight      uint64
}

func (v *verifier) BanffAbortBlock(b *block.BanffAbortBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// Must be the last validity check on the block

func (v *verifier) BanffCommitBlock(b *block.BanffCommitBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// Must be the last validity check on the block

func (v *verifier) BanffProposalBlock(b *block.BanffProposalBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// Advance the time to [nextChainTime].

// Must be the last validity check on the block

func (v *verifier) BanffStandardBlock(b *block.BanffStandardBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// Advance the time to [b.Timestamp()].

// Must be the last validity check on the block

func (v *verifier) ApricotAbortBlock(b *block.ApricotAbortBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// Must be the last validity check on the block

func (v *verifier) ApricotCommitBlock(b *block.ApricotCommitBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// Must be the last validity check on the block

func (v *verifier) ApricotProposalBlock(b *block.ApricotProposalBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// Must be the last validity check on the block

func (v *verifier) ApricotStandardBlock(b *block.ApricotStandardBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// Must be the last validity check on the block

func (v *verifier) ApricotAtomicBlock(b *block.ApricotAtomicBlock) error {
	_ = "STUB: not implemented"
	// We call [commonBlock] here rather than [apricotCommonBlock] because below
	// this check we perform the more strict check that ApricotPhase5 isn't
	// activated.
	return nil
}

// cache tx as dropped

func (v *verifier) banffOptionBlock(b block.BanffBlock) error {
	_ = "STUB: not implemented"
	return nil
}

// Banff option blocks must be uniquely generated from the
// BanffProposalBlock. This means that the timestamp must be
// standardized to a specific value. Therefore, we require the timestamp to
// be equal to the parents timestamp.

func (v *verifier) banffNonOptionBlock(b block.BanffBlock) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *verifier) apricotCommonBlock(b block.Block) error {
	_ = "STUB: not implemented"
	// We can use the parent timestamp here, because we are guaranteed that the
	// parent was verified. Apricot blocks only update the timestamp with
	// AdvanceTimeTxs. This means that this block's timestamp will be equal to
	// the parent block's timestamp; unless this is a CommitBlock. In order for
	// the timestamp of the CommitBlock to be after the Banff activation,
	// the parent ApricotProposalBlock must include an AdvanceTimeTx with a
	// timestamp after the Banff timestamp. This is verified not to occur
	// during the verification of the ProposalBlock.
	return nil
}

func (v *verifier) commonBlock(b block.Block) error { _ = "STUB: not implemented"; return nil }

// abortBlock populates the state of this block if [nil] is returned.
//
// Invariant: The call to abortBlock must be the last validity check on the
// block. If this function returns [nil], the block is cached as valid.
func (v *verifier) abortBlock(b block.Block) error { _ = "STUB: not implemented"; return nil }

// commitBlock populates the state of this block if [nil] is returned.
//
// Invariant: The call to commitBlock must be the last validity check on the
// block. If this function returns [nil], the block is cached as valid.
func (v *verifier) commitBlock(b block.Block) error { _ = "STUB: not implemented"; return nil }

// proposalBlock populates the state of this block if [nil] is returned.
//
// Invariant: The call to proposalBlock must be the last validity check on the
// block. If this function returns [nil], the block is cached as valid.
func (v *verifier) proposalBlock(
	b block.Block,
	tx *txs.Tx,
	onDecisionState *state.Diff,
	gasConsumed gas.Gas,
	onCommitState *state.Diff,
	onAbortState *state.Diff,
	feeCalculator txfee.Calculator,
	inputs set.Set[ids.ID],
	atomicRequests map[ids.ID]*atomic.Requests,
	onAcceptFunc func(),
) error {
	_ = "STUB: not implemented"
	return nil
}

// cache tx as dropped

// It is safe to use [b.onAbortState] here because the timestamp will
// never be modified by an Apricot Abort block and the timestamp will
// always be the same as the Banff Proposal Block.

// standardBlock populates the state of this block if [nil] is returned.
//
// Invariant: The call to standardBlock must be the last validity check on the
// block. If this function returns [nil], the block is cached as valid.
func (v *verifier) standardBlock(
	b block.Block,
	txs []*txs.Tx,
	feeCalculator txfee.Calculator,
	onAcceptState *state.Diff,
	changedDuringAdvanceTime bool,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify that the block performs changes. If it does not, it never should
// have been issued.

func (v *verifier) processStandardTxs(txs []*txs.Tx, feeCalculator txfee.Calculator, diff *state.Diff, parentID ids.ID) (
	set.Set[ids.ID],
	map[ids.ID]*atomic.Requests,
	func(),
	gas.Gas,
	bool,
	error,
) {
	_ = "STUB: not implemented"
	// Complexity is limited first to avoid processing too large of a block.
	return nil, nil, nil, *new(gas.Gas), false, nil
}

// If this block exceeds the available capacity, ConsumeGas will return
// an error.

// Updating the fee state prior to executing the transactions is fine
// because the fee calculator was already created.

// cache tx as dropped

// ensure it doesn't overlap with current input batch

// Add UTXOs to batch

// Add/merge in the atomic requests represented by [tx]

// After processing all the transactions, deactivate any L1 validators that
// might not have sufficient fee to pay for the next second.
//
// This ensures that L1 validators are not undercharged for the next second.

func calculateBlockMetrics(
	config *config.Internal,
	blk block.Block,
	s state.Chain,
	gasConsumed gas.Gas,
) metrics.Block {
	_ = "STUB: not implemented"
	return *new(metrics.Block)
}

// deactivateLowBalanceL1Validators deactivates any L1 validators that might not
// have sufficient fees to pay for the next second. The returned bool will be
// true if at least one L1 validator was deactivated.
func deactivateLowBalanceL1Validators(
	config validatorfee.Config,
	diff *state.Diff,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// 1 second

// Invariant: Proposal transactions do not impact L1 validator state.

// If the validator has exactly the right amount of fee for the next
// second we should not remove them here.
//
// GetActiveL1ValidatorsIterator iterates in order of increasing
// EndAccumulatedFee, so we can break early.

// The iterator must be released prior to attempting to write to the
// diff.
