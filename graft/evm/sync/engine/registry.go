// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package engine

import (
	"context"
	"errors"

	"golang.org/x/sync/errgroup"

	"github.com/ava-labs/avalanchego/graft/evm/message"
	"github.com/ava-labs/avalanchego/graft/evm/sync/types"
)

var errSyncerAlreadyRegistered = errors.New("syncer already registered")

// syncerTask represents a single syncer with its name for identification.
type syncerTask struct {
	name   string
	syncer types.Syncer
}

// SyncerRegistry manages a collection of syncers for sequential execution.
type SyncerRegistry struct {
	syncers       []syncerTask
	registeredIDs map[string]bool // Track registered IDs to prevent duplicates.
}

// NewSyncerRegistry creates a new empty syncer registry.
func NewSyncerRegistry() *SyncerRegistry { _ = "STUB: not implemented"; return nil }

// Register adds a syncer to the registry.
// Returns an error if a syncer with the same name is already registered.
func (r *SyncerRegistry) Register(syncer types.Syncer) error { _ = "STUB: not implemented"; return nil }

// RunSyncerTasks executes all registered syncers synchronously.
func (r *SyncerRegistry) RunSyncerTasks(ctx context.Context, summary message.Syncable) error {
	_ = "STUB: not implemented"
	// Ensure finalization runs regardless of how this function exits.
	// This guarantees cleanup even on early returns or panics.
	return nil
}

// Early return if context is already canceled (e.g., during shutdown).

// StartAsync launches all registered syncers and returns an [errgroup.Group]
// whose Wait() completes when all syncers exit. The context returned will be
// cancelled when any syncer fails, propagating shutdown to the others.
func (r *SyncerRegistry) StartAsync(ctx context.Context, summary message.Syncable) *errgroup.Group {
	_ = "STUB: not implemented"
	return nil
}

// Context cancellation during shutdown is expected.

// FinalizeAll iterates over all registered syncers and calls Finalize on those that implement the Finalizer interface.
// Errors are logged but not returned to ensure best-effort cleanup of all syncers.
func (r *SyncerRegistry) FinalizeAll(summary message.Syncable) { _ = "STUB: not implemented"; return }
