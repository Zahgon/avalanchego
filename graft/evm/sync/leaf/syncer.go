// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package leaf

import (
	"context"
	"errors"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/evm/message"
	"github.com/ava-labs/avalanchego/graft/evm/sync/types"
)

var ErrFailedToFetchLeafs = errors.New("failed to fetch leafs")

// SyncTask represents a complete task to be completed by the leaf syncer.
// Note: each SyncTask is processed on its own goroutine and there will
// not be concurrent calls to the callback methods. Implementations should return
// the same value for Root, Account, Start, and NodeType throughout the sync.
// The value returned by End can change between calls to OnLeafs.
type SyncTask interface {
	Root() common.Hash                                      // Root of the trie to sync
	Account() common.Hash                                   // Account hash of the trie to sync (only applicable to storage tries)
	Start() []byte                                          // Starting key to request new leaves
	End() []byte                                            // End key to request new leaves
	NodeType() message.NodeType                             // Specifies the message type (atomic/state trie) for the leaf syncer to send
	OnStart() (bool, error)                                 // Callback when tasks begins, returns true if work can be skipped
	OnLeafs(ctx context.Context, keys, vals [][]byte) error // Callback when new leaves are received from the network
	OnFinish(ctx context.Context) error                     // Callback when there are no more leaves in the trie to sync or when we reach End()
}

type SyncerConfig struct {
	RequestSize      uint16                   // Number of leafs to request from a peer at a time
	NumWorkers       int                      // Number of workers to process leaf sync tasks
	LeafsRequestType message.LeafsRequestType // Type of leafs request to use
}

type CallbackSyncer struct {
	config *SyncerConfig
	client types.LeafClient
	tasks  <-chan SyncTask
}

// NewCallbackSyncer creates a new syncer object to perform leaf sync of tries.
func NewCallbackSyncer(client types.LeafClient, tasks <-chan SyncTask, config *SyncerConfig) *CallbackSyncer {
	_ = "STUB: not implemented"
	return nil
}

// workerLoop reads from [c.tasks] and calls [c.syncTask] until [ctx] is finished
// or [c.tasks] is closed.
func (c *CallbackSyncer) workerLoop(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// syncTask performs [task], requesting the leaves of the trie corresponding to [task.Root]
// starting at [task.Start] and invoking the callbacks as necessary.
func (c *CallbackSyncer) syncTask(ctx context.Context, task SyncTask) error {
	_ = "STUB: not implemented"
	return nil
}

// If [ctx] has finished, return early.

// End is intentionally nil, because VerifyRangeProof does not handle empty responses with non-empty end key.

// resize [leafsResponse.Keys] and [leafsResponse.Vals] in case
// the response includes any keys past [End()].
// Note: We truncate the response here as opposed to sending End
// in the request, as [VerifyRangeProof] does not handle empty
// responses correctly with a non-empty end key for the range.

// If we have completed syncing this task, invoke [OnFinish] and mark the task
// as complete.

// Update start to be one bit past the last returned key for the next request.
// Note: since more was true, this cannot cause an overflow.

// Sync launches [numWorkers] worker goroutines to process LeafSyncTasks from [c.tasks].
func (c *CallbackSyncer) Sync(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Start the worker threads with the desired context.
	return nil
}
