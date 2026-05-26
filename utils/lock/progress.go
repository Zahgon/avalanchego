// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package lock

import (
	"cmp"
	"context"
	"sync"
)

// ProgressSubscription allows a caller to subscribe to progress updates.
type ProgressSubscription[T cmp.Ordered] struct {
	signal   *Cond
	lock     sync.Mutex
	progress T
}

// NewProgressSubscription returns a new ProgressSubscription with the given initial progress.
func NewProgressSubscription[T cmp.Ordered](initialProgress T) *ProgressSubscription[T] {
	_ = "STUB: not implemented"
	return nil
}

// SetProgress updates the progress of this ProgressSubscription to the given value.
// This will unblock any calls to WaitForProgress that are waiting for a progress value above the given value.
// It is assumed that progress values are monotonically increasing.
func (ps *ProgressSubscription[T]) SetProgress(progress T) { _ = "STUB: not implemented"; return }

// WaitForProgress blocks until the progress of this ProgressSubscription is above the given value,
// or until the given context is cancelled.
func (ps *ProgressSubscription[T]) WaitForProgress(ctx context.Context, pos T) error {
	_ = "STUB: not implemented"
	return nil
}
