// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package validators

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	_ Manager = (*manager)(nil)

	ErrZeroWeight        = errors.New("weight must be non-zero")
	ErrMissingValidators = errors.New("missing validators")
)

type ManagerCallbackListener interface {
	OnValidatorAdded(subnetID ids.ID, nodeID ids.NodeID, pk *bls.PublicKey, txID ids.ID, weight uint64)
	OnValidatorRemoved(subnetID ids.ID, nodeID ids.NodeID, weight uint64)
	OnValidatorWeightChanged(subnetID ids.ID, nodeID ids.NodeID, oldWeight, newWeight uint64)
}

type SetCallbackListener interface {
	OnValidatorAdded(nodeID ids.NodeID, pk *bls.PublicKey, txID ids.ID, weight uint64)
	OnValidatorRemoved(nodeID ids.NodeID, weight uint64)
	OnValidatorWeightChanged(nodeID ids.NodeID, oldWeight, newWeight uint64)
}

// Manager holds the validator set of each subnet
type Manager interface {
	fmt.Stringer

	// Add a new staker to the subnet.
	// Returns an error if:
	// - [weight] is 0
	// - [nodeID] is already in the validator set
	// If an error is returned, the set will be unmodified.
	AddStaker(subnetID ids.ID, nodeID ids.NodeID, pk *bls.PublicKey, txID ids.ID, weight uint64) error

	// AddWeight to an existing staker to the subnet.
	// Returns an error if:
	// - [weight] is 0
	// - [nodeID] is not already in the validator set
	// If an error is returned, the set will be unmodified.
	// AddWeight can result in a total weight that overflows uint64.
	// In this case no error will be returned for this call.
	// However, the next TotalWeight call will return an error.
	AddWeight(subnetID ids.ID, nodeID ids.NodeID, weight uint64) error

	// GetWeight retrieves the validator weight from the subnet.
	GetWeight(subnetID ids.ID, nodeID ids.NodeID) uint64

	// GetValidator returns the validator tied to the specified ID in subnet.
	// If the validator doesn't exist, returns false.
	GetValidator(subnetID ids.ID, nodeID ids.NodeID) (*Validator, bool)

	// GetValidatorIDs returns the validator IDs in the subnet.
	GetValidatorIDs(subnetID ids.ID) []ids.NodeID

	// SubsetWeight returns the sum of the weights of the validators in the subnet.
	// Returns err if subset weight overflows uint64.
	SubsetWeight(subnetID ids.ID, validatorIDs set.Set[ids.NodeID]) (uint64, error)

	// RemoveWeight from a staker in the subnet. If the staker's weight becomes 0, the staker
	// will be removed from the subnet set.
	// Returns an error if:
	// - [weight] is 0
	// - [nodeID] is not already in the subnet set
	// - the weight of the validator would become negative
	// If an error is returned, the set will be unmodified.
	RemoveWeight(subnetID ids.ID, nodeID ids.NodeID, weight uint64) error

	// NumSubnets returns the number of subnets with non-zero weight.
	NumSubnets() int

	// NumValidators returns the number of validators currently in the subnet.
	NumValidators(subnetID ids.ID) int

	// TotalWeight returns the cumulative weight of all validators in the subnet.
	// Returns err if total weight overflows uint64.
	TotalWeight(subnetID ids.ID) (uint64, error)

	// Sample returns a collection of validatorIDs in the subnet, potentially with duplicates.
	// If sampling the requested size isn't possible, an error will be returned.
	Sample(subnetID ids.ID, size int) ([]ids.NodeID, error)

	// GetAllMaps returns a copy of all validators of all subnets
	GetAllMaps() map[ids.ID]map[ids.NodeID]*GetValidatorOutput

	// Map of the validators in this subnet
	GetMap(subnetID ids.ID) map[ids.NodeID]*GetValidatorOutput

	// When a validator is added, removed, or its weight changes, the listener
	// will be notified of the event.
	RegisterCallbackListener(listener ManagerCallbackListener)

	// When a validator is added, removed, or its weight changes on [subnetID],
	// the listener will be notified of the event.
	RegisterSetCallbackListener(subnetID ids.ID, listener SetCallbackListener)
}

// NewManager returns a new, empty manager
func NewManager() Manager { _ = "STUB: not implemented"; return *new(Manager) }

type manager struct {
	lock sync.RWMutex

	// Key: Subnet ID
	// Value: The validators that validate the subnet
	subnetToVdrs      map[ids.ID]*vdrSet
	callbackListeners []ManagerCallbackListener
}

func (m *manager) AddStaker(subnetID ids.ID, nodeID ids.NodeID, pk *bls.PublicKey, txID ids.ID, weight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) AddWeight(subnetID ids.ID, nodeID ids.NodeID, weight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) GetWeight(subnetID ids.ID, nodeID ids.NodeID) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (m *manager) GetValidator(subnetID ids.ID, nodeID ids.NodeID) (*Validator, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (m *manager) SubsetWeight(subnetID ids.ID, validatorIDs set.Set[ids.NodeID]) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *manager) RemoveWeight(subnetID ids.ID, nodeID ids.NodeID, weight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// If this was the last validator in the subnet and no callback listeners
// are registered, remove the subnet

func (m *manager) NumSubnets() int { _ = "STUB: not implemented"; return 0 }

func (m *manager) NumValidators(subnetID ids.ID) int { _ = "STUB: not implemented"; return 0 }

func (m *manager) TotalWeight(subnetID ids.ID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *manager) Sample(subnetID ids.ID, size int) ([]ids.NodeID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *manager) GetAllMaps() map[ids.ID]map[ids.NodeID]*GetValidatorOutput {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) GetMap(subnetID ids.ID) map[ids.NodeID]*GetValidatorOutput {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) RegisterCallbackListener(listener ManagerCallbackListener) {
	_ = "STUB: not implemented"
	return
}

func (m *manager) RegisterSetCallbackListener(subnetID ids.ID, listener SetCallbackListener) {
	_ = "STUB: not implemented"
	return
}

func (m *manager) String() string { _ = "STUB: not implemented"; return "" }

func (m *manager) GetValidatorIDs(subnetID ids.ID) []ids.NodeID {
	_ = "STUB: not implemented"
	return nil
}
