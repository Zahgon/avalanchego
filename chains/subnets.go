// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package chains

import (
	"errors"
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/subnets"
)

var ErrNoPrimaryNetworkConfig = errors.New("no subnet config for primary network found")

// Subnets holds the currently running subnets on this node
type Subnets struct {
	nodeID  ids.NodeID
	configs map[ids.ID]subnets.Config

	lock    sync.RWMutex
	subnets map[ids.ID]subnets.Subnet
}

// GetOrCreate returns a subnet running on this node, or creates one if it was
// not running before. Returns the subnet and if the subnet was created.
func (s *Subnets) GetOrCreate(subnetID ids.ID) (subnets.Subnet, bool) {
	_ = "STUB: not implemented"
	return *new(subnets.Subnet), false
}

// Default to the primary network config if a subnet config was not
// specified

// Bootstrapping returns the subnetIDs of any chains that are still
// bootstrapping.
func (s *Subnets) Bootstrapping() []ids.ID { _ = "STUB: not implemented"; return nil }

// NewSubnets returns an instance of Subnets
func NewSubnets(
	nodeID ids.NodeID,
	configs map[ids.ID]subnets.Config,
) (*Subnets, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
