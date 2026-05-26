// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package node

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/set"
)

var _ validators.Manager = (*overriddenManager)(nil)

// newOverriddenManager returns a Manager that overrides of all calls to the
// underlying Manager to only operate on the validators in [subnetID].
func newOverriddenManager(subnetID ids.ID, manager validators.Manager) *overriddenManager {
	_ = "STUB: not implemented"
	return nil
}

// overriddenManager is a wrapper around a Manager that overrides of all calls
// to the underlying Manager to only operate on the validators in [subnetID].
// subnetID here is typically the primary network ID, as it has the superset of
// all subnet validators.
type overriddenManager struct {
	manager  validators.Manager
	subnetID ids.ID
}

func (o *overriddenManager) AddStaker(_ ids.ID, nodeID ids.NodeID, pk *bls.PublicKey, txID ids.ID, weight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *overriddenManager) AddWeight(_ ids.ID, nodeID ids.NodeID, weight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *overriddenManager) GetWeight(_ ids.ID, nodeID ids.NodeID) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func (o *overriddenManager) GetValidator(_ ids.ID, nodeID ids.NodeID) (*validators.Validator, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (o *overriddenManager) SubsetWeight(_ ids.ID, nodeIDs set.Set[ids.NodeID]) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *overriddenManager) RemoveWeight(_ ids.ID, nodeID ids.NodeID, weight uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (o *overriddenManager) NumSubnets() int { _ = "STUB: not implemented"; return 0 }

func (o *overriddenManager) NumValidators(ids.ID) int { _ = "STUB: not implemented"; return 0 }

func (o *overriddenManager) TotalWeight(ids.ID) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o *overriddenManager) Sample(_ ids.ID, size int) ([]ids.NodeID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (o *overriddenManager) GetAllMaps() map[ids.ID]map[ids.NodeID]*validators.GetValidatorOutput {
	_ = "STUB: not implemented"
	return nil
}

func (o *overriddenManager) GetMap(ids.ID) map[ids.NodeID]*validators.GetValidatorOutput {
	_ = "STUB: not implemented"
	return nil
}

func (o *overriddenManager) RegisterCallbackListener(listener validators.ManagerCallbackListener) {
	_ = "STUB: not implemented"
	return
}

func (o *overriddenManager) RegisterSetCallbackListener(_ ids.ID, listener validators.SetCallbackListener) {
	_ = "STUB: not implemented"
	return
}

func (o *overriddenManager) String() string { _ = "STUB: not implemented"; return "" }

func (o *overriddenManager) GetValidatorIDs(ids.ID) []ids.NodeID {
	_ = "STUB: not implemented"
	return nil
}
