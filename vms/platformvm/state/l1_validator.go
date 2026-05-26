// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"errors"

	"github.com/google/btree"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/iterator"
	"github.com/ava-labs/avalanchego/utils/maybe"
)

var (
	_ btree.LessFunc[L1Validator] = L1Validator.Less
	_ utils.Sortable[L1Validator] = L1Validator{}

	ErrMutatedL1Validator     = errors.New("L1 validator contains mutated constant fields")
	ErrConflictingL1Validator = errors.New("L1 validator contains conflicting subnetID + nodeID pair")
	ErrDuplicateL1Validator   = errors.New("L1 validator contains duplicate subnetID + nodeID pair")
)

type L1Validators interface {
	// GetActiveL1ValidatorsIterator returns an iterator of all the active L1
	// validators in increasing order of EndAccumulatedFee.
	//
	// It is the caller's responsibility to call [Release] on the iterator after
	// use.
	//
	// It is not guaranteed to be safe to modify the state while using the
	// iterator. After releasing the iterator, the state may be safely modified.
	GetActiveL1ValidatorsIterator() (iterator.Iterator[L1Validator], error)

	// NumActiveL1Validators returns the number of currently active L1
	// validators.
	NumActiveL1Validators() int

	// WeightOfL1Validators returns the total active and inactive weight of L1
	// validators on [subnetID].
	WeightOfL1Validators(subnetID ids.ID) (uint64, error)

	// GetL1Validator returns the validator with [validationID] if it exists. If
	// the validator does not exist, [err] will equal [database.ErrNotFound].
	GetL1Validator(validationID ids.ID) (L1Validator, error)

	// HasL1Validator returns the validator with [validationID] if it exists.
	HasL1Validator(subnetID ids.ID, nodeID ids.NodeID) (bool, error)

	// PutL1Validator inserts [l1Validator] as a validator. If the weight of the
	// validator is 0, the validator is removed.
	//
	// If inserting this validator attempts to modify any of the constant fields
	// of the L1 validator struct, an error will be returned.
	//
	// If inserting this validator would cause the total weight of L1 validators
	// on a subnet to overflow MaxUint64, an error will be returned.
	//
	// If inserting this validator would cause there to be multiple validators
	// with the same subnetID and nodeID pair to exist at the same time, an
	// error will be returned.
	//
	// If an L1 validator is added with the same validationID as a previously
	// removed L1 validator, the behavior is undefined.
	PutL1Validator(l1Validator L1Validator) error
}

// L1Validator defines an ACP-77 validator. For a given ValidationID, it is
// expected for SubnetID, NodeID, PublicKey, RemainingBalanceOwner,
// DeactivationOwner, and StartTime to be constant.
type L1Validator struct {
	// ValidationID is not serialized because it is used as the key in the
	// database, so it doesn't need to be stored in the value.
	ValidationID ids.ID

	SubnetID ids.ID     `serialize:"true"`
	NodeID   ids.NodeID `serialize:"true"`

	// PublicKey is the uncompressed BLS public key of the validator. It is
	// guaranteed to be populated.
	PublicKey []byte `serialize:"true"`

	// RemainingBalanceOwner is the owner that will be used when returning the
	// balance of the validator after removing accrued fees.
	RemainingBalanceOwner []byte `serialize:"true"`

	// DeactivationOwner is the owner that can manually deactivate the
	// validator.
	DeactivationOwner []byte `serialize:"true"`

	// StartTime is the unix timestamp, in seconds, when this validator was
	// added to the set.
	StartTime uint64 `serialize:"true"`

	// Weight of this validator. It can be updated when the MinNonce is
	// increased. If the weight is being set to 0, the validator is being
	// removed.
	Weight uint64 `serialize:"true"`

	// MinNonce is the smallest nonce that can be used to modify this
	// validator's weight. It is initially set to 0 and is set to one higher
	// than the last nonce used. It is not valid to use nonce MaxUint64 unless
	// the weight is being set to 0, which removes the validator from the set.
	MinNonce uint64 `serialize:"true"`

	// EndAccumulatedFee is the amount of accumulated fees per validator that
	// can accrue before this validator must be deactivated. It is equal to the
	// amount of fees this validator is willing to pay plus the total amount of
	// fees a validator would have needed to pay from the activation of the Etna
	// upgrade until this validator was registered. Note that this relies on the
	// fact that every validator is charged the same fee for each unit of time.
	//
	// If this value is 0, the validator is inactive.
	EndAccumulatedFee uint64 `serialize:"true"`
}

