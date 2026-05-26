// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bloom

import (
	"errors"
	"sync"
)

const (
	minHashes  = 1
	maxHashes  = 16 // Supports a false positive probability of 2^-16 when using optimal size values
	minEntries = 1

	bitsPerByte    = 8
	bytesPerUint64 = 8
	hashRotation   = 17
)

var (
	errInvalidNumHashes = errors.New("invalid num hashes")
	errTooFewHashes     = errors.New("too few hashes")
	errTooManyHashes    = errors.New("too many hashes")
	errTooFewEntries    = errors.New("too few entries")
)

type Filter struct {
	// numBits is always equal to [bitsPerByte * len(entries)]
	numBits uint64

	lock      sync.RWMutex
	hashSeeds []uint64
	entries   []byte
	count     int
}

// New creates a new Filter with the specified number of hashes and bytes for
// entries. The returned bloom filter is safe for concurrent usage.
func New(numHashes, numEntries int) (*Filter, error) { _ = "STUB: not implemented"; return nil, nil }

// Add adds the provided hash to the bloom filter. It returns true if the hash
// was not already present in the bloom filter.
func (f *Filter) Add(hash uint64) bool { _ = "STUB: not implemented"; return false }

// hint to the compiler that numBits is not 0

// Count returns the number of elements that have been added to the bloom
// filter.
func (f *Filter) Count() int { _ = "STUB: not implemented"; return 0 }

func (f *Filter) Contains(hash uint64) bool { _ = "STUB: not implemented"; return false }

func (f *Filter) Marshal() []byte { _ = "STUB: not implemented"; return nil }

func newHashSeeds(count int) ([]uint64, error) { _ = "STUB: not implemented"; return nil, nil }

func contains(hashSeeds []uint64, entries []byte, hash uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// hint to the compiler that numBits is not 0

func marshal(hashSeeds []uint64, entries []byte) []byte { _ = "STUB: not implemented"; return nil }
