// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gossip

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/bloom"
)

// NewBloomFilter returns a new instance of a bloom filter with at least [minTargetElements] elements
// anticipated at any moment, and a false positive probability of [targetFalsePositiveProbability]. If the
// false positive probability exceeds [resetFalsePositiveProbability], the bloom filter will be reset.
//
// Invariant: The returned bloom filter is not safe to reset concurrently with
// other operations. However, it is otherwise safe to access concurrently.
//
// Deprecated: [BloomSet] should be used to manage bloom filters.
func NewBloomFilter(
	registerer prometheus.Registerer,
	namespace string,
	minTargetElements int,
	targetFalsePositiveProbability,
	resetFalsePositiveProbability float64,
) (*BloomFilter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Deprecated: [BloomSet] should be used to manage bloom filters.
type BloomFilter struct {
	minTargetElements              int
	targetFalsePositiveProbability float64
	resetFalsePositiveProbability  float64

	metrics *bloom.Metrics

	maxCount int
	bloom    *bloom.Filter
	// salt is provided to eventually unblock collisions in Bloom. It's possible
	// that conflicting Gossipable items collide in the bloom filter, so a salt
	// is generated to eventually resolve collisions.
	salt ids.ID
}

func (b *BloomFilter) Add(gossipable Gossipable) { _ = "STUB: not implemented"; return }

func (b *BloomFilter) Has(gossipable Gossipable) bool { _ = "STUB: not implemented"; return false }

func (b *BloomFilter) BloomFilter() (*bloom.Filter, ids.ID) {
	_ = "STUB: not implemented"
	return nil,

		// ResetBloomFilterIfNeeded resets a bloom filter if it breaches [targetFalsePositiveProbability].
		//
		// If [targetElements] exceeds [minTargetElements], the size of the bloom filter will grow to maintain
		// the same [targetFalsePositiveProbability].
		//
		// Returns true if the bloom filter was reset.
		//
		// Deprecated: [BloomSet] should be used to manage bloom filters.
		*new(ids.ID)
}

func ResetBloomFilterIfNeeded(
	bloomFilter *BloomFilter,
	targetElements int,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func resetBloomFilter(
	bloomFilter *BloomFilter,
	targetElements int,
	targetFalsePositiveProbability,
	resetFalsePositiveProbability float64,
) error {
	_ = "STUB: not implemented"
	return nil
}
