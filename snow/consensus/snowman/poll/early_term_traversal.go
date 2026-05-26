// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package poll

import (
	"errors"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/bag"
)

var (
	errPollDurationVectorMetrics = errors.New("failed to register poll_duration vector metrics")
	errPollCountVectorMetrics    = errors.New("failed to register poll_count vector metrics")

	terminationReason = "reason"
	exhaustedReason   = "exhausted"
	earlyFailReason   = "early_fail"
	earlyAlphaReason  = "early_alpha"

	exhaustedLabel = prometheus.Labels{
		terminationReason: exhaustedReason,
	}
	earlyFailLabel = prometheus.Labels{
		terminationReason: earlyFailReason,
	}
	earlyAlphaLabel = prometheus.Labels{
		terminationReason: earlyAlphaReason,
	}
)

type earlyTermMetrics struct {
	durExhaustedPolls  prometheus.Gauge
	durEarlyFailPolls  prometheus.Gauge
	durEarlyAlphaPolls prometheus.Gauge

	countExhaustedPolls  prometheus.Counter
	countEarlyFailPolls  prometheus.Counter
	countEarlyAlphaPolls prometheus.Counter
}

func newEarlyTermMetrics(reg prometheus.Registerer) (*earlyTermMetrics, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *earlyTermMetrics) observeExhausted(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (m *earlyTermMetrics) observeEarlyFail(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

func (m *earlyTermMetrics) observeEarlyAlpha(duration time.Duration) {
	_ = "STUB: not implemented"
	return
}

type earlyTermTraversalFactory struct {
	alphaPreference int
	alphaConfidence int
	bt              BlockTraversal
	metrics         *earlyTermMetrics
}

type BlockTraversal interface {
	GetParent(id ids.ID) (ids.ID, bool)
}

// NewEarlyTermFactory returns a factory that returns polls with early termination.
func NewEarlyTermFactory(
	alphaPreference int,
	alphaConfidence int,
	reg prometheus.Registerer,
	bt BlockTraversal,
) (Factory, error) {
	_ = "STUB: not implemented"
	return *new(Factory), nil
}

func (f *earlyTermTraversalFactory) New(vdrs bag.Bag[ids.NodeID]) Poll {
	_ = "STUB: not implemented"
	return *new(Poll)
}

// earlyTermPoll finishes when any remaining validators can't change
// the result of the poll for all the votes and transitive votes.
type earlyTermPoll struct {
	votes           bag.Bag[ids.ID]
	polled          bag.Bag[ids.NodeID]
	alphaPreference int
	alphaConfidence int
	bt              BlockTraversal
	metrics         *earlyTermMetrics
	start           time.Time
	finished        bool
}

// Vote registers a response for this poll
func (p *earlyTermPoll) Vote(vdr ids.NodeID, vote ids.ID) { _ = "STUB: not implemented"; return }

// make sure that a validator can't respond multiple times

// track the votes the validator responded with

// Drop any future response for this poll
func (p *earlyTermPoll) Drop(vdr ids.NodeID) { _ = "STUB: not implemented"; return }

// Finished returns true when one of the following conditions is met.
//
//  1. There are no outstanding votes.
//  2. It is impossible for the poll to achieve an alphaPreference majority
//     after applying transitive voting.
//  3. A single element has achieved an alphaPreference majority and it is
//     impossible for it to achieve an alphaConfidence majority after applying
//     transitive voting.
//  4. A single element has achieved an alphaConfidence majority.
func (p *earlyTermPoll) Finished() bool { _ = "STUB: not implemented"; return false }

// Case 1

// Case 2

//    v
//   /
//  u
// We build a vote graph where each vertex represents a block ID.
// A vertex 'v' is a parent of vertex 'u' if the ID of 'u' corresponds
// to a block that is the successive block of the corresponding block for 'v'.

// If vertex 'v' is a parent of vertex 'u', then a vote for the ID of vertex 'u'
// should also be considered as a vote for the ID of the vertex 'v'.

//     v
//   /   \
//  u     w
// If two competing blocks 'u', 'w' are potential successors to a block 'v',
// snowman would instantiate a unary snowflake instance on the prefix of 'u' and 'w'.
// The prefix inherits the votes for the IDs of 'u' and 'w'.
// We therefore compute the transitive votes for all prefixes of IDs
// for each bifurcation in the transitive vote graph.

// We wish to compute the votes for snowflake instances, no matter if they correspond to an actual block ID,
// or a unary snowflake instance for a shared prefix between a bifurcation of two competing blocks.
// For that, only the number of votes and existence of such snowflake instances matters.

// Given the aforementioned votes, we wish to see whether there exists a snowflake instance
// that can benefit from waiting for more invocations of Vote().
// We therefore check each amount of votes separately and see if voting for that snowflake instance
// should terminate, as it cannot be improved by further voting.

// If we have no votes, we may be able to improve the poll on some ID.

// Consider the votes for each ID or prefix of IDs,
// if we shouldn't terminate in one of them, then we should not terminate this poll now.

// We should terminate the poll only when votes for all IDs or prefixes cannot be improved.

func (p *earlyTermPoll) shouldTerminate(freq int, remaining int) bool {
	_ = "STUB: not implemented"
	return false
}

// Case 2
// Case 3
// Case 4

// Result returns the result of this poll
func (p *earlyTermPoll) Result() bag.Bag[ids.ID] { _ = "STUB: not implemented"; return nil }

func (p *earlyTermPoll) PrefixedString(prefix string) string { _ = "STUB: not implemented"; return "" }

func (p *earlyTermPoll) String() string { _ = "STUB: not implemented"; return "" }

func aggregateVotesFromPrefixesAndIDs(transitiveVotesForPrefixes []int, transitiveVotes bag.Bag[ids.ID]) []int {
	_ = "STUB: not implemented"
	return nil
}

func computeTransitiveVotesForPrefixes(votesGraph *voteGraph, transitiveVotes bag.Bag[ids.ID]) []int {
	_ = "STUB: not implemented"
	return nil
}

// Each shared prefix is associated with a bunch of IDs.
// Sum up all the transitive votes for these blocks,
// and return all such shared prefixes indexed by the underlying transitive descendant IDs.

func descendantIDsOfVertex(v *voteVertex) []ids.ID { _ = "STUB: not implemented"; return nil }

func sumVotesFromIDs(ids []ids.ID, transitiveVotes bag.Bag[ids.ID]) int {
	_ = "STUB: not implemented"
	return 0
}
