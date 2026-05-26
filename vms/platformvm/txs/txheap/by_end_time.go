// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txheap

import (
	"time"
)

var _ TimedHeap = (*byEndTime)(nil)

type TimedHeap interface {
	Heap

	Timestamp() time.Time
}

type byEndTime struct {
	txHeap
}

func NewByEndTime() TimedHeap { _ = "STUB: not implemented"; return *new(TimedHeap) }

func (h *byEndTime) Timestamp() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }
