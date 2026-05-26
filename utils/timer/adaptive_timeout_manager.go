// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package timer

import (
	"errors"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/heap"
	"github.com/ava-labs/avalanchego/utils/math"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
)

var (
	errNonPositiveHalflife        = errors.New("timeout halflife must be positive")
	errInitialTimeoutAboveMaximum = errors.New("initial timeout cannot be greater than maximum timeout")
	errInitialTimeoutBelowMinimum = errors.New("initial timeout cannot be less than minimum timeout")
	errTooSmallTimeoutCoefficient = errors.New("timeout coefficient must be >= 1")

	_ AdaptiveTimeoutManager = (*adaptiveTimeoutManager)(nil)
)

type adaptiveTimeout struct {
	id             ids.RequestID // Unique ID of this timeout
	handler        func()        // Function to execute if timed out
	duration       time.Duration // How long this timeout was set for
	deadline       time.Time     // When this timeout should be fired
	measureLatency bool          // Whether this request should impact latency
}

// AdaptiveTimeoutConfig contains the parameters provided to the
// adaptive timeout manager.
type AdaptiveTimeoutConfig struct {
	InitialTimeout time.Duration `json:"initialTimeout"`
	MinimumTimeout time.Duration `json:"minimumTimeout"`
	MaximumTimeout time.Duration `json:"maximumTimeout"`
	// Timeout is [timeoutCoefficient] * average response time
	// [timeoutCoefficient] must be > 1
	TimeoutCoefficient float64 `json:"timeoutCoefficient"`
	// Larger halflife --> less volatile timeout
	// [timeoutHalfLife] must be positive
	TimeoutHalflife time.Duration `json:"timeoutHalflife"`
}

type AdaptiveTimeoutManager interface {
	// Start the timeout manager.
	// Must be called before any other method.
	// Must only be called once.
	Dispatch()
	// Stop the timeout manager.
	// Must only be called once.
	Stop()
	// Returns the current network timeout duration.
	TimeoutDuration() time.Duration
	// Registers a timeout for the item with the given [id].
	// If the timeout occurs before the item is Removed, [timeoutHandler] is called.
	Put(id ids.RequestID, measureLatency bool, timeoutHandler func())
	// Remove the timeout associated with [id].
	// Its timeout handler will not be called.
	Remove(id ids.RequestID)
	// ObserveLatency manually registers a response latency.
	// We use this to pretend that it a query to a benched validator
	// timed out when actually, we never even sent them a request.
	ObserveLatency(latency time.Duration)
}

type adaptiveTimeoutManager struct {
	lock sync.Mutex
	// Tells the time. Can be faked for testing.
	clock                            mockable.Clock
	networkTimeoutMetric, avgLatency prometheus.Gauge
	numTimeouts                      prometheus.Counter
	numPendingTimeouts               prometheus.Gauge
	// Averages the response time from all peers
	averager math.Averager
	// Timeout is [timeoutCoefficient] * average response time
	// [timeoutCoefficient] must be > 1
	timeoutCoefficient float64
	minimumTimeout     time.Duration
	maximumTimeout     time.Duration
	currentTimeout     time.Duration // Amount of time before a timeout
	timeoutHeap        heap.Map[ids.RequestID, *adaptiveTimeout]
	timer              *Timer // Timer that will fire to clear the timeouts
}

func NewAdaptiveTimeoutManager(
	config *AdaptiveTimeoutConfig,
	reg prometheus.Registerer,
) (AdaptiveTimeoutManager, error) {
	_ = "STUB: not implemented"
	return *new(AdaptiveTimeoutManager), nil
}

func (tm *adaptiveTimeoutManager) TimeoutDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (tm *adaptiveTimeoutManager) Dispatch() { _ = "STUB: not implemented"; return }

func (tm *adaptiveTimeoutManager) Stop() { _ = "STUB: not implemented"; return }

func (tm *adaptiveTimeoutManager) Put(id ids.RequestID, measureLatency bool, timeoutHandler func()) {
	_ = "STUB: not implemented"
	return
}

// Assumes [tm.lock] is held
func (tm *adaptiveTimeoutManager) put(id ids.RequestID, measureLatency bool, handler func()) {
	_ = "STUB: not implemented"
	return
}

func (tm *adaptiveTimeoutManager) Remove(id ids.RequestID) { _ = "STUB: not implemented"; return }

// Assumes [tm.lock] is held
func (tm *adaptiveTimeoutManager) remove(id ids.RequestID, now time.Time) {
	_ = "STUB: not implemented"
	// Observe the response time to update average network response time.
	return
}

// Assumes [tm.lock] is not held.
func (tm *adaptiveTimeoutManager) timeout() { _ = "STUB: not implemented"; return }

// getNextTimeoutHandler returns nil once there is nothing left to remove

// Don't execute a callback with a lock held

func (tm *adaptiveTimeoutManager) ObserveLatency(latency time.Duration) {
	_ = "STUB: not implemented"
	return
}

// Assumes [tm.lock] is held
func (tm *adaptiveTimeoutManager) observeLatencyAndUpdateTimeout(latency time.Duration, now time.Time) {
	_ = "STUB: not implemented"
	return
}

// Update the metrics

// Returns the handler function associated with the next timeout.
// If there are no timeouts, or if the next timeout is after [now],
// returns nil.
// Assumes [tm.lock] is held
func (tm *adaptiveTimeoutManager) getNextTimeoutHandler(now time.Time) func() {
	_ = "STUB: not implemented"
	return nil
}

// Calculate the time of the next timeout and set
// the timer to fire at that time.
func (tm *adaptiveTimeoutManager) setNextTimeoutTime() { _ = "STUB: not implemented"; return }

// There are no pending timeouts
