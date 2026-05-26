// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vmtest

import (
	"testing"

	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/database/prefixdb"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/extension"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/enginetest"
	"github.com/ava-labs/avalanchego/upgrade/upgradetest"
	"github.com/ava-labs/avalanchego/vms/evm/sync/customrawdb"

	avalancheatomic "github.com/ava-labs/avalanchego/chains/atomic"
	commoneng "github.com/ava-labs/avalanchego/snow/engine/common"
)

var Schemes = []string{rawdb.HashScheme, customrawdb.FirewoodScheme}

type TestVMConfig struct {
	IsSyncing bool
	Fork      *upgradetest.Fork
	// If genesisJSON is empty, defaults to the genesis corresponding to the
	// fork.
	GenesisJSON string
	ConfigJSON  string
	// DB scheme, defaults to HashScheme
	Scheme string
}

type TestVMSuite struct {
	VM           commoneng.VM
	DB           *prefixdb.Database
	AtomicMemory *avalancheatomic.Memory
	AppSender    *enginetest.Sender
	Ctx          *snow.Context
}

// SetupTestVM initializes a VM for testing. It sets up the genesis and returns the
// issuer channel, database, atomic memory, app sender, and context.
// Expects the passed VM to be a uninitialized VM.
func SetupTestVM(t *testing.T, vm commoneng.VM, config TestVMConfig) *TestVMSuite {
	_ = "STUB: not implemented"
	return nil
}

// ResetMetrics resets the vm avalanchego metrics, and allows
// for the VM to be re-initialized in tests.
func ResetMetrics(snowCtx *snow.Context) { _ = "STUB: not implemented"; return }

func OverrideSchemeConfig(scheme string, configJSON string) (string, error) {
	_ = "STUB: not implemented"
	// If the scheme is not Firewood, return the configJSON as is
	return "", nil
}

// Parse existing config into a map to preserve only non-zero values

// Set Firewood-specific configuration flags (these will override any existing values)

// Marshal back to JSON

func IssueTxsAndBuild(txs []*types.Transaction, vm extension.InnerVM) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

func IssueTxsAndSetPreference(txs []*types.Transaction, vm extension.InnerVM) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}
