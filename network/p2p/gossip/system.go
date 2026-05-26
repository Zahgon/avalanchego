// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gossip

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/utils/logging"
)

const defaultRequestPeriod = time.Second

// SystemConfig provides all the configurations needed to create a gossip
// system.
//
// All fields are optional.
type SystemConfig struct {
	Log       logging.Logger
	Registry  prometheus.Registerer
	Namespace string

	HandlerID uint64 // Defaults to [p2p.TxGossipHandlerID]

	TargetMessageSize int // Defaults to 20 KiB

	ThrottlingPeriod time.Duration // Defaults to one hour
	RequestPeriod    time.Duration // Defaults to one request per second

	PushGossipParams   BranchingFactor // Defaults to 100 validators and top 90% of stake
	PushRegossipParams BranchingFactor // Defaults to 10 validators

	DiscardedPushCacheSize int           // Defaults to 16,384
	RegossipPeriod         time.Duration // Defaults to 30 seconds
}

func (c *SystemConfig) setDefaults() { _ = "STUB: not implemented"; return }

// SystemSet is the backend interface required to construct a gossip system.
type SystemSet[T Gossipable] interface {
	HandlerSet[T]
	PullGossiperSet[T]
	PushGossiperSet
}

// NewSystem is a helper to construct the senders and receivers for a gossip
// protocol.
func NewSystem[T Gossipable](
	nodeID ids.NodeID,
	network *p2p.Network,
	validatorPeers *p2p.Validators,
	set SystemSet[T],
	marshaller Marshaller[T],
	c SystemConfig,
) (
	p2p.Handler,
	*ValidatorGossiper,
	*PushGossiper[T],
	error,
) {
	_ = "STUB: not implemented"
	return *new(p2p.Handler), nil, nil, nil
}

// Pull requests are filtered by validators and are throttled to prevent
// spamming. Push messages are not filtered.
