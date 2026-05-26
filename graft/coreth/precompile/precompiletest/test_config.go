// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package precompiletest

import (
	"testing"

	"github.com/ava-labs/avalanchego/graft/coreth/precompile/precompileconfig"
)

// ConfigVerifyTest is a test case for verifying a config
type ConfigVerifyTest struct {
	Config        precompileconfig.Config
	ChainConfig   precompileconfig.ChainConfig
	ExpectedError error
}

// ConfigEqualTest is a test case for comparing two configs
type ConfigEqualTest struct {
	Config   precompileconfig.Config
	Other    precompileconfig.Config
	Expected bool
}

func RunVerifyTests(t *testing.T, tests map[string]ConfigVerifyTest) {
	_ = "STUB: not implemented"
	return
}

func RunEqualTests(t *testing.T, tests map[string]ConfigEqualTest) {
	_ = "STUB: not implemented"
	return
}
