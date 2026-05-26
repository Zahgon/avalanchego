// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"crypto/tls"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/ava-labs/avalanchego/chains"
	"github.com/ava-labs/avalanchego/config/node"
	"github.com/ava-labs/avalanchego/genesis"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network"
	"github.com/ava-labs/avalanchego/snow/consensus/simplex"
	"github.com/ava-labs/avalanchego/snow/consensus/snowball"
	"github.com/ava-labs/avalanchego/snow/networking/benchlist"
	"github.com/ava-labs/avalanchego/snow/networking/router"
	"github.com/ava-labs/avalanchego/snow/networking/tracker"
	"github.com/ava-labs/avalanchego/subnets"
	"github.com/ava-labs/avalanchego/trace"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/profiler"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/timer"
	"github.com/ava-labs/avalanchego/vms/platformvm/reward"
)

const (
	chainConfigFileName  = "config"
	chainUpgradeFileName = "upgrade"
	subnetConfigFileExt  = ".json"

	maxDiskSpaceThreshold = 50
)

type consensusMode int

const (
	modeSimplex consensusMode = iota
	modeSnow
	modeSnowFromDeprecated
	modeDefaultSnow
)

var (
	// Deprecated key --> deprecation message (i.e. which key replaces it)
	// TODO: deprecate "BootstrapIDsKey" and "BootstrapIPsKey"
	deprecatedKeys = map[string]string{
		SystemTrackerRequiredAvailableDiskSpaceKey:         fmt.Sprintf("Use %s instead", SystemTrackerRequiredAvailableDiskSpacePercentageKey),
		SystemTrackerWarningThresholdAvailableDiskSpaceKey: fmt.Sprintf("Use %s instead", SystemTrackerWarningAvailableDiskSpacePercentageKey),
	}

	errConflictingACPOpinion                  = errors.New("supporting and objecting to the same ACP")
	errConflictingImplicitACPOpinion          = errors.New("objecting to enabled ACP")
	errSybilProtectionDisabledStakerWeights   = errors.New("sybil protection disabled weights must be positive")
	errSybilProtectionDisabledOnPublicNetwork = errors.New("sybil protection disabled on public network")
	errInvalidUptimeRequirement               = errors.New("uptime requirement must be in the range [0, 1]")
	errMinValidatorStakeAboveMax              = errors.New("minimum validator stake can't be greater than maximum validator stake")
	errInvalidDelegationFee                   = errors.New("delegation fee must be in the range [0, 1,000,000]")
	errInvalidMinStakeDuration                = errors.New("min stake duration must be > 0")
	errMinStakeDurationAboveMax               = errors.New("max stake duration can't be less than min stake duration")
	errStakeMaxConsumptionTooLarge            = fmt.Errorf("max stake consumption must be less than or equal to %d", reward.PercentDenominator)
	errStakeMaxConsumptionBelowMin            = errors.New("stake max consumption can't be less than min stake consumption")
	errStakeMintingPeriodBelowMin             = errors.New("stake minting period can't be less than max stake duration")
	errCannotTrackPrimaryNetwork              = errors.New("cannot track primary network")
	errStakingKeyContentUnset                 = fmt.Errorf("%s key not set but %s set", StakingTLSKeyContentKey, StakingCertContentKey)
	errStakingCertContentUnset                = fmt.Errorf("%s key set but %s not set", StakingTLSKeyContentKey, StakingCertContentKey)
	errPluginDirNotADirectory                 = errors.New("plugin dir is not a directory")
	errCannotReadDirectory                    = errors.New("cannot read directory")
	errUnmarshalling                          = errors.New("unmarshalling failed")
	errFileDoesNotExist                       = errors.New("file does not exist")
	errInvalidSignerConfig                    = fmt.Errorf("only one of the following flags can be set: %s, %s, %s, %s", StakingEphemeralSignerEnabledKey, StakingSignerKeyContentKey, StakingSignerKeyPathKey, StakingRPCSignerEndpointKey)
	errDiskSpaceOutOfRange                    = fmt.Errorf("out of range [0,%d]", maxDiskSpaceThreshold)
	errDiskWarnAfterFatal                     = errors.New("warning disk space threshold cannot be greater than fatal threshold")
)

func getPrimaryNetworkSnowConfig(v *viper.Viper) *snowball.Parameters {
	_ = "STUB: not implemented"
	return nil
}

// applySnowballParameterDefaults populates unset fields in a snowball.Parameters
// struct with values from viper. It is intended for subnet configurations that
// have been unmarshalled into a snowball.Parameters struct where some fields
// may be omitted.
//
// If a field is zero-valued, it is treated as unset and defaulted from the
// corresponding configuration key.
//
// This function also handles the deprecated Alpha field for backward
// compatibility by mapping it to both AlphaPreference and AlphaConfidence.
func applySnowballParameterDefaults(config *snowball.Parameters, v *viper.Viper) {
	_ = "STUB: not implemented"
	return
}

