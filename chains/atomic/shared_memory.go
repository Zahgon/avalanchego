// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomic

import (
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
)

var _ SharedMemory = (*sharedMemory)(nil)

type Requests struct {
	RemoveRequests [][]byte   `serialize:"true"`
	PutRequests    []*Element `serialize:"true"`

	peerChainID ids.ID
}

type Element struct {
	Key    []byte   `serialize:"true"`
	Value  []byte   `serialize:"true"`
	Traits [][]byte `serialize:"true"`
}

type SharedMemory interface {
	// Get fetches the values corresponding to [keys] that have been sent from
	// [peerChainID]
	//
	// Invariant: Get guarantees that the resulting values array is the same
	//            length as keys.
	Get(peerChainID ids.ID, keys [][]byte) (values [][]byte, err error)
	// Indexed returns a paginated result of values that possess any of the
	// given traits and were sent from [peerChainID].
	Indexed(
		peerChainID ids.ID,
		traits [][]byte,
		startTrait,
		startKey []byte,
		limit int,
	) (
		values [][]byte,
		lastTrait,
		lastKey []byte,
		err error,
	)
	// Apply performs the requested set of operations by atomically applying
	// [requests] to their respective chainID keys in the map along with the
	// batches on the underlying DB.
	//
	// Invariant: The underlying database of [batches] must be the same as the
	//            underlying database for SharedMemory.
	Apply(requests map[ids.ID]*Requests, batches ...database.Batch) error
}

// sharedMemory provides the API for a blockchain to interact with shared memory
// of another blockchain
type sharedMemory struct {
	m           *Memory
	thisChainID ids.ID
}

func (sm *sharedMemory) Get(peerChainID ids.ID, keys [][]byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sm *sharedMemory) Indexed(
	peerChainID ids.ID,
	traits [][]byte,
	startTrait,
	startKey []byte,
	limit int,
) ([][]byte, []byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (sm *sharedMemory) Apply(requests map[ids.ID]*Requests, batches ...database.Batch) error {
	_ = "STUB: not implemented"
	// Sorting here introduces an ordering over the locks to prevent any
	// deadlocks
	return nil
}

// Make sure all operations are committed atomically

// Perform any remove requests on the inbound database

// Add Put requests to the outbound database.

// Commit the operations on shared memory atomically with the contents of
// [batches].
