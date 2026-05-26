// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// TODO: move to network

package evm

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/ava-labs/libevm/core/types"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/graft/coreth/core"
	"github.com/ava-labs/avalanchego/graft/coreth/core/txpool"
	"github.com/ava-labs/avalanchego/graft/coreth/eth"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p/gossip"
	"github.com/ava-labs/avalanchego/utils/bloom"
)

const pendingTxsBuffer = 10

var (
	_ gossip.Gossipable               = (*GossipEthTx)(nil)
	_ gossip.Marshaller[*GossipEthTx] = (*GossipEthTxMarshaller)(nil)
	_ gossip.SystemSet[*GossipEthTx]  = (*GossipEthTxPool)(nil)

	_ eth.PushGossiper = (*EthPushGossiper)(nil)
)

func NewGossipEthTxPool(mempool *txpool.TxPool, registerer prometheus.Registerer) (*GossipEthTxPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GossipEthTxPool struct {
	mempool    *txpool.TxPool
	pendingTxs chan core.NewTxsEvent

	bloom *gossip.BloomFilter
	lock  sync.RWMutex

	// subscribed is set to true when the gossip subscription is active
	// mostly used for testing
	subscribed atomic.Bool
}

// IsSubscribed returns whether or not the gossip subscription is active.
func (g *GossipEthTxPool) IsSubscribed() bool { _ = "STUB: not implemented"; return false }

func (g *GossipEthTxPool) Subscribe(ctx context.Context) { _ = "STUB: not implemented"; return }

// Add enqueues the transaction to the mempool. Subscribe should be called
// to receive an event if tx is actually added to the mempool or not.
func (g *GossipEthTxPool) Add(tx *GossipEthTx) error { _ = "STUB: not implemented"; return nil }

// Has should just return whether or not the [txID] is still in the mempool,
// not whether it is in the mempool AND pending.
func (g *GossipEthTxPool) Has(txID ids.ID) bool { _ = "STUB: not implemented"; return false }

func (g *GossipEthTxPool) Iterate(f func(tx *GossipEthTx) bool) { _ = "STUB: not implemented"; return }

func (g *GossipEthTxPool) BloomFilter() (*bloom.Filter, ids.ID) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ID)
}

type GossipEthTxMarshaller struct{}

func (GossipEthTxMarshaller) MarshalGossip(tx *GossipEthTx) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (GossipEthTxMarshaller) UnmarshalGossip(bytes []byte) (*GossipEthTx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GossipEthTx struct {
	Tx *types.Transaction
}

func (tx *GossipEthTx) GossipID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

// EthPushGossiper is used by the ETH backend to push transactions issued over
// the RPC and added to the mempool to peers.
type EthPushGossiper struct {
	vm *VM
}

func (e *EthPushGossiper) Add(tx *types.Transaction) {
	_ = "STUB: not implemented"
	// eth.Backend is initialized before the [ethTxPushGossiper] is created, so
	// we just ignore any gossip requests until it is set.
	return
}
