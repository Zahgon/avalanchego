// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package txgossip provides a mempool for [Streaming Asynchronous Execution],
// which is also compatible with AvalancheGo's [gossip] mechanism.
//
// [Streaming Asynchronous Execution]: https://github.com/avalanche-foundation/ACPs/tree/main/ACPs/194-streaming-asynchronous-execution
package txgossip

import (
	"github.com/ava-labs/libevm/core/txpool"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p/gossip"
)

var _ gossip.Gossipable = Transaction{}

// A Transaction is a [gossip.Gossipable] wrapper for a [types.Transaction].
type Transaction struct {
	*types.Transaction
}

// GossipID returns the transaction hash.
func (tx Transaction) GossipID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

var _ gossip.Marshaller[Transaction] = Marshaller{}

// A Marshaller implements [gossip.Marshaller] for [Transaction], based on RLP
// encoding.
type Marshaller struct{}

// MarshalGossip returns the [rlp] encoding of the underlying
// [types.Transaction].
func (Marshaller) MarshalGossip(tx Transaction) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UnmarshalGossip [rlp] decodes the buffer into a [types.Transaction].
func (Marshaller) UnmarshalGossip(buf []byte) (Transaction, error) {
	_ = "STUB: not implemented"
	return *new(Transaction), nil
}

// Set couples a [gossip.BloomSet] with a [txpool.TxPool] that acts as the
// backing for the set.
type Set struct {
	*gossip.BloomSet[Transaction]
	Pool *txpool.TxPool

	set    *txSet
	pushTo []func(...Transaction)
}

// NewSet returns a new Set. Use [gossip.BloomSet.Add] or [Set.SendTx] to add
// transactions to the pool, which SHOULD NOT be populated directly.
func NewSet(pool *txpool.TxPool, config gossip.BloomSetConfig) (*Set, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var _ gossip.Set[Transaction] = (*txSet)(nil)

type txSet struct {
	pool *txpool.TxPool
}

func (s *txSet) Add(tx Transaction) error { _ = "STUB: not implemented"; return nil }

func (s *txSet) addToPool(local bool, txs ...*types.Transaction) []error {
	_ = "STUB: not implemented"
	return nil
}

/*sync*/

func (s *txSet) Has(id ids.ID) bool { _ = "STUB: not implemented"; return false }

func (s *txSet) Iterate(fn func(Transaction) bool) {
	_ = "STUB: not implemented"
	// TODO(arr4n) implement a method on libevm's [txpool.TxPool] that returns
	// a more efficient iterator.
	return
}

func (s *txSet) Len() int { _ = "STUB: not implemented"; return 0 }
