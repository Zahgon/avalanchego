// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package extras

import (
	"errors"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/common/math"

	ethparams "github.com/ava-labs/libevm/params"
)

var (
	errStateUpgradeNilTimestamp          = errors.New("state upgrade config block timestamp cannot be nil")
	errStateUpgradeTimestampZero         = errors.New("state upgrade config block timestamp must be greater than 0")
	errStateUpgradeTimestampNotMonotonic = errors.New("state upgrade config block timestamp must be greater than previous timestamp")
)

// StateUpgrade describes the modifications to be made to the state during
// a state upgrade.
type StateUpgrade struct {
	BlockTimestamp *uint64 `json:"blockTimestamp,omitempty"`

	// map from account address to the modification to be made to the account.
	StateUpgradeAccounts map[common.Address]StateUpgradeAccount `json:"accounts"`
}

// StateUpgradeAccount describes the modifications to be made to an account during
// a state upgrade.
type StateUpgradeAccount struct {
	Code          hexutil.Bytes               `json:"code,omitempty"`
	Storage       map[common.Hash]common.Hash `json:"storage,omitempty"`
	BalanceChange *math.HexOrDecimal256       `json:"balanceChange,omitempty"`
}

func (s *StateUpgrade) Equal(other *StateUpgrade) bool { _ = "STUB: not implemented"; return false }

// verifyStateUpgrades checks [c.StateUpgrades] is well formed:
// - the specified blockTimestamps must monotonically increase
func (c *ChainConfig) verifyStateUpgrades() error { _ = "STUB: not implemented"; return nil }

// Verify the upgrade's timestamp is equal 0 (to avoid confusion with genesis).

// Verify specified timestamps are strictly monotonically increasing.

// GetActivatingStateUpgrades returns all state upgrades configured to activate during the
// state transition from a block with timestamp [from] to a block with timestamp [to].
func (*ChainConfig) GetActivatingStateUpgrades(from *uint64, to uint64, upgrades []StateUpgrade) []StateUpgrade {
	_ = "STUB: not implemented"
	return nil
}

// checkStateUpgradesCompatible checks if [stateUpgrades] are compatible with [c] at [headTimestamp].
func (c *ChainConfig) checkStateUpgradesCompatible(stateUpgrades []StateUpgrade, lastTimestamp uint64) *ethparams.ConfigCompatError {
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
