// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gossip

import (
	"errors"
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/bloom"
)

const (
	DefaultMinTargetElements              = 1000
	DefaultTargetFalsePositiveProbability = .01
	DefaultResetFalsePositiveProbability  = .05
)

var (
	_ Set[Gossipable]             = (*BloomSet[Gossipable])(nil)
	_ PullGossiperSet[Gossipable] = (*BloomSet[Gossipable])(nil)

	ErrBloomReset = errors.New("bloom reset")
)

type BloomSetConfig struct {
	// Metrics allows exposing the current state of the bloom filter across
	// additions and resets. If nil, no metrics are recorded.
	Metrics *bloom.Metrics
	// MinTargetElements is the minimum number of elements to target when
	// creating a new bloom filter. If zero, [DefaultMinTargetElements] is used.
	MinTargetElements int
	// TargetFalsePositiveProbability is the target false positive probability
	// when creating a new bloom filter. If zero,
	// [DefaultTargetFalsePositiveProbability] is used.
	TargetFalsePositiveProbability float64
	// ResetFalsePositiveProbability is the false positive probability at
	// which the bloom filter is reset. If zero,
	// [DefaultResetFalsePositiveProbability] is used.
	ResetFalsePositiveProbability float64
}

func (c *BloomSetConfig) fillDefaults() { _ = "STUB: not implemented"; return }

type Set[T Gossipable] interface {
	HandlerSet[T]
	PushGossiperSet
	// Len returns the number of items in the set.
	//
	// This value should match the number of items that can be iterated over
	// with a call to Iterate.
	Len() int
}

// NewBloomSet wraps the [Set] with a bloom filter. It is expected for all
// future additions to the provided set to go through the returned [BloomSet].
func NewBloomSet[T Gossipable](
	set Set[T],
	c BloomSetConfig,
) (*BloomSet[T], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type BloomSet[T Gossipable] struct {
	set Set[T]
	c   BloomSetConfig

	lock     sync.RWMutex
	maxCount int
	bloom    *bloom.Filter
	// salt is provided to eventually unblock collisions in Bloom. It's possible
	// that conflicting Gossipable items collide in the bloom filter, so a salt
	// is generated to eventually resolve collisions.
	salt ids.ID
}

// Add adds v to the set and bloom filter.
//
// If adding the inner set succeeds and resetting the bloom filter fails,
// [ErrBloomReset] is returned. However, it is still guaranteed that v has
// been added to the bloom filter.
func (s *BloomSet[T]) Add(v T) error { _ = "STUB: not implemented"; return nil }

// addToBloom adds the provided ID to the bloom filter.
//
// Even if an error is returned, the ID has still been added to the bloom
// filter.
func (s *BloomSet[T]) addToBloom(h ids.ID) error { _ = "STUB: not implemented"; return nil }

// Bloom filter was already reset by another thread

// shouldReset expects either a read lock or a write lock to be held.
func (s *BloomSet[T]) shouldReset() bool { _ = "STUB: not implemented"; return false }

// resetBloom attempts to generate a new bloom filter and fill it with the
// current entries in the set.
//
// If an error is returned, the bloom filter and salt are unchanged.
//
// resetBloom expects a write lock to be held.
func (s *BloomSet[T]) resetBloom() error { _ = "STUB: not implemented"; return nil }

func (s *BloomSet[T]) Has(h ids.ID) bool { _ = "STUB: not implemented"; return false }

func (s *BloomSet[T]) Iterate(f func(T) bool) { _ = "STUB: not implemented"; return }

func (s *BloomSet[T]) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *BloomSet[_]) BloomFilter() (*bloom.Filter, ids.ID) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ID)
}
