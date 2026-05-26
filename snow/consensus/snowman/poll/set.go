// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package poll

import (
	"errors"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/bag"
	"github.com/ava-labs/avalanchego/utils/linked"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/metric"
)

var (
	errFailedPollsMetric         = errors.New("failed to register polls metric")
	errFailedPollDurationMetrics = errors.New("failed to register poll_duration metrics")
)

type pollHolder interface {
	GetPoll() Poll
	StartTime() time.Time
}

type poll struct {
	Poll
	start time.Time
}

func (p poll) GetPoll() Poll { _ = "STUB: not implemented"; return *new(Poll) }

func (p poll) StartTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

type set struct {
	log      logging.Logger
	numPolls prometheus.Gauge
	durPolls metric.Averager
	factory  Factory
	// maps requestID -> poll
	polls *linked.Hashmap[uint32, pollHolder]
}

// NewSet returns a new empty set of polls
func NewSet(
	factory Factory,
	log logging.Logger,
	reg prometheus.Registerer,
) (Set, error) {
	_ = "STUB: not implemented"
	return *new(Set), nil
}

// Add to the current set of polls
// Returns true if the poll was registered correctly and the network sample
// should be made.
func (s *set) Add(requestID uint32, vdrs bag.Bag[ids.NodeID]) bool {
	_ = "STUB: not implemented"
	return false
}

// create the new poll

// increase the metrics

// Vote registers the connections response to a query for [id]. If there was no
// query, or the response has already be registered, nothing is performed.
func (s *set) Vote(requestID uint32, vdr ids.NodeID, vote ids.ID) []bag.Bag[ids.ID] {
	_ = "STUB: not implemented"
	return nil
}

// processFinishedPolls checks for other dependent finished polls and returns them all if finished
func (s *set) processFinishedPolls() []bag.Bag[ids.ID] { _ = "STUB: not implemented"; return nil }

// iterate from oldest to newest

// since we're iterating from oldest to newest, if the next poll has not finished,
// we can break and return what we have so far

// decrease the metrics

// only gets here if the poll has finished
// results will have values if this and other newer polls have finished

// Drop registers the connections response to a query for [id]. If there was no
// query, or the response has already be registered, nothing is performed.
func (s *set) Drop(requestID uint32, vdr ids.NodeID) []bag.Bag[ids.ID] {
	_ = "STUB: not implemented"
	return nil
}

// Len returns the number of outstanding polls
func (s *set) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *set) String() string { _ = "STUB: not implemented"; return "" }
