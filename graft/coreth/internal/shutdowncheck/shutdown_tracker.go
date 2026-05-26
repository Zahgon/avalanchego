// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package shutdowncheck

import (
	"github.com/ava-labs/libevm/ethdb"
)

// ShutdownTracker is a service that reports previous unclean shutdowns
// upon start. It needs to be started after a successful start-up and stopped
// after a successful shutdown, just before the db is closed.
type ShutdownTracker struct {
	db     ethdb.Database
	stopCh chan struct{}
}

// NewShutdownTracker creates a new ShutdownTracker instance and has
// no other side-effect.
func NewShutdownTracker(db ethdb.Database) *ShutdownTracker { _ = "STUB: not implemented"; return nil }

// MarkStartup is to be called in the beginning when the node starts. It will:
// - Push a new startup marker to the db
// - Report previous unclean shutdowns
func (t *ShutdownTracker) MarkStartup() { _ = "STUB: not implemented"; return }

// Start runs an event loop that updates the current marker's timestamp every 5 minutes.
func (t *ShutdownTracker) Start() { _ = "STUB: not implemented"; return }

// Stop will stop the update loop and clear the current marker.
func (t *ShutdownTracker) Stop() {
	_ = "STUB: not implemented"
	// Stop update loop.
	return
}

// Clear last marker.
