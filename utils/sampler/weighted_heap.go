// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sampler

import (
	"github.com/ava-labs/avalanchego/utils"
)

var (
	_ Weighted                            = (*weightedHeap)(nil)
	_ utils.Sortable[weightedHeapElement] = weightedHeapElement{}
)

type weightedHeapElement struct {
	weight           uint64
	cumulativeWeight uint64
	index            int
}

// Compare the elements. Weight is in decreasing order. Index is in increasing
// order.
func (e weightedHeapElement) Compare(other weightedHeapElement) int {
	_ = "STUB: not implemented"
	// By accounting for the initial index of the weights, this results in a
	// stable sort. We do this rather than using `sort.Stable` because of the
	// reported change in performance of the sort used.
	return 0
}

// Sampling is performed by executing a search over a tree of elements in the
// order of their probabilistic occurrence.
//
// Initialization takes O(n * log(n)) time, where n is the number of elements
// that can be sampled.
// Sampling can take up to O(log(n)) time. As the distribution becomes more
// biased, sampling will become faster in expectation.
type weightedHeap struct {
	heap []weightedHeapElement
}

func (s *weightedHeap) Initialize(weights []uint64) error { _ = "STUB: not implemented"; return nil }

// Optimize so that the most probable values are at the top of the heap

// Initialize the heap

// Explicitly performing a shift here allows the compiler to avoid
// checking for negative numbers, which saves a couple cycles

func (s *weightedHeap) Sample(value uint64) (int, bool) { _ = "STUB: not implemented"; return 0, false }

// We shouldn't return the root, so check the left child

// If the weight is greater than the left weight, you should move to
// the right child
