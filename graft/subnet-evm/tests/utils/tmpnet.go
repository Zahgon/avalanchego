// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
)

var DefaultChainConfig = map[string]any{
	"log-level":         "debug",
	"warp-api-enabled":  true,
	"local-txs-enabled": true,
}

func NewTmpnetNetwork(owner string, nodes []*tmpnet.Node, flags tmpnet.FlagsMap, subnets ...*tmpnet.Subnet) *tmpnet.Network {
	_ = "STUB: not implemented"
	return nil
}

// Create the configuration that will enable creation and access to a
// subnet created on a temporary network.
func NewTmpnetSubnet(name string, genesis []byte, chainConfig map[string]any, nodes ...*tmpnet.Node) *tmpnet.Subnet {
	_ = "STUB: not implemented"
	return nil
}
