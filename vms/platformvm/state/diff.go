// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/iterator"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/platformvm/fx"
	"github.com/ava-labs/avalanchego/vms/platformvm/status"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var (
	_ Chain    = (*Diff)(nil)
	_ Versions = stateGetter{}

	ErrMissingParentState = errors.New("missing parent state")
)

// Diff is a copy-on-write layer on top of a parent [Chain]. It records
// mutations locally and applies them to the parent via [Diff.Apply].
type Diff struct {
	parentID      ids.ID
	stateVersions Versions

	timestamp                   time.Time
	feeState                    gas.State
	l1ValidatorExcess           gas.Gas
	accruedFees                 uint64
	parentNumActiveL1Validators int

	// Subnet ID --> supply of native asset of the subnet
	currentSupply map[ids.ID]uint64

	expiryDiff       *expiryDiff
	l1ValidatorsDiff *l1ValidatorsDiff

	currentStakerDiffs diffStakers
	// map of subnetID -> nodeID -> staking info
	modifiedStakingInfo map[ids.ID]map[ids.NodeID]StakingInfo
	pendingStakerDiffs  diffStakers

	addedSubnetIDs []ids.ID
	// Subnet ID --> Owner of the subnet
	subnetOwners map[ids.ID]fx.Owner
	// Subnet ID --> Conversion of the subnet
	subnetToL1Conversions map[ids.ID]SubnetToL1Conversion
	// Subnet ID --> Tx that transforms the subnet
	transformedSubnets map[ids.ID]*txs.Tx

	addedChains map[ids.ID][]*txs.Tx

	addedRewardUTXOs map[ids.ID][]*avax.UTXO

	addedTxs map[ids.ID]*txAndStatus

	// map of modified UTXOID -> *UTXO if the UTXO is nil, it has been removed
	modifiedUTXOs map[ids.ID]*avax.UTXO
}

