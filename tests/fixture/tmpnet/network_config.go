// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
)

// The Network type is defined in this file (reading/writing configuration) and network.go
// (orchestration).

var errMissingNetworkDir = errors.New("failed to write network: missing network directory")

// Read network and node configuration from disk.
func (n *Network) Read(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Write network configuration to disk.
func (n *Network) Write() error { _ = "STUB: not implemented"; return nil }

// Read network configuration from disk.
func (n *Network) readNetwork() error { _ = "STUB: not implemented"; return nil }

// Read the nodes associated with the network from disk.
func (n *Network) readNodes(ctx context.Context) error {
	_ = "STUB: not implemented"

	// Node configuration is stored in child directories
	return nil
}

// If no config file exists, assume this is not the path of a node

func (n *Network) writeNodes() error { _ = "STUB: not implemented"; return nil }

// For consumption outside of avalanchego. Needs to be kept exported.
func (n *Network) GetGenesisPath() string { _ = "STUB: not implemented"; return "" }

func (n *Network) readGenesis() error { _ = "STUB: not implemented"; return nil }

func (n *Network) writeGenesis() error { _ = "STUB: not implemented"; return nil }

func (n *Network) getConfigPath() string { _ = "STUB: not implemented"; return "" }

func (n *Network) readConfig() error { _ = "STUB: not implemented"; return nil }

// The subset of network fields to store in the network config file.
type serializedNetworkConfig struct {
	UUID                 string                  `json:"uuid,omitempty"`
	Owner                string                  `json:"owner,omitempty"`
	NetworkID            uint32                  `json:"networkID,omitempty"`
	PrimarySubnetConfig  ConfigMap               `json:"primarySubnetConfig,omitempty"`
	PrimaryChainConfigs  map[string]ConfigMap    `json:"primaryChainConfigs,omitempty"`
	DefaultFlags         FlagsMap                `json:"defaultFlags,omitempty"`
	DefaultRuntimeConfig NodeRuntimeConfig       `json:"defaultRuntimeConfig,omitempty"`
	PreFundedKeys        []*secp256k1.PrivateKey `json:"preFundedKeys,omitempty"`
}

func (n *Network) writeNetworkConfig() error { _ = "STUB: not implemented"; return nil }

func (n *Network) EnvFilePath() string { _ = "STUB: not implemented"; return "" }

func (n *Network) EnvFileContents() string { _ = "STUB: not implemented"; return "" }

// Write an env file that sets the network dir env when sourced.
func (n *Network) writeEnvFile() error { _ = "STUB: not implemented"; return nil }

func (n *Network) GetSubnetDir() string { _ = "STUB: not implemented"; return "" }

func (n *Network) readSubnets() error { _ = "STUB: not implemented"; return nil }
