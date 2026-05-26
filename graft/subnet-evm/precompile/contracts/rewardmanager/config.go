// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Code generated
// This file is a generated precompile contract with stubbed abstract functions.

package rewardmanager

import (
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/contract"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompileconfig"

	"github.com/ava-labs/libevm/common"
)

var _ precompileconfig.Config = (*Config)(nil)

type InitialRewardConfig struct {
	AllowFeeRecipients bool           `json:"allowFeeRecipients"`
	RewardAddress      common.Address `json:"rewardAddress,omitempty"`
}

func (i *InitialRewardConfig) Equal(other *InitialRewardConfig) bool {
	_ = "STUB: not implemented"
	return false
}

func (i *InitialRewardConfig) Verify() error { _ = "STUB: not implemented"; return nil }

func (i *InitialRewardConfig) Configure(state contract.StateDB) {
	_ = "STUB: not implemented"
	// enable allow fee recipients
	return
}

// if reward address is empty and allow fee recipients is false
// then disable rewards

// set reward address

// Config implements the StatefulPrecompileConfig interface while adding in the
// RewardManager specific precompile config.
type Config struct {
	allowlist.AllowListConfig
	precompileconfig.Upgrade
	InitialRewardConfig *InitialRewardConfig `json:"initialRewardConfig,omitempty"`
}

// NewConfig returns a config for a network upgrade at [blockTimestamp] that enables
// RewardManager with the given [admins], [enableds] and [managers] as members of the allowlist with [initialConfig] as initial rewards config if specified.
func NewConfig(blockTimestamp *uint64, admins []common.Address, enableds []common.Address, managers []common.Address, initialConfig *InitialRewardConfig) *Config {
	_ = "STUB: not implemented"
	return nil
}

// NewDisableConfig returns config for a network upgrade at [blockTimestamp]
// that disables RewardManager.
func NewDisableConfig(blockTimestamp *uint64) *Config { _ = "STUB: not implemented"; return nil }

// Key returns the key for the Contract precompileconfig.
// This should be the same key as used in the precompile module.
func (*Config) Key() string {
	_ = "STUB: not implemented"

	// Verify tries to verify Config and returns an error accordingly.
	return ""
}

func (c *Config) Verify(chainConfig precompileconfig.ChainConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// Equal returns true if [cfg] is a [*RewardManagerConfig] and it has been configured identical to [c].
func (c *Config) Equal(cfg precompileconfig.Config) bool {
	_ = "STUB: not implemented"
	// typecast before comparison
	return false
}
