// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package window

import (
	"sync"
	"time"

	"github.com/ava-labs/avalanchego/utils/buffer"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
)

var _ Window[struct{}] = (*window[struct{}])(nil)

// Window is an interface which represents a sliding window of elements.
type Window[T any] interface {
	Add(value T)
	Oldest() (T, bool)
	Length() int
}

type window[T any] struct {
	// mocked clock for unit testing
	clock *mockable.Clock
	// time-to-live for elements in the window
	ttl time.Duration
	// max amount of elements allowed in the window
	maxSize int
	// min amount of elements required in the window before allowing removal
	// based on time
	minSize int

	// mutex for synchronization
	lock sync.Mutex
	// elements in the window
	elements buffer.Deque[node[T]]
}

// Config exposes parameters for Window
type Config struct {
	Clock   *mockable.Clock
	MaxSize int
	MinSize int
	TTL     time.Duration
}

// New returns an instance of window
func New[T any](config Config) Window[T] { _ = "STUB: not implemented"; return nil }

// Add adds an element to a window and also evicts any elements if they've been
// present in the window beyond the configured time-to-live
func (w *window[T]) Add(value T) { _ = "STUB: not implemented"; return }

// add the new block id

// Oldest returns the oldest element in the window.
func (w *window[T]) Oldest() (T, bool) { _ = "STUB: not implemented"; return *new(T), false }

// Length returns the number of elements in the window.
func (w *window[T]) Length() int { _ = "STUB: not implemented"; return 0 }

// removeStaleNodes removes any nodes beyond the configured ttl of a window node.
func (w *window[T]) removeStaleNodes() {
	_ = "STUB: not implemented"
	// If we're beyond the expiry threshold, removeStaleNodes this node from our
	// window. Nodes are guaranteed to be strictly increasing in entry time,
	// so we can break this loop once we find the first non-stale one.
	return
}

// helper struct to represent elements in the window
type node[T any] struct {
	value     T
	entryTime time.Time
}
