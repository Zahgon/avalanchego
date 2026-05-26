// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"sync"
	"sync/atomic"
)

type BoundedWorkers struct {
	workerCount        atomic.Int32
	workerSpawner      chan struct{}
	outstandingWorkers sync.WaitGroup

	work      chan func()
	workClose sync.Once
}

// NewBoundedWorkers returns an instance of [BoundedWorkers] that
// will spawn up to count goroutines.
func NewBoundedWorkers(count int) *BoundedWorkers { _ = "STUB: not implemented"; return nil }

// startWorker creates a new goroutine to execute [f] immediately and then keeps the goroutine
// alive to continue executing new work.
func (b *BoundedWorkers) startWorker(f func()) { _ = "STUB: not implemented"; return }

// Execute the given function on an existing goroutine waiting for more work, a new goroutine,
// or return if the context is canceled.
//
// Execute must not be called after Wait, otherwise it might panic.
func (b *BoundedWorkers) Execute(f func()) {
	_ = "STUB: not implemented"
	// Ensure we feed idle workers first
	return
}

// Fallback to waiting for an idle worker or allocating
// a new worker (if we aren't yet at max concurrency)

// Wait returns after all enqueued work finishes and all goroutines to exit.
// Wait returns the number of workers that were spawned during the run.
//
// Wait can only be called after ALL calls to [Execute] have returned.
//
// It is safe to call Wait multiple times but not safe to call [Execute]
// after [Wait] has been called.
func (b *BoundedWorkers) Wait() int { _ = "STUB: not implemented"; return 0 }
