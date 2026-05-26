// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package indexer

import (
	"errors"
	"fmt"
	"sync"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
)

// Maximum number of containers IDs that can be fetched at a time in a call to
// GetContainerRange
const MaxFetchedByRange = 1024

var (
	// Maps to the byte representation of the next accepted index
	nextAcceptedIndexKey   = []byte{0x00}
	indexToContainerPrefix = []byte{0x01}
	containerToIDPrefix    = []byte{0x02}
	errNoneAccepted        = errors.New("no containers have been accepted")
	errNumToFetchInvalid   = fmt.Errorf("numToFetch must be in [1,%d]", MaxFetchedByRange)
	errNoContainerAtIndex  = errors.New("no container at index")

	_ snow.Acceptor = (*index)(nil)
)

// index indexes containers in their order of acceptance
//
// Invariant: index is thread-safe.
// Invariant: index assumes that Accept is called, before the container is
// committed to the database of the VM, in the order they were accepted.
type index struct {
	clock mockable.Clock
	lock  sync.RWMutex
	// The index of the next accepted transaction
	nextAcceptedIndex uint64
	// When [baseDB] is committed, writes to [baseDB]
	vDB    *versiondb.Database
	baseDB database.Database
	// Both [indexToContainer] and [containerToIndex] have [vDB] underneath
	// Index --> Container
	indexToContainer database.Database
	// Container ID --> Index
	containerToIndex database.Database
	log              logging.Logger
}

// Create a new thread-safe index.
//
// Invariant: Closes [baseDB] on close.
func newIndex(
	baseDB database.Database,
	log logging.Logger,
	clock mockable.Clock,
) (*index, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get next accepted index from db

// Close this index
func (i *index) Close() error { _ = "STUB: not implemented"; return nil }

// Index that the given transaction is accepted
// Returned error should be treated as fatal; the VM should not commit [containerID]
// or any new containers as accepted.
func (i *index) Accept(ctx *snow.ConsensusContext, containerID ids.ID, containerBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// It may be the case that in a previous run of this node, this index committed [containerID]
// as accepted and then the node shut down before the VM committed [containerID] as accepted.
// In that case, when the node restarts Accept will be called with the same container.
// Make sure we don't index the same container twice in that event.

// Persist index --> Container

// Persist container ID --> index

// Persist next accepted index

// Atomically commit [i.vDB], [i.indexToContainer], [i.containerToIndex] to [i.baseDB]

// Returns the ID of the [index]th accepted container and the container itself.
// For example, if [index] == 0, returns the first accepted container.
// If [index] == 1, returns the second accepted container, etc.
// Returns an error if there is no container at the given index.
func (i *index) GetContainerByIndex(index uint64) (Container, error) {
	_ = "STUB: not implemented"
	return *new(Container), nil
}

// Assumes [i.lock] is held
func (i *index) getContainerByIndex(index uint64) (Container, error) {
	_ = "STUB: not implemented"
	return *new(Container), nil
}

// [indexBytes] is the byte representation of the index to fetch.
// Assumes [i.lock] is held
func (i *index) getContainerByIndexBytes(indexBytes []byte) (Container, error) {
	_ = "STUB: not implemented"
	return *new(Container), nil
}

// GetContainerRange returns the IDs of containers at indices
// [startIndex], [startIndex+1], ..., [startIndex+numToFetch-1].
// [startIndex] should be <= i.lastAcceptedIndex().
// [numToFetch] should be in [0, MaxFetchedByRange]
func (i *index) GetContainerRange(startIndex, numToFetch uint64) ([]Container, error) {
	_ = "STUB: not implemented"
	// Check arguments for validity
	return nil, nil
}

// Calculate the last index we will fetch

// [lastIndex] is always >= [startIndex] so this is safe.
// [numToFetch] is limited to [MaxFetchedByRange] so [containers] is bounded in size.

// Returns database.ErrNotFound if the container is not indexed as accepted
func (i *index) GetIndex(id ids.ID) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func (i *index) GetContainerByID(id ids.ID) (Container, error) {
	_ = "STUB: not implemented"
	return *new(Container), nil
}

// Read index from database

// GetLastAccepted returns the last accepted container.
// Returns an error if no containers have been accepted.
func (i *index) GetLastAccepted() (Container, error) {
	_ = "STUB: not implemented"
	return *new(Container), nil
}

// Assumes i.lock is held
// Returns:
//
//  1. The index of the most recently accepted transaction, or 0 if no
//     transactions have been accepted
//  2. Whether at least 1 transaction has been accepted
func (i *index) lastAcceptedIndex() (uint64, bool) { _ = "STUB: not implemented"; return 0, false }
