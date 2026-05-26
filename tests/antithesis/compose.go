// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package antithesis

import (
	"errors"

	"github.com/compose-spec/compose-go/types"

	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
)

const bootstrapIndex = 0

const (
	targetPathEnvName = "TARGET_PATH"
	imageTagEnvName   = "IMAGE_TAG"
)

var (
	errTargetPathEnvVarNotSet = errors.New(targetPathEnvName + " environment variable not set")
	errImageTagEnvVarNotSet   = errors.New(imageTagEnvName + " environment variable not set")
	errAvalancheGoEvVarNotSet = errors.New(tmpnet.AvalancheGoPathEnvName + " environment variable not set")
	errPluginDirEnvVarNotSet  = errors.New(tmpnet.AvalancheGoPluginDirEnvName + " environment variable not set")
)

// Creates docker compose configuration for an antithesis test setup. Configuration is via env vars to
// simplify usage by main entrypoints. If the provided network includes a subnet, the initial DB state for
// the subnet will be created and written to the target path.
func GenerateComposeConfig(network *tmpnet.Network, baseImageName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Subnet testing requires creating an initial db state for the bootstrap node

// Plugin dir configured here is only used for initializing the bootstrap db.

// Initialize the given path with the docker compose configuration (compose file and
// volumes) needed for an Antithesis test setup.
func initComposeConfig(
	network *tmpnet.Network,
	nodeImageName string,
	workloadImageName string,
	targetPath string,
) error {
	_ = "STUB: not implemented"
	// Generate a compose project for the specified network
	return nil
}

// Write the compose file

// Create the volume paths

// Create a new docker compose project for an antithesis test setup
// for the provided network configuration.
func newComposeProject(network *tmpnet.Network, nodeImageName string, workloadImageName string) (*types.Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Apply configuration appropriate to a test network

// DB volume for bootstrap node will need to initialized with the subnet

// The env is defined with the keys and then converted to env
// vars because only the keys are available as constants.

// Collect URIs for the workload container

// Convert a mapping of avalanche config keys to a mapping of env vars
func keyMapToEnvVarMap(keyMap types.Mapping) types.Mapping {
	_ = "STUB: not implemented"
	return *new(types.Mapping)
}

// Retrieve the service name for a node at the given index. Common to
// GenerateComposeConfig and InitDBVolumes to ensure consistency
// between db volumes configuration and volume paths.
func getServiceName(index int) string { _ = "STUB: not implemented"; return "" }
