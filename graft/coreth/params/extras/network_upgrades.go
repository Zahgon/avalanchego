// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package extras

import (
	"github.com/ava-labs/avalanchego/upgrade"

	ethparams "github.com/ava-labs/libevm/params"
)

// NetworkUpgrades tracks the timestamps of all the Avalanche upgrades.
//
// For each upgrade, a nil value means the fork hasn't happened and is not
// scheduled. A pointer to 0 means the fork has already activated.
type NetworkUpgrades struct {
	ApricotPhase1BlockTimestamp *uint64 `json:"apricotPhase1BlockTimestamp,omitempty"` // Apricot Phase 1 Block Timestamp
	// Apricot Phase 2 Block Timestamp includes a modified version of the Berlin
	// Hard Fork.
	ApricotPhase2BlockTimestamp *uint64 `json:"apricotPhase2BlockTimestamp,omitempty"`
	// Apricot Phase 3 introduces dynamic fees and a modified version of the
	// London Hard Fork.
	ApricotPhase3BlockTimestamp *uint64 `json:"apricotPhase3BlockTimestamp,omitempty"`
	// Apricot Phase 4 introduces the notion of a block fee to the dynamic fee
	// algorithm.
	ApricotPhase4BlockTimestamp *uint64 `json:"apricotPhase4BlockTimestamp,omitempty"`
	// Apricot Phase 5 introduces a batch of atomic transactions with a maximum
	// atomic gas limit per block.
	ApricotPhase5BlockTimestamp *uint64 `json:"apricotPhase5BlockTimestamp,omitempty"`
	// Apricot Phase Pre-6 deprecates the NativeAssetCall precompile (soft).
	ApricotPhasePre6BlockTimestamp *uint64 `json:"apricotPhasePre6BlockTimestamp,omitempty"`
	// Apricot Phase 6 deprecates the NativeAssetBalance and NativeAssetCall
	// precompiles.
	ApricotPhase6BlockTimestamp *uint64 `json:"apricotPhase6BlockTimestamp,omitempty"`
	// Apricot Phase Post-6 deprecates the NativeAssetCall precompile (soft).
	ApricotPhasePost6BlockTimestamp *uint64 `json:"apricotPhasePost6BlockTimestamp,omitempty"`
	// Banff restricts import/export transactions to AVAX.
	BanffBlockTimestamp *uint64 `json:"banffBlockTimestamp,omitempty"`
	// Cortina increases the block gas limit to 15M.
	CortinaBlockTimestamp *uint64 `json:"cortinaBlockTimestamp,omitempty"`
	// Durango activates Avalanche Warp Messaging and the Shanghai Execution
	// Spec Upgrade (https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/shanghai.md#included-eips).
	//
	// Note: EIP-4895 is excluded since withdrawals are not relevant to the
	// Avalanche C-Chain or Subnets running the EVM.
	DurangoBlockTimestamp *uint64 `json:"durangoBlockTimestamp,omitempty"`
	// Etna activates Cancun (https://github.com/ethereum/execution-specs/blob/master/network-upgrades/mainnet-upgrades/cancun.md#included-eips)
	// and reduces the min base fee.
	// Note: EIP-4844 BlobTxs are not enabled in the mempool and blocks are not
	// allowed to contain them. For details see https://github.com/avalanche-foundation/ACPs/pull/131
	EtnaTimestamp *uint64 `json:"etnaTimestamp,omitempty"`
	// Fortuna modifies the gas price mechanism based on ACP-176
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

// IsApricotPhase1 returns whether [time] represents a block
// with a timestamp after the Apricot Phase 1 upgrade time.
func (n NetworkUpgrades) IsApricotPhase1(time uint64) bool { _ = "STUB: not implemented"; return false }

// IsApricotPhase2 returns whether [time] represents a block
// with a timestamp after the Apricot Phase 2 upgrade time.
func (n NetworkUpgrades) IsApricotPhase2(time uint64) bool { _ = "STUB: not implemented"; return false }

// IsApricotPhase3 returns whether [time] represents a block
// with a timestamp after the Apricot Phase 3 upgrade time.
func (n *NetworkUpgrades) IsApricotPhase3(time uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// IsApricotPhase4 returns whether [time] represents a block
// with a timestamp after the Apricot Phase 4 upgrade time.
func (n NetworkUpgrades) IsApricotPhase4(time uint64) bool { _ = "STUB: not implemented"; return false }

// IsApricotPhase5 returns whether [time] represents a block
// with a timestamp after the Apricot Phase 5 upgrade time.
func (n NetworkUpgrades) IsApricotPhase5(time uint64) bool { _ = "STUB: not implemented"; return false }

// IsApricotPhasePre6 returns whether [time] represents a block
// with a timestamp after the Apricot Phase Pre 6 upgrade time.
func (n NetworkUpgrades) IsApricotPhasePre6(time uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// IsApricotPhase6 returns whether [time] represents a block
// with a timestamp after the Apricot Phase 6 upgrade time.
func (n NetworkUpgrades) IsApricotPhase6(time uint64) bool { _ = "STUB: not implemented"; return false }

// IsApricotPhasePost6 returns whether [time] represents a block
// with a timestamp after the Apricot Phase 6 Post upgrade time.
func (n NetworkUpgrades) IsApricotPhasePost6(time uint64) bool {
	_ = "STUB: not implemented"
	return false
}

// IsBanff returns whether [time] represents a block
// with a timestamp after the Banff upgrade time.
func (n NetworkUpgrades) IsBanff(time uint64) bool { _ = "STUB: not implemented"; return false }

// IsCortina returns whether [time] represents a block
// with a timestamp after the Cortina upgrade time.
func (n NetworkUpgrades) IsCortina(time uint64) bool { _ = "STUB: not implemented"; return false }

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

func (n NetworkUpgrades) Description() string { _ = "STUB: not implemented"; return "" }

func GetNetworkUpgrades(agoUpgrade upgrade.Config) NetworkUpgrades {
	_ = "STUB: not implemented"
	return *new(NetworkUpgrades)
}

type AvalancheRules struct {
	IsApricotPhase1, IsApricotPhase2, IsApricotPhase3, IsApricotPhase4, IsApricotPhase5 bool
	IsApricotPhasePre6, IsApricotPhase6, IsApricotPhasePost6                            bool
	IsBanff                                                                             bool
	IsCortina                                                                           bool
	IsDurango                                                                           bool
	IsEtna                                                                              bool
	IsFortuna                                                                           bool
	IsGranite                                                                           bool
	IsHelicon                                                                           bool
}

// IsGraniteActivated is used by the warp precompile to determine which gas costs to use.
func (a AvalancheRules) IsGraniteActivated() bool {
	_ = "STUB: not implemented"

	// IsDurangoActivated is used by the precompileconfig to determine if the Durango upgrade is activated.
	// For Coreth we don't really need this, but need to implement it to satisfy the interface.
	return false
}

func (a AvalancheRules) IsDurangoActivated() bool { _ = "STUB: not implemented"; return false }

func (n *NetworkUpgrades) GetAvalancheRules(timestamp uint64) AvalancheRules {
	_ = "STUB: not implemented"
	return *new(AvalancheRules)
}

func ptrToString(val *uint64) string { _ = "STUB: not implemented"; return "" }
