// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package timer

import (
	"sync"
	"time"
)

// Timer wraps a timer object. This allows a user to specify a handler. Once
// specifying the handler, the dispatch thread can be called. The dispatcher
// will only return after calling Stop. SetTimeoutIn will result in calling the
// handler in the specified amount of time.
type Timer struct {
	handler func()
	timeout chan struct{}

	lock                    sync.Mutex
	wg                      sync.WaitGroup
	finished, shouldExecute bool
	duration                time.Duration
}

// NewTimer creates a new timer object
func NewTimer(handler func()) *Timer { _ = "STUB: not implemented"; return nil }

// SetTimeoutIn will set the timer to fire the handler in [duration]
func (t *Timer) SetTimeoutIn(duration time.Duration) { _ = "STUB: not implemented"; return }

// Cancel the currently scheduled event
func (t *Timer) Cancel() { _ = "STUB: not implemented"; return }

// Stop this timer from executing any more.
func (t *Timer) Stop() { _ = "STUB: not implemented"; return }

func (t *Timer) Dispatch() { _ = "STUB: not implemented"; return }

// t.finished needs to be thread safe

func (t *Timer) reset() { _ = "STUB: not implemented"; return }
