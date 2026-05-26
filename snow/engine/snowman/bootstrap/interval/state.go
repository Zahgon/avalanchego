// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package interval

import (
	"errors"

	"github.com/ava-labs/avalanchego/database"
)

const (
	intervalPrefixByte byte = iota
	blockPrefixByte

	prefixLen = 1
)

var (
	intervalPrefix = []byte{intervalPrefixByte}
	blockPrefix    = []byte{blockPrefixByte}

	errInvalidKeyLength = errors.New("invalid key length")
)

func GetIntervals(db database.Iteratee) ([]*Interval, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PutInterval(db database.KeyValueWriter, upperBound uint64, lowerBound uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteInterval(db database.KeyValueDeleter, upperBound uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// makeIntervalKey uses the upperBound rather than the lowerBound because blocks
// are fetched from tip towards genesis. This means that it is more common for
// the lowerBound to change than the upperBound. Modifying the lowerBound only
// requires a single write rather than a write and a delete when modifying the
// upperBound.
func makeIntervalKey(upperBound uint64) []byte { _ = "STUB: not implemented"; return nil }

// GetBlockIterator returns a block iterator that will produce values
// corresponding to persisted blocks in order of increasing height.
func GetBlockIterator(db database.Iteratee) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

// GetBlockIteratorWithStart returns a block iterator that will produce values
// corresponding to persisted blocks in order of increasing height starting at
// [height].
func GetBlockIteratorWithStart(db database.Iteratee, height uint64) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func GetBlock(db database.KeyValueReader, height uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func PutBlock(db database.KeyValueWriter, height uint64, bytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func DeleteBlock(db database.KeyValueDeleter, height uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// makeBlockKey ensures that the returned key maintains the same sorted order as
// the height. This ensures that database iteration of block keys will iterate
// from lower height to higher height.
func makeBlockKey(height uint64) []byte { _ = "STUB: not implemented"; return nil }
