// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesistest

import (
	"testing"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/snowtest"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/platformvm/reward"

	platformvmgenesis "github.com/ava-labs/avalanchego/vms/platformvm/genesis"
)

const (
	DefaultValidatorDuration = 28 * 24 * time.Hour
	DefaultValidatorWeight   = 5 * units.MilliAvax
	DefaultInitialBalance    = 1 * units.Avax

	ValidatorDelegationShares = reward.PercentDenominator
	XChainName                = "x"
	InitialSupply             = 360 * units.MegaAvax
)

var (
	AVAXAsset = avax.Asset{ID: snowtest.AVAXAssetID}

	DefaultValidatorStartTime     = upgrade.InitiallyActiveTime
	DefaultValidatorStartTimeUnix = uint64(DefaultValidatorStartTime.Unix())
	DefaultValidatorEndTime       = DefaultValidatorStartTime.Add(DefaultValidatorDuration)
	DefaultValidatorEndTimeUnix   = uint64(DefaultValidatorEndTime.Unix())
)

var (
	// Keys that are funded in the genesis
	DefaultFundedKeys = secp256k1.TestKeys()

	// Node IDs of genesis validators
	DefaultNodeIDs []ids.NodeID
)

func init() {
	DefaultNodeIDs = make([]ids.NodeID, len(DefaultFundedKeys))
	for i := range DefaultFundedKeys {
		DefaultNodeIDs[i] = ids.GenerateTestNodeID()
	}
}

type Config struct {
	NetworkID          uint32
	NodeIDs            []ids.NodeID
	ValidatorWeight    uint64
	ValidatorStartTime time.Time
	ValidatorEndTime   time.Time

	FundedKeys     []*secp256k1.PrivateKey
	InitialBalance uint64
}

func New(t testing.TB, c Config) *platformvmgenesis.Genesis { _ = "STUB: not implemented"; return nil }

func NewBytes(t testing.TB, c Config) []byte { _ = "STUB: not implemented"; return nil }
