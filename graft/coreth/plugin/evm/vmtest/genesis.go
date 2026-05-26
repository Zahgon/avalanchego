// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vmtest

import (
	"math/big"
	"testing"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/database/prefixdb"
	"github.com/ava-labs/avalanchego/graft/coreth/core"
	"github.com/ava-labs/avalanchego/graft/coreth/params"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/upgrade/ap3"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/upgrade/upgradetest"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"

	avalancheatomic "github.com/ava-labs/avalanchego/chains/atomic"
)

var (
	TestKeys         = secp256k1.TestKeys()[:3]
	TestEthAddrs     []common.Address // testEthAddrs[i] corresponds to testKeys[i]
	TestShortIDAddrs []ids.ShortID
	InitialBaseFee   = big.NewInt(ap3.InitialBaseFee)
	InitialFund      = new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(10))
)

func init() {
	for _, pk := range TestKeys {
		TestEthAddrs = append(TestEthAddrs, pk.EthAddress())
		TestShortIDAddrs = append(TestShortIDAddrs, pk.Address())
	}
}

// GenesisJSON returns the JSON representation of the genesis block
// for the given chain configuration, with pre-funded accounts.
func GenesisJSON(cfg *params.ChainConfig) string { _ = "STUB: not implemented"; return "" }

func NewTestGenesis(cfg *params.ChainConfig) *core.Genesis { _ = "STUB: not implemented"; return nil }

// Use chainId: 43111, so that it does not overlap with any Avalanche ChainIDs, which may have their
// config overridden in vm.Initialize.

// After Durango, an additional account is funded in tests to use
// with warp messages.

// Fund the test keys

func NewPrefundedGenesis(
	balance int,
	addresses ...common.Address,
) *core.Genesis {
	_ = "STUB: not implemented"
	return nil
}

// SetupGenesis sets up the genesis
func SetupGenesis(
	t *testing.T,
	fork upgradetest.Fork,
) (*snow.Context,
	*prefixdb.Database,
	[]byte,
	*avalancheatomic.Memory,
) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// initialize the atomic memory

// NB: this lock is intentionally left locked when this function returns.
// The caller of this function is responsible for unlocking.
