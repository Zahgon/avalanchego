// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package params

import (
	"errors"

	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/upgrade"
)

const (
	maxJSONLen = 64 * 1024 * 1024 // 64MB

	// TODO: Value to pass to geth's Rules by default where the appropriate
	// context is not available in the avalanche code. (similar to context.TODO())
	IsMergeTODO = true
)

var (
	initiallyActive       = uint64(upgrade.InitiallyActiveTime.Unix())
	unscheduledActivation = uint64(upgrade.UnscheduledActivationTime.Unix())

	errInvalidUpgradeTime = errors.New("invalid upgrade time")
)

// SetEthUpgrades enables Ethereum network upgrades using the same time as
// the Avalanche network upgrade that enables them.
func SetEthUpgrades(c *ChainConfig) error {
	_ = "STUB: not implemented"
	// Set Ethereum block upgrades to initially activated as they were already
	// activated on launch.
	return nil
}

// Because Fuji and Mainnet have already accepted the Berlin and London
// blocks, it is assumed that they are scheduled for activation.

// https://testnet.snowtrace.io/block/184985?chainid=43113, AP2 activation block
// https://testnet.snowtrace.io/block/805078?chainid=43113, AP3 activation block

// https://snowtrace.io/block/1640340?chainid=43114, AP2 activation block
// https://snowtrace.io/block/3308552?chainid=43114, AP3 activation block

// In testing or local networks, we only support enabling Berlin and
// London at the initially active time. This corresponds to an intended
// block number of 0.

// We only mark Shanghai and Cancun as enabled if we have marked them as
// scheduled.

func GetExtra(c *ChainConfig) *extras.ChainConfig { _ = "STUB: not implemented"; return nil }

func Copy(c *ChainConfig) ChainConfig { _ = "STUB: not implemented"; return *new(ChainConfig) }

// WithExtra sets the extra payload on `c` and returns the modified argument.
func WithExtra(c *ChainConfig, extra *extras.ChainConfig) *ChainConfig {
	_ = "STUB: not implemented"
	return nil
}

type ChainConfigWithUpgradesJSON struct {
	ChainConfig
	UpgradeConfig extras.UpgradeConfig `json:"upgrades,omitempty"`
}

// MarshalJSON implements json.Marshaler. This is a workaround for the fact that
// the embedded ChainConfig struct has a MarshalJSON method, which prevents
// the default JSON marshalling from working for UpgradeConfig.
// TODO: consider removing this method by allowing external tag for the embedded
// ChainConfig struct.
func (cu ChainConfigWithUpgradesJSON) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// embed the ChainConfig struct into the response
	return nil, nil
}

// merge the two JSON objects

func (cu *ChainConfigWithUpgradesJSON) UnmarshalJSON(input []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// ToWithUpgradesJSON converts the ChainConfig to ChainConfigWithUpgradesJSON with upgrades explicitly displayed.
// ChainConfig does not include upgrades in its JSON output.
// This is a workaround for showing upgrades in the JSON output.
func ToWithUpgradesJSON(c *ChainConfig) *ChainConfigWithUpgradesJSON {
	_ = "STUB: not implemented"
	return nil
}
