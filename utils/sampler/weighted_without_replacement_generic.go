// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sampler

type weightedWithoutReplacementGeneric struct {
	u Uniform
	w Weighted
}

func (s *weightedWithoutReplacementGeneric) Initialize(weights []uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *weightedWithoutReplacementGeneric) Sample(count int) ([]int, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
