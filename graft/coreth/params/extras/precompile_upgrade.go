// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package extras

import (
	"errors"

	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/coreth/precompile/precompileconfig"

	ethparams "github.com/ava-labs/libevm/params"
)

var errNoKey = errors.New("PrecompileUpgrade cannot be empty")

// PrecompileUpgrade is a helper struct embedded in UpgradeConfig.
// It is used to unmarshal the json into the correct precompile config type
// based on the key. Keys are defined in each precompile module, and registered in
// precompile/registry/registry.go.
type PrecompileUpgrade struct {
	precompileconfig.Config
}

// UnmarshalJSON unmarshals the json into the correct precompile config type
// based on the key. Keys are defined in each precompile module, and registered in
// precompile/registry/registry.go.
// Ex: {"feeManagerConfig": {...}} where "feeManagerConfig" is the key
func (u *PrecompileUpgrade) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON marshal the precompile config into json based on the precompile key.
// Ex: {"feeManagerConfig": {...}} where "feeManagerConfig" is the key
func (u *PrecompileUpgrade) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// verifyPrecompileUpgrades checks [c.PrecompileUpgrades] is well formed:
//   - [upgrades] must specify exactly one key per PrecompileUpgrade
//   - the specified blockTimestamps must monotonically increase
//   - the specified blockTimestamps must be compatible with those
//     specified in the chainConfig by genesis.
//   - check a precompile is disabled before it is re-enabled
func (c *ChainConfig) verifyPrecompileUpgrades() error {
	_ = "STUB: not implemented"
	// Store this struct to keep track of the last upgrade for each precompile key.
	// Required for timestamp and disabled checks.
	return nil
}

// next range over upgrades to verify correct use of disabled and blockTimestamps.
// previousUpgradeTimestamp is used to verify monotonically increasing timestamps.

// lastUpgradeByKey is the previous processed upgrade for this precompile key.

// Verify specified timestamps are monotonically increasing across all precompile keys.
// Note: It is OK for multiple configs of DIFFERENT keys to specify the same timestamp.

// Verify specified timestamps are monotonically increasing across same precompile keys.
// Note: It is NOT OK for multiple configs of the SAME key to specify the same timestamp.

// GetActivePrecompileConfig returns the most recent precompile config corresponding to [address].
// If none have occurred, returns nil.
func (c *ChainConfig) GetActivePrecompileConfig(address common.Address, timestamp uint64) precompileconfig.Config {
	_ = "STUB: not implemented"
	return *new(precompileconfig.Config)
}

// return the most recent config

// GetActivatingPrecompileConfigs returns all precompile upgrades configured to activate during the
// state transition from a block with timestamp [from] to a block with timestamp [to].
func (*ChainConfig) GetActivatingPrecompileConfigs(address common.Address, from *uint64, to uint64, upgrades []PrecompileUpgrade) []precompileconfig.Config {
	_ = "STUB: not implemented"
	// Get key from address.
	return nil
}

// Loop over all upgrades checking for the requested precompile config.

// Check if the precompile activates in the specified range.

// checkPrecompilesCompatible checks if [precompileUpgrades] are compatible with [c] at [headTimestamp].
// Returns a ConfigCompatError if upgrades already activated at [headTimestamp] are missing from
// [precompileUpgrades]. Upgrades not already activated may be modified or absent from [precompileUpgrades].
// Returns nil if [precompileUpgrades] is compatible with [c].
// Assumes given timestamp is the last accepted block timestamp.
// This ensures that as long as the node has not accepted a block with a different rule set it will allow a
// new upgrade to be applied as long as it activates after the last accepted block.
//
//nolint:unused
func (c *ChainConfig) checkPrecompilesCompatible(precompileUpgrades []PrecompileUpgrade, time uint64) *ethparams.ConfigCompatError {
	_ = "STUB: not implemented"
	return nil
}

// checkPrecompileCompatible verifies that the precompile specified by [address] is compatible between [c]
// and [precompileUpgrades] at [headTimestamp].
// Returns an error if upgrades already activated at [headTimestamp] are missing from [precompileUpgrades].
// Upgrades that have already gone into effect cannot be modified or absent from [precompileUpgrades].
//
//nolint:unused
func (c *ChainConfig) checkPrecompileCompatible(address common.Address, precompileUpgrades []PrecompileUpgrade, time uint64) *ethparams.ConfigCompatError {
	_ = "STUB: not implemented"
	// All active upgrades (from nil to [lastTimestamp]) must match.
	return nil
}

// Check activated upgrades are still present.

// missing upgrade

// All upgrades that have activated must be identical.

// then, make sure newUpgrades does not have additional upgrades
// that are already activated. (cannot perform retroactive upgrade)

// this indexes to the first element in newUpgrades after the end of activeUpgrades

// EnabledStatefulPrecompiles returns current stateful precompile configs that are enabled at [blockTimestamp].
func (c *ChainConfig) EnabledStatefulPrecompiles(blockTimestamp uint64) Precompiles {
	_ = "STUB: not implemented"
	return *new(Precompiles)
}
