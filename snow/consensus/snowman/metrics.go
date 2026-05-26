// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/linked"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/math"
	"github.com/ava-labs/avalanchego/utils/metric"
)

type processingStart struct {
	time       time.Time
	pollNumber uint64
}

type metrics struct {
	log logging.Logger

	currentMaxVerifiedHeight uint64
	maxVerifiedHeight        prometheus.Gauge

	lastAcceptedHeight    prometheus.Gauge
	lastAcceptedTimestamp prometheus.Gauge

	// processingBlocks keeps track of the [processingStart] that each block was
	// issued into the consensus instance. This is used to calculate the amount
	// of time to accept or reject the block.
	processingBlocks *linked.Hashmap[ids.ID, processingStart]

	// numProcessing keeps track of the number of processing blocks
	numProcessing prometheus.Gauge

	blockSizeAcceptedSum prometheus.Gauge
	// pollsAccepted tracks the number of polls that a block was in processing
	// for before being accepted
	pollsAccepted metric.Averager
	// latAccepted tracks the number of nanoseconds that a block was processing
	// before being accepted
	latAccepted          metric.Averager
	consensusLatencies   prometheus.Histogram
	buildLatencyAccepted prometheus.Gauge

	blockSizeRejectedSum prometheus.Gauge
	// pollsRejected tracks the number of polls that a block was in processing
	// for before being rejected
	pollsRejected metric.Averager
	// latRejected tracks the number of nanoseconds that a block was processing
	// before being rejected
	latRejected metric.Averager

	// numFailedPolls keeps track of the number of polls that failed
	numFailedPolls prometheus.Counter

	// numSuccessfulPolls keeps track of the number of polls that succeeded
	numSuccessfulPolls prometheus.Counter

	// avgAcceptanceLatency tracks the average acceptance time
	avgAcceptanceLatency math.Averager
}

func newMetrics(
	log logging.Logger,
	reg prometheus.Registerer,
	lastAcceptedHeight uint64,
	lastAcceptedTime time.Time,
) (*metrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initially set the metrics for the last accepted block.

func (m *metrics) Issued(blkID ids.ID, pollNumber uint64) { _ = "STUB: not implemented"; return }

func (m *metrics) Verified(height uint64) { _ = "STUB: not implemented"; return }

func (m *metrics) Accepted(
	blkID ids.ID,
	height uint64,
	timestamp time.Time,
	pollNumber uint64,
	blockSize int,
) {
	_ = "STUB: not implemented"
	return
}

func (m *metrics) Rejected(blkID ids.ID, pollNumber uint64, blockSize int) {
	_ = "STUB: not implemented"
	return
}

func (m *metrics) MeasureAndGetOldestDuration() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (m *metrics) SuccessfulPoll() { _ = "STUB: not implemented"; return }

func (m *metrics) FailedPoll() { _ = "STUB: not implemented"; return }

func (m *metrics) GetAverageAcceptanceTime() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}
