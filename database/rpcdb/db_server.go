// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcdb

import (
	"context"
	"errors"
	"sync"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/utils/units"

	rpcdbpb "github.com/ava-labs/avalanchego/proto/pb/rpcdb"
)

const iterationBatchSize = 128 * units.KiB

var errUnknownIterator = errors.New("unknown iterator")

// DatabaseServer is a database that is managed over RPC.
type DatabaseServer struct {
	rpcdbpb.UnsafeDatabaseServer

	db database.Database

	// iteratorLock protects [nextIteratorID] and [iterators] from concurrent
	// modifications. Similarly to [batchLock], [iteratorLock] does not protect
	// the actual Iterator. Iterators are documented as not being safe for
	// concurrent use. Therefore, it is up to the client to respect this
	// invariant.
	iteratorLock   sync.RWMutex
	nextIteratorID uint64
	iterators      map[uint64]database.Iterator
}

// NewServer returns a database instance that is managed remotely
func NewServer(db database.Database) *DatabaseServer { _ = "STUB: not implemented"; return nil }

// Has delegates the Has call to the managed database and returns the result
func (db *DatabaseServer) Has(_ context.Context, req *rpcdbpb.HasRequest) (*rpcdbpb.HasResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get delegates the Get call to the managed database and returns the result
func (db *DatabaseServer) Get(_ context.Context, req *rpcdbpb.GetRequest) (*rpcdbpb.GetResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put delegates the Put call to the managed database and returns the result
func (db *DatabaseServer) Put(_ context.Context, req *rpcdbpb.PutRequest) (*rpcdbpb.PutResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete delegates the Delete call to the managed database and returns the
// result
func (db *DatabaseServer) Delete(_ context.Context, req *rpcdbpb.DeleteRequest) (*rpcdbpb.DeleteResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Compact delegates the Compact call to the managed database and returns the
// result
func (db *DatabaseServer) Compact(_ context.Context, req *rpcdbpb.CompactRequest) (*rpcdbpb.CompactResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close delegates the Close call to the managed database and returns the result
func (db *DatabaseServer) Close(context.Context, *rpcdbpb.CloseRequest) (*rpcdbpb.CloseResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HealthCheck performs a heath check against the underlying database.
func (db *DatabaseServer) HealthCheck(ctx context.Context, _ *emptypb.Empty) (*rpcdbpb.HealthCheckResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WriteBatch takes in a set of key-value pairs and atomically writes them to
// the internal database
func (db *DatabaseServer) WriteBatch(_ context.Context, req *rpcdbpb.WriteBatchRequest) (*rpcdbpb.WriteBatchResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewIteratorWithStartAndPrefix allocates an iterator and returns the iterator
// ID
func (db *DatabaseServer) NewIteratorWithStartAndPrefix(_ context.Context, req *rpcdbpb.NewIteratorWithStartAndPrefixRequest) (*rpcdbpb.NewIteratorWithStartAndPrefixResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IteratorNext attempts to call next on the requested iterator
func (db *DatabaseServer) IteratorNext(_ context.Context, req *rpcdbpb.IteratorNextRequest) (*rpcdbpb.IteratorNextResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IteratorError attempts to report any errors that occurred during iteration
func (db *DatabaseServer) IteratorError(_ context.Context, req *rpcdbpb.IteratorErrorRequest) (*rpcdbpb.IteratorErrorResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IteratorRelease attempts to release the resources allocated to an iterator
func (db *DatabaseServer) IteratorRelease(_ context.Context, req *rpcdbpb.IteratorReleaseRequest) (*rpcdbpb.IteratorReleaseResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
