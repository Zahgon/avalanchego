// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package flags

import (
	"github.com/spf13/pflag"

	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
)

var validRuntimes = []string{
	processRuntime,
	kubeRuntime,
}

type RuntimeConfigVars struct {
	runtime            string
	processRuntimeVars processRuntimeVars
	kubeRuntimeVars    kubeRuntimeVars
}

// NewRuntimeConfigFlagVars registers runtime config flag variables for stdlib flag
func NewRuntimeConfigFlagVars() *RuntimeConfigVars { _ = "STUB: not implemented"; return nil }

// NewRuntimeConfigFlagSetVars registers runtime config flag variables for pflag
func NewRuntimeConfigFlagSetVars(flagSet *pflag.FlagSet) *RuntimeConfigVars {
	_ = "STUB: not implemented"
	return nil
}

func (v *RuntimeConfigVars) register(stringVar varFunc[string]) { _ = "STUB: not implemented"; return }

// GetNodeRuntimeConfig returns the validated node runtime config
func (v *RuntimeConfigVars) GetNodeRuntimeConfig() (*tmpnet.NodeRuntimeConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
