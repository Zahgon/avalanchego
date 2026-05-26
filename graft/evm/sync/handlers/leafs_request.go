// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package handlers

import (
	"context"
	"time"

	"github.com/ava-labs/libevm/ethdb/memorydb"
	"github.com/ava-labs/libevm/trie"
	"github.com/ava-labs/libevm/triedb"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/graft/evm/core/state/snapshot"
	"github.com/ava-labs/avalanchego/graft/evm/message"
	"github.com/ava-labs/avalanchego/graft/evm/sync/handlers/stats"
	"github.com/ava-labs/avalanchego/ids"
)

var _ LeafRequestHandler = (*leafsRequestHandler)(nil)

// LeafRequestHandler handles incoming leaf requests from peers.
type LeafRequestHandler interface {
	OnLeafsRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, leafsRequest message.LeafsRequest) ([]byte, error)
}

// SnapshotProvider provides access to the snapshot tree.
type SnapshotProvider interface {
	Snapshots() *snapshot.Tree
}

const (
	// Maximum number of leaves to return in a message.LeafsResponse
	// This parameter overrides any other Limit specified
	// in message.LeafsRequest if it is greater than this value
	maxLeavesLimit = uint16(1024)

	// Maximum percent of the time left to deadline to spend on optimistically
	// reading the snapshot to find the response
	maxSnapshotReadTimePercent = 75

	segmentLen = 64 // divide data from snapshot to segments of this size
)

// leafsRequestHandler is a peer.RequestHandler for types.LeafsRequest
// serving requested trie data
type leafsRequestHandler struct {
	trieDB           *triedb.Database
	snapshotProvider SnapshotProvider
	codec            codec.Manager
	stats            stats.LeafsRequestHandlerStats
	trieKeyLength    int
}

func NewLeafsRequestHandler(trieDB *triedb.Database, trieKeyLength int, snapshotProvider SnapshotProvider, codec codec.Manager, syncerStats stats.LeafsRequestHandlerStats) *leafsRequestHandler {
	_ = "STUB: not implemented"
	return nil
}

// OnLeafsRequest returns encoded message.LeafsResponse for a given message.LeafsRequest
// Returns leaves with proofs for specified (Start-End) (both inclusive) ranges
// Returned message.LeafsResponse may contain partial leaves within requested Start and End range if:
// - ctx expired while fetching leafs
// - number of leaves read is greater than Limit (message.LeafsRequest)
// Specified Limit in message.LeafsRequest is overridden to maxLeavesLimit if it is greater than maxLeavesLimit
// Expects returned errors to be treated as FATAL
// Never returns errors
// Returns nothing if NodeType is invalid or requested trie root is not found
// Assumes ctx is active
func (lrh *leafsRequestHandler) OnLeafsRequest(ctx context.Context, nodeID ids.NodeID, requestID uint32, leafsRequest message.LeafsRequest) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: We should know the state root that accounts correspond to,
// as this information will be necessary to access storage tries when
// the trie is path based.
// stateRoot := common.Hash{}

// override limit if it is greater than the configured maxLeavesLimit

// pass snapshot to responseBuilder if non-nil snapshot getter provided

// ensure metrics are captured properly on all return paths

type responseBuilder struct {
	request   message.LeafsRequest
	response  *message.LeafsResponse
	t         *trie.Trie
	snap      *snapshot.Tree
	keyLength int
	limit     uint16

	// stats
	trieReadTime time.Duration
	proofTime    time.Duration
	stats        stats.LeafsRequestHandlerStats
}

func (rb *responseBuilder) handleRequest(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Read from snapshot if a [snapshot.Tree] was provided in initialization
	return nil
}

// reset the proof if we will iterate the trie further

// more indicates whether there are more leaves in the trie

// omit proof via early return

// Generate the proof and add it to the response.

// closing memdb does not error

