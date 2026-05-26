// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	DefaultHTTPPort    = 9650
	DefaultStakingPort = 9651

	AvalancheGoDataDirVar    = "AVALANCHEGO_DATA_DIR"
	defaultUnexpandedDataDir = "$" + AvalancheGoDataDirVar

	DefaultProcessContextFilename = "process.json"
)

var (
	// [defaultUnexpandedDataDir] will be expanded when reading the flags
	defaultDataDir              = filepath.Join("$HOME", ".avalanchego")
	defaultDBDir                = filepath.Join(defaultUnexpandedDataDir, "db")
	defaultLogDir               = filepath.Join(defaultUnexpandedDataDir, "logs")
	defaultProfileDir           = filepath.Join(defaultUnexpandedDataDir, "profiles")
	defaultStakingPath          = filepath.Join(defaultUnexpandedDataDir, "staking")
	defaultStakingTLSKeyPath    = filepath.Join(defaultStakingPath, "staker.key")
	defaultStakingCertPath      = filepath.Join(defaultStakingPath, "staker.crt")
	defaultStakingSignerKeyPath = filepath.Join(defaultStakingPath, "signer.key")
	defaultConfigDir            = filepath.Join(defaultUnexpandedDataDir, "configs")
	defaultChainConfigDir       = filepath.Join(defaultConfigDir, "chains")
	defaultVMConfigDir          = filepath.Join(defaultConfigDir, "vms")
	defaultVMAliasFilePath      = filepath.Join(defaultVMConfigDir, "aliases.json")
	defaultChainAliasFilePath   = filepath.Join(defaultChainConfigDir, "aliases.json")
	defaultSubnetConfigDir      = filepath.Join(defaultConfigDir, "subnets")
	defaultPluginDir            = filepath.Join(defaultUnexpandedDataDir, "plugins")
	defaultChainDataDir         = filepath.Join(defaultUnexpandedDataDir, "chainData")
	defaultProcessContextPath   = filepath.Join(defaultUnexpandedDataDir, DefaultProcessContextFilename)
)

func deprecateFlags(fs *pflag.FlagSet) error { _ = "STUB: not implemented"; return nil }

func addProcessFlags(fs *pflag.FlagSet) {
	_ = "STUB: not implemented"
	// If true, print the version and quit.
	return
}

func addNodeFlags(fs *pflag.FlagSet) {
	_ = "STUB: not implemented"
	// Home directory
	return
}

// System

// Plugin directory

// Config File

// Genesis

// Upgrade

// Network ID

// ACP flagging

// AVAX fees:
// Validator fees:

// Dynamic fees:

// Static fees:

// Database

// Logging

// Peer List Gossip

// Public IP Resolution

// Inbound Connection Throttling

// Outbound Connection Throttling

// Timeouts

// Note: The default value is set to false here because the default
// networkID is mainnet. The real default value of NetworkAllowPrivateIPs is
// based on the networkID.

// The PROXY protocol specification recommends setting this value to be at
// least 3 seconds to cover a TCP retransmit.
// Ref: https://www.haproxy.org/download/2.3/doc/proxy-protocol.txt
// Specifying a timeout of 0 will actually result in a timeout of 200ms, but
// a timeout of 0 should generally not be provided.

// Benchlist

// Router

// Inbound Throttling

// Outbound Throttling

// HTTP APIs

// Enable/Disable APIs

// Health Checks

// Network Layer Health

// Router Health

// Staking
// Bind to all interfaces by default.

// Uptime Requirement

// Minimum Stake required to validate the Primary Network

// Maximum Stake that can be staked and delegated to a validator on the Primary Network

// Minimum Stake that can be delegated on the Primary Network

// Minimum Stake Duration

// Maximum Stake Duration

// Stake Reward Configs

// Subnets

// State syncing

// Bootstrapping
// TODO: combine "BootstrapIPsKey" and "BootstrapIDsKey" into one flag

// Snow Consensus

// Simplex Consensus

// ProposerVM

// Metrics

// Indexer

// Config Directories

// Chain Data Directory

// Profiles

// Aliasing

// Delays

// System resource trackers

// CPU management

// Disk management

// Opentelemetry tracing

// BuildFlagSet returns a complete set of flags for avalanchego
func BuildFlagSet() *pflag.FlagSet { _ = "STUB: not implemented"; return nil }

// getExpandedArg gets the string in viper corresponding to [key] and expands
// any variables using the OS env. If the [AvalancheGoDataDirVar] var is used,
// we expand the value of the variable with the string in viper corresponding to
// [DataDirKey].
func getExpandedArg(v *viper.Viper, key string) string { _ = "STUB: not implemented"; return "" }
