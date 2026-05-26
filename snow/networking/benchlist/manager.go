// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package benchlist

import (
	"sync"

	"github.com/ava-labs/avalanchego/api/metrics"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/validators"
)

var _ Manager = (*manager)(nil)

// Manager provides an interface for a benchlist to register whether
// queries have been successful or unsuccessful and place validators with
// consistently failing queries on a benchlist to prevent waiting up to
// the full network timeout for their responses.
type Manager interface {
	// RegisterResponse registers that we receive a request response from
	// [nodeID] regarding [chainID] within the timeout
	RegisterResponse(chainID ids.ID, nodeID ids.NodeID)
	// RegisterFailure registers that a request to [nodeID] regarding
	// [chainID] timed out
	RegisterFailure(chainID ids.ID, nodeID ids.NodeID)

	// RegisterChain registers a new chain with metrics under [namespace]
	RegisterChain(ctx *snow.ConsensusContext) error
	// IsBenched returns true if messages to [nodeID] regarding chain [chainID]
	// should not be sent over the network and should immediately fail.
	// Returns false if such messages should be sent, or if the chain is unknown.
	IsBenched(chainID ids.ID, nodeID ids.NodeID) bool
	// GetBenched returns an array of chainIDs where the specified
	// [nodeID] is benched. If called on an id.ShortID that does
	// not map to a validator, it will return an empty array.
	GetBenched(nodeID ids.NodeID) []ids.ID
	// Shutdown stops all chain benchlists.
	Shutdown()
}

type manager struct {
	benchable Benchable
	vdrs      validators.Manager
	reg       metrics.MultiGatherer
	config    Config

	lock   sync.RWMutex
	chains map[ids.ID]*benchlist

	shutdownOnce sync.Once
	shutdown     bool
}

// NewManager returns a manager for chain-specific query benchlisting
func NewManager(
	benchable Benchable,
	vdrs validators.Manager,
	reg metrics.MultiGatherer,
	config Config,
) Manager {
	_ = "STUB: not implemented"
	return *new(Manager)
}

// IsBenched returns true if messages to [nodeID] regarding [chainID]
// should not be sent over the network and should immediately fail.
func (m *manager) IsBenched(chainID ids.ID, nodeID ids.NodeID) bool {
	_ = "STUB: not implemented"
	return false
}

// GetBenched returns an array of chainIDs where the specified
// [nodeID] is benched. If called on an id.ShortID that does
// not map to a validator, it will return an empty array.
func (m *manager) GetBenched(nodeID ids.NodeID) []ids.ID { _ = "STUB: not implemented"; return nil }

func (m *manager) RegisterChain(ctx *snow.ConsensusContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) RegisterResponse(chainID ids.ID, nodeID ids.NodeID) {
	_ = "STUB: not implemented"
	return
}

func (m *manager) RegisterFailure(chainID ids.ID, nodeID ids.NodeID) {
	_ = "STUB: not implemented"
	return
}

func (m *manager) Shutdown() { _ = "STUB: not implemented"; return }

type noBenchlist struct{}

// NewNoBenchlist returns an empty benchlist that will never stop any queries
func NewNoBenchlist() Manager { _ = "STUB: not implemented"; return *new(Manager) }

func (noBenchlist) RegisterChain(*snow.ConsensusContext) error {
	_ = "STUB: not implemented"
	return nil
}

func (noBenchlist) RegisterResponse(ids.ID, ids.NodeID) { _ = "STUB: not implemented"; return }

func (noBenchlist) RegisterFailure(ids.ID, ids.NodeID) { _ = "STUB: not implemented"; return }

func (noBenchlist) IsBenched(ids.ID, ids.NodeID) bool { _ = "STUB: not implemented"; return false }

func (noBenchlist) GetBenched(ids.NodeID) []ids.ID { _ = "STUB: not implemented"; return nil }

func (noBenchlist) Shutdown() { _ = "STUB: not implemented"; return }