// applySimplexDefaults sets the default values for any unset fields in the
// simplex.Parameters.
func applySimplexDefaults(config *simplex.Parameters, v *viper.Viper) {
	_ = "STUB: not implemented"
	return
}

func resolveConsensusMode(config *subnets.Config) consensusMode {
	_ = "STUB: not implemented"
	return *new(consensusMode)
}

// applySubnetConfigDefaults sets the default values for any unset fields in the subnets.Config.
func applySubnetConfigDefaults(config *subnets.Config, v *viper.Viper) {
	_ = "STUB: not implemented"
	return
}

// getSubnetConfigFromBytes unmarshals a subnet config from rawBytes and validates it.
// It also sets any unset fields to their default values for the provided subnet config.
func getSubnetConfigFromBytes(rawBytes []byte, v *viper.Viper) (subnets.Config, error) {
	_ = "STUB: not implemented"
	return *new(subnets.Config), nil
}

// Ensure that at most one consensus parameter type is set

// set unset fields

// validate parameters

func getLoggingConfig(v *viper.Viper) (logging.Config, error) {
	_ = "STUB: not implemented"
	return *new(logging.Config), nil
}

func getHTTPConfig(v *viper.Viper) (node.HTTPConfig, error) {
	_ = "STUB: not implemented"
	return *new(node.HTTPConfig), nil
}

func getRouterHealthConfig(v *viper.Viper, halflife time.Duration) (router.HealthConfig, error) {
	_ = "STUB: not implemented"
	return *new(router.HealthConfig), nil
}

func getAdaptiveTimeoutConfig(v *viper.Viper) (timer.AdaptiveTimeoutConfig, error) {
	_ = "STUB: not implemented"
	return *new(timer.AdaptiveTimeoutConfig), nil
}

func getNetworkConfig(
	v *viper.Viper,
	networkID uint32,
	sybilProtectionEnabled bool,
	halflife time.Duration,
) (network.Config, error) {
	_ = "STUB: not implemented"
	// Set the max number of recent inbound connections upgraded to be
	// equal to the max number of inbound connections per second.
	return *new(network.Config), nil
}

// Because this node version has scheduled these ACPs, we should notify
// peers that we support these upgrades.

// To decrease unnecessary network traffic, peers will not be notified of
// objection or support of activated ACPs.

func getBenchlistConfig(v *viper.Viper, snowballParameters *snowball.Parameters) (benchlist.Config, error) {
	_ = "STUB: not implemented"
	return *new(benchlist.Config), nil
}

// AlphaConfidence is used here to ensure that benching can't cause a
// liveness failure. If AlphaPreference were used, the benchlist may grow to
// a point that committing would be extremely unlikely to happen.

func getStateSyncConfig(v *viper.Viper) (node.StateSyncConfig, error) {
	_ = "STUB: not implemented"
	return *new(node.StateSyncConfig), nil
}

func getBootstrapConfig(v *viper.Viper, networkID uint32) (node.BootstrapConfig, error) {
	_ = "STUB: not implemented"
	return *new(node.BootstrapConfig), nil
}

// TODO: Add a "BootstrappersKey" flag to more clearly enforce ID and IP
// length equality.

// ID is populated below

func getIPConfig(v *viper.Viper) (node.IPConfig, error) {
	_ = "STUB: not implemented"
	return *new(node.IPConfig), nil
}

func getProfilerConfig(v *viper.Viper) (profiler.Config, error) {
	_ = "STUB: not implemented"
	return *new(profiler.Config), nil
}

func getStakingTLSCertFromFlag(v *viper.Viper) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

func getStakingTLSCertFromFile(v *viper.Viper) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	// Parse the staking key/cert paths and expand environment variables
	return *new(tls.Certificate), nil
}

// If staking key/cert locations are specified but not found, error

// Create the staking key/cert if [stakingKeyPath] and [stakingCertPath] don't exist

// Load and parse the staking key/cert

func getStakingTLSCert(v *viper.Viper) (tls.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(tls.Certificate), nil
}

// Use an ephemeral staking key/cert

func getStakingConfig(v *viper.Viper, networkID uint32) (node.StakingConfig, error) {
	_ = "STUB: not implemented"
	return *new(node.StakingConfig), nil
}

func getStakingSignerConfig(v *viper.Viper) (node.StakingSignerConfig, error) {
	_ = "STUB: not implemented"
	// A maximum of one signer option can be set
	return *new(node.StakingSignerConfig), nil
}

// Set signerKeyPath only none of the other signer options are set

