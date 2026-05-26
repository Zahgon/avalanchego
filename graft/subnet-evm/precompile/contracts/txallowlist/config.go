// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txallowlist

import (
	"github.com/ava-labs/libevm/common"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompileconfig"
)

var _ precompileconfig.Config = (*Config)(nil)

// Config implements the StatefulPrecompileConfig interface while adding in the
// TxAllowList specific precompile config.
type Config struct {
	allowlist.AllowListConfig
	precompileconfig.Upgrade
}

// NewConfig returns a config for a network upgrade at [blockTimestamp] that enables
// TxAllowList with the given [admins], [enableds] and [managers] as members of the allowlist.
func NewConfig(blockTimestamp *uint64, admins []common.Address, enableds []common.Address, managers []common.Address) *Config {
	_ = "STUB: not implemented"
	return nil
}

// NewDisableConfig returns config for a network upgrade at [blockTimestamp]
// that disables TxAllowList.
func NewDisableConfig(blockTimestamp *uint64) *Config { _ = "STUB: not implemented"; return nil }

func (*Config) Key() string {
	_ = "STUB: not implemented"

	// Equal returns true if [cfg] is a [*TxAllowListConfig] and it has been configured identical to [c].
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
