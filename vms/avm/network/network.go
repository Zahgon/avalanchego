// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/network/p2p/gossip"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/avm/txs"
	"github.com/ava-labs/avalanchego/vms/txs/mempool"
)

var (
	_ common.AppHandler    = (*Network)(nil)
	_ validators.Connector = (*Network)(nil)
)

type Network struct {
	*p2p.Network

	log     logging.Logger
	mempool *gossipMempool

	txPushGossiper        *gossip.PushGossiper[*txs.Tx]
	txPushGossipFrequency time.Duration
	txPullGossiper        gossip.Gossiper
	txPullGossipFrequency time.Duration
}

func New(
	log logging.Logger,
	nodeID ids.NodeID,
	subnetID ids.ID,
	vdrs validators.State,
	parser txs.Parser,
	txVerifier TxVerifier,
	mempool mempool.Mempool[*txs.Tx],
	appSender common.AppSender,
	registerer prometheus.Registerer,
	config Config,
) (*Network, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *Network) PushGossip(ctx context.Context) { _ = "STUB: not implemented"; return }

func (n *Network) PullGossip(ctx context.Context) { _ = "STUB: not implemented"; return }

// IssueTxFromRPC attempts to add a tx to the mempool, after verifying it. If
// the tx is added to the mempool, it will attempt to push gossip the tx to
// random peers in the network.
//
// If the tx is already in the mempool, mempool.ErrDuplicateTx will be
// returned.
// If the tx is not added to the mempool, an error will be returned.
func (n *Network) IssueTxFromRPC(tx *txs.Tx) error { _ = "STUB: not implemented"; return nil }

// IssueTxFromRPCWithoutVerification attempts to add a tx to the mempool,
// without first verifying it. If the tx is added to the mempool, it will
// attempt to push gossip the tx to random peers in the network.
//
// If the tx is already in the mempool, mempool.ErrDuplicateTx will be
// returned.
// If the tx is not added to the mempool, an error will be returned.
func (n *Network) IssueTxFromRPCWithoutVerification(tx *txs.Tx) error {
	_ = "STUB: not implemented"
	return nil
}
