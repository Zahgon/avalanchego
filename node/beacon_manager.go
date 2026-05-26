// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package node

import (
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/networking/router"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/version"
)

var _ router.ExternalHandler = (*beaconManager)(nil)

type beaconManager struct {
	router.ExternalHandler
	beacons                     validators.Manager
	requiredConns               int64
	numConns                    int64
	onSufficientlyConnected     chan struct{}
	onceOnSufficientlyConnected sync.Once
}

func (b *beaconManager) Connected(nodeID ids.NodeID, nodeVersion *version.Application, subnetID ids.ID) {
	_ = "STUB: not implemented"
	return
}

func (b *beaconManager) Disconnected(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }
