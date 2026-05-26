// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"errors"

	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/fee"
)

// TODO: Before Etna, ensure that the maximum number of expiries to track is
// limited to a reasonable number by this window.
const (
	second                            = 1
	minute                            = 60 * second
	hour                              = 60 * minute
	day                               = 24 * hour
	RegisterL1ValidatorTxExpiryWindow = day
)

var (
	_ txs.Visitor = (*standardTxExecutor)(nil)

	errEmptyNodeID                      = errors.New("validator nodeID cannot be empty")
	errMaxStakeDurationTooLarge         = errors.New("max stake duration must be less than or equal to the global max stake duration")
	errMissingStartTimePreDurango       = errors.New("staker transactions must have a StartTime pre-Durango")
	errEtnaUpgradeNotActive             = errors.New("attempting to use an Etna-upgrade feature prior to activation")
	errTransformSubnetTxPostEtna        = errors.New("TransformSubnetTx is not permitted post-Etna")
	errMaxNumActiveValidators           = errors.New("already at the max number of active validators")
	errCouldNotLoadSubnetToL1Conversion = errors.New("could not load subnet conversion")
	errWrongWarpMessageSourceChainID    = errors.New("wrong warp message source chain ID")
	errWrongWarpMessageSourceAddress    = errors.New("wrong warp message source address")
	errWarpMessageExpired               = errors.New("warp message expired")
	errWarpMessageNotYetAllowed         = errors.New("warp message not yet allowed")
	errWarpMessageAlreadyIssued         = errors.New("warp message already issued")
	errCouldNotLoadL1Validator          = errors.New("could not load L1 validator")
	errWarpMessageContainsStaleNonce    = errors.New("warp message contains stale nonce")
	errRemovingLastValidator            = errors.New("attempting to remove the last L1 validator from a converted subnet")
	errStateCorruption                  = errors.New("state corruption")
)

