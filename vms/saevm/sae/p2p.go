// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sae

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
)

// newNetwork creates the P2P network with a registered validator set.
func newNetwork(
	snowCtx *snow.Context,
	sender common.AppSender,
	reg *prometheus.Registry,
) (
	*p2p.Network,
	*p2p.Peers,
	*p2p.Validators,
	error,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}
