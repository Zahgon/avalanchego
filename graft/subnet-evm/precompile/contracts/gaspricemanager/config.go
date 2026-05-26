// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gaspricemanager

import (
	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/commontype"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompileconfig"
)

var _ precompileconfig.Config = (*Config)(nil)

// Config is the configuration for the gas price manager precompile.
// It specifies:
//   - when the precompile is activated ([precompileconfig.Upgrade])
//   - who may call it ([allowlist.AllowListConfig])
//   - an optional initial gas price config ([commontype.GasPriceConfig]) to write to contract storage on activation.
type Config struct {
	allowlist.AllowListConfig
	precompileconfig.Upgrade
	InitialGasPriceConfig *commontype.GasPriceConfig `json:"initialGasPriceConfig,omitempty"` // activated immediately on precompile enable if provided
}

// NewConfig returns a config that enables GasPriceManager at `blockTimestamp`.
func NewConfig(blockTimestamp *uint64, admins []common.Address, enabled []common.Address, managers []common.Address, initialConfig *commontype.GasPriceConfig) *Config {
	_ = "STUB: not implemented"
	return nil
}

// NewDisableConfig returns a config that disables GasPriceManager at `blockTimestamp`.
func NewDisableConfig(blockTimestamp *uint64) *Config { _ = "STUB: not implemented"; return nil }

// Key must match ConfigKey used in the precompile module.
func (*Config) Key() string {
	_ = "STUB: not implemented"

	// Equal returns true if [cfg] is a *Config identical to [c].
	return ""
}

func (c *Config) Equal(cfg precompileconfig.Config) bool { _ = "STUB: not implemented"; return false }

// Verify validates the allow list config and, if set, the initial gas price config.
func (c *Config) Verify(chainConfig precompileconfig.ChainConfig) error {
	_ = "STUB: not implemented"
	return nil
}
