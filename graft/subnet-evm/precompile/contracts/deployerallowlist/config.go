// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package deployerallowlist

import (
	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompileconfig"
)

var _ precompileconfig.Config = (*Config)(nil)

// Config contains the configuration for the ContractDeployerAllowList precompile,
// consisting of the initial allowlist and the timestamp for the network upgrade.
type Config struct {
	allowlist.AllowListConfig
	precompileconfig.Upgrade
}

// NewConfig returns a config for a network upgrade at [blockTimestamp] that enables
// ContractDeployerAllowList with [admins], [enableds] and [managers] as members of the allowlist.
func NewConfig(blockTimestamp *uint64, admins []common.Address, enableds []common.Address, managers []common.Address) *Config {
	_ = "STUB: not implemented"
	return nil
}

// NewDisableConfig returns config for a network upgrade at [blockTimestamp]
// that disables ContractDeployerAllowList.
func NewDisableConfig(blockTimestamp *uint64) *Config { _ = "STUB: not implemented"; return nil }

func (*Config) Key() string {
	_ = "STUB: not implemented"

	// Equal returns true if [cfg] is a [*ContractDeployerAllowListConfig] and it has been configured identical to [c].
	return ""
}

func (c *Config) Equal(cfg precompileconfig.Config) bool {
	_ = "STUB: not implemented"
	// typecast before comparison
	return false
}

func (c *Config) Verify(chainConfig precompileconfig.ChainConfig) error {
	_ = "STUB: not implemented"
	return nil
}
