// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package allowlisttest

import (
	"testing"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/allowlist"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/modules"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompileconfig"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompiletest"
)

// mkConfigWithAllowList creates a new config with the correct type for [module]
// by marshalling [cfg] to JSON and then unmarshalling it into the config.
func mkConfigWithAllowList(module modules.Module, cfg *allowlist.AllowListConfig) precompileconfig.Config {
	_ = "STUB: not implemented"
	return *new(precompileconfig.Config)
}

func mkConfigWithUpgradeAndAllowList(module modules.Module, cfg *allowlist.AllowListConfig, update precompileconfig.Upgrade) precompileconfig.Config {
	_ = "STUB: not implemented"
	return *new(precompileconfig.Config)
}

func AllowListConfigVerifyTests(t testing.TB, module modules.Module) map[string]precompiletest.ConfigVerifyTest {
	_ = "STUB: not implemented"
	return nil
}

func AllowListConfigEqualTests(_ testing.TB, module modules.Module) map[string]precompiletest.ConfigEqualTest {
	_ = "STUB: not implemented"
	return nil
}

func VerifyPrecompileWithAllowListTests(t *testing.T, module modules.Module, verifyTests map[string]precompiletest.ConfigVerifyTest) {
	_ = "STUB: not implemented"
	return
}

// Add the contract specific tests to the map of tests to run.

func EqualPrecompileWithAllowListTests(t *testing.T, module modules.Module, equalTests map[string]precompiletest.ConfigEqualTest) {
	_ = "STUB: not implemented"
	return
}

// Add the contract specific tests to the map of tests to run.
