// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package meter

import (
	"time"
)

var (
	_ Factory = (*ContinuousFactory)(nil)
	_ Meter   = (*continuousMeter)(nil)
)

// ContinuousFactory implements the Factory interface by returning a continuous
// time meter.
type ContinuousFactory struct{}

func (ContinuousFactory) New(halflife time.Duration) Meter {
	_ = "STUB: not implemented"
	return *new(Meter)
}

type continuousMeter struct {
	halflife float64
	value    float64

	numCoresRunning float64
	lastUpdated     time.Time
}

// NewMeter returns a new Meter with the provided halflife
func NewMeter(halflife time.Duration) Meter { _ = "STUB: not implemented"; return *new(Meter) }

func (a *continuousMeter) Inc(now time.Time, numCores float64) { _ = "STUB: not implemented"; return }

func (a *continuousMeter) Dec(now time.Time, numCores float64) { _ = "STUB: not implemented"; return }

func (a *continuousMeter) Read(now time.Time) float64 { _ = "STUB: not implemented"; return 0 }

func (a *continuousMeter) TimeUntil(now time.Time, value float64) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Note that [factor] >= 1

// Note that [numHalfLives] >= 0

// Overflow protection
