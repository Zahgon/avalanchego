// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package timer

import (
	"time"
)

// ProgressFromHash returns the progress out of MaxUint64 assuming [b] is a key
// in a uniformly distributed sequence that is being iterated lexicographically.
func ProgressFromHash(b []byte) uint64 {
	_ = "STUB: not implemented"
	// binary.BigEndian.Uint64 will panic if the input length is less than 8, so
	// pad 0s as needed.
	return 0
}

// A sample represents a completed amount and the timestamp of the sample
type sample struct {
	completed uint64
	timestamp time.Time
}

// EtaTracker tracks the ETA of a job
type EtaTracker struct {
	samples        []sample
	samplePosition uint8
	totalSamples   uint64
	slowdownFactor float64
}

// NewEtaTracker creates a new EtaTracker with the given maximum number of samples
// and a slowdown factor. The slowdown factor is a multiplier that is added to the ETA
// based on the percentage completed.
//
// The adjustment works as follows:
//   - At 0% progress: ETA is multiplied by slowdownFactor
//   - At 100% progress: ETA is the raw estimate (no adjustment)
//   - Between 0% and 100%: Adjustment decreases linearly with progress
//
// Example: With slowdownFactor = 2.0:
//   - At 0% progress: ETA = raw_estimate * 2.0
//   - At 50% progress: ETA = raw_estimate * 1.5
//   - At 100% progress: ETA = raw_estimate * 1.0
//
// If maxSamples is less than 1, it will default to 5
func NewEtaTracker(maxSamples uint8, slowdownFactor float64) *EtaTracker {
	_ = "STUB: not implemented"
	return nil
}

// AddSample adds a sample to the EtaTracker
// It returns the remaining time to complete the target and the percent complete
// The returned values are rounded to the nearest second and 2 decimal places respectively
// This function can return a nil time.Duration indicating that there are not yet enough
// samples to calculate an accurate ETA.
//
// The first sample should be at 0% progress to establish a baseline
func (t *EtaTracker) AddSample(completed uint64, target uint64, timestamp time.Time) (*time.Duration, float64) {
	_ = "STUB: not implemented"
	return nil, 0
}

// save the oldest sample; this will not be used if we don't have enough samples

// If we don't have enough samples, return nil

// Calculate the time and progress since the oldest sample

// Check if target is already completed or exceeded

// scale to 0.00 to 100.00

// EstimateETA calculates ETA from start time and current progress.
//
// Deprecated: use EtaTracker instead
func EstimateETA(startTime time.Time, progress, end uint64) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