func (v L1Validator) Less(o L1Validator) bool { _ = "STUB: not implemented"; return false }

// Compare determines a canonical ordering of L1 validators based on their
// EndAccumulatedFees and ValidationIDs. Lower EndAccumulatedFees result in an
// earlier ordering.
func (v L1Validator) Compare(o L1Validator) int { _ = "STUB: not implemented"; return 0 }

// immutableFieldsAreUnmodified returns true if two versions of the same
// validator are valid. Either because the validationID has changed or because
// no unexpected fields have been modified.
func (v L1Validator) immutableFieldsAreUnmodified(o L1Validator) bool {
	_ = "STUB: not implemented"
	return false
}

func (v L1Validator) isDeleted() bool { _ = "STUB: not implemented"; return false }

func (v L1Validator) IsActive() bool { _ = "STUB: not implemented"; return false }

func (v L1Validator) effectiveValidationID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (v L1Validator) effectiveNodeID() ids.NodeID {
	_ = "STUB: not implemented"
	return *new(ids.NodeID)
}

func (v L1Validator) effectivePublicKey() *bls.PublicKey { _ = "STUB: not implemented"; return nil }

func (v L1Validator) effectivePublicKeyBytes() []byte { _ = "STUB: not implemented"; return nil }

func getL1Validator(
	cache cache.Cacher[ids.ID, maybe.Maybe[L1Validator]],
	db database.KeyValueReader,
	validationID ids.ID,
) (L1Validator, error) {
	_ = "STUB: not implemented"
	return *new(L1Validator), nil
}

func putL1Validator(
	db database.KeyValueWriter,
	cache cache.Cacher[ids.ID, maybe.Maybe[L1Validator]],
	l1Validator L1Validator,
) error {
	_ = "STUB: not implemented"
	return nil
}

func deleteL1Validator(
	db database.KeyValueDeleter,
	cache cache.Cacher[ids.ID, maybe.Maybe[L1Validator]],
	validationID ids.ID,
) error {
	_ = "STUB: not implemented"
	return nil
}

type l1ValidatorsDiff struct {
	netAddedActive      int               // May be negative
	modifiedTotalWeight map[ids.ID]uint64 // subnetID -> totalWeight
	modified            map[ids.ID]L1Validator
	modifiedHasNodeIDs  map[subnetIDNodeID]bool
	active              *btree.BTreeG[L1Validator]
}

func newL1ValidatorsDiff() *l1ValidatorsDiff { _ = "STUB: not implemented"; return nil }

// getActiveL1ValidatorsIterator takes in the parent iterator, removes all
// modified validators, and then adds all modified active validators.
func (d *l1ValidatorsDiff) getActiveL1ValidatorsIterator(parentIterator iterator.Iterator[L1Validator]) iterator.Iterator[L1Validator] {
	_ = "STUB: not implemented"
	return nil
}

func (d *l1ValidatorsDiff) hasL1Validator(subnetID ids.ID, nodeID ids.NodeID) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (d *l1ValidatorsDiff) putL1Validator(state Chain, l1Validator L1Validator) error {
	_ = "STUB: not implemented"
	return nil
}

// Verify that there is not a legacy subnet validator with the same
// subnetID+nodeID as this L1 validator.

type activeL1Validators struct {
	lookup map[ids.ID]L1Validator
	tree   *btree.BTreeG[L1Validator]
}

func newActiveL1Validators() *activeL1Validators { _ = "STUB: not implemented"; return nil }

func (a *activeL1Validators) get(validationID ids.ID) (L1Validator, bool) {
	_ = "STUB: not implemented"
	return *new(L1Validator), false
}

func (a *activeL1Validators) put(l1Validator L1Validator) { _ = "STUB: not implemented"; return }

func (a *activeL1Validators) delete(validationID ids.ID) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *activeL1Validators) len() int { _ = "STUB: not implemented"; return 0 }

func (a *activeL1Validators) newIterator() iterator.Iterator[L1Validator] {
	_ = "STUB: not implemented"
	return nil
}

func (a *activeL1Validators) addStakersToValidatorManager(vdrs validators.Manager) error {
	_ = "STUB: not implemented"
	return nil
}

func addL1ValidatorToValidatorManager(vdrs validators.Manager, l1Validator L1Validator) error {
	_ = "STUB: not implemented"
	return nil
}
