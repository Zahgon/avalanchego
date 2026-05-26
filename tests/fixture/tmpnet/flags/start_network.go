// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package flags

import (
	"github.com/spf13/pflag"

	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
)

type StartNetworkVars struct {
	// Accessible directly
	RootNetworkDir string
	NetworkOwner   string

	// Accessible via a validating method
	nodeCount   int
	runtimeVars *RuntimeConfigVars

	defaultNetworkOwner string
	defaultNodeCount    int
}

func NewStartNetworkFlagVars(defaultNetworkOwner string, defaultNodeCount int) *StartNetworkVars {
	_ = "STUB: not implemented"
	return nil
}

func NewStartNetworkFlagSetVars(flagSet *pflag.FlagSet, defaultNetworkOwner string, defaultNodeCount int) *StartNetworkVars {
	_ = "STUB: not implemented"
	return nil
}

func (v *StartNetworkVars) register(stringVar varFunc[string], intVar varFunc[int]) {
	_ = "STUB: not implemented"
	return
}

// An empty string prompts the use of the default path.

func (v *StartNetworkVars) GetNodeCount() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (v *StartNetworkVars) GetNodeRuntimeConfig() (*tmpnet.NodeRuntimeConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
