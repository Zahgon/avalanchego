// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package database

import (
	"errors"

	"github.com/ava-labs/libevm/ethdb"

	avalanchegodb "github.com/ava-labs/avalanchego/database"
)

var (
	errSnapshotNotSupported = errors.New("snapshot is not supported")
	errStatNotSupported     = errors.New("stat is not supported")

	_ ethdb.Batch         = (*batch)(nil)
	_ ethdb.KeyValueStore = (*database)(nil)
)

type database struct {
	db avalanchegodb.Database
}

func New(db avalanchegodb.Database) ethdb.KeyValueStore {
	_ = "STUB: not implemented"
	return *new(ethdb.KeyValueStore)
}

func (database) Stat(string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (db database) NewBatch() ethdb.Batch { _ = "STUB: not implemented"; return *new(ethdb.Batch) }

func (db database) Has(key []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (db database) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (db database) Put(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (db database) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (db database) Compact(start, limit []byte) error { _ = "STUB: not implemented"; return nil }

func (db database) Close() error { _ = "STUB: not implemented"; return nil }

func (db database) NewBatchWithSize(int) ethdb.Batch {
	_ = "STUB: not implemented"
	return *new(ethdb.Batch)
}

func (database) NewSnapshot() (ethdb.Snapshot, error) {
	_ = "STUB: not implemented"
	return *new(ethdb.Snapshot), nil
}

func (db database) NewIterator(prefix []byte, start []byte) ethdb.Iterator {
	_ = "STUB: not implemented"
	return *new(ethdb.Iterator)
}

type batch struct {
	batch avalanchegodb.Batch
}

func (b batch) Put(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (b batch) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (b batch) ValueSize() int { _ = "STUB: not implemented"; return 0 }

func (b batch) Write() error { _ = "STUB: not implemented"; return nil }

func (b batch) Reset() { _ = "STUB: not implemented"; return }

func (b batch) Replay(w ethdb.KeyValueWriter) error { _ = "STUB: not implemented"; return nil }
