// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package code

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/libevm/options"

	"github.com/ava-labs/avalanchego/graft/evm/sync/types"
)

const defaultQueueCapacity = 5000

var (
	_ types.Finalizer = (*Queue)(nil)

	ErrQueueClosed = errors.New("code queue is closed")
)

// Queue is a fan-in/fan-out bridge between code hash producers (leaf sync workers)
// and the code syncer consumer. Producers call [Queue.AddCode] which persists durable
// disk markers and appends hashes to an internal queue. A single background goroutine
// forwards them to the output channel. [Queue.AddCode] never blocks the caller.
//
// Deduplication and local-code checks are the consumer's responsibility.
type Queue struct {
	db  ethdb.Database
	out chan common.Hash // output to consumer

	cancel      context.CancelFunc
	done        <-chan struct{} // cancelled on Shutdown
	forwardDone chan struct{}   // closed when forward() exits

	closeMu     sync.RWMutex
	closeInOnce sync.Once
	closed      bool

	pendingMu sync.Mutex
	pending   []common.Hash
	in        chan struct{} // producer signal, buffered to 1

	capacity int
}

type QueueOption = options.Option[Queue]

// WithCapacity overrides the queue buffer capacity.
func WithCapacity(n int) QueueOption { _ = "STUB: not implemented"; return *new(QueueOption) }

// NewQueue creates a code queue. Call [Queue.Finalize] for normal completion
// or [Queue.Shutdown] for cancellation. Both are safe to call in any order.
func NewQueue(db ethdb.Database, opts ...QueueOption) (*Queue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CodeHashes returns the receive-only channel consumed by the code syncer.
func (q *Queue) CodeHashes() <-chan common.Hash {
	_ = "STUB: not implemented"

	// AddCode persists code hashes as durable disk markers and enqueues them
	// for the forwarder goroutine. Never blocks the caller.
	// Returns [ErrQueueClosed] after [Queue.Shutdown] or [Queue.Finalize].
	return nil
}

func (q *Queue) AddCode(ctx context.Context, codeHashes []common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// Persist to-fetch markers keyed by code hash (idempotent overwrites).
// The consumer deletes markers after fetching or if code is already present.

// Signal coalescing: skip if the forwarder is already notified.

// Finalize waits for all pending hashes to be sent, then closes out.
// Blocks if no consumer is draining [Queue.CodeHashes]. Idempotent with [Queue.Shutdown].
func (q *Queue) Finalize() error { _ = "STUB: not implemented"; return nil }

// Shutdown cancels the forwarder, waits for exit, then closes out.
// Unsent hashes are safe as disk markers and will be recovered on restart.
// Idempotent with [Queue.Finalize].
func (q *Queue) Shutdown() { _ = "STUB: not implemented"; return }

func (q *Queue) markClosed() { _ = "STUB: not implemented"; return }

// stop waits for in-flight AddCode calls (via write lock), optionally cancels
// the forwarder, signals no more work, and waits for the forwarder to exit.
func (q *Queue) stop(shouldCancel bool) { _ = "STUB: not implemented"; return }

// forward moves hashes from pending to `q.out`. It owns `q.out` and closes it on exit.
func (q *Queue) forward() { _ = "STUB: not implemented"; return }

// drainPending sends all accumulated pending hashes to out.
// Returns true if cancelled via done.
func (q *Queue) drainPending() bool { _ = "STUB: not implemented"; return false }

// init recovers persisted code markers from disk and re-enqueues them.
// AddCode will re-persist the same markers, which is a harmless redundancy
// that only happens on resume after restart.
func (q *Queue) init() error {
	dbCodeHashes, err := recoverUnfetchedCodeHashes(q.db)
	if err != nil {
		return fmt.Errorf("unable to recover previous sync state: %w", err)
	}

	// context.Background: init runs during construction before sync starts,
	// the queue is not closed yet so AddCode will always succeed.
	if err := q.AddCode(context.Background(), dbCodeHashes); err != nil {
		return fmt.Errorf("unable to resume previous sync: %w", err)
	}

	return nil
}

// recoverUnfetchedCodeHashes returns persisted code markers that still need fetching
// and deletes markers for code already present locally.
func recoverUnfetchedCodeHashes(db ethdb.Database) ([]common.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