// StandardTx executes the standard transaction [tx].
//
// [state] is modified to represent the state of the chain after the execution
// of [tx].
//
// Returns:
//   - The IDs of any import UTXOs consumed.
//   - The, potentially nil, atomic requests that should be performed against
//     shared memory when this transaction is accepted.
//   - A, potentially nil, function that should be called when this transaction
//     is accepted.
func StandardTx(
	backend *Backend,
	feeCalculator fee.Calculator,
	tx *txs.Tx,
	state *state.Diff,
) (set.Set[ids.ID], map[ids.ID]*atomic.Requests, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

type standardTxExecutor struct {
	// inputs, to be filled before visitor methods are called
	backend       *Backend
	state         *state.Diff // state is expected to be modified
	feeCalculator fee.Calculator
	tx            *txs.Tx

	// outputs of visitor execution
	onAccept       func() // may be nil
	inputs         set.Set[ids.ID]
	atomicRequests map[ids.ID]*atomic.Requests // may be nil
}

func (*standardTxExecutor) AdvanceTimeTx(*txs.AdvanceTimeTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (*standardTxExecutor) RewardValidatorTx(*txs.RewardValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *standardTxExecutor) AddValidatorTx(tx *txs.AddValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *standardTxExecutor) AddSubnetValidatorTx(tx *txs.AddSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *standardTxExecutor) AddDelegatorTx(tx *txs.AddDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *standardTxExecutor) CreateChainTx(tx *txs.CreateChainTx) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify the flowcheck

// Consume the UTXOS

// Produce the UTXOS

// Add the new chain to the database

// If this proposal is committed and this node is a member of the subnet
// that validates the blockchain, create the blockchain

func (e *standardTxExecutor) CreateSubnetTx(tx *txs.CreateSubnetTx) error {
	_ = "STUB: not implemented"
	// Make sure this transaction is well formed.
	return nil
}

// Verify the flowcheck

// Consume the UTXOS

// Produce the UTXOS

// Add the new subnet to the database

func (e *standardTxExecutor) ImportTx(tx *txs.ImportTx) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip verification of the shared memory inputs if the other primary
// network chains are not guaranteed to be up-to-date.

// Verify the flowcheck

// Consume the UTXOS

// Produce the UTXOS

// Note: We apply atomic requests even if we are not verifying atomic
// requests to ensure the shared state will be correct if we later start
// verifying the requests.

func (e *standardTxExecutor) ExportTx(tx *txs.ExportTx) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify the flowcheck

// Consume the UTXOS

// Produce the UTXOS

// Note: We apply atomic requests even if we are not verifying atomic
// requests to ensure the shared state will be correct if we later start
// verifying the requests.

// Verifies a [*txs.RemoveSubnetValidatorTx] and, if it passes, executes it on
// [e.State]. For verification rules, see [verifyRemoveSubnetValidatorTx]. This
// transaction will result in [tx.NodeID] being removed as a validator of
// [tx.SubnetID].
// Note: [tx.NodeID] may be either a current or pending validator.
func (e *standardTxExecutor) RemoveSubnetValidatorTx(tx *txs.RemoveSubnetValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

// Invariant: There are no permissioned subnet delegators to remove.

func (e *standardTxExecutor) TransformSubnetTx(tx *txs.TransformSubnetTx) error {
	_ = "STUB: not implemented"
	return nil
}

// Note: math.MaxInt32 * time.Second < math.MaxInt64 - so this can never
// overflow.

// Verify the flowcheck

// Invariant: [tx.AssetID != e.Ctx.AVAXAssetID]. This prevents the first
//            entry in this map literal from being overwritten by the
//            second entry.

// Consume the UTXOS

// Produce the UTXOS

// Transform the new subnet in the database

func (e *standardTxExecutor) AddPermissionlessValidatorTx(tx *txs.AddPermissionlessValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *standardTxExecutor) AddPermissionlessDelegatorTx(tx *txs.AddPermissionlessDelegatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

// Verifies a [*txs.TransferSubnetOwnershipTx] and, if it passes, executes it on
// [e.State]. For verification rules, see [verifyTransferSubnetOwnershipTx].
// This transaction will result in the ownership of [tx.Subnet] being transferred
// to [tx.Owner].
func (e *standardTxExecutor) TransferSubnetOwnershipTx(tx *txs.TransferSubnetOwnershipTx) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *standardTxExecutor) BaseTx(tx *txs.BaseTx) error { _ = "STUB: not implemented"; return nil }

// Verify the tx is well-formed

/*=isDurangoActive*/

// Verify the flowcheck

// Consume the UTXOS

// Produce the UTXOS

func (e *standardTxExecutor) ConvertSubnetToL1Tx(tx *txs.ConvertSubnetToL1Tx) error {
	_ = "STUB: not implemented"
	return nil
}

/*=isDurangoActive*/

// If Balance is 0, this is 0

// We are attempting to add an active validator

// Consume the UTXOS

// Produce the UTXOS

// Track the subnet conversion in the database

func (e *standardTxExecutor) RegisterL1ValidatorTx(tx *txs.RegisterL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

/*=isDurangoActive*/

// Verify the flowcheck

// Parse the warp message.

// Verify that the warp message was sent from the expected chain and
// address.

// Verify that the message contains a valid expiry time.

// Verify that this warp message isn't being replayed.

// Verify proof of possession provided by the transaction against the public
// key provided by the warp message.

// Create the L1 validator.

// If Balance is 0, this is will remain 0

// If the balance is non-zero, this validator should be initially active.

// Verify that there is space for an active validator.

// Mark the validator as active.

// Consume the UTXOS

// Produce the UTXOS

// Prevent this warp message from being replayed

func (e *standardTxExecutor) SetL1ValidatorWeightTx(tx *txs.SetL1ValidatorWeightTx) error {
	_ = "STUB: not implemented"
	return nil
}

/*=isDurangoActive*/

// Verify the flowcheck

// Parse the warp message.

// Verify that the message contains a valid nonce for a current validator.

// Verify that the warp message was sent from the expected chain and
// address.

// Check if we are removing the validator.

// Verify that we are not removing the last validator.

// If the validator is currently active, we need to refund the remaining
// balance.

// This check should be unreachable. However, it prevents AVAX
// from being minted due to state corruption. This also prevents
// invalid UTXOs from being created (with 0 value).

// If the weight is being set to 0, it is possible for the nonce increment
// to overflow. However, the validator is being removed and the nonce
// doesn't matter. If weight is not 0, [msg.Nonce] is enforced by
// [msg.Verify()] to be less than MaxUInt64 and can therefore be incremented
// without overflow.

// Consume the UTXOS

// Produce the UTXOS

func (e *standardTxExecutor) IncreaseL1ValidatorBalanceTx(tx *txs.IncreaseL1ValidatorBalanceTx) error {
	_ = "STUB: not implemented"
	return nil
}

/*=isDurangoActive*/

// Verify the flowcheck

// If the validator is currently inactive, we are activating it.

// Consume the UTXOS

// Produce the UTXOS

func (e *standardTxExecutor) DisableL1ValidatorTx(tx *txs.DisableL1ValidatorTx) error {
	_ = "STUB: not implemented"
	return nil
}

/*=isDurangoActive*/

// Verify the flowcheck

// Consume the UTXOS

// Produce the UTXOS

// If the validator is already disabled, there is nothing to do.

// This check should be unreachable. However, including it ensures
// that AVAX can't get minted out of thin air due to state
// corruption.

// Disable the validator

// Creates the staker as defined in [stakerTx] and adds it to [e.State].
func (e *standardTxExecutor) putStaker(stakerTx txs.Staker) error {
	_ = "STUB: not implemented"
	return nil
}

// Pre-Durango, stakers set a future [StartTime] and are added to the
// pending staker set. They are promoted to the current staker set once
// the chain time reaches [StartTime].

// Only calculate the potentialReward for permissionless stakers.
// Recall that we only need to check if this is a permissioned
// validator as there are no permissioned delegators

// Post-Durango, stakers are immediately added to the current staker
// set. Their [StartTime] is the current chain time.

// verifyL1Conversion verifies that the L1 conversion of [subnetID] references
// the [expectedChainID] and [expectedAddress].
func verifyL1Conversion(
	state state.Chain,
	subnetID ids.ID,
	expectedChainID ids.ID,
	expectedAddress []byte,
) error {
	_ = "STUB: not implemented"
	return nil
}
