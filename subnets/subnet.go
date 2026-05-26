// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package subnets

import (
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/set"
)

var _ Subnet = (*subnet)(nil)

type Allower interface {
	// IsAllowed filters out nodes that are not allowed to connect to this subnet
	IsAllowed(nodeID ids.NodeID, isValidator bool) bool
}

// Subnet keeps track of the currently bootstrapping chains in a subnet. If no
// chains in the subnet are currently bootstrapping, the subnet is considered
// bootstrapped.
type Subnet interface {
	common.BootstrapTracker

	// AddChain adds a chain to this Subnet
	AddChain(chainID ids.ID) bool

	// Config returns config of this Subnet
	Config() Config

	Allower
}

type subnet struct {
	lock            sync.RWMutex
	bootstrapping   set.Set[ids.ID]
	bootstrapped    set.Set[ids.ID]
	config          Config
	myNodeID        ids.NodeID
	bootstrapSignal common.PreemptionSignal
}

func New(myNodeID ids.NodeID, config Config) Subnet { _ = "STUB: not implemented"; return *new(Subnet) }

func (s *subnet) AllBootstrapped() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (s *subnet) IsBootstrapped() bool { _ = "STUB: not implemented"; return false }

func (s *subnet) Bootstrapped(chainID ids.ID) { _ = "STUB: not implemented"; return }

func (s *subnet) AddChain(chainID ids.ID) bool { _ = "STUB: not implemented"; return false }

func (s *subnet) Config() Config { _ = "STUB: not implemented"; return *new(Config) }

func (s *subnet) IsAllowed(nodeID ids.NodeID, isValidator bool) bool {
	_ = "STUB: not implemented"
	// Case 1: NodeID is this node
	// Case 2: This subnet is not validator-only subnet
	// Case 3: NodeID is a validator for this chain
	// Case 4: NodeID is explicitly allowed whether it's subnet validator or not
	return false
}
