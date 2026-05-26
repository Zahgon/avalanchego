// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nativeminter

import (
	"errors"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/math"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompileconfig"
)

var (
	_ precompileconfig.Config = (*Config)(nil)

	ErrInitialMintNilAmount     = errors.New("initial mint cannot contain nil amount")
	ErrInitialMintInvalidAmount = errors.New("initial mint cannot contain invalid amount")
)

// Config implements the precompileconfig.Config interface while adding in the
// ContractNativeMinter specific precompile config.
type Config struct {
	allowlist.AllowListConfig
	precompileconfig.Upgrade
	InitialMint map[common.Address]*math.HexOrDecimal256 `json:"initialMint,omitempty"` // addresses to receive the initial mint mapped to the amount to mint
}

// NewConfig returns a config for a network upgrade at [blockTimestamp] that enables
// ContractNativeMinter with the given [admins], [enableds] and [managers] as members of the allowlist.
// Also mints balances according to [initialMint] when the upgrade activates.
func NewConfig(blockTimestamp *uint64, admins []common.Address, enableds []common.Address, managers []common.Address, initialMint map[common.Address]*math.HexOrDecimal256) *Config {
	_ = "STUB: not implemented"
	return nil
}

// NewDisableConfig returns config for a network upgrade at [blockTimestamp]
// that disables ContractNativeMinter.
func NewDisableConfig(blockTimestamp *uint64) *Config { _ = "STUB: not implemented"; return nil }

// Key returns the key for the ContractNativeMinter precompileconfig.
// This should be the same key as used in the precompile module.
func (*Config) Key() string {
	_ = "STUB: not implemented"

	// Equal returns true if [cfg] is a [*ContractNativeMinterConfig] and it has been configured identical to [c].
	return ""
}

func (c *Config) Equal(cfg precompileconfig.Config) bool {
	_ = "STUB: not implemented"
	// typecast before comparison
	return false
}

func (c *Config) Verify(chainConfig precompileconfig.ChainConfig) error {
	_ = "STUB: not implemented"
	// ensure that all of the initial mint values in the map are non-nil positive values
	return nil
}
