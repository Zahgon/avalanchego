// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sampler

import (
	"sync"
)

var globalRNG = newRNG()

func newRNG() *rng {
	_ = "STUB: not implemented"
	// We don't use a cryptographically secure source of randomness here, as
	// there's no need to ensure a truly random sampling.
	return nil
}

type rng struct {
	lock sync.Mutex
	rng  Source
}

type Source interface {
	// Uint64 returns a random number in [0, MaxUint64] and advances the
	// generator's state.
	Uint64() uint64
}

// Uint64Inclusive returns a pseudo-random number in [0,n].
//
// Invariant: The result of this function is stored in chain state, so any
// modifications are considered breaking.
func (r *rng) Uint64Inclusive(n uint64) uint64 {
	_ = "STUB: not implemented"

	// n+1 is power of two, so we can just mask
	//
	// Note: This does work for MaxUint64 as overflow is explicitly part of the
	// compiler specification: https://go.dev/ref/spec#Integer_overflow
	return 0
}

// n is greater than MaxUint64/2 so we need to just iterate until we get a
// number in the requested range.

// n is less than MaxUint64/2 so we generate a number in the range
// [0, k*(n+1)) where k is the largest integer such that k*(n+1) is less
// than or equal to MaxUint64/2. We can't easily find k such that k*(n+1) is
// less than or equal to MaxUint64 because the calculation would overflow.
//
// ref: https://github.com/golang/go/blob/ce10e9d84574112b224eae88dc4e0f43710808de/src/math/rand/rand.go#L127-L132

// uint63 returns a random number in [0, MaxInt64]
func (r *rng) uint63() uint64 { _ = "STUB: not implemented"; return 0 }

// uint64 returns a random number in [0, MaxUint64]
func (r *rng) uint64() uint64 {
	_ = "STUB: not implemented"
	// Note: We must grab a write lock here because rng.Uint64 internally
	// modifies state.
	return 0
}