// NewDiff returns a new [Diff] whose parent is identified by parentID within
// stateVersions.
func NewDiff(
	parentID ids.ID,
	stateVersions Versions,
	allowAddingStakerAfterDeletion StakerAdditionAfterDeletionLegality,
) (*Diff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type stateGetter struct {
	state Chain
}

func (s stateGetter) GetState(ids.ID) (Chain, bool) {
	_ = "STUB: not implemented"
	return *new(Chain), false
}

func NewDiffOn(parentState Chain, allowAddingStakerAfterDeletion StakerAdditionAfterDeletionLegality) (*Diff, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Diff) GetTimestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (d *Diff) SetTimestamp(timestamp time.Time) { _ = "STUB: not implemented"; return }

func (d *Diff) GetFeeState() gas.State { _ = "STUB: not implemented"; return *new(gas.State) }

func (d *Diff) SetFeeState(feeState gas.State) { _ = "STUB: not implemented"; return }

func (d *Diff) GetL1ValidatorExcess() gas.Gas { _ = "STUB: not implemented"; return *new(gas.Gas) }

func (d *Diff) SetL1ValidatorExcess(excess gas.Gas) { _ = "STUB: not implemented"; return }

func (d *Diff) GetAccruedFees() uint64 { _ = "STUB: not implemented"; return 0 }

func (d *Diff) SetAccruedFees(accruedFees uint64) { _ = "STUB: not implemented"; return }

func (d *Diff) GetCurrentSupply(subnetID ids.ID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If the subnet supply wasn't modified in this diff, ask the parent state.

func (d *Diff) SetCurrentSupply(subnetID ids.ID, currentSupply uint64) {
	_ = "STUB: not implemented"
	return
}

func (d *Diff) GetExpiryIterator() (iterator.Iterator[ExpiryEntry], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Diff) HasExpiry(entry ExpiryEntry) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (d *Diff) PutExpiry(entry ExpiryEntry) { _ = "STUB: not implemented"; return }

func (d *Diff) DeleteExpiry(entry ExpiryEntry) { _ = "STUB: not implemented"; return }

func (d *Diff) GetActiveL1ValidatorsIterator() (iterator.Iterator[L1Validator], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Diff) NumActiveL1Validators() int { _ = "STUB: not implemented"; return 0 }

func (d *Diff) WeightOfL1Validators(subnetID ids.ID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Diff) GetL1Validator(validationID ids.ID) (L1Validator, error) {
	_ = "STUB: not implemented"
	return *new(L1Validator), nil
}

func (d *Diff) HasL1Validator(subnetID ids.ID, nodeID ids.NodeID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (d *Diff) PutL1Validator(l1Validator L1Validator) error { _ = "STUB: not implemented"; return nil }

func (d *Diff) GetCurrentValidator(subnetID ids.ID, nodeID ids.NodeID) (*Staker, error) {
	_ = "STUB: not implemented"
	// If the validator was modified in this diff, return the modified
	// validator.
	return nil, nil
}

// If the validator wasn't modified in this diff, ask the parent state.

func (d *Diff) SetStakingInfo(subnetID ids.ID, nodeID ids.NodeID, stakingInfo StakingInfo) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Diff) GetStakingInfo(subnetID ids.ID, nodeID ids.NodeID) (StakingInfo, error) {
	_ = "STUB: not implemented"
	return *new(StakingInfo), nil
}

func (d *Diff) PutCurrentValidator(staker *Staker) error { _ = "STUB: not implemented"; return nil }

func (d *Diff) DeleteCurrentValidator(staker *Staker) error { _ = "STUB: not implemented"; return nil }

func (d *Diff) GetCurrentDelegatorIterator(subnetID ids.ID, nodeID ids.NodeID) (iterator.Iterator[*Staker], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Diff) PutCurrentDelegator(staker *Staker) error { _ = "STUB: not implemented"; return nil }

func (d *Diff) DeleteCurrentDelegator(staker *Staker) error { _ = "STUB: not implemented"; return nil }

func (d *Diff) GetCurrentStakerIterator() (iterator.Iterator[*Staker], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Diff) GetPendingValidator(subnetID ids.ID, nodeID ids.NodeID) (*Staker, error) {
	_ = "STUB: not implemented"
	// If the validator was modified in this diff, return the modified
	// validator.
	return nil, nil
}

// If the validator wasn't modified in this diff, ask the parent state.

func (d *Diff) PutPendingValidator(staker *Staker) error { _ = "STUB: not implemented"; return nil }

func (d *Diff) DeletePendingValidator(staker *Staker) { _ = "STUB: not implemented"; return }

func (d *Diff) GetPendingDelegatorIterator(subnetID ids.ID, nodeID ids.NodeID) (iterator.Iterator[*Staker], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Diff) PutPendingDelegator(staker *Staker) { _ = "STUB: not implemented"; return }

func (d *Diff) DeletePendingDelegator(staker *Staker) { _ = "STUB: not implemented"; return }

func (d *Diff) GetPendingStakerIterator() (iterator.Iterator[*Staker], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Diff) AddSubnet(subnetID ids.ID) { _ = "STUB: not implemented"; return }

func (d *Diff) GetSubnetOwner(subnetID ids.ID) (fx.Owner, error) {
	_ = "STUB: not implemented"
	return *new(fx.Owner), nil
}

// If the subnet owner was not assigned in this diff, ask the parent state.

func (d *Diff) SetSubnetOwner(subnetID ids.ID, owner fx.Owner) { _ = "STUB: not implemented"; return }

func (d *Diff) GetSubnetToL1Conversion(subnetID ids.ID) (SubnetToL1Conversion, error) {
	_ = "STUB: not implemented"
	return *new(SubnetToL1Conversion), nil
}

// If the subnet conversion was not assigned in this diff, ask the parent state.

func (d *Diff) SetSubnetToL1Conversion(subnetID ids.ID, c SubnetToL1Conversion) {
	_ = "STUB: not implemented"
	return
}

func (d *Diff) GetSubnetTransformation(subnetID ids.ID) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the subnet wasn't transformed in this diff, ask the parent state.

func (d *Diff) AddSubnetTransformation(transformSubnetTxIntf *txs.Tx) {
	_ = "STUB: not implemented"
	return
}

func (d *Diff) AddChain(createChainTx *txs.Tx) { _ = "STUB: not implemented"; return }

func (d *Diff) GetTx(txID ids.ID) (*txs.Tx, status.Status, error) {
	_ = "STUB: not implemented"
	return nil, *new(status.Status), nil
}

func (d *Diff) AddTx(tx *txs.Tx, status status.Status) { _ = "STUB: not implemented"; return }

func (d *Diff) AddRewardUTXO(txID ids.ID, utxo *avax.UTXO) { _ = "STUB: not implemented"; return }

func (d *Diff) GetUTXO(utxoID ids.ID) (*avax.UTXO, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *Diff) AddUTXO(utxo *avax.UTXO) { _ = "STUB: not implemented"; return }

func (d *Diff) DeleteUTXO(utxoID ids.ID) { _ = "STUB: not implemented"; return }

func (d *Diff) Apply(baseState Chain) error { _ = "STUB: not implemented"; return nil }

// Ensure that all l1Validator deletions happen before any l1Validator
// additions. This ensures that a subnetID+nodeID pair that was deleted and
// then re-added in a single diff can't get reordered into the addition
// happening first; which would return an error.

// Delegators must be removed before their respective validators

// We might have removed the validator and then added it in the same diff.
// We therefore first delete and then only after add it.

// Delegators must be added after validators are added

// We might have removed the validator and then added it in the same diff.
// We therefore first delete and then only after add it.

// addCurrentDelegators adds all delegators for validator to baseState
func addCurrentDelegators(state Chain, validator *diffValidator) error {
	_ = "STUB: not implemented"
	return nil
}
