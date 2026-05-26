// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tracker

import (
	"context"
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/version"
)

var (
	_ Peers = (*lockedPeers)(nil)
	_ Peers = (*meteredPeers)(nil)
	_ Peers = (*peerData)(nil)
)

type Peers interface {
	validators.SetCallbackListener
	validators.Connector

	// ConnectedWeight returns the currently connected stake weight
	ConnectedWeight() uint64
	// ConnectedPercent returns the currently connected stake percentage [0, 1]
	ConnectedPercent() float64
	// SampleValidator returns a randomly selected connected validator. If there
	// are no currently connected validators then it will return false.
	SampleValidator() (ids.NodeID, bool)
	// GetValidators returns the set of all validators
	// known to this peer manager
	GetValidators() set.Set[ids.NodeID]
	// ConnectedValidators returns the set of all validators
	// that are currently connected
	ConnectedValidators() set.Set[ids.NodeID]
}

type lockedPeers struct {
	lock  sync.RWMutex
	peers Peers
}

func NewPeers() Peers { _ = "STUB: not implemented"; return *new(Peers) }

func (p *lockedPeers) OnValidatorAdded(nodeID ids.NodeID, pk *bls.PublicKey, txID ids.ID, weight uint64) {
	_ = "STUB: not implemented"
	return
}

func (p *lockedPeers) OnValidatorRemoved(nodeID ids.NodeID, weight uint64) {
	_ = "STUB: not implemented"
	return
}

func (p *lockedPeers) OnValidatorWeightChanged(nodeID ids.NodeID, oldWeight, newWeight uint64) {
	_ = "STUB: not implemented"
	return
}

func (p *lockedPeers) Connected(ctx context.Context, nodeID ids.NodeID, version *version.Application) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *lockedPeers) Disconnected(ctx context.Context, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *lockedPeers) ConnectedWeight() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *lockedPeers) ConnectedPercent() float64 { _ = "STUB: not implemented"; return 0 }

func (p *lockedPeers) SampleValidator() (ids.NodeID, bool) {
	_ = "STUB: not implemented"
	return *new(ids.NodeID), false
}

func (p *lockedPeers) GetValidators() set.Set[ids.NodeID] { _ = "STUB: not implemented"; return nil }

func (p *lockedPeers) ConnectedValidators() set.Set[ids.NodeID] {
	_ = "STUB: not implemented"
	return nil
}

type meteredPeers struct {
	Peers

	percentConnected prometheus.Gauge
	numValidators    prometheus.Gauge
	totalWeight      prometheus.Gauge
}

func NewMeteredPeers(reg prometheus.Registerer) (Peers, error) {
	_ = "STUB: not implemented"
	return *new(Peers), nil
}

func (p *meteredPeers) OnValidatorAdded(nodeID ids.NodeID, pk *bls.PublicKey, txID ids.ID, weight uint64) {
	_ = "STUB: not implemented"
	return
}

func (p *meteredPeers) OnValidatorRemoved(nodeID ids.NodeID, weight uint64) {
	_ = "STUB: not implemented"
	return
}

func (p *meteredPeers) OnValidatorWeightChanged(nodeID ids.NodeID, oldWeight, newWeight uint64) {
	_ = "STUB: not implemented"
	return
}

func (p *meteredPeers) Connected(ctx context.Context, nodeID ids.NodeID, version *version.Application) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *meteredPeers) Disconnected(ctx context.Context, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

type peerData struct {
	// validators maps nodeIDs to their current stake weight
	validators map[ids.NodeID]uint64
	// totalWeight is the total weight of all validators
	totalWeight uint64
	// connectedWeight contains the sum of all connected validator weights
	connectedWeight uint64
	// connectedValidators is the set of currently connected peers with a
	// non-zero stake weight
	connectedValidators set.Set[ids.NodeID]
	// connectedPeers is the set of all connected peers
	connectedPeers set.Set[ids.NodeID]
}

func (p *peerData) OnValidatorAdded(nodeID ids.NodeID, _ *bls.PublicKey, _ ids.ID, weight uint64) {
	_ = "STUB: not implemented"
	return
}

func (p *peerData) OnValidatorRemoved(nodeID ids.NodeID, weight uint64) {
	_ = "STUB: not implemented"
	return
}

func (p *peerData) OnValidatorWeightChanged(nodeID ids.NodeID, oldWeight, newWeight uint64) {
	_ = "STUB: not implemented"
	return
}

func (p *peerData) Connected(_ context.Context, nodeID ids.NodeID, _ *version.Application) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *peerData) Disconnected(_ context.Context, nodeID ids.NodeID) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *peerData) ConnectedWeight() uint64 { _ = "STUB: not implemented"; return 0 }

func (p *peerData) ConnectedPercent() float64 { _ = "STUB: not implemented"; return 0 }

func (p *peerData) SampleValidator() (ids.NodeID, bool) {
	_ = "STUB: not implemented"
	return *new(ids.NodeID), false
}

func (p *peerData) GetValidators() set.Set[ids.NodeID] { _ = "STUB: not implemented"; return nil }

func (p *peerData) ConnectedValidators() set.Set[ids.NodeID] {
	_ = "STUB: not implemented"
	// The set is copied to avoid future changes from being reflected in the
	// returned set.
	return nil
}
