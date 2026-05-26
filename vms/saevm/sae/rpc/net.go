// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"github.com/ava-labs/libevm/common/hexutil"

	"github.com/ava-labs/avalanchego/network/p2p"
)

// netAPI offers the `net` RPCs.
type netAPI struct {
	peers   *p2p.Peers
	chainID string
}

func newNetAPI(peers *p2p.Peers, chainID uint64) *netAPI { _ = "STUB: not implemented"; return nil }

func (*netAPI) Listening() bool {
	_ = "STUB: not implemented"
	// The node is always listening for p2p connections.
	return false
}

func (s *netAPI) PeerCount() hexutil.Uint { _ = "STUB: not implemented"; return *new(hexutil.Uint) }

// Peers includes ourself, so we subtract one.

func (s *netAPI) Version() string { _ = "STUB: not implemented"; return "" }