// fillFromSnapshot reads data from snapshot and returns true if the response is complete.
// Otherwise, the caller should attempt to iterate the trie and determine if a range proof
// should be added to the response.
func (rb *responseBuilder) fillFromSnapshot(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Optimistically read leafs from the snapshot, assuming they have not been
// modified since the requested root. If this assumption can be verified with
// range proofs and data from the trie, we can skip iterating the trie as
// an optimization.
// Since we are performing this read optimistically, we use a separate context
// with reduced timeout so there is enough time to read the trie if the snapshot
// read does not contain up-to-date data.

// Update read snapshot time here, so that we include the case that an error occurred.

// Check if the entire range read from the snapshot is valid according to the trie.

// closing memdb does not error

// omit proof via early return

// The data from the snapshot could not be validated as a whole. It is still likely
// most of the data from the snapshot is useable, so we try to validate smaller
// segments of the data and use them in the response.

// we don't need this proof

// segment is not valid

// segment is valid

// if there is a gap between valid segments, fill the gap with data from the trie

// remove the last key added since it is snapKeys[i] and will be added back
// Note: this is safe because we were able to verify the range proof that
// shows snapKeys[i] is part of the trie.

// all the key/vals in the segment are valid, but possibly shorten segmentEnd
// here to respect limit. this is necessary in case the number of leafs we read
// from the trie is more than the length of a segment which cannot be validated. limit

// generateRangeProof returns a range proof for the range specified by [start] and [keys] using [t].
func (rb *responseBuilder) generateRangeProof(start []byte, keys [][]byte) (*memorydb.Database, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If [start] is empty, populate it with the appropriate length key starting at 0.

// closing memdb does not error

// If there is a non-zero number of keys, set [end] for the range proof to the last key.

// closing memdb does not error

// verifyRangeProof verifies the provided range proof with [keys/vals], starting at [start].
// Returns a boolean indicating if there are more leaves to the right of the last key in the trie and a nil error if the range proof is successfully verified.
func (rb *responseBuilder) verifyRangeProof(keys, vals [][]byte, start []byte, proof *memorydb.Database) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// If [start] is empty, populate it with the appropriate length key starting at 0.

// iterateVals returns the values contained in [db]
func iterateVals(db *memorydb.Database) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// iterate db into [][]byte and return

// isRangeValid generates and verifies a range proof, returning true if keys/vals are
// part of the trie. If [hasGap] is true, the range is validated independent of the
// existing response. If [hasGap] is false, the range proof begins at a key which
// guarantees the range can be appended to the response.
// Additionally returns a boolean indicating if there are more leaves in the trie.
func (rb *responseBuilder) isRangeValid(keys, vals [][]byte, hasGap bool) (*memorydb.Database, bool, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, false, nil
}

// nextKey returns the nextKey that could potentially be part of the response.
func (rb *responseBuilder) nextKey() []byte { _ = "STUB: not implemented"; return nil }

// fillFromTrie iterates key/values from the response builder's trie and appends
// them to the response. Iteration begins from the last key already in the response,
// or the request start if the response is empty. Iteration ends at [end] or if
// the number of leafs reaches the builder's limit.
// Returns true if there are more keys in the trie.
func (rb *responseBuilder) fillFromTrie(ctx context.Context, end []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// create iterator to iterate the trie

// if we're at the end, break this loop

// If we've returned enough data or run out of time, set the more flag and exit
// this flag will determine if the proof is generated or not

// append key/vals to the response

// readLeafsFromSnapshot iterates the storage snapshot of the requested account
// (or the main account trie if account is empty). Returns up to [rb.limit] key/value
// pairs for keys that are in the request's range (inclusive).
func (rb *responseBuilder) readLeafsFromSnapshot(ctx context.Context) ([][]byte, [][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Get an iterator into the storage or the main account snapshot.

// if we're at the end, break this loop

// If we've returned enough data or run out of time, set the more flag and exit
// this flag will determine if the proof is generated or not
