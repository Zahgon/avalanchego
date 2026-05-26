// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sampler

type defaultMap map[uint64]uint64

func (m defaultMap) get(key uint64, defaultVal uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// uniformReplacer allows for sampling over a uniform distribution without
// replacement.
//
// Sampling is performed by lazily performing an array shuffle of the array
// [0, 1, ..., length - 1]. By performing the first count swaps of this shuffle,
// we can create an array of length count with elements sampled with uniform
// probability.
//
// Initialization takes O(1) time.
//
// Sampling is performed in O(count) time and O(count) space.
type uniformReplacer struct {
	rng        *rng
	length     uint64
	drawn      defaultMap
	drawsCount uint64
}

func (s *uniformReplacer) Initialize(length uint64) { _ = "STUB: not implemented"; return }

func (s *uniformReplacer) Sample(count int) ([]uint64, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *uniformReplacer) Reset() { _ = "STUB: not implemented"; return }

func (s *uniformReplacer) Next() (uint64, bool) { _ = "STUB: not implemented"; return 0, false }
