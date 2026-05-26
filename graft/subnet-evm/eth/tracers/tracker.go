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
// Copyright 2022 The go-ethereum Authors
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

package tracers

import (
	"sync"
)

// stateTracker is an auxiliary tool used to cache the release functions of all
// used trace states, and to determine whether the creation of trace state needs
// to be paused in case there are too many states waiting for tracing.
type stateTracker struct {
	limit    int                // Maximum number of states allowed waiting for tracing
	oldest   uint64             // The number of the oldest state which is still using for trace
	used     []bool             // List of flags indicating whether the trace state has been used up
	releases []StateReleaseFunc // List of trace state release functions waiting to be called
	cond     *sync.Cond
	lock     *sync.RWMutex
}

// newStateTracker initializes the tracker with provided state limits and
// the number of the first state that will be used.
func newStateTracker(limit int, oldest uint64) *stateTracker { _ = "STUB: not implemented"; return nil }

// releaseState marks the state specified by the number as released and caches
// the corresponding release functions internally.
func (t *stateTracker) releaseState(number uint64, release StateReleaseFunc) {
	_ = "STUB: not implemented"
	return
}

// Set the state as used, the corresponding flag is indexed by
// the distance between the specified state and the oldest state
// which is still using for trace.

// If the oldest state is used up, update the oldest marker by moving
// it to the next state which is not used up.

// Clean up the array tail since they are useless now.

// Fire the signal to all waiters that oldest marker is updated.

// callReleases invokes all cached release functions.
func (t *stateTracker) callReleases() { _ = "STUB: not implemented"; return }

// wait blocks until the accumulated trace states are less than the limit.
func (t *stateTracker) wait(number uint64) error { _ = "STUB: not implemented"; return nil }

// number is now within limit, wait over
