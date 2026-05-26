// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package database

import (
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
)

const (
	Uint64Size = 8 // bytes
	BoolSize   = 1 // bytes
	BoolFalse  = 0x00
	BoolTrue   = 0x01

	// kvPairOverhead is an estimated overhead for a kv pair in a database.
	kvPairOverhead = 8 // bytes
)

var (
	boolFalseKey = []byte{BoolFalse}
	boolTrueKey  = []byte{BoolTrue}

	errWrongSize = errors.New("value has unexpected size")
)

func PutID(db KeyValueWriter, key []byte, val ids.ID) error { _ = "STUB: not implemented"; return nil }

func GetID(db KeyValueReader, key []byte) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func PutUInt64(db KeyValueWriter, key []byte, val uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func GetUInt64(db KeyValueReader, key []byte) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func PackUInt64(val uint64) []byte { _ = "STUB: not implemented"; return nil }

func ParseUInt64(b []byte) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func PutUInt32(db KeyValueWriter, key []byte, val uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func GetUInt32(db KeyValueReader, key []byte) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func PackUInt32(val uint32) []byte { _ = "STUB: not implemented"; return nil }

func ParseUInt32(b []byte) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

func PutTimestamp(db KeyValueWriter, key []byte, val time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func GetTimestamp(db KeyValueReader, key []byte) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func ParseTimestamp(b []byte) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func PutBool(db KeyValueWriter, key []byte, b bool) error { _ = "STUB: not implemented"; return nil }

func GetBool(db KeyValueReader, key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// WithDefault returns the value at [key] in [db]. If the key doesn't exist, it
// returns [def].
func WithDefault[V any](
	get func(KeyValueReader, []byte) (V, error),
	db KeyValueReader,
	key []byte,
	def V,
) (V, error) {
	_ = "STUB: not implemented"
	return *new(V), nil
}

func Count(db Iteratee) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func Size(db Iteratee) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func AtomicClear(readerDB Iteratee, deleterDB KeyValueDeleter) error {
	_ = "STUB: not implemented"
	return nil
}

// AtomicClearPrefix deletes from [deleterDB] all keys in [readerDB] that have the given [prefix].
func AtomicClearPrefix(readerDB Iteratee, deleterDB KeyValueDeleter, prefix []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Remove all key-value pairs from [db].
// Writes each batch when it reaches [writeSize].
func Clear(db Database, writeSize int) error { _ = "STUB: not implemented"; return nil }

// Removes all keys with the given [prefix] from [db].
// Writes each batch when it reaches [writeSize].
func ClearPrefix(db Database, prefix []byte, writeSize int) error {
	_ = "STUB: not implemented"
	return nil
}

// Defer the release of the iterator inside a closure to guarantee that the
// latest, not the first, iterator is released on return.

// Avoid too much memory pressure by periodically writing to the
// database.

// Reset the iterator to release references to now deleted keys.
