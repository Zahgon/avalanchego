// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"errors"
	"net"
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/networking/router"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/subnets"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/set"
)

var (
	errClosed = errors.New("closed")

	_ net.Listener    = (*noopListener)(nil)
	_ subnets.Allower = (*nodeIDConnector)(nil)
)

type noopListener struct {
	once   sync.Once
	closed chan struct{}
}

func newNoopListener() net.Listener { _ = "STUB: not implemented"; return *new(net.Listener) }

func (l *noopListener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (l *noopListener) Close() error { _ = "STUB: not implemented"; return nil }

func (*noopListener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func NewTestNetworkConfig(
	metrics prometheus.Registerer,
	networkID uint32,
	currentValidators validators.Manager,
	trackedSubnets set.Set[ids.ID],
) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO actually monitor usage
// TestNetwork doesn't use disk so we don't need to track it, but we should
// still have guardrails around cpu/memory usage.

func NewTestNetwork(
	log logging.Logger,
	metrics prometheus.Registerer,
	cfg *Config,
	router router.ExternalHandler,
) (Network, error) {
	_ = "STUB: not implemented"
	return *new(Network), nil
}

// Must be updated for each network upgrade

type nodeIDConnector struct {
	nodeID ids.NodeID
}

func newNodeIDConnector(nodeID ids.NodeID) *nodeIDConnector { _ = "STUB: not implemented"; return nil }

func (f *nodeIDConnector) IsAllowed(nodeID ids.NodeID, _ bool) bool {
	_ = "STUB: not implemented"
	return false
}
