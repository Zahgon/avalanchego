// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/fee"
)

var (
	ErrWeightTooSmall                  = errors.New("weight of this validator is too low")
	ErrWeightTooLarge                  = errors.New("weight of this validator is too large")
	ErrInsufficientDelegationFee       = errors.New("staker charges an insufficient delegation fee")
	ErrStakeTooShort                   = errors.New("staking period is too short")
	ErrStakeTooLong                    = errors.New("staking period is too long")
	ErrFlowCheckFailed                 = errors.New("flow check failed")
	ErrNotValidator                    = errors.New("isn't a current or pending validator")
	ErrRemovePermissionlessValidator   = errors.New("attempting to remove permissionless validator")
	ErrStakeOverflow                   = errors.New("validator stake exceeds limit")
	ErrPeriodMismatch                  = errors.New("proposed staking period is not inside dependent staking period")
	ErrOverDelegated                   = errors.New("validator would be over delegated")
	ErrIsNotTransformSubnetTx          = errors.New("is not a transform subnet tx")
	ErrTimestampNotBeforeStartTime     = errors.New("chain timestamp not before start time")
	ErrAlreadyValidator                = errors.New("already a validator")
	ErrDuplicateValidator              = errors.New("duplicate validator")
	ErrDelegateToPermissionedValidator = errors.New("delegation to permissioned validator")
	ErrWrongStakedAssetID              = errors.New("incorrect staked assetID")
	ErrDurangoUpgradeNotActive         = errors.New("attempting to use a Durango-upgrade feature prior to activation")
	ErrAddValidatorTxPostDurango       = errors.New("AddValidatorTx is not permitted post-Durango")
	ErrAddDelegatorTxPostDurango       = errors.New("AddDelegatorTx is not permitted post-Durango")
)

// verifySubnetValidatorPrimaryNetworkRequirements verifies the primary
// network requirements for [subnetValidator]. An error is returned if they
// are not fulfilled.
func verifySubnetValidatorPrimaryNetworkRequirements(
	isDurangoActive bool,
	chainState state.Chain,
	subnetValidator txs.Validator,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure that the period this validator validates the specified subnet
// is a subset of the time they validate the primary network.

// verifyAddValidatorTx carries out the validation for an AddValidatorTx.
// It returns the tx outputs that should be returned if this validator is not
// added to the staking set.
func verifyAddValidatorTx(
	backend *Backend,
	feeCalculator fee.Calculator,
	chainState state.Chain,
	sTx *txs.Tx,
	tx *txs.AddValidatorTx,
) (
	[]*avax.TransferableOutput,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify the tx is well-formed

/*=isDurangoActive*/

// Ensure validator is staking at least the minimum amount

// Ensure validator isn't staking too much

// Ensure the validator fee is at least the minimum amount

// Ensure staking length is not too short

// Ensure staking length is not too long

/*=isDurangoActive*/

// Verify the flowcheck

// verifyAddSubnetValidatorTx carries out the validation for an
// AddSubnetValidatorTx.
func verifyAddSubnetValidatorTx(
	backend *Backend,
	feeCalculator fee.Calculator,
	chainState state.Chain,
	sTx *txs.Tx,
	tx *txs.AddSubnetValidatorTx,
) error {
	_ = "STUB: not implemented"
	// Verify the tx is well-formed
	return nil
}

// Ensure staking length is not too short

// Ensure staking length is not too long

// Verify the flowcheck

// Returns the representation of [tx.NodeID] validating [tx.Subnet].
// Returns true if [tx.NodeID] is a current validator of [tx.Subnet].
// Returns an error if the given tx is invalid.
// The transaction is valid if:
// * [tx.NodeID] is a current/pending PoA validator of [tx.Subnet].
// * [sTx]'s creds authorize it to spend the stated inputs.
// * [sTx]'s creds authorize it to remove a validator from [tx.Subnet].
// * The flow checker passes.
func verifyRemoveSubnetValidatorTx(
	backend *Backend,
	feeCalculator fee.Calculator,
	chainState state.Chain,
	sTx *txs.Tx,
	tx *txs.RemoveSubnetValidatorTx,
) (*state.Staker, bool, error) {
	_ = "STUB: not implemented"
	// Verify the tx is well-formed
	return nil, false, nil
}

// It isn't a current or pending validator.

// Not bootstrapped yet -- don't need to do full verification.

// verifyAddDelegatorTx carries out the validation for an AddDelegatorTx.
// It returns the tx outputs that should be returned if this delegator is not
// added to the staking set.
func verifyAddDelegatorTx(
	backend *Backend,
	feeCalculator fee.Calculator,
	chainState state.Chain,
	sTx *txs.Tx,
	tx *txs.AddDelegatorTx,
) (
	[]*avax.TransferableOutput,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Verify the tx is well-formed

/*=isDurangoActive*/

// Ensure staking length is not too short

// Ensure staking length is not too long

// Ensure validator is staking at least the minimum amount

/*=isDurangoActive*/

// Verify the flowcheck

// verifyAddPermissionlessValidatorTx carries out the validation for an
// AddPermissionlessValidatorTx.
func verifyAddPermissionlessValidatorTx(
	backend *Backend,
	feeCalculator fee.Calculator,
	chainState state.Chain,
	sTx *txs.Tx,
	tx *txs.AddPermissionlessValidatorTx,
) error {
	_ = "STUB: not implemented"
	// Verify the tx is well-formed
	return nil
}

// Ensure validator is staking at least the minimum amount

// Ensure validator isn't staking too much

// Ensure the validator fee is at least the minimum amount

// Ensure staking length is not too short

// Ensure staking length is not too long

// Wrong assetID used

// Verify the flowcheck

// verifyAddPermissionlessDelegatorTx carries out the validation for an
// AddPermissionlessDelegatorTx.
func verifyAddPermissionlessDelegatorTx(
	backend *Backend,
	feeCalculator fee.Calculator,
	chainState state.Chain,
	sTx *txs.Tx,
	tx *txs.AddPermissionlessDelegatorTx,
) error {
	_ = "STUB: not implemented"
	// Verify the tx is well-formed
	return nil
}

// Ensure delegator is staking at least the minimum amount

// Ensure staking length is not too short

// Ensure staking length is not too long

// Wrong assetID used

// Invariant: Delegators must only be able to reference validator
//            transactions that implement [txs.ValidatorTx]. All
//            validator transactions implement this interface except the
//            AddSubnetValidatorTx. AddSubnetValidatorTx is the only
//            permissioned validator, so we verify this delegator is
//            pointing to a permissionless validator.

// Verify the flowcheck

// Returns an error if the given tx is invalid.
// The transaction is valid if:
// * [sTx]'s creds authorize it to spend the stated inputs.
// * [sTx]'s creds authorize it to transfer ownership of [tx.Subnet].
// * The flow checker passes.
func verifyTransferSubnetOwnershipTx(
	backend *Backend,
	feeCalculator fee.Calculator,
	chainState state.Chain,
	sTx *txs.Tx,
	tx *txs.TransferSubnetOwnershipTx,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify the tx is well-formed

/*=isDurangoActive*/

// Not bootstrapped yet -- don't need to do full verification.

// Verify the flowcheck

// Ensure the proposed validator starts after the current time
func verifyStakerStartTime(isDurangoActive bool, chainTime, stakerTime time.Time) error {
	_ = "STUB: not implemented"
	// Pre Durango activation, start time must be after current chain time.
	// Post Durango activation, start time is not validated
	return nil
}
