// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"errors"

	"github.com/google/btree"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/iterator"
)

var (
	ErrAddingStakerAfterDeletion = errors.New("attempted to add a staker after deleting it")
	errUnexpectedStaker          = errors.New("unexpected staker")
)

// StakerAdditionAfterDeletionLegality specifies whether a staker can be added after being deleted in the same diff.
// Pre Helicon it is forbidden, and post Helicon it is allowed.
type StakerAdditionAfterDeletionLegality bool

const (
	StakerAdditionAfterDeletionAllowed   StakerAdditionAfterDeletionLegality = true
	StakerAdditionAfterDeletionForbidden StakerAdditionAfterDeletionLegality = false
)

type Stakers interface {
	CurrentStakers
	PendingStakers
}

type CurrentStakers interface {
	// GetCurrentValidator returns the Staker describing the validator on subnetID with nodeID.
	// [database.ErrNotFound] is returned if the validator is not in the validator set.
	GetCurrentValidator(subnetID ids.ID, nodeID ids.NodeID) (*Staker, error)

	// PutCurrentValidator adds the Staker to the validator set.
	//
	// This returns an error if staker is already in the validator set.
	PutCurrentValidator(staker *Staker) error

	// DeleteCurrentValidator removes the Staker from the validator set.
	//
	// This returns an error if staker is not already in the validator set or if there are delegators
	// for staker still present.
	DeleteCurrentValidator(staker *Staker) error

	// SetStakingInfo updates the mutable staking info for nodeID on subnetID.
	//
	// This returns an error if the validator is not in the validator set.
	// TODO should support sets in the same block that a validator is added.
	SetStakingInfo(subnetID ids.ID, nodeID ids.NodeID, stakingInfo StakingInfo) error

	// GetStakingInfo returns the mutable staking info for nodeID on subnetID.
	//
	// This returns an error if the validator is not in the validator set.
	// TODO should support gets in the same block that a validator is added.
	GetStakingInfo(subnetID ids.ID, nodeID ids.NodeID) (StakingInfo, error)

	// GetCurrentDelegatorIterator returns the delegators associated with the
	// validator on subnetID with nodeID. Delegators are sorted by their
	// removal from current staker set (i.e. Staker.NextTime).
	//
	// This returns an empty iterator if the validator is not in the validator set.
	GetCurrentDelegatorIterator(subnetID ids.ID, nodeID ids.NodeID) (iterator.Iterator[*Staker], error)

	// PutCurrentDelegator adds the staker describing a delegator to the
	// staker set.
	//
	// This returns an error if the validator is not in the validator set.
	//
	// Invariant: staker is not currently a CurrentDelegator
	// TODO error if the delegator is already present
	PutCurrentDelegator(staker *Staker) error

	// DeleteCurrentDelegator removes the staker describing a delegator from
	// the staker set.
	//
	// This returns an error if the validator is not in the validator set.
	//
	// Invariant: staker is currently a CurrentDelegator
	// TODO error if the delegator was not present
	DeleteCurrentDelegator(staker *Staker) error

	// GetCurrentStakerIterator returns stakers in order of their removal from
	// the current staker set.
	GetCurrentStakerIterator() (iterator.Iterator[*Staker], error)
}

type PendingStakers interface {
	// GetPendingValidator returns the Staker describing the validator on
	// [subnetID] with [nodeID]. If the validator does not exist,
	// [database.ErrNotFound] is returned.
	GetPendingValidator(subnetID ids.ID, nodeID ids.NodeID) (*Staker, error)

	// PutPendingValidator adds the [staker] describing a validator to the
	// staker set.
	PutPendingValidator(staker *Staker) error

	// DeletePendingValidator removes the [staker] describing a validator from
	// the staker set.
	DeletePendingValidator(staker *Staker)

	// GetPendingDelegatorIterator returns the delegators associated with the
	// validator on [subnetID] with [nodeID]. Delegators are sorted by their
	// removal from pending staker set.
	GetPendingDelegatorIterator(subnetID ids.ID, nodeID ids.NodeID) (iterator.Iterator[*Staker], error)

	// PutPendingDelegator adds the [staker] describing a delegator to the
	// staker set.
	PutPendingDelegator(staker *Staker)

	// DeletePendingDelegator removes the [staker] describing a delegator from
	// the staker set.
	DeletePendingDelegator(staker *Staker)

	// GetPendingStakerIterator returns stakers in order of their removal from
	// the pending staker set.
	GetPendingStakerIterator() (iterator.Iterator[*Staker], error)
}

type baseStakers struct {
	// subnetID --> nodeID --> current state for the validator of the subnet
	validators map[ids.ID]map[ids.NodeID]*baseStaker
	stakers    *btree.BTreeG[*Staker]
	// subnetID --> nodeID --> diff for that validator since the last db write
	validatorDiffs map[ids.ID]map[ids.NodeID]*diffValidator
}

