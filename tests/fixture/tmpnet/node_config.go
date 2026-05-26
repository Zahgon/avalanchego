// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"context"
)

// The Node type is defined in this file node_config.go
// (reading/writing configuration) and node.go (orchestration).

// For consumption outside of avalanchego. Needs to be kept exported.
func (n *Node) GetFlagsPath() string { _ = "STUB: not implemented"; return "" }

func (n *Node) getConfigPath() string { _ = "STUB: not implemented"; return "" }

func (n *Node) readConfig() error { _ = "STUB: not implemented"; return nil }

type serializedNodeConfig struct {
	IsEphemeral   bool               `json:"isEphemeral,omitempty"`
	Flags         FlagsMap           `json:"flags,omitempty"`
	RuntimeConfig *NodeRuntimeConfig `json:"runtimeConfig,omitempty"`
}

func (n *Node) writeConfig() error { _ = "STUB: not implemented"; return nil }

func (n *Node) Read(ctx context.Context, network *Network, dataDir string) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *Node) Write() error { _ = "STUB: not implemented"; return nil }

func (n *Node) writeMetricsSnapshot(data []byte) error { _ = "STUB: not implemented"; return nil }

// Create a compatible filesystem from the current timestamp
