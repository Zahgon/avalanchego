// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package core

import (
	"math/big"

	"github.com/ava-labs/libevm/core/state"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/params"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contract"
)

// ApplyPrecompileActivations checks if any of the precompiles specified by the chain config are enabled or disabled by the block
// transition from `parentTimestamp` to the timestamp set in `blockContext`. If this is the case, it calls [modules.Module]'s Configure
// to apply the necessary state transitions for the upgrade.
// This function is called within genesis setup to configure the starting state for precompiles enabled at genesis.
// In block processing and building, [ApplyUpgrades] is called instead which also applies state upgrades.
func ApplyPrecompileActivations(c *params.ChainConfig, parentTimestamp *uint64, blockContext contract.ConfigurationBlockContext, statedb *state.StateDB) error {
	_ = "STUB: not implemented"
	return nil
}

// Note: [modules.RegisteredModules] returns precompiles sorted by module addresses.
// This ensures:
// - the order we call [modules.Module]'s Configure for each precompile is consistent
// - even if precompiles read/write state other than their own they will observe
//   an identical global state in a deterministic order when they are configured.

// If this transition activates the upgrade, configure the stateful precompile.
// (or deconfigure it if it is being disabled.)

// Calling [state.StateDB]'s Finalise here effectively commits the SelfDestruct call and wipes the contract state.
// This enables re-configuration of the same contract state in the same block.
// Without an immediate Finalise call after the SelfDestruct, a reconfigured precompiled state can be wiped out
// since SelfDestruct will be committed after the reconfiguration.

// Set the nonce of the precompile's address (as is done when a contract is created) to ensure
// that it is marked as non-empty and will not be cleaned up when the statedb is finalized.

// Set the code of the precompile's address to a non-zero length byte slice to ensure that the precompile
// can be called from within Solidity contracts. Solidity adds a check before invoking a contract to ensure
// that it does not attempt to invoke a non-existent contract.

// applyStateUpgrades checks if any of the state upgrades specified by the chain config are activated by the block
// transition from [parentTimestamp] to the timestamp set in [header]. If this is the case, it calls [Configure]
// to apply the necessary state transitions for the upgrade.
func applyStateUpgrades(c *params.ChainConfig, parentTimestamp *uint64, blockContext contract.ConfigurationBlockContext, statedb *state.StateDB) error {
	_ = "STUB: not implemented"
	// Apply state upgrades
	return nil
}

// ApplyUpgrades checks if any of the precompile or state upgrades specified by the chain config are activated by the block
// transition from [parentTimestamp] to the timestamp set in [header]. If this is the case, it calls [Configure]
// to apply the necessary state transitions for the upgrade.
// This function is called:
// - in block processing to update the state when processing a block.
// - in the miner to apply the state upgrades when producing a block.
func ApplyUpgrades(c *params.ChainConfig, parentTimestamp *uint64, blockContext contract.ConfigurationBlockContext, statedb *state.StateDB) error {
	_ = "STUB: not implemented"
	return nil
}

// BlockContext implements [contract.ConfigurationBlockContext].
type BlockContext struct {
	number    *big.Int
	timestamp uint64
}

// NewBlockContext creates a [BlockContext] using the block number
// and block timestamp provided. This function is usually necessary to convert
// a `*types.Block` to be passed as a [contract.ConfigurationBlockContext]
// to [ApplyUpgrades].
func NewBlockContext(number *big.Int, timestamp uint64) *BlockContext {
	_ = "STUB: not implemented"
	return nil
}

func (bc *BlockContext) Number() *big.Int  { _ = "STUB: not implemented"; return nil }
func (bc *BlockContext) Timestamp() uint64 { _ = "STUB: not implemented"; return 0 }
