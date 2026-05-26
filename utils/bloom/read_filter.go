// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bloom

var (
	EmptyFilter = &ReadFilter{
		hashSeeds: make([]uint64, minHashes),
		entries:   make([]byte, minEntries),
	}
	FullFilter = &ReadFilter{
		hashSeeds: make([]uint64, minHashes),
		entries:   make([]byte, minEntries),
	}
)

func init() {
	for i := range FullFilter.entries {
		FullFilter.entries[i] = 0xFF
	}
}

type ReadFilter struct {
	hashSeeds []uint64
	entries   []byte
}

// Parse [bytes] into a read-only bloom filter.
func Parse(bytes []byte) (*ReadFilter, error) { _ = "STUB: not implemented"; return nil, nil }

// numEntries = len(bytes) - entriesOffset

func (f *ReadFilter) Contains(hash uint64) bool { _ = "STUB: not implemented"; return false }

func (f *ReadFilter) Marshal() []byte { _ = "STUB: not implemented"; return nil }
