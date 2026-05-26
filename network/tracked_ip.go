// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"net/netip"
	"sync"
	"time"
)

type trackedIP struct {
	delayLock sync.RWMutex
	delay     time.Duration

	ip netip.AddrPort

	stopTrackingOnce sync.Once
	onStopTracking   chan struct{}
}

func newTrackedIP(ip netip.AddrPort) *trackedIP { _ = "STUB: not implemented"; return nil }

func (ip *trackedIP) trackNewIP(newIP netip.AddrPort) *trackedIP {
	_ = "STUB: not implemented"
	return nil
}

func (ip *trackedIP) getDelay() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (ip *trackedIP) increaseDelay(initialDelay, maxDelay time.Duration) {
	_ = "STUB: not implemented"
	return
}

// If the timeout was previously 0, ensure that there is a reasonable delay.

// Randomization is only performed here to distribute reconnection
// attempts to a node that previously shut down. This doesn't
// require cryptographically secure random number generation.
// set the timeout to [1, 2) * timeout
// #nosec G404

// set the timeout to [.75, 1) * maxDelay
// #nosec G404

func (ip *trackedIP) stopTracking() { _ = "STUB: not implemented"; return }
