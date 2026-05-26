// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package throttling

import (
	"context"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/time/rate"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/metric"
)

var _ bandwidthThrottler = (*bandwidthThrottlerImpl)(nil)

// Returns a bandwidth throttler that uses a token bucket
// model, where each token is 1 byte, to rate-limit bandwidth usage.
// See https://pkg.go.dev/golang.org/x/time/rate#Limiter
type bandwidthThrottler interface {
	// Blocks until [nodeID] can read a message of size [msgSize].
	// AddNode([nodeID], ...) must have been called since
	// the last time RemoveNode([nodeID]) was called, if any.
	// It's safe for multiple goroutines to concurrently call Acquire.
	// Returns immediately if [ctx] is canceled.
	Acquire(ctx context.Context, msgSize uint64, nodeID ids.NodeID)

	// Add a new node to this throttler.
	// Must be called before Acquire(..., [nodeID]) is called.
	// RemoveNode([nodeID]) must have been called since the last time
	// AddNode([nodeID], ...) was called, if any.
	// Its bandwidth allocation refills at a rate of [refillRate].
	// Its bandwidth allocation can hold up to [maxBurstSize] at a time.
	// [maxBurstSize] must be at least the maximum message size.
	// It's safe for multiple goroutines to concurrently call AddNode.
	AddNode(nodeID ids.NodeID)

	// Remove a node from this throttler.
	// AddNode([nodeID], ...) must have been called since
	// the last time RemoveNode([nodeID]) was called, if any.
	// Must be called when we stop reading messages from [nodeID].
	// It's safe for multiple goroutines to concurrently call RemoveNode.
	RemoveNode(nodeID ids.NodeID)
}

type BandwidthThrottlerConfig struct {
	// Rate at which the inbound bandwidth consumable by a peer replenishes
	RefillRate uint64 `json:"bandwidthRefillRate"`
	// Max amount of consumable bandwidth that can accumulate for a given peer
	MaxBurstSize uint64 `json:"bandwidthMaxBurstRate"`
}

func newBandwidthThrottler(
	log logging.Logger,
	registerer prometheus.Registerer,
	config BandwidthThrottlerConfig,
) (bandwidthThrottler, error) {
	_ = "STUB: not implemented"
	return *new(bandwidthThrottler), nil
}

type bandwidthThrottlerMetrics struct {
	acquireLatency  metric.Averager
	awaitingAcquire prometheus.Gauge
}

type bandwidthThrottlerImpl struct {
	BandwidthThrottlerConfig
	metrics bandwidthThrottlerMetrics
	log     logging.Logger
	lock    sync.RWMutex
	// Node ID --> token bucket based rate limiter where each token
	// is a byte of bandwidth.
	limiters map[ids.NodeID]*rate.Limiter
}

// See BandwidthThrottler.
func (t *bandwidthThrottlerImpl) Acquire(
	ctx context.Context,
	msgSize uint64,
	nodeID ids.NodeID,
) {
	_ = "STUB: not implemented"
	return
}

// This should never happen. If it is, the caller is misusing this struct.

// This should only happen on shutdown.

// See BandwidthThrottler.
func (t *bandwidthThrottlerImpl) AddNode(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

// See BandwidthThrottler.
func (t *bandwidthThrottlerImpl) RemoveNode(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }
