// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcdb

import (
	"context"
	"sync"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/utils"

	rpcdbpb "github.com/ava-labs/avalanchego/proto/pb/rpcdb"
)

var (
	_ database.Database = (*DatabaseClient)(nil)
	_ database.Batch    = (*batch)(nil)
	_ database.Iterator = (*iterator)(nil)
)

// DatabaseClient is an implementation of database that talks over RPC.
type DatabaseClient struct {
	client rpcdbpb.DatabaseClient

	closed utils.Atomic[bool]
}

// NewClient returns a database instance connected to a remote database instance
func NewClient(client rpcdbpb.DatabaseClient) *DatabaseClient {
	_ = "STUB: not implemented"
	return nil
}

// Has attempts to return if the database has a key with the provided value.
func (db *DatabaseClient) Has(key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Get attempts to return the value that was mapped to the key that was provided
func (db *DatabaseClient) Get(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put attempts to set the value this key maps to
func (db *DatabaseClient) Put(key, value []byte) error { _ = "STUB: not implemented"; return nil }

// Delete attempts to remove any mapping from the key
func (db *DatabaseClient) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

// NewBatch returns a new batch
func (db *DatabaseClient) NewBatch() database.Batch {
	_ = "STUB: not implemented"
	return *new(database.Batch)
}

func (db *DatabaseClient) NewIterator() database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func (db *DatabaseClient) NewIteratorWithStart(start []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

func (db *DatabaseClient) NewIteratorWithPrefix(prefix []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

// NewIteratorWithStartAndPrefix returns a new empty iterator
func (db *DatabaseClient) NewIteratorWithStartAndPrefix(start, prefix []byte) database.Iterator {
	_ = "STUB: not implemented"
	return *new(database.Iterator)
}

// Compact attempts to optimize the space utilization in the provided range
func (db *DatabaseClient) Compact(start, limit []byte) error { _ = "STUB: not implemented"; return nil }

// Close attempts to close the database
func (db *DatabaseClient) Close() error { _ = "STUB: not implemented"; return nil }

func (db *DatabaseClient) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type batch struct {
	database.BatchOps

	db *DatabaseClient
}

func (b *batch) Write() error { _ = "STUB: not implemented"; return nil }

func (b *batch) Inner() database.Batch { _ = "STUB: not implemented"; return *new(database.Batch) }

type iterator struct {
	db *DatabaseClient
	id uint64

	data        []*rpcdbpb.PutRequest
	fetchedData chan []*rpcdbpb.PutRequest

	errLock sync.RWMutex
	err     error

	reqUpdateError chan chan struct{}

	once     sync.Once
	onClose  chan struct{}
	onClosed chan struct{}
}

func newIterator(db *DatabaseClient, id uint64) *iterator { _ = "STUB: not implemented"; return nil }

// Invariant: fetch is the only thread with access to send requests to the
// server's iterator. This is needed because iterators are not thread safe and
// the server expects the client (us) to only ever issue one request at a time
// for a given iterator id.
func (it *iterator) fetch() { _ = "STUB: not implemented"; return }

// Next attempts to move the iterator to the next element and returns if this
// succeeded
func (it *iterator) Next() bool { _ = "STUB: not implemented"; return false }

// Error returns any that occurred while iterating
func (it *iterator) Error() error { _ = "STUB: not implemented"; return nil }

// Key returns the key of the current element
func (it *iterator) Key() []byte { _ = "STUB: not implemented"; return nil }

// Value returns the value of the current element
func (it *iterator) Value() []byte { _ = "STUB: not implemented"; return nil }

// Release frees any resources held by the iterator
func (it *iterator) Release() { _ = "STUB: not implemented"; return }

func (it *iterator) updateError() { _ = "STUB: not implemented"; return }

func (it *iterator) setError(err error) { _ = "STUB: not implemented"; return }

func (it *iterator) getError() error { _ = "STUB: not implemented"; return nil }
