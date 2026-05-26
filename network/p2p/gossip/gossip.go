// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gossip

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/cache/lru"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/utils/bloom"
	"github.com/ava-labs/avalanchego/utils/buffer"
	"github.com/ava-labs/avalanchego/utils/logging"
)

const (
	ioLabel    = "io"
	sentIO     = "sent"
	receivedIO = "received"

	typeLabel  = "type"
	pushType   = "push"
	pullType   = "pull"
	unsentType = "unsent"
	sentType   = "sent"

	defaultGossipableCount = 64
)

var (
	_ Gossiper = (*ValidatorGossiper)(nil)
	_ Gossiper = (*PullGossiper[Gossipable])(nil)
	_ Gossiper = (*PushGossiper[Gossipable])(nil)

	ioTypeLabels   = []string{ioLabel, typeLabel}
	sentPushLabels = prometheus.Labels{
		ioLabel:   sentIO,
		typeLabel: pushType,
	}
	receivedPushLabels = prometheus.Labels{
		ioLabel:   receivedIO,
		typeLabel: pushType,
	}
	sentPullLabels = prometheus.Labels{
		ioLabel:   sentIO,
		typeLabel: pullType,
	}
	receivedPullLabels = prometheus.Labels{
		ioLabel:   receivedIO,
		typeLabel: pullType,
	}
	typeLabels   = []string{typeLabel}
	unsentLabels = prometheus.Labels{
		typeLabel: unsentType,
	}
	sentLabels = prometheus.Labels{
		typeLabel: sentType,
	}

	ErrInvalidNumValidators     = errors.New("num validators cannot be negative")
	ErrInvalidNumNonValidators  = errors.New("num non-validators cannot be negative")
	ErrInvalidNumPeers          = errors.New("num peers cannot be negative")
	ErrInvalidNumToGossip       = errors.New("must gossip to at least one peer")
	ErrInvalidDiscardedSize     = errors.New("discarded size cannot be negative")
	ErrInvalidTargetGossipSize  = errors.New("target gossip size cannot be negative")
	ErrInvalidRegossipFrequency = errors.New("re-gossip frequency cannot be negative")
)

// Gossipable is an item that can be gossiped across the network
type Gossipable interface {
	GossipID() ids.ID
}

// Marshaller handles parsing logic for a concrete Gossipable type
type Marshaller[T Gossipable] interface {
	MarshalGossip(T) ([]byte, error)
	UnmarshalGossip([]byte) (T, error)
}

// Gossiper gossips Gossipables to other nodes
type Gossiper interface {
	// Gossip runs a cycle of gossip. Returns an error if we failed to gossip.
	Gossip(ctx context.Context) error
}

// ValidatorGossiper only calls [Gossip] if the given node is a validator
type ValidatorGossiper struct {
	Gossiper

	NodeID     ids.NodeID
	Validators p2p.ValidatorSet
}

// Metrics that are tracked across a gossip protocol. A given protocol should
// only use a single instance of Metrics.
type Metrics struct {
	count                   *prometheus.CounterVec
	bytes                   *prometheus.CounterVec
	tracking                *prometheus.GaugeVec
	trackingLifetimeAverage prometheus.Gauge
	topValidators           *prometheus.GaugeVec
	bloomFilterHitRate      prometheus.Histogram
}

// NewMetrics returns a common set of metrics
func NewMetrics(
	metrics prometheus.Registerer,
	namespace string,
) (Metrics, error) {
	_ = "STUB: not implemented"
	return *new(Metrics), nil
}

// Buckets are (-∞, 0], (0, 25%], (25%, 50%], (50%, 75%], (75%, ∞).
// 0% is placed into its own bucket so that useless bloom filters
// can be inspected individually.

func (m *Metrics) observeMessage(labels prometheus.Labels, count int, bytes int) error {
	_ = "STUB: not implemented"
	return nil
}

