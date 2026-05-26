// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p/gossip"
	"github.com/ava-labs/avalanchego/utils/bloom"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/mempool"
)

var (
	_ gossip.Marshaller[*txs.Tx] = (*txMarshaller)(nil)
	_ gossip.Gossipable          = (*txs.Tx)(nil)
)

// bloomChurnMultiplier is the number used to multiply the size of the mempool
// to determine how large of a bloom filter to create.
const bloomChurnMultiplier = 3

type txMarshaller struct{}

func (txMarshaller) MarshalGossip(tx *txs.Tx) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (txMarshaller) UnmarshalGossip(bytes []byte) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newGossipMempool(
	mempool *mempool.Mempool,
	registerer prometheus.Registerer,
	log logging.Logger,
	txVerifier TxVerifier,
	minTargetElements int,
	targetFalsePositiveProbability,
	resetFalsePositiveProbability float64,
) (*gossipMempool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type gossipMempool struct {
	*mempool.Mempool
	log        logging.Logger
	txVerifier TxVerifier

	lock  sync.RWMutex
	bloom *gossip.BloomFilter
}

func (g *gossipMempool) Add(tx *txs.Tx) error { _ = "STUB: not implemented"; return nil }

// If the tx is being dropped - just ignore it
//
// TODO: Should we allow re-verification of the transaction even if it
// failed previously?

func (g *gossipMempool) Has(txID ids.ID) bool { _ = "STUB: not implemented"; return false }

func (g *gossipMempool) BloomFilter() (*bloom.Filter, ids.ID) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ID)
}
