// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package uptimetracker

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/uptime"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
)

var ErrValidationIDNotFound = errors.New("validationID not found")

// UptimeTracker tracks uptime information for validators
type UptimeTracker struct {
	validatorState validators.State
	subnetID       ids.ID
	manager        uptime.Manager

	lock                sync.Mutex
	height              uint64
	state               *state
	synced              bool
	connectedValidators set.Set[ids.NodeID]
	// Deactivated validators are treated as being offline
	deactivatedValidators set.Set[ids.NodeID]
}

// New returns a new instance of UptimeTracker
func New(
	validatorState validators.State,
	subnetID ids.ID,
	db database.Database,
	clock *mockable.Clock,
) (*UptimeTracker, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetUptime returns the uptime of the validator corresponding to validationID
func (u *UptimeTracker) GetUptime(validationID ids.ID) (
	time.Duration,
	time.Time,
	error,
) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Time), nil
}

// Connect starts tracking a node. Nodes that are activated and connected will
// be treated as online.
func (u *UptimeTracker) Connect(nodeID ids.NodeID) error { _ = "STUB: not implemented"; return nil }

// Disconnect stops tracking a node. Disconnected nodes are treated as being
// offline.
func (u *UptimeTracker) Disconnect(nodeID ids.NodeID) error { _ = "STUB: not implemented"; return nil }

// Sync updates the validator set and writes our state. Sync starts tracking
// uptimes for all active validators the first time it is called.
func (u *UptimeTracker) Sync(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Update validator set if we're behind

// Initialize uptimes if this is the first time Sync has been called

func (u *UptimeTracker) update(
	height uint64,
	currentValidatorSet map[ids.ID]*validators.GetCurrentValidatorOutput,
) error {
	_ = "STUB: not implemented"
	// We are behind and need to update our local state
	return nil
}

// This validator is still in the latest update

// Add or update validators

// Check if there was a status change

// If there was a status change we need to activate/deactivate the
// validator

// This validator is now active and is treated as online

// This validator is no longer active and is treated as offline

// This is a new validator

// This validator is not active and is treated is offline

// activate treats nodeID as online
func (u *UptimeTracker) activate(nodeID ids.NodeID) error { _ = "STUB: not implemented"; return nil }

// deactivate treats nodeID as offline
func (u *UptimeTracker) deactivate(nodeID ids.NodeID) error { _ = "STUB: not implemented"; return nil }

// Shutdown stops tracking uptimes and writes our state.
func (u *UptimeTracker) Shutdown() error { _ = "STUB: not implemented"; return nil }
