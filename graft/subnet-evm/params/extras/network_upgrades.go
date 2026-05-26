// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package extras

import (
	"errors"

	"github.com/ava-labs/avalanchego/upgrade"

	ethparams "github.com/ava-labs/libevm/params"
)

var (
	unscheduledActivation = uint64(upgrade.UnscheduledActivationTime.Unix())
	initiallyActiveTime   = uint64(upgrade.InitiallyActiveTime.Unix())

	errCannotBeNil             = errors.New("timestamp cannot be nil")
	errTimestampTooEarly       = errors.New("provided timestamp must be greater than or equal to the default timestamp")
	errUnsupportedForkOrdering = errors.New("unsupported fork ordering")
)

// NetworkUpgrades contains timestamps that enable network upgrades.
// Avalanche specific network upgrades are also included here.
// (nil = no fork, 0 = already activated)
type NetworkUpgrades struct {
	// SubnetEVMTimestamp is a placeholder that activates Avalanche Upgrades prior to ApricotPhase6
	SubnetEVMTimestamp *uint64 `json:"subnetEVMTimestamp,omitempty"`
	// Durango activates the Shanghai Execution Spec Upgrade from Ethereum (https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/shanghai.md#included-eips)
	// and Avalanche Warp Messaging.
	// Note: EIP-4895 is excluded since withdrawals are not relevant to the Avalanche C-Chain or Subnets running the EVM.
	DurangoTimestamp *uint64 `json:"durangoTimestamp,omitempty"`
	// Placeholder for EtnaTimestamp
	EtnaTimestamp *uint64 `json:"etnaTimestamp,omitempty"`
	// Fortuna has no effect on Subnet-EVM by itself, but is included for completeness.
	FortunaTimestamp *uint64 `json:"fortunaTimestamp,omitempty"`
	// Granite adds a millisecond timestamp, precompile updates, and P-Chain epochs
	GraniteTimestamp *uint64 `json:"graniteTimestamp,omitempty"`
	// Helicon is a placeholder for the next upgrade
	HeliconTimestamp *uint64 `json:"heliconTimestamp,omitempty"`
}

func (n *NetworkUpgrades) Equal(other *NetworkUpgrades) bool {
	_ = "STUB: not implemented"
	return false
}

func (n *NetworkUpgrades) checkNetworkUpgradesCompatible(newcfg *NetworkUpgrades, time uint64) *ethparams.ConfigCompatError {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkUpgrades) forkOrder() []fork { _ = "STUB: not implemented"; return nil }

// SetDefaults sets the default values for the network upgrades.
// This overrides deactivating the network upgrade by providing a timestamp of nil value.
func (n *NetworkUpgrades) SetDefaults(agoUpgrades upgrade.Config) {
	_ = "STUB: not implemented"
	return
}

// If the network upgrade is not set, set it to the default value.
// If the network upgrade is set to 0, we also treat it as nil and set it default.
// Invariant: This is because in prior versions, upgrades were not modifiable and were directly set to their default values.
// Most of the tools and configurations just provide these as 0, so it is safer to treat 0 as nil and set to default
// to prevent premature activations of the network upgrades for live networks.

// verifyNetworkUpgrades checks that the network upgrades are well formed.
func (n *NetworkUpgrades) verifyNetworkUpgrades(agoUpgrades upgrade.Config) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *NetworkUpgrades) Override(o *NetworkUpgrades) { _ = "STUB: not implemented"; return }

// IsSubnetEVM returns whether [time] represents a block
// with a timestamp after the SubnetEVM upgrade time.
func (n NetworkUpgrades) IsSubnetEVM(time uint64) bool { _ = "STUB: not implemented"; return false }

// IsDurango returns whether [time] represents a block
// with a timestamp after the Durango upgrade time.
func (n NetworkUpgrades) IsDurango(time uint64) bool { _ = "STUB: not implemented"; return false }

// IsEtna returns whether [time] represents a block
// with a timestamp after the Etna upgrade time.
func (n NetworkUpgrades) IsEtna(time uint64) bool { _ = "STUB: not implemented"; return false }

// IsFortuna returns whether [time] represents a block
// with a timestamp after the Fortuna upgrade time.
func (n *NetworkUpgrades) IsFortuna(time uint64) bool { _ = "STUB: not implemented"; return false }

// IsGranite returns whether [time] represents a block
// with a timestamp after the Granite upgrade time.
func (n *NetworkUpgrades) IsGranite(time uint64) bool { _ = "STUB: not implemented"; return false }

// IsHelicon returns whether [time] represents a block
// with a timestamp after the Helicon upgrade time.
func (n *NetworkUpgrades) IsHelicon(time uint64) bool { _ = "STUB: not implemented"; return false }

func (n *NetworkUpgrades) Description() string { _ = "STUB: not implemented"; return "" }

type AvalancheRules struct {
	IsSubnetEVM bool
	IsDurango   bool
	IsEtna      bool
	IsFortuna   bool
	IsGranite   bool
	IsHelicon   bool
}

// IsGraniteActivated is used by the warp precompile to determine which gas costs to use.
func (a AvalancheRules) IsGraniteActivated() bool {
	_ = "STUB: not implemented"

	// IsDurangoActivated is used by the warp precompile to determine which gas costs to use.
	return false
}

func (a AvalancheRules) IsDurangoActivated() bool { _ = "STUB: not implemented"; return false }

func (n *NetworkUpgrades) GetAvalancheRules(time uint64) AvalancheRules {
	_ = "STUB: not implemented"
	return *new(AvalancheRules)
}

// GetNetworkUpgrades returns the network upgrades for the specified avalanchego upgrades.
// Nil values are used to indicate optional upgrades.
func GetNetworkUpgrades(agoUpgrade upgrade.Config) NetworkUpgrades {
	_ = "STUB: not implemented"
	return *new(NetworkUpgrades)
}

// Fortuna is optional and has no effect on Subnet-EVM

// verifyWithDefault checks that the provided timestamp is greater than or equal to the default timestamp.
func verifyWithDefault(configTimestamp *uint64, defaultTimestamp *uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// handle avalanche edge-cases:
// nil -> error unless default is unscheduled
// 0  -> allowed for initially-active defaults
// non-zero -> must be >= default.

func ptrToString(val *uint64) string { _ = "STUB: not implemented"; return "" }