type baseStaker struct {
	validator  *Staker
	delegators *btree.BTreeG[*Staker]
}

func newBaseStakers() *baseStakers { _ = "STUB: not implemented"; return nil }

func (v *baseStakers) GetValidator(subnetID ids.ID, nodeID ids.NodeID) (*Staker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *baseStakers) PutValidator(staker *Staker) { _ = "STUB: not implemented"; return }

func (v *baseStakers) DeleteValidator(staker *Staker) { _ = "STUB: not implemented"; return }

func (v *baseStakers) GetDelegatorIterator(subnetID ids.ID, nodeID ids.NodeID) iterator.Iterator[*Staker] {
	_ = "STUB: not implemented"
	return nil
}

func (v *baseStakers) PutDelegator(staker *Staker) { _ = "STUB: not implemented"; return }

func (v *baseStakers) DeleteDelegator(staker *Staker) { _ = "STUB: not implemented"; return }

func (v *baseStakers) GetStakerIterator() iterator.Iterator[*Staker] {
	_ = "STUB: not implemented"
	return nil
}

func (v *baseStakers) getOrCreateValidator(subnetID ids.ID, nodeID ids.NodeID) *baseStaker {
	_ = "STUB: not implemented"
	return nil
}

// pruneValidator assumes that the named validator is currently in the
// [validators] map.
func (v *baseStakers) pruneValidator(subnetID ids.ID, nodeID ids.NodeID) {
	_ = "STUB: not implemented"
	return
}

func (v *baseStakers) getOrCreateValidatorDiff(subnetID ids.ID, nodeID ids.NodeID) *diffValidator {
	_ = "STUB: not implemented"
	return nil
}

type diffStakers struct {
	// isAdditionAfterDeletionAllowed specifies whether a staker can be added after being deleted in the same diff.
	// This is done to preserve the pre-Helicon invariant that a staker cannot be added after being deleted,
	// while allowing post-Helicon diffs to do that.
	isAdditionAfterDeletionAllowed StakerAdditionAfterDeletionLegality
	// subnetID --> nodeID --> diff for that validator
	validatorDiffs map[ids.ID]map[ids.NodeID]*diffValidator
	addedStakers   *btree.BTreeG[*Staker]
	deletedStakers map[ids.ID]*Staker
}

type diffValidator struct {
	// added represents a validator that was added in this diff, or nil if no
	// validator was added. Can be non-nil at the same time as removed to represent a replacement.
	added *Staker
	// removed represents a validator that was removed in this diff, or nil if no
	// validator was removed. Can be non-nil at the same time as added to represent a replacement.
	removed           *Staker
	addedDelegators   *btree.BTreeG[*Staker]
	deletedDelegators map[ids.ID]*Staker
}

// weightChanges returns the total weight added to and removed from this
// validator by this diff. The added weight includes the added validator and all
// added delegators. The removed weight includes the removed validator and all
// deleted delegators.
func (d *diffValidator) weightChanges() (addedWeight uint64, removedWeight uint64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (d *diffValidator) WeightDiff() (ValidatorWeightDiff, error) {
	_ = "STUB: not implemented"
	return *new(ValidatorWeightDiff), nil
}

// GetValidator attempts to fetch the validator with the given subnetID and
// nodeID.
func (s *diffStakers) GetValidator(subnetID ids.ID, nodeID ids.NodeID) (*Staker, diffValidatorStatus) {
	_ = "STUB: not implemented"
	return nil, *new(diffValidatorStatus)
}

func (s *diffStakers) PutValidator(staker *Staker) error { _ = "STUB: not implemented"; return nil }

// Enforce the invariant that a validator cannot be added after being
// deleted.

// We set the removed field when we delete the validator that was not added in this diff before.
// So if we reached here, it means we removed it first and now either re-adding it.
// If we're re-adding the exact same validator, we should remove it from the deleted stakers set since it's no longer deleted.

// If we're re-adding the exact same validator that was removed,
// the two operations cancel out.

func (s *diffStakers) DeleteValidator(staker *Staker) { _ = "STUB: not implemented"; return }

// This validator was added in this diff. Rollback the addition.

func (s *diffStakers) GetDelegatorIterator(
	parentIterator iterator.Iterator[*Staker],
	subnetID ids.ID,
	nodeID ids.NodeID,
) iterator.Iterator[*Staker] {
	_ = "STUB: not implemented"
	return nil
}

func (s *diffStakers) PutDelegator(staker *Staker) { _ = "STUB: not implemented"; return }

func (s *diffStakers) DeleteDelegator(staker *Staker) { _ = "STUB: not implemented"; return }

func (s *diffStakers) GetStakerIterator(parentIterator iterator.Iterator[*Staker]) iterator.Iterator[*Staker] {
	_ = "STUB: not implemented"
	return nil
}

func (s *diffStakers) getOrCreateDiff(subnetID ids.ID, nodeID ids.NodeID) *diffValidator {
	_ = "STUB: not implemented"
	return nil
}
