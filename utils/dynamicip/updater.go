// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package dynamicip

import (
	"context"
	"net/netip"
	"time"

	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/logging"
)

const ipResolutionTimeout = 10 * time.Second

var _ Updater = (*updater)(nil)

// Updater periodically updates this node's public IP.
// Dispatch() and Stop() should only be called once.
type Updater interface {
	// Start periodically resolving and updating our public IP.
	// Doesn't return until after Stop() is called.
	// Should be called in a goroutine.
	Dispatch(log logging.Logger)
	// Stop resolving and updating our public IP.
	Stop()
}

type updater struct {
	// The IP we periodically modify.
	dynamicIP *utils.Atomic[netip.AddrPort]
	// Used to find out what our public IP is.
	resolver Resolver
	// The parent of all contexts passed into resolver.Resolve().
	// Cancelling causes Dispatch() to eventually return.
	rootCtx context.Context
	// Cancelling causes Dispatch() to eventually return.
	// All in-flight calls to resolver.Resolve() will be cancelled.
	rootCtxCancel context.CancelFunc
	// Closed when Dispatch() has returned.
	doneChan chan struct{}
	// How often we update the public IP.
	updateFreq time.Duration
}

// Returns a new Updater that updates [dynamicIP]
// every [updateFreq]. Uses [resolver] to find
// out what our public IP is.
func NewUpdater(
	dynamicIP *utils.Atomic[netip.AddrPort],
	resolver Resolver,
	updateFreq time.Duration,
) Updater {
	_ = "STUB: not implemented"
	return *new(Updater)
}

// Start updating [u.dynamicIP] every [u.updateFreq].
// Stops when [dynamicIP.stopChan] is closed.
func (u *updater) Dispatch(log logging.Logger) { _ = "STUB: not implemented"; return }

func (u *updater) Stop() {
	_ = "STUB: not implemented"
	// Cause Dispatch() to return and cancel all
	// in-flight calls to resolver.Resolve().
	return
}

// Wait until Dispatch() has returned.
