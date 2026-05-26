// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vmtest

import (
	"context"
	"sync"
	"testing"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/rawdb"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb"

	"github.com/ava-labs/avalanchego/graft/coreth/consensus/dummy"
	"github.com/ava-labs/avalanchego/graft/coreth/core"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/extension"
	"github.com/ava-labs/avalanchego/graft/evm/utils/utilstest"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/enginetest"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/vms/evm/sync/customrawdb"

	avalancheatomic "github.com/ava-labs/avalanchego/chains/atomic"
	avalanchedatabase "github.com/ava-labs/avalanchego/database"
)

var schemes = []string{rawdb.HashScheme, customrawdb.FirewoodScheme}

type SyncerVMTest struct {
	Name     string
	TestFunc func(
		t *testing.T,
		testSetup *SyncTestSetup,
	)
}

var SyncerVMTests = []SyncerVMTest{
	{
		Name:     "SkipStateSyncTest",
		TestFunc: SkipStateSyncTest,
	},
	{
		Name:     "StateSyncFromScratchTest",
		TestFunc: StateSyncFromScratchTest,
	},
	{
		Name:     "StateSyncFromScratchExceedParentTest",
		TestFunc: StateSyncFromScratchExceedParentTest,
	},
	{
		Name:     "StateSyncToggleEnabledToDisabledTest",
		TestFunc: StateSyncToggleEnabledToDisabledTest,
	},
	{
		Name:     "VMShutdownWhileSyncingTest",
		TestFunc: VMShutdownWhileSyncingTest,
	},
}

func SkipStateSyncTest(t *testing.T, testSetup *SyncTestSetup) { _ = "STUB: not implemented"; return }

// must be greater than [syncableInterval] to skip sync

func StateSyncFromScratchTest(t *testing.T, testSetup *SyncTestSetup) {
	_ = "STUB: not implemented"
	return
}

// must be less than [syncableInterval] to perform sync

func StateSyncFromScratchExceedParentTest(t *testing.T, testSetup *SyncTestSetup) {
	_ = "STUB: not implemented"
	return
}

// must be less than [syncableInterval] to perform sync

func StateSyncToggleEnabledToDisabledTest(t *testing.T, testSetup *SyncTestSetup) {
	_ = "STUB: not implemented"
	return
}

func stateSyncToggleEnabledToDisabledTest(t *testing.T, testSetup *SyncTestSetup, scheme string) {
	_ = "STUB: not implemented"
	return
}

// must be less than SyncableInterval to perform sync

// Fail all requests after maxRequestsBeforeShutdown to interrupt the sync
// TODO(alarso16): Changing this value may cause the test to fail.
// Syncer cannot know whether it failed or not, so it relies on state root matching.

// Perform sync resulting in early termination.

// Process some blocks to prove state is valid

// Verify the snapshot disk layer matches the last block root

func VMShutdownWhileSyncingTest(t *testing.T, testSetup *SyncTestSetup) {
	_ = "STUB: not implemented"
	return
}

// must be less than SyncableInterval to perform sync

// Shutdown the VM after maxRequests to interrupt the sync

// Note this verifies the VM shutdown does not time out while syncing.

// Perform sync resulting in early termination.

type SyncTestSetup struct {
	NewVM             func() (extension.InnerVM, dummy.ConsensusCallbacks) // should not be initialized
	AfterInit         func(t *testing.T, testParams SyncTestParams, vmSetup SyncVMSetup, isServer bool)
	GenFn             func(i int, vm extension.InnerVM, gen *core.BlockGen)
	ExtraSyncerVMTest func(t *testing.T, syncerVM SyncVMSetup)
}

func initSyncServerAndClientVMs(t *testing.T, test SyncTestParams, numBlocks int, testSetup *SyncTestSetup) *testSyncVMSetup {
	_ = "STUB: not implemented"
	return nil

	// override commitInterval so the call to trie creates a commit at the height [syncableInterval].
	// This is necessary to support fetching a state summary.
}

// make some accounts

// must be set for FillAccountsWithStorageAndCode to work

