// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"errors"

	"github.com/ava-labs/avalanchego/ids"

	pchaintxs "github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

const (
	configChainIDAlias = "X"
)

var (
	errStakeDurationTooHigh            = errors.New("initial stake duration larger than maximum configured")
	errNoInitiallyStakedFunds          = errors.New("initial staked funds cannot be empty")
	errNoSupply                        = errors.New("initial supply must be > 0")
	errNoStakeDuration                 = errors.New("initial stake duration must be > 0")
	errNoStakers                       = errors.New("initial stakers must be > 0")
	errNoCChainGenesis                 = errors.New("C-Chain genesis cannot be empty")
	errNoTxs                           = errors.New("genesis creates no transactions")
	errNoAllocationToStake             = errors.New("no allocation to stake")
	errDuplicateInitiallyStakedAddress = errors.New("duplicate initially staked address")
	errConflictingNetworkIDs           = errors.New("conflicting networkIDs")
	errFutureStartTime                 = errors.New("startTime cannot be in the future")
	errInitialStakeDurationTooLow      = errors.New("initial stake duration is too low")
	errOverridesStandardNetworkConfig  = errors.New("overrides standard network genesis config")
	errAllocationsLockedAmountTooLow   = errors.New("total allocations locked amount is too low")
)

// validateInitialStakedFunds ensures all staked
// funds have allocations and that all staked
// funds are unique.
//
// This function assumes that NetworkID in *Config has already
// been checked for correctness.
func validateInitialStakedFunds(config *Config) error { _ = "STUB: not implemented"; return nil }

// It is ok to have duplicates as different
// ethAddrs could claim to the same avaxAddr.

// validateAllocationsLockedAmount ensures that the sum of all locked
// allocation amounts is at least the number of initial stakers.
func validateAllocationsLockedAmount(config *Config) error { _ = "STUB: not implemented"; return nil }

// validateConfig returns an error if the provided
// *Config is not considered valid.
func validateConfig(networkID uint32, config *Config, stakingCfg *StakingConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// We don't impose any restrictions on the minimum
// stake duration to enable complex testing configurations
// but recommend setting a minimum duration of at least
// 15 minutes.

// Initial stake duration of genesis validators must be
// not larger than maximal stake duration specified for any validator.

// FromFile returns the genesis data of the Platform Chain.
//
// Since an Avalanche network has exactly one Platform Chain, and the Platform
// Chain defines the genesis state of the network (who is staking, which chains
// exist, etc.), defining the genesis state of the Platform Chain is the same as
// defining the genesis state of the network.
//
// FromFile accepts:
// 1) The ID of the new network. [networkID]
// 2) The location of a custom genesis config to load. [filepath]
//
// If [filepath] is empty or the given network ID is Mainnet, Testnet, or Local, returns error.
// If [filepath] is non-empty and networkID isn't Mainnet, Testnet, or Local,
// loads the network genesis data from the config at [filepath].
//
// FromFile returns:
//
//  1. The byte representation of the genesis state of the platform chain
//     (ie the genesis state of the network)
//  2. The asset ID of AVAX
func FromFile(networkID uint32, filepath string, stakingCfg *StakingConfig) ([]byte, ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ID), nil
}

// FromFlag returns the genesis data of the Platform Chain.
//
// Since an Avalanche network has exactly one Platform Chain, and the Platform
// Chain defines the genesis state of the network (who is staking, which chains
// exist, etc.), defining the genesis state of the Platform Chain is the same as
// defining the genesis state of the network.
//
// FromFlag accepts:
// 1) The ID of the new network. [networkID]
// 2) The content of a custom genesis config to load. [genesisContent]
//
// If [genesisContent] is empty or the given network ID is Mainnet, Testnet, or Local, returns error.
// If [genesisContent] is non-empty and networkID isn't Mainnet, Testnet, or Local,
// loads the network genesis data from [genesisContent].
//
// FromFlag returns:
//
//  1. The byte representation of the genesis state of the platform chain
//     (ie the genesis state of the network)
//  2. The asset ID of AVAX
func FromFlag(networkID uint32, genesisContent string, stakingCfg *StakingConfig) ([]byte, ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ID), nil
}

// FromConfig returns:
//
//  1. The byte representation of the genesis state of the platform chain
//     (ie the genesis state of the network)
//  2. The asset ID of AVAX
func FromConfig(config *Config) ([]byte, ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ID), nil
}

// Specify the genesis state of the AVM

// The AVM starts out with one asset: AVAX

// Build UTXOs for the Platform Chain

// Build validators for the Platform Chain

// Specify the chains that exist upon this network's creation

func splitAllocations(allocations []Allocation, numSplits int) [][]Allocation {
	_ = "STUB: not implemented"
	return nil
}

// Already added to the X-chain

// Going to be added until the correct amount is reached

// Already added to the X-chain

// Going to be added until the correct amount is reached

func VMGenesis(genesisBytes []byte, vmID ids.ID) (*pchaintxs.Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AVAXAssetID(avmGenesisBytes []byte) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}
