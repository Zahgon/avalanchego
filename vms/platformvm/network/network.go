// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package network

import (
	"context"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/network/p2p/gossip"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/platformvm/config"
	"github.com/ava-labs/avalanchego/vms/platformvm/state"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs/mempool"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
)

type Network struct {
	*p2p.Network

	log                       logging.Logger
	mempool                   *gossipMempool
	partialSyncPrimaryNetwork bool

	txPushGossiper        *gossip.PushGossiper[*txs.Tx]
	txPushGossipFrequency time.Duration
	txPullGossiper        gossip.Gossiper
	txPullGossipFrequency time.Duration
	peers                 *p2p.Peers
}

func New(
	log logging.Logger,
	nodeID ids.NodeID,
	subnetID ids.ID,
	vdrs validators.State,
	txVerifier TxVerifier,
	mempool *mempool.Mempool,
	partialSyncPrimaryNetwork bool,
	appSender common.AppSender,
	stateLock sync.Locker,
	state state.Chain,
	signer warp.Signer,
	registerer prometheus.Registerer,
	config config.Network,
) (*Network, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We allow all peers to request warp messaging signatures

func (n *Network) PushGossip(ctx context.Context) { _ = "STUB: not implemented"; return }

func (n *Network) PullGossip(ctx context.Context) {
	_ = "STUB: not implemented"
	// If the node is running partial sync, we do not perform any pull gossip
	// because we should never be a validator.
	return
}

func (n *Network) AppGossip(ctx context.Context, nodeID ids.NodeID, msgBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Network) IssueTxFromRPC(tx *txs.Tx) error { _ = "STUB: not implemented"; return nil }

func (n *Network) Peers() *p2p.Peers { _ = "STUB: not implemented"; return nil }
