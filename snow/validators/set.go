// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"errors"
	"math/big"
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/sampler"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	errDuplicateValidator   = errors.New("duplicate validator")
	errMissingValidator     = errors.New("missing validator")
	errTotalWeightNotUint64 = errors.New("total weight is not a uint64")
	errInsufficientWeight   = errors.New("insufficient weight")
)

// newSet returns a new, empty set of validators.
func newSet(subnetID ids.ID, callbackListeners []ManagerCallbackListener) *vdrSet {
	_ = "STUB: not implemented"
	return nil
}

type vdrSet struct {
	subnetID ids.ID

	lock        sync.RWMutex
	vdrs        map[ids.NodeID]*Validator
	vdrSlice    []*Validator
	weights     []uint64
	totalWeight *big.Int

	samplerInitialized bool
	sampler            sampler.WeightedWithoutReplacement

	managerCallbackListeners []ManagerCallbackListener
	setCallbackListeners     []SetCallbackListener
}

func (s *vdrSet) Add(nodeID ids.NodeID, pk *bls.PublicKey, txID ids.ID, weight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *vdrSet) add(nodeID ids.NodeID, pk *bls.PublicKey, txID ids.ID, weight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *vdrSet) AddWeight(nodeID ids.NodeID, weight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *vdrSet) addWeight(nodeID ids.NodeID, weight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *vdrSet) GetWeight(nodeID ids.NodeID) uint64 { _ = "STUB: not implemented"; return 0 }

func (s *vdrSet) getWeight(nodeID ids.NodeID) uint64 { _ = "STUB: not implemented"; return 0 }

func (s *vdrSet) SubsetWeight(subset set.Set[ids.NodeID]) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *vdrSet) subsetWeight(subset set.Set[ids.NodeID]) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *vdrSet) RemoveWeight(nodeID ids.NodeID, weight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *vdrSet) removeWeight(nodeID ids.NodeID, weight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// We first calculate the new weight of the validator, as this guarantees
// that none of the following operations can underflow.

// Get the last element

// Move element at last index --> index of removed validator

// Remove validator

func (s *vdrSet) Get(nodeID ids.NodeID) (*Validator, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *vdrSet) get(nodeID ids.NodeID) (*Validator, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *vdrSet) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *vdrSet) len() int { _ = "STUB: not implemented"; return 0 }

func (s *vdrSet) HasCallbackRegistered() bool { _ = "STUB: not implemented"; return false }

func (s *vdrSet) Map() map[ids.NodeID]*GetValidatorOutput { _ = "STUB: not implemented"; return nil }

func (s *vdrSet) Sample(size int) ([]ids.NodeID, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *vdrSet) sample(size int) ([]ids.NodeID, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *vdrSet) TotalWeight() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *vdrSet) String() string { _ = "STUB: not implemented"; return "" }

func (s *vdrSet) PrefixedString(prefix string) string { _ = "STUB: not implemented"; return "" }

func (s *vdrSet) prefixedString(prefix string) string { _ = "STUB: not implemented"; return "" }

func (s *vdrSet) RegisterManagerCallbackListener(callbackListener ManagerCallbackListener) {
	_ = "STUB: not implemented"
	return
}

func (s *vdrSet) RegisterCallbackListener(callbackListener SetCallbackListener) {
	_ = "STUB: not implemented"
	return
}

// Assumes [s.lock] is held
func (s *vdrSet) callWeightChangeCallbacks(node ids.NodeID, oldWeight, newWeight uint64) {
	_ = "STUB: not implemented"
	return
}

// Assumes [s.lock] is held
func (s *vdrSet) callValidatorAddedCallbacks(node ids.NodeID, pk *bls.PublicKey, txID ids.ID, weight uint64) {
	_ = "STUB: not implemented"
	return
}

// Assumes [s.lock] is held
func (s *vdrSet) callValidatorRemovedCallbacks(node ids.NodeID, weight uint64) {
	_ = "STUB: not implemented"
	return
}

func (s *vdrSet) GetValidatorIDs() []ids.NodeID { _ = "STUB: not implemented"; return nil }
