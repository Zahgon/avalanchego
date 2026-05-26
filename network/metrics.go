// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/peer"
	"github.com/ava-labs/avalanchego/utils/set"
)

type metrics struct {
	// trackedSubnets does not include the primary network ID
	trackedSubnets set.Set[ids.ID]

	numTracked                   prometheus.Gauge
	numPeers                     prometheus.Gauge
	numSubnetPeers               *prometheus.GaugeVec
	timeSinceLastMsgSent         prometheus.Gauge
	timeSinceLastMsgReceived     prometheus.Gauge
	sendFailRate                 prometheus.Gauge
	connected                    prometheus.Counter
	disconnected                 prometheus.Counter
	acceptFailed                 prometheus.Counter
	inboundConnRateLimited       prometheus.Counter
	inboundConnAllowed           prometheus.Counter
	tlsConnRejected              prometheus.Counter
	numUselessPeerListBytes      prometheus.Counter
	nodeUptimeWeightedAverage    prometheus.Gauge
	nodeUptimeRewardingStake     prometheus.Gauge
	peerConnectedLifetimeAverage prometheus.Gauge
	lock                         sync.RWMutex
	peerConnectedStartTimes      map[ids.NodeID]float64
	peerConnectedStartTimesSum   float64
}

func newMetrics(
	registerer prometheus.Registerer,
	trackedSubnets set.Set[ids.ID],
) (*metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// init subnet tracker metrics with tracked subnets

// initialize to 0

func (m *metrics) markConnected(peer *peer.Peer) { _ = "STUB: not implemented"; return }

func (m *metrics) markDisconnected(peer *peer.Peer) { _ = "STUB: not implemented"; return }

func (m *metrics) updatePeerConnectionLifetimeMetrics() { _ = "STUB: not implemented"; return }
