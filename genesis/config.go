// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/vms/platformvm/signer"
)

const localNetworkUpdateStartTimePeriod = 9 * 30 * 24 * time.Hour // 9 months

var (
	_ utils.Sortable[Allocation] = Allocation{}

	errInvalidGenesisJSON = errors.New("could not unmarshal genesis JSON")
)

type LockedAmount struct {
	Amount   uint64 `json:"amount"`
	Locktime uint64 `json:"locktime"`
}

type Allocation struct {
	ETHAddr        ids.ShortID    `json:"ethAddr"`
	AVAXAddr       ids.ShortID    `json:"avaxAddr"`
	InitialAmount  uint64         `json:"initialAmount"`
	UnlockSchedule []LockedAmount `json:"unlockSchedule"`
}

func (a Allocation) Unparse(networkID uint32) (UnparsedAllocation, error) {
	_ = "STUB: not implemented"
	return *new(UnparsedAllocation), nil
}

func (a Allocation) Compare(other Allocation) int { _ = "STUB: not implemented"; return 0 }

type Staker struct {
	NodeID        ids.NodeID                `json:"nodeID"`
	RewardAddress ids.ShortID               `json:"rewardAddress"`
	DelegationFee uint32                    `json:"delegationFee"`
	Signer        *signer.ProofOfPossession `json:"signer,omitempty"`
}

func (s Staker) Unparse(networkID uint32) (UnparsedStaker, error) {
	_ = "STUB: not implemented"
	return *new(UnparsedStaker), nil
}

// Config contains the genesis addresses used to construct a genesis
type Config struct {
	NetworkID uint32 `json:"networkID"`

	Allocations []Allocation `json:"allocations"`

	StartTime                  uint64        `json:"startTime"`
	InitialStakeDuration       uint64        `json:"initialStakeDuration"`
	InitialStakeDurationOffset uint64        `json:"initialStakeDurationOffset"`
	InitialStakedFunds         []ids.ShortID `json:"initialStakedFunds"`
	InitialStakers             []Staker      `json:"initialStakers"`

	CChainGenesis string `json:"cChainGenesis"`

	Message string `json:"message"`
}

func (c Config) Unparse() (UnparsedConfig, error) {
	_ = "STUB: not implemented"
	return *new(UnparsedConfig), nil
}

func (c *Config) InitialSupply() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

var (
	// MainnetConfig is the config that should be used to generate the mainnet
	// genesis.
	MainnetConfig Config

	// FujiConfig is the config that should be used to generate the fuji
	// genesis.
	FujiConfig Config

	// LocalConfig is the config that should be used to generate a local
	// genesis.
	LocalConfig Config

	// unmodifiedLocalConfig is the LocalConfig before advancing the StartTime
	// to a recent value.
	unmodifiedLocalConfig Config
)

func init() {
	unparsedMainnetConfig := UnparsedConfig{}
	unparsedFujiConfig := UnparsedConfig{}
	unparsedLocalConfig := UnparsedConfig{}

	err := errors.Join(
		json.Unmarshal(mainnetGenesisConfigJSON, &unparsedMainnetConfig),
		json.Unmarshal(fujiGenesisConfigJSON, &unparsedFujiConfig),
		json.Unmarshal(localGenesisConfigJSON, &unparsedLocalConfig),
	)
	if err != nil {
		panic(err)
	}

	MainnetConfig, err = unparsedMainnetConfig.Parse()
	if err != nil {
		panic(err)
	}

	FujiConfig, err = unparsedFujiConfig.Parse()
	if err != nil {
		panic(err)
	}

	unmodifiedLocalConfig, err = unparsedLocalConfig.Parse()
	if err != nil {
		panic(err)
	}

	// Renew the staking start time of the local config if required
	definedStartTime := time.Unix(int64(unmodifiedLocalConfig.StartTime), 0)
	recentStartTime := getRecentStartTime(
		definedStartTime,
		time.Now(),
		localNetworkUpdateStartTimePeriod,
	)

	LocalConfig = unmodifiedLocalConfig
	LocalConfig.StartTime = uint64(recentStartTime.Unix())
}

func GetConfig(networkID uint32) *Config { _ = "STUB: not implemented"; return nil }

// GetConfigFile loads a *Config from a provided filepath.
func GetConfigFile(fp string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }

// GetConfigContent loads a *Config from a provided environment variable
func GetConfigContent(genesisContent string) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseGenesisJSONBytesToConfig(bytes []byte) (*Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getRecentStartTime advances [definedStartTime] in chunks of [period]. It
// returns the latest startTime that isn't after [now].
func getRecentStartTime(
	definedStartTime time.Time,
	now time.Time,
	period time.Duration,
) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
