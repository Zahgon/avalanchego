// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evmstate

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/libevm/options"
	"github.com/ava-labs/libevm/triedb"

	"github.com/ava-labs/avalanchego/graft/evm/core/state/snapshot"
	"github.com/ava-labs/avalanchego/graft/evm/message"
	"github.com/ava-labs/avalanchego/graft/evm/sync/client"
	"github.com/ava-labs/avalanchego/graft/evm/sync/code"
	"github.com/ava-labs/avalanchego/graft/evm/sync/leaf"
	"github.com/ava-labs/avalanchego/graft/evm/sync/types"
)

const (
	segmentThreshold       = 500_000 // if we estimate trie to have greater than this number of leafs, split it
	numStorageTrieSegments = 4
	numMainTrieSegments    = 8
	defaultNumWorkers      = 8
)

var (
	_                           types.Syncer = (*stateSync)(nil)
	errCodeRequestQueueRequired              = errors.New("code request queue is required")
	errLeafsRequestSizeRequired              = errors.New("leafs request size must be > 0")
)

// stateSync keeps the state of the entire state sync operation.
type stateSync struct {
	db               ethdb.Database            // database we are syncing
	root             common.Hash               // root of the EVM state we are syncing to
	trieDB           *triedb.Database          // trieDB on top of db we are syncing. used to restore any existing tries.
	snapshot         snapshot.SnapshotIterable // used to access the database we are syncing as a snapshot.
	batchSize        uint                      // write batches when they reach this size
	leafsRequestType message.LeafsRequestType  // type of leafs request to use (coreth or subnet-evm wire format)
	segments         chan leaf.SyncTask        // channel of tasks to sync
	syncer           *leaf.CallbackSyncer      // performs the sync, looping over each task's range and invoking specified callbacks
	codeQueue        *code.Queue               // queue that manages the asynchronous download and batching of code hashes
	trieQueue        *trieQueue                // manages a persistent list of storage tries we need to sync and any segments that are created for them

	// track the main account trie specifically to commit its root at the end of the operation
	mainTrie *trieToSync

	// track the tries currently being synced
	lock            sync.RWMutex
	triesInProgress map[common.Hash]*trieToSync

	// track completion and progress of work
	mainTrieDone       chan struct{}
	storageTriesDone   chan struct{}
	triesInProgressSem chan struct{}
	stats              *trieSyncStats

	// syncCompleted is set to true when the sync completes successfully.
	// This provides an explicit success signal for Finalize().
	syncCompleted atomic.Bool
}

// SyncerOption configures the state syncer via functional options.
type SyncerOption = options.Option[stateSync]

// WithBatchSize sets the database batch size for writes.
func WithBatchSize(n uint) SyncerOption { _ = "STUB: not implemented"; return *new(SyncerOption) }

func NewSyncer(client client.Client, db ethdb.Database, root common.Hash, codeQueue *code.Queue, leafsRequestSize uint16, leafsRequestType message.LeafsRequestType, opts ...SyncerOption) (types.Syncer, error) {
	_ = "STUB: not implemented"
	return *new(types.Syncer), nil
}

// Construct with defaults, then apply options directly to stateSync.

// [triesInProgressSem] is used to keep the number of tries syncing
// less than or equal to [defaultNumWorkers].

// Each [trieToSync] will have a maximum of [numSegments] segments.
// We set the capacity of [segments] such that [defaultNumWorkers]
// storage tries can sync concurrently.

// Apply functional options.

// create a trieToSync for the main trie and mark it as in progress.

// Use context.Background() for initialization since we don't have a sync context yet.
// This is safe because startSyncing is called before Sync() starts.

// Name returns the human-readable name for this sync task.
func (*stateSync) Name() string { _ = "STUB: not implemented"; return "" }

// ID returns the stable identifier for this sync task.
func (*stateSync) ID() string { _ = "STUB: not implemented"; return "" }

func (t *stateSync) Sync(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Start the leaf syncer and storage trie producer.
	return nil
}

// Note: code fetcher should already be initialized.

// The errgroup wait will take care of returning the first error that occurs, or returning
// nil if syncing finish without an error.

// onStorageTrieFinished is called after a storage trie finishes syncing.
func (t *stateSync) onStorageTrieFinished(root common.Hash) error {
	_ = "STUB: not implemented"
	return nil
	// allow another trie to start (release the semaphore)
}

// mark the storage trie as done in trieQueue

// track the completion of this storage trie

// when the last storage trie finishes, close the segments channel

// onMainTrieFinished is called after the main trie finishes syncing.
func (t *stateSync) onMainTrieFinished() error { _ = "STUB: not implemented"; return nil }

// count the number of storage tries we need to sync for eta purposes.

// mark the main trie done

// onSyncComplete is called after the account trie and
// all storage tries have completed syncing. We persist
// [mainTrie]'s batch last to avoid persisting the state
// root before all storage tries are done syncing.
func (t *stateSync) onSyncComplete() error { _ = "STUB: not implemented"; return nil }

// storageTrieProducer waits for the main trie to finish
// syncing then starts to add storage trie roots along
// with their corresponding accounts to the segments channel.
// returns nil if all storage tries were iterated and an
// error if one occurred or the context expired.
func (t *stateSync) storageTrieProducer(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Wait for main trie to finish to ensure when this thread terminates
	// there are no more storage tries to sync
	return nil
}

// check ctx here to exit the loop early

// If there are no storage tries, then root will be the empty hash on the first pass.

// acquire semaphore (to keep number of tries in progress limited)

// Arbitrarily use the first account for making requests to the server.
// Note: getNextTrie guarantees that if a non-nil storage root is returned, then the
// slice of account hashes is non-empty.

// create a trieToSync for the storage trie and mark it as in progress.

// start syncing after tracking the trie as in progress

// addTrieInProgress tracks the root as being currently synced.
func (t *stateSync) addTrieInProgress(root common.Hash, trie *trieToSync) {
	_ = "STUB: not implemented"
	return
}

// removeTrieInProgress removes root from the set of tracked tries in progress
// and returns the number of tries in progress after the removal.
func (t *stateSync) removeTrieInProgress(root common.Hash) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Finalize flushes in-progress trie batches to disk to preserve progress on failure.
func (t *stateSync) Finalize() error { _ = "STUB: not implemented"; return nil }
