// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/fee"
)

const (
	// Maximum future start time for staking/delegating
	MaxFutureStartTime = 24 * 7 * 2 * time.Hour

	// SyncBound is the synchrony bound used for safe decision making
	SyncBound = 10 * time.Second

	MaxValidatorWeightFactor = 5
)

var (
	_ txs.Visitor = (*proposalTxExecutor)(nil)

	ErrRemoveStakerTooEarly          = errors.New("attempting to remove staker before their end time")
	ErrRemoveWrongStaker             = errors.New("attempting to remove wrong staker")
	ErrInvalidState                  = errors.New("generated output isn't valid state")
	ErrShouldBePermissionlessStaker  = errors.New("expected permissionless staker")
	ErrWrongTxType                   = errors.New("wrong transaction type")
	ErrInvalidID                     = errors.New("invalid ID")
	ErrProposedAddStakerTxAfterBanff = errors.New("staker transaction proposed after Banff")
	ErrAdvanceTimeTxIssuedAfterBanff = errors.New("AdvanceTimeTx issued after Banff")
)

// ProposalTx executes the proposal transaction [tx].
//
// [onCommitState] will be modified to reflect the changes made to the state if
// the proposal is committed.
//
// [onAbortState] will be modified to reflect the changes made to the state if
// the proposal is aborted.
//
// Invariant: It is assumed that [onCommitState] and [onAbortState] represent
// the same state when passed into this function.
func ProposalTx(
	backend *Backend,
	feeCalculator fee.Calculator,
	tx *txs.Tx,
	onCommitState *state.Diff,
	onAbortState *state.Diff,
) error {
	_ = "STUB: not implemented"
	return nil
}

type proposalTxExecutor struct {
	// inputs, to be filled before visitor methods are called
	backend       *Backend
	feeCalculator fee.Calculator
	tx            *txs.Tx
	// [onCommitState] is the state used for validation.
	// [onCommitState] is modified by this struct's methods to
	// reflect changes made to the state if the proposal is committed.
	onCommitState *state.Diff
	// [onAbortState] is modified by this struct's methods to
	// reflect changes made to the state if the proposal is aborted.
	onAbortState *state.Diff
}

func (*proposalTxExecutor) CreateChainTx(*txs.CreateChainTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*proposalTxExecutor) CreateSubnetTx(*txs.CreateSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*proposalTxExecutor) ImportTx(*txs.ImportTx) error { _ = "STUB: not implemented"; return nil }

func (*proposalTxExecutor) ExportTx(*txs.ExportTx) error { _ = "STUB: not implemented"; return nil }

func (*proposalTxExecutor) RemoveSubnetValidatorTx(*txs.RemoveSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*proposalTxExecutor) TransformSubnetTx(*txs.TransformSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*proposalTxExecutor) AddPermissionlessValidatorTx(*txs.AddPermissionlessValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*proposalTxExecutor) AddPermissionlessDelegatorTx(*txs.AddPermissionlessDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*proposalTxExecutor) TransferSubnetOwnershipTx(*txs.TransferSubnetOwnershipTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*proposalTxExecutor) BaseTx(*txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

func (*proposalTxExecutor) ConvertSubnetToL1Tx(*txs.ConvertSubnetToL1Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*proposalTxExecutor) RegisterL1ValidatorTx(*txs.RegisterL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*proposalTxExecutor) SetL1ValidatorWeightTx(*txs.SetL1ValidatorWeightTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*proposalTxExecutor) IncreaseL1ValidatorBalanceTx(*txs.IncreaseL1ValidatorBalanceTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*proposalTxExecutor) DisableL1ValidatorTx(*txs.DisableL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *proposalTxExecutor) AddValidatorTx(tx *txs.AddValidatorTx) error {
	_ = "STUB: not implemented"
	// AddValidatorTx is a proposal transaction until the Banff fork
	// activation. Following the activation, AddValidatorTxs must be issued into
	// StandardBlocks.
	return nil
}

// Set up the state if this tx is committed
// Consume the UTXOs

// Produce the UTXOs

// Set up the state if this tx is aborted
// Consume the UTXOs

// Produce the UTXOs

func (e *proposalTxExecutor) AddSubnetValidatorTx(tx *txs.AddSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	// AddSubnetValidatorTx is a proposal transaction until the Banff fork
	// activation. Following the activation, AddSubnetValidatorTxs must be
	// issued into StandardBlocks.
	return nil
}

// Set up the state if this tx is committed
// Consume the UTXOs

// Produce the UTXOs

// Set up the state if this tx is aborted
// Consume the UTXOs

// Produce the UTXOs

func (e *proposalTxExecutor) AddDelegatorTx(tx *txs.AddDelegatorTx) error {
	_ = "STUB: not implemented"
	// AddDelegatorTx is a proposal transaction until the Banff fork
	// activation. Following the activation, AddDelegatorTxs must be issued into
	// StandardBlocks.
	return nil
}

// Set up the state if this tx is committed
// Consume the UTXOs

// Produce the UTXOs

// Set up the state if this tx is aborted
// Consume the UTXOs

// Produce the UTXOs

func (e *proposalTxExecutor) AdvanceTimeTx(tx *txs.AdvanceTimeTx) error {
	_ = "STUB: not implemented"
	return nil
}

// Validate [newChainTime]

// Note that state doesn't change if this proposal is aborted

func (e *proposalTxExecutor) RewardValidatorTx(tx *txs.RewardValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify that the chain's timestamp is the validator's end time

// Invariant: A [txs.DelegatorTx] does not also implement the
//            [txs.ValidatorTx] interface.

// Handle staker lifecycle.

// Handle staker lifecycle.

// Invariant: Permissioned stakers are removed by the advancement of
//            time and the current chain timestamp is == this staker's
//            EndTime. This means only permissionless stakers should be
//            left in the staker set.

// If the reward is aborted, then the current supply should be decreased.

func (e *proposalTxExecutor) rewardValidatorTx(uValidatorTx txs.ValidatorTx, validator *state.Staker) error {
	_ = "STUB: not implemented"
	return nil
}

// Invariant: The staked asset must be equal to the reward asset.

// Refund the stake only when validator is about to leave
// the staking set

// Provide the reward here

// Provide the accrued delegatee rewards from successful delegations here.

// Note: There is no [offset] if the RewardValidatorTx is
// aborted, because the validator reward is not awarded.

func (e *proposalTxExecutor) rewardDelegatorTx(uDelegatorTx txs.DelegatorTx, delegator *state.Staker) error {
	_ = "STUB: not implemented"
	return nil
}

// Invariant: The staked asset must be equal to the reward asset.

// Refund the stake only when delegator is about to leave
// the staking set

// We're (possibly) rewarding a delegator, so we need to fetch
// the validator they are delegated to.

// Invariant: Delegators must only be able to reference validator
//            transactions that implement [txs.ValidatorTx]. All
//            validator transactions implement this interface except the
//            AddSubnetValidatorTx.

// Calculate split of reward between delegator/delegatee

// Reward the delegator here

// Reward the delegatee here

// Invariant: The rewards calculator can never return a
//            [potentialReward] that would overflow the
//            accumulated rewards.

// For any validators starting after [CortinaTime], we defer rewarding the
// [reward] until their staking period is over.

// For any validators who started prior to [CortinaTime], we issue the
// [delegateeReward] immediately.
