// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package math

import (
	"sync"
	"time"
)

type syncAverager struct {
	lock     sync.RWMutex
	averager Averager
}

func NewSyncAverager(averager Averager) Averager { _ = "STUB: not implemented"; return *new(Averager) }

func (a *syncAverager) Observe(value float64, currentTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func (a *syncAverager) Read() float64 { _ = "STUB: not implemented"; return 0 }