// patch serverVM's lastAcceptedBlock to have the new root
// and update the vm's state so the trie with accounts will
// be returned by StateSyncGetLastSummary

// initialise [syncerVM] with blank genesis state
// we also override [syncerVM]'s commit interval so the atomic trie works correctly.

// override [serverVM]'s SendAppResponse function to trigger AppResponse on [syncerVM]

// connect peer to [syncerVM]

// override syncerVM's SendAppRequest function to trigger AppRequest on serverVM

// testSyncVMSetup contains the required set up for a client VM to perform state sync
// off of a server VM.
type testSyncVMSetup struct {
	serverVM SyncVMSetup
	syncerVM syncerVMSetup

	fundedAccounts map[*utilstest.Key]*types.StateAccount
}

type SyncVMSetup struct {
	VM                 extension.InnerVM
	SnowCtx            *snow.Context
	ConsensusCallbacks dummy.ConsensusCallbacks
	DB                 avalanchedatabase.Database
	AtomicMemory       *avalancheatomic.Memory
	AppSender          *enginetest.Sender
}

type syncerVMSetup struct {
	SyncVMSetup
	shutdownOnceSyncerVM *shutdownOnceVM
}

type shutdownOnceVM struct {
	extension.InnerVM
	shutdownOnce sync.Once
}

func (vm *shutdownOnceVM) Shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// SyncTestParams contains both the actual VMs as well as the parameters with the expected output.
type SyncTestParams struct {
	responseIntercept  func(vm extension.InnerVM, nodeID ids.NodeID, requestID uint32, response []byte)
	StateSyncMinBlocks uint64
	SyncableInterval   uint64
	SyncMode           block.StateSyncMode
	StateScheme        string
	expectedErr        error
}

func testSyncerVM(t *testing.T, testSyncVMSetup *testSyncVMSetup, test SyncTestParams, extraSyncerVMTest func(t *testing.T, syncerVMSetup SyncVMSetup)) {
	_ = "STUB: not implemented"
	return
}

// get last summary and test related methods

// If the test is expected to error, require that the correct error is returned and finish the test.

// Note we re-open the database here to avoid a closed error when the test is for a shutdown VM.
// TODO: this avoids circular dependencies but is not ideal.

// set [syncerVM] to bootstrapping and verify the last accepted block has been updated correctly
// and that we can bootstrap and process some blocks.

// check the last block is indexed

// tail should be the last block synced

// arbitrary choice

// tail should be the minimum last synced block, since we skipped it to the last block

// check we can transition to [NormalOp] state and continue to process blocks.

// Generate blocks after we have entered normal consensus as well

// tail should be the minimum last synced block, since we skipped it to the last block

// patchBlock returns a copy of [blk] with [root] and updates [db] to
// include the new block as canonical for [blk]'s height.
// This breaks the digestibility of the chain since after this call
// [blk] does not necessarily define a state transition from its parent
// state to the new state root.
func patchBlock(blk *types.Block, root common.Hash, db ethdb.Database) *types.Block {
	_ = "STUB: not implemented"
	return nil
}

// generateAndAcceptBlocks uses [core.GenerateChain] to generate blocks, then
// calls Verify and Accept on each generated block
// TODO: consider using this helper function in vm_test.go and elsewhere in this package to clean up tests
func generateAndAcceptBlocks(t *testing.T, vm extension.InnerVM, numBlocks int, gen func(int, extension.InnerVM, *core.BlockGen), accepted func(*types.Block), cb dummy.ConsensusCallbacks) {
	_ = "STUB: not implemented"
	return
}

// acceptExternalBlock defines a function to parse, verify, and accept a block once it has been
// generated by GenerateChain

// We must not commit this state to disk, as it will make it impossible to verify/accept the generated blocks in the same db for Firewood.
// Firewood requires a linear and in-order history. All unused states will be cleaned at commit time.

// necessary for syntactic validation of the block

// requireSyncPerformedHeight verifies the latest sync performed height matches expectations.
// Pass 0 to verify no sync was performed.
func requireSyncPerformedHeight(t *testing.T, db ethdb.KeyValueStore, expected uint64) {
	_ = "STUB: not implemented"
	return
}
