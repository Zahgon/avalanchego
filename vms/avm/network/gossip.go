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
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/txs/mempool"
)

var (
	_ gossip.Set[*txs.Tx]        = (*gossipMempool)(nil)
	_ gossip.Marshaller[*txs.Tx] = (*txParser)(nil)
)

// bloomChurnMultiplier is the number used to multiply the size of the mempool
// to determine how large of a bloom filter to create.
const bloomChurnMultiplier = 3

type txParser struct {
	parser txs.Parser
}

func (*txParser) MarshalGossip(tx *txs.Tx) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *txParser) UnmarshalGossip(bytes []byte) (*txs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newGossipMempool(
	mempool mempool.Mempool[*txs.Tx],
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
	mempool.Mempool[*txs.Tx]
	log        logging.Logger
	txVerifier TxVerifier

	lock  sync.RWMutex
	bloom *gossip.BloomFilter
}

// Add is called by the p2p SDK when handling transactions that were pushed to
// us and when handling transactions that were pulled from a peer. If this
// returns a nil error while handling push gossip, the p2p SDK will queue the
// transaction to push gossip as well.
func (g *gossipMempool) Add(tx *txs.Tx) error { _ = "STUB: not implemented"; return nil }

// If the tx is being dropped - just ignore it
//
// TODO: Should we allow re-verification of the transaction even if it
// failed previously?

// Verify the tx at the currently preferred state

func (g *gossipMempool) Has(txID ids.ID) bool { _ = "STUB: not implemented"; return false }

func (g *gossipMempool) AddWithoutVerification(tx *txs.Tx) error {
	_ = "STUB: not implemented"
	return nil
}

func (g *gossipMempool) Iterate(f func(*txs.Tx) bool) { _ = "STUB: not implemented"; return }

func (g *gossipMempool) BloomFilter() (*bloom.Filter, ids.ID) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ID)
}