func getTxFeeConfig(v *viper.Viper, networkID uint32) genesis.TxFeeConfig {
	_ = "STUB: not implemented"
	return *new(genesis.TxFeeConfig)
}

func getUpgradeConfig(v *viper.Viper, networkID uint32) (upgrade.Config, error) {
	_ = "STUB: not implemented"
	return *new(upgrade.Config), nil
}

func getGenesisData(v *viper.Viper, networkID uint32, stakingCfg *genesis.StakingConfig) ([]byte, ids.ID, error) {
	_ = "STUB: not implemented"
	// try first loading genesis content directly from flag/env-var
	return nil, *new(ids.ID), nil
}

// if content is not specified go for the file

// finally if file is not specified/readable go for the predefined config

func getTrackedSubnets(v *viper.Viper) (set.Set[ids.ID], error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDatabaseConfig(v *viper.Viper, networkID uint32) (node.DatabaseConfig, error) {
	_ = "STUB: not implemented"
	return *new(node.DatabaseConfig), nil
}

func getAliases(v *viper.Viper, name string, contentKey string, fileKey string) (map[ids.ID][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getVMAliases(v *viper.Viper) (map[ids.ID][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getChainAliases(v *viper.Viper) (map[ids.ID][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getPathFromDirKey reads flag value from viper instance and then checks the folder existence
func getPathFromDirKey(v *viper.Viper, configKey string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// user specified a config dir explicitly, but dir does not exist.

func getChainConfigsFromFlag(v *viper.Viper) (map[string]chains.ChainConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getChainConfigsFromDir(v *viper.Viper) (map[string]chains.ChainConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getChainConfigs reads & puts chainConfigs to node config
func getChainConfigs(v *viper.Viper) (map[string]chains.ChainConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// readChainConfigPath reads chain config files from static directories and returns map with contents,
// if successful.
func readChainConfigPath(chainConfigPath string) (map[string]chains.ChainConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// chainconfigdir/chainId/config.*

// chainconfigdir/chainId/upgrade.*

// getSubnetConfigs reads subnet configs from the correct place
// (flag or file) and returns a non-nil map.
func getSubnetConfigs(v *viper.Viper, subnetIDs []ids.ID) (map[ids.ID]subnets.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getSubnetConfigsFromFlags(v *viper.Viper, subnetIDs []ids.ID) (map[ids.ID]subnets.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// partially parse configs to be filled by defaults later

// getSubnetConfigsFromDir reads SubnetConfigs to node config map
func getSubnetConfigsFromDir(v *viper.Viper, subnetIDs []ids.ID) (map[ids.ID]subnets.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// reads subnet config files from a path and given subnetIDs and returns a map.

// subnet config path does not exist but not explicitly specified, so ignore it

// this subnet config does not exist, the default configuration will be used

// subnetConfigDir/subnetID.json

func getPrimaryNetworkConfig(v *viper.Viper) subnets.Config {
	_ = "STUB: not implemented"
	return *new(subnets.Config)
}

func getCPUTargeterConfig(v *viper.Viper) (tracker.TargeterConfig, error) {
	_ = "STUB: not implemented"
	return *new(tracker.TargeterConfig), nil
}

func getDiskSpaceConfig(v *viper.Viper) (
	requiredAvailableDiskSpacePercentage uint64,
	warningAvailableDiskSpacePercentage uint64,
	err error,
) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func getDiskTargeterConfig(v *viper.Viper) (tracker.TargeterConfig, error) {
	_ = "STUB: not implemented"
	return *new(tracker.TargeterConfig), nil
}

func getTraceConfig(v *viper.Viper) (trace.Config, error) {
	_ = "STUB: not implemented"
	return *new(trace.Config), nil
}

// Returns the path to the directory that contains VM binaries.
func getPluginDir(v *viper.Viper) (string, error) { _ = "STUB: not implemented"; return "", nil }

// If the flag was given, assert it exists and is a directory

// If the flag wasn't given, make sure the default location exists.

func GetNodeConfig(v *viper.Viper) (node.Config, error) {
	_ = "STUB: not implemented"
	return *new(node.Config), nil
}

// Gossiping

// App handling

// Logging

// Network ID

// Database

// IP configuration

// Staking

// Tracked Subnets

// HTTP APIs

// Health

// Halflife of continuous averager used in health checks

// Router

// Metrics

// Adaptive Timeout Config

// Upgrade config

// Network Config

// Subnet Configs

// Benchlist

// File Descriptor Limit

// Tx Fee

// Genesis Data

// StateSync Configs

// Bootstrap Configs

// Chain Configs

// Profiler

// VM Aliases

// Chain aliases

func providedFlags(v *viper.Viper) map[string]interface{} { _ = "STUB: not implemented"; return nil }