func (v ValidatorGossiper) Gossip(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func NewPullGossiper[T Gossipable](
	log logging.Logger,
	marshaller Marshaller[T],
	set PullGossiperSet[T],
	client *p2p.Client,
	metrics Metrics,
	pollSize int,
) *PullGossiper[T] {
	_ = "STUB: not implemented"
	return nil
}

// PullGossiperSet exposes the current bloom filter and allows adding new items
// that were not included in the filter.
//
// TODO: Consider naming this interface based on what it provides rather than
// how its used.
type PullGossiperSet[T Gossipable] interface {
	// Add adds a value to the set. Returns an error if v was not added.
	Add(v T) error
	// BloomFilter returns the bloom filter and its corresponding salt.
	BloomFilter() (bloom *bloom.Filter, salt ids.ID)
}

type PullGossiper[T Gossipable] struct {
	log        logging.Logger
	marshaller Marshaller[T]
	set        PullGossiperSet[T]
	client     *p2p.Client
	metrics    Metrics
	pollSize   int
}

func (p *PullGossiper[_]) Gossip(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (p *PullGossiper[_]) handleResponse(
	_ context.Context,
	nodeID ids.NodeID,
	responseBytes []byte,
	err error,
) {
	_ = "STUB: not implemented"
	return
}

// NewPushGossiper returns an instance of PushGossiper
func NewPushGossiper[T Gossipable](
	marshaller Marshaller[T],
	set PushGossiperSet,
	validators p2p.ValidatorSubset,
	client *p2p.Client,
	metrics Metrics,
	gossipParams BranchingFactor,
	regossipParams BranchingFactor,
	discardedSize int,
	targetGossipSize int,
	maxRegossipFrequency time.Duration,
) (*PushGossiper[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PushGossiperSet exposes whether hashes are still included in a set.
//
// TODO: Consider naming this interface based on what it provides rather than
// how its used.
type PushGossiperSet interface {
	// Has returns true if the hash is in the set.
	Has(h ids.ID) bool
}

// PushGossiper broadcasts gossip to peers randomly in the network
type PushGossiper[T Gossipable] struct {
	marshaller Marshaller[T]
	set        PushGossiperSet
	validators p2p.ValidatorSubset
	client     *p2p.Client
	metrics    Metrics

	gossipParams         BranchingFactor
	regossipParams       BranchingFactor
	targetGossipSize     int
	maxRegossipFrequency time.Duration

	lock         sync.Mutex
	tracking     map[ids.ID]*tracking
	addedTimeSum float64 // unix nanoseconds
	toGossip     buffer.Deque[T]
	toRegossip   buffer.Deque[T]
	discarded    *lru.Cache[ids.ID, struct{}] // discarded attempts to avoid overgossiping transactions that are frequently dropped
}

type BranchingFactor struct {
	// StakePercentage determines the percentage of stake that should have
	// gossip sent to based on the inverse CDF of stake weights. This value does
	// not account for the connectivity of the nodes.
	StakePercentage float64
	// Validators specifies the number of connected validators, in addition to
	// any validators sent from the StakePercentage parameter, to send gossip
	// to. These validators are sampled uniformly rather than by stake.
	Validators int
	// NonValidators specifies the number of connected non-validators to send
	// gossip to.
	NonValidators int
	// Peers specifies the number of connected validators or non-validators, in
	// addition to the number sent due to other configs, to send gossip to.
	Peers int
}

func (b *BranchingFactor) Verify() error { _ = "STUB: not implemented"; return nil }

type tracking struct {
	addedTime    float64 // unix nanoseconds
	lastGossiped time.Time
}

// Gossip flushes any queued gossipables.
func (p *PushGossiper[T]) Gossip(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Don't mark dropped unsent transactions as discarded

// Mark dropped sent transactions as discarded

func (p *PushGossiper[T]) gossip(
	ctx context.Context,
	now time.Time,
	gossipParams BranchingFactor,
	toGossip buffer.Deque[T],
	toRegossip buffer.Deque[T],
	discarded cache.Cacher[ids.ID, struct{}],
	metricsLabels prometheus.Labels,
) error {
	_ = "STUB: not implemented"
	return nil
}

// Ensure item is still in the set before we gossip.

// Cache that the item was dropped

// Ensure we don't attempt to send a gossipable too frequently.

// Put the gossipable on the front of the queue to keep items sorted
// by last issuance time.

// If there is nothing to gossip, we can exit early.

// Send gossipables to peers

// Add enqueues new gossipables to be pushed. If a gossipable is already tracked,
// it is not added again.
func (p *PushGossiper[T]) Add(gossipables ...T) { _ = "STUB: not implemented"; return }

// Add new gossipables to be sent.

// Pretend that recently discarded transactions were just gossiped.

func (p *PushGossiper[_]) updateMetrics(nowUnixNano float64) { _ = "STUB: not implemented"; return }

// Every calls [Gossip] every [period] amount of time.
func Every(ctx context.Context, log logging.Logger, gossiper Gossiper, period time.Duration) {
	_ = "STUB: not implemented"
	return
}
