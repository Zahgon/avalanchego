// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package math

import (
	"time"
)

type continuousAverager struct {
	halflife    float64
	weightedSum float64
	normalizer  float64
	lastUpdated time.Time
}

// NewUninitializedAverager creates a new averager with the given halflife. If
// [Read] is called before [Observe], the zero value will be returned. When
// [Observe] is called the first time, the averager will be initialized with
// [value] at that time.
func NewUninitializedAverager(halfLife time.Duration) Averager {
	_ = "STUB: not implemented"
	// Use 0 as the initialPrediction and 0 as the currentTime, so that when the
	// first observation occurs (at a non-zero time) the initial prediction's
	// weight will become negligible.
	return *new(Averager)
}

func NewAverager(
	initialPrediction float64,
	halflife time.Duration,
	currentTime time.Time,
) Averager {
	_ = "STUB: not implemented"
	return *new(Averager)
}

func (a *continuousAverager) Observe(value float64, currentTime time.Time) {
	_ = "STUB: not implemented"
	return
}

// If the times are called in order, scale the previous values to keep the
// sizes manageable

// If this is called multiple times at the same wall clock time, no
// scaling needs to occur

// If the times are called out of order, don't scale the previous values

func (a *continuousAverager) Read() float64 { _ = "STUB: not implemented"; return 0 }
