// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"time"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/wrappers"
)

// preDelegateeRewardSize is the size of codec marshalling
// [preDelegateeRewardMetadata].
//
// CodecVersionLen + UpDurationLen + LastUpdatedLen + PotentialRewardLen
const preDelegateeRewardSize = codec.VersionSize + 3*wrappers.LongLen

type preDelegateeRewardMetadata struct {
	UpDuration      time.Duration `v0:"true"`
	LastUpdated     uint64        `v0:"true"` // Unix time in seconds
	PotentialReward uint64        `v0:"true"`
}

type validatorMetadata struct {
	UpDuration               time.Duration `v0:"true"`
	LastUpdated              uint64        `v0:"true"` // Unix time in seconds
	PotentialReward          uint64        `v0:"true"`
	PotentialDelegateeReward uint64        `v0:"true"`
	StakerStartTime          uint64        `          v1:"true"`

	txID        ids.ID
	lastUpdated time.Time
}

// Permissioned validators originally wrote their values as nil.
// With Banff we wrote the potential reward.
// With Cortina we wrote the potential reward with the potential delegatee reward.
// We now write the uptime, reward, and delegatee reward together.
func parseValidatorMetadata(bytes []byte, metadata *validatorMetadata) error {
	_ = "STUB: not implemented"
	return nil

	// nothing was stored
}

// only potential reward was stored

// potential reward and uptime was stored but potential delegatee reward
// was not

// everything was stored

// StakingInfo holds mutable validator data that can be modified.
type StakingInfo struct {
	DelegateeReward uint64
}

func stakingInfoFromMetadata(vdrMetadata *validatorMetadata) StakingInfo {
	_ = "STUB: not implemented"
	return *new(StakingInfo)
}

type validatorState struct {
	metadata map[ids.NodeID]map[ids.ID]*validatorMetadata // vdrID -> subnetID -> metadata
	// updatedMetadata tracks (vdrID, subnetID) -> txIDs needing DB sync since the
	// last WriteValidatorMetadata.
	updatedMetadata map[ids.NodeID]map[ids.ID]set.Set[ids.ID]
}

func newValidatorState() *validatorState { _ = "STUB: not implemented"; return nil }

// LoadValidatorMetadata sets the `uptime` of `vdrID` on `subnetID`.
// [GetUptime] and [SetUptime] will return an error if `vdrID` and
// `subnetID` hasn't been loaded. This call will not result in a write to
// disk.
func (vs *validatorState) LoadValidatorMetadata(
	vdrID ids.NodeID,
	subnetID ids.ID,
	uptime *validatorMetadata,
) {
	_ = "STUB: not implemented"
	return
}

// AddValidatorMetadata loads the metadata and marks it as updated so it will
// be written to disk on the next call to [WriteValidatorMetadata].
func (vs *validatorState) AddValidatorMetadata(
	vdrID ids.NodeID,
	subnetID ids.ID,
	vm *validatorMetadata,
) {
	_ = "STUB: not implemented"
	return
}

// GetUptime returns the current uptime measurements of `vdrID` on
// `subnetID`.
func (vs *validatorState) GetUptime(
	vdrID ids.NodeID,
	subnetID ids.ID,
) (time.Duration, time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Time), nil
}

// SetUptime updates the uptime measurements of `vdrID` on `subnetID`.
// Unless these measurements are deleted first, the next call to
// [WriteValidatorMetadata] will write this update to disk.
//
// This is called by the consensus layer to track validator connection times.
func (vs *validatorState) SetUptime(
	vdrID ids.NodeID,
	subnetID ids.ID,
	upDuration time.Duration,
	lastUpdated time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

// GetStakingInfo returns the mutable staking info for the validator on [subnetID] with [vdrID].
func (vs *validatorState) GetStakingInfo(
	subnetID ids.ID,
	vdrID ids.NodeID,
) (StakingInfo, error) {
	_ = "STUB: not implemented"
	return *new(StakingInfo), nil
}

// SetStakingInfo updates the mutable staking info of `vdrID` on `subnetID`.
// Unless deleted first, the next call to [WriteValidatorMetadata] will write this update to disk.
//
// This is called by execution layer to update mutable staking info.
func (vs *validatorState) SetStakingInfo(
	subnetID ids.ID,
	vdrID ids.NodeID,
	stakingInfo StakingInfo,
) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteValidatorMetadata removes in-memory references to the metadata of
// `vdrID` on `subnetID`. The txID is recorded for deletion from disk on the
// next [WriteValidatorMetadata]. Any staged updates from [SetUptime] or
// [SetStakingInfo] are dropped.
func (vs *validatorState) DeleteValidatorMetadata(vdrID ids.NodeID, subnetID ids.ID) {
	_ = "STUB: not implemented"
	return
}

// WriteValidatorMetadata persists all entries in updatedMetadata to disk. For
// each (vdrID, subnetID) and txID in the set: if metadata exists and its txID
// matches, write it to disk; otherwise delete the txID from disk.
func (vs *validatorState) WriteValidatorMetadata(
	dbPrimary database.KeyValueWriterDeleter,
	dbSubnet database.KeyValueWriterDeleter,
	codecVersion uint16,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (vs *validatorState) addUpdatedTxID(vdrID ids.NodeID, subnetID ids.ID, txID ids.ID) {
	_ = "STUB: not implemented"
	return
}
