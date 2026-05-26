// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tracker

import (
	"context"
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/version"
)

var _ Startup = (*startup)(nil)

type Startup interface {
	Peers

	ShouldStart() bool
}

type startup struct {
	Peers

	lock          sync.RWMutex
	startupWeight uint64
	shouldStart   bool
}

func NewStartup(peers Peers, startupWeight uint64) Startup {
	_ = "STUB: not implemented"
	return *new(Startup)
}

func (s *startup) OnValidatorAdded(nodeID ids.NodeID, pk *bls.PublicKey, txID ids.ID, weight uint64) {
	_ = "STUB: not implemented"
	return
}

func (s *startup) OnValidatorWeightChanged(nodeID ids.NodeID, oldWeight, newWeight uint64) {
	_ = "STUB: not implemented"
	return
}

func (s *startup) Connected(ctx context.Context, nodeID ids.NodeID, nodeVersion *version.Application) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *startup) ShouldStart() bool { _ = "STUB: not implemented"; return false }
