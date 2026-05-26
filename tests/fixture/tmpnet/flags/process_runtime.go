// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package flags

import (
	"fmt"

	"github.com/spf13/pflag"

	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
)

const (
	processRuntime   = "process"
	processDocPrefix = "[process runtime] "

	avalanchegoPathFlag = "avalanchego-path"
)

var errAvalancheGoRequired = fmt.Errorf("--%s or %s are required", avalanchegoPathFlag, tmpnet.AvalancheGoPathEnvName)

type processRuntimeVars struct {
	config tmpnet.ProcessRuntimeConfig
}

func (v *processRuntimeVars) registerWithFlag() { _ = "STUB: not implemented"; return }

func (v *processRuntimeVars) registerWithFlagSet(flagSet *pflag.FlagSet) {
	_ = "STUB: not implemented"
	return
}

func (v *processRuntimeVars) register(stringVar varFunc[string], boolVar varFunc[bool]) {
	_ = "STUB: not implemented"
	return
}

func (v *processRuntimeVars) getProcessRuntimeConfig() (*tmpnet.ProcessRuntimeConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *processRuntimeVars) validate() error { _ = "STUB: not implemented"; return nil }

// A relative path must be resolvable to an absolute path

// The absolute path must exist
