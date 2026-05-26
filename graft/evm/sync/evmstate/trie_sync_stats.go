// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evmstate

import (
	"sync"
	"time"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/metrics"

	safemath "github.com/ava-labs/avalanchego/utils/math"
)

const (
	updateFrequency  = 1 * time.Minute
	leafRateHalfLife = 1 * time.Minute
	epsilon          = 1e-6 // added to avoid division by 0
)

// trieSyncStats keeps track of the total number of leafs and tries
// completed during a sync.
type trieSyncStats struct {
	lock sync.Mutex

	lastUpdated time.Time
	leafsRate   safemath.Averager

	triesRemaining   int
	triesSynced      int
	triesStartTime   time.Time
	leafsSinceUpdate uint64

	remainingLeafs map[*trieSegment]uint64

	// metrics
	totalLeafs     metrics.Counter
	triesSegmented metrics.Counter
	leafsRateGauge metrics.Gauge
}

func newTrieSyncStats() *trieSyncStats { _ = "STUB: not implemented"; return nil }

// metrics

// incTriesSegmented increases the metric for segmented tries.
func (t *trieSyncStats) incTriesSegmented() { _ = "STUB: not implemented"; return }

// safe to be called concurrently

// incLeafs takes a lock and adds [count] to the total number of leafs synced.
// periodically outputs a log message with the number of leafs and tries.
func (t *trieSyncStats) incLeafs(segment *trieSegment, count uint64, remaining uint64) {
	_ = "STUB: not implemented"
	return
}

// estimateSegmentsInProgressTime returns the ETA for all trie segments
// in progress to finish (uses the one with most remaining leafs to estimate).
func (t *trieSyncStats) estimateSegmentsInProgressTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// if there are no tries in progress, return 0

// trieDone takes a lock and adds one to the total number of tries synced.
func (t *trieSyncStats) trieDone(root common.Hash) { _ = "STUB: not implemented"; return }

// updateETA calculates and logs and ETA based on the number of leafs
// currently in progress and the number of tries remaining.
// assumes lock is held.
func (t *trieSyncStats) updateETA(sinceUpdate time.Duration, now time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// provide a separate ETA for the account trie syncing step since we
// don't know the total number of storage tries yet.

func (t *trieSyncStats) setTriesRemaining(triesRemaining int) { _ = "STUB: not implemented"; return }

// roundETA rounds [d] to a minute and chops off the "0s" suffix
// returns "<1m" if [d] rounds to 0 minutes.
func roundETA(d time.Duration) string { _ = "STUB: not implemented"; return "" }
