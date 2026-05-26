// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/constants"
	"github.com/ava-labs/avalanchego/vms/nftfx"
	"github.com/ava-labs/avalanchego/vms/propertyfx"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

var (
	PChainAliases = []string{"P", "platform"}
	XChainAliases = []string{"X", "avm"}
	CChainAliases = []string{"C", "evm"}
	VMAliases     = map[ids.ID][]string{
		constants.PlatformVMID: {"platform"},
		constants.AVMID:        {"avm"},
		constants.EVMID:        {"evm"},
		secp256k1fx.ID:         {"secp256k1fx"},
		nftfx.ID:               {"nftfx"},
		propertyfx.ID:          {"propertyfx"},
	}
)

// Aliases returns the default aliases based on the network ID
func Aliases(genesisBytes []byte) (map[string][]string, map[ids.ID][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// TODO let's not re-create genesis to do aliasing
