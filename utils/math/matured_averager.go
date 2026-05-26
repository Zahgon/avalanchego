// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package math

import "time"

// maturedAverager wraps an Averager and returns zero until the averager has
// been running for at least the half-life duration.
type maturedAverager struct {
	Averager
	halfLife     time.Duration
	startTime    time.Time
	lastObserved time.Time
	hasObserved  bool
}

var _ Averager = (*maturedAverager)(nil)

// NewMaturedAverager creates a new matured averager that wraps the provided
// `averager`. It returns zero from [Averager.Read] until at least `halfLife` duration has
// passed since the first [Averager.Observe] call.
func NewMaturedAverager(halfLife time.Duration, averager Averager) Averager {
	_ = "STUB: not implemented"
	return *new(Averager)
}

func (a *maturedAverager) Observe(value float64, currentTime time.Time) {
	_ = "STUB: not implemented"
	return
}

func (a *maturedAverager) Read() float64 {
	_ = "STUB: not implemented"
	// If we haven't observed anything yet, return zero
	return 0
}

// Check if enough time has passed since the first observation
