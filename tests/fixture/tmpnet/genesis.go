// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"errors"
	"math/big"

	"github.com/ava-labs/avalanchego/genesis"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/units"
)

const (
	defaultGasLimit = uint64(100_000_000) // Gas limit is arbitrary

	// Arbitrarily large amount of AVAX to fund keys on the X-Chain for testing
	defaultFundedKeyXChainAmount = 30 * units.MegaAvax
)

var (
	// Arbitrarily large amount of AVAX (10^12) to fund keys on the C-Chain for testing
	defaultFundedKeyCChainAmount = new(big.Int).Exp(big.NewInt(10), big.NewInt(30), nil)

	errNoKeysForGenesis           = errors.New("no keys to fund for genesis")
	errInvalidNetworkIDForGenesis = errors.New("network ID can't be mainnet, testnet or local network ID for genesis")
	errMissingStakersForGenesis   = errors.New("no stakers provided for genesis")
)

// Helper type to simplify configuring X-Chain genesis balances
type XChainBalanceMap map[ids.ShortID]uint64

// Create a genesis struct valid for bootstrapping a test
// network. Note that many of the genesis fields (e.g. reward
// addresses) are randomly generated or hard-coded.
func NewTestGenesis(
	networkID uint32,
	nodes []*Node,
	keysToFund []*secp256k1.PrivateKey,
) (*genesis.UnparsedConfig, error) {
	_ = "STUB: not implemented"
	// Validate inputs
	return nil, nil
}

// Address that controls stake doesn't matter -- generate it randomly

// Ensure the total stake allows a MegaAvax per staker

// The eth address is only needed to link pre-mainnet assets. Until that capability
// becomes necessary for testing, use a bogus address.
//
// Reference: https://github.com/ava-labs/avalanchego/issues/1365#issuecomment-1511508767

// Provides stake to validators

// 1 Week

// 1 year
// 90 minutes

// Ensure pre-funded keys have arbitrary large balances on both chains to support testing

// Set X-Chain balances

// 1 Week

// Define C-Chain genesis

// The rest of the config is set in coreth on VM initialization
// Difficulty is a mandatory field
// This time enables Avalanche upgrades by default

// Returns staker configuration for the given set of nodes.
func stakersForNodes(networkID uint32, nodes []*Node) ([]genesis.UnparsedStaker, error) {
	_ = "STUB: not implemented"
	// Give staking rewards for initial validators to a random address. Any testing of staking rewards
	// will be easier to perform with nodes other than the initial validators since the timing of
	// staking can be more easily controlled.
	return nil, nil
}

// Configure provided nodes as initial stakers
