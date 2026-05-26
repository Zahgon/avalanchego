// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package throttling

import (
	"context"
	"net"

	"golang.org/x/time/rate"
)

var _ net.Listener = (*throttledListener)(nil)

// Wraps [listener] and returns a net.Listener that will accept at most
// [maxConnsPerSec] connections per second.
// [maxConnsPerSec] must be non-negative.
func NewThrottledListener(listener net.Listener, maxConnsPerSec float64) net.Listener {
	_ = "STUB: not implemented"
	return *new(net.Listener)
}

// [throttledListener] is a net.Listener that rate-limits
// acceptance of incoming connections.
// Note that InboundConnUpgradeThrottler rate-limits _upgrading_ of
// inbound connections, whereas throttledListener rate-limits
// _acceptance_ of inbound connections.
type throttledListener struct {
	// [ctx] is cancelled when Close() is called
	ctx context.Context
	// [ctxCancelFunc] cancels [ctx] when it's called
	ctxCancelFunc func()
	// The underlying listener
	listener net.Listener
	// Handles rate-limiting
	limiter *rate.Limiter
}

func (l *throttledListener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	// Wait until the rate-limiter says to accept the
	// next incoming connection. If l.Close() is called,
	// Wait will return immediately.
	return *new(net.Conn), nil
}

func (l *throttledListener) Close() error {
	_ = "STUB: not implemented"
	// Cancel [l.ctx] so Accept() will return immediately
	return nil
}

func (l *throttledListener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }
