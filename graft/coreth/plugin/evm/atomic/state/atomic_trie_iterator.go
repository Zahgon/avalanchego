// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"errors"

	"github.com/ava-labs/libevm/trie"

	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
)

var errKeyLength = errors.New("atomic trie key length invalid")

type atomicTrieIterator struct {
	trieIterator *trie.Iterator // underlying trie.Iterator
	codec        codec.Manager
	key          []byte
	atomicOps    *atomic.Requests // atomic operation entries at this iteration
	blockchainID ids.ID           // blockchain ID
	blockNumber  uint64           // block number at this iteration
	err          error            // error if any has occurred
}

func NewAtomicTrieIterator(trieIterator *trie.Iterator, codec codec.Manager) *atomicTrieIterator {
	_ = "STUB: not implemented"
	return nil
}

// Error returns error, if any encountered during this iteration
func (a *atomicTrieIterator) Error() error {
	_ = "STUB: not implemented"

	// Next returns whether there are more nodes to iterate over
	// On success, this function sets the blockNumber and atomicOps fields
	// In case of an error during this iteration, it sets the error value and resets the above fields.
	// It is the responsibility of the caller to check the result of Error() after an iterator reports
	// having no more elements to iterate.
	return nil
}

func (a *atomicTrieIterator) Next() bool { _ = "STUB: not implemented"; return false }

// if the underlying iterator has data to iterate over, parse and set the fields
// key is [blockNumberBytes]+[blockchainIDBytes] = 8+32=40 bytes

// If the key has an unexpected length, set the error and stop the iteration since the data is
// no longer reliable.

// The value in the iterator should be the atomic requests serialized the the codec.

// Success, update the struct fields

// trieIterator.Key is already newly allocated so copy is not needed here

// resetFields resets the value fields of the iterator to their nil values and sets the error value to [err].
func (a *atomicTrieIterator) resetFields(err error) { _ = "STUB: not implemented"; return }

// BlockNumber returns the current block number
func (a *atomicTrieIterator) BlockNumber() uint64 { _ = "STUB: not implemented"; return 0 }

// BlockchainID returns the current blockchain ID at the current block number
func (a *atomicTrieIterator) BlockchainID() ids.ID {
	_ = "STUB: not implemented"
	return *

	// AtomicOps returns atomic requests for the blockchainID at the current block number
	// returned object can be freely modified
	new(ids.ID)
}

func (a *atomicTrieIterator) AtomicOps() *atomic.Requests {
	_ = "STUB: not implemented"

	// Key returns the current database key that the iterator is iterating
	// returned []byte can be freely modified
	return nil
}

func (a *atomicTrieIterator) Key() []byte {
	_ = "STUB: not implemented"

	// Value returns the current database value that the iterator is iterating
	return nil
}

func (a *atomicTrieIterator) Value() []byte { _ = "STUB: not implemented"; return nil }
