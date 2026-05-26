// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package main

import (
	"context"
	"flag"
	"fmt"
	"maps"
	"os"
	"os/exec"
	"slices"
	"strings"
	"time"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm"
	"github.com/ava-labs/avalanchego/tests"
)

var (
	blockDirArg        string
	currentStateDirArg string
	startBlockArg      uint64
	endBlockArg        uint64
	minWaitTimeArg     time.Duration
	maxWaitTimeArg     time.Duration
	configNameArg      string
	configBytesArg     []byte

	predefinedConfigs = map[string]string{
		"firewood": `{
			"state-scheme": "firewood",
			"snapshot-cache": 0,
			"pruning-enabled": true,
			"state-sync-enabled": false
		}`,
		"firewood-archive": `{
			"state-scheme": "firewood",
			"snapshot-cache": 0,
			"pruning-enabled": false,
			"state-sync-enabled": false
		}`,
	}
)

func init() {
	evm.RegisterAllLibEVMExtras()

	flag.StringVar(&blockDirArg, "block-dir", blockDirArg, "Block DB directory to read from during re-execution.")
	flag.StringVar(&currentStateDirArg, "current-state-dir", currentStateDirArg, "Current state directory including VM DB and Chain Data Directory for re-execution.")
	flag.Uint64Var(&startBlockArg, "start-block", 101, "Start block to begin execution (exclusive).")
	flag.Uint64Var(&endBlockArg, "end-block", 200, "End block to end execution (inclusive).")
	flag.DurationVar(&minWaitTimeArg, "min-wait-time", 20*time.Second, "Minimum amount of time to wait before crashing.")
	flag.DurationVar(&maxWaitTimeArg, "max-wait-time", 30*time.Second, "Maximum amount of time to wait before crashing.")

	predefinedConfigKeys := slices.Collect(maps.Keys(predefinedConfigs))
	predefinedConfigOptionsStr := fmt.Sprintf("[%s]", strings.Join(predefinedConfigKeys, ", "))
	flag.StringVar(&configNameArg, "config", configNameArg, fmt.Sprintf("Specifies the predefined config to use for the VM. Options include %s.", predefinedConfigOptionsStr))

	flag.Parse()

	predefinedConfigStr, ok := predefinedConfigs[configNameArg]
	if !ok {
		fmt.Fprintf(os.Stderr, "invalid config name %q. Valid options include %s.\n", configNameArg, predefinedConfigOptionsStr)
		os.Exit(1)
	}
	configBytesArg = []byte(predefinedConfigStr)
}

func main() {
	tc := tests.NewTestContext(tests.NewDefaultLogger("chaos-test"))
	tc.SetDefaultContextParent(context.Background())
	tc.RecoverAndExit()

	run(
		tc,
		minWaitTimeArg,
		maxWaitTimeArg,
		blockDirArg,
		currentStateDirArg,
		startBlockArg,
		endBlockArg,
		configNameArg,
		configBytesArg,
	)
}

// run executes a chaos test that simulates an application crash during C-Chain
// block reexecution that uses Firewood. It verifies that the VM can recover from
// an unexpected termination and resume processing from the correct block height
// using persisted state.
//
// Running the chaos test involves a few steps:
//  1. Start a reexecution test process using the Firewood state scheme
//  2. Allow the reexecution test to run for the specified wait duration
//  3. Forcefully terminate the process with SIGKILL to simulate a crash
//  4. Open the VM database to read the last accepted block height from persisted state
//  5. Restart the reexecution test from the recovered height to verify state consistency
func run(
	tc tests.TestContext,
	minWaitTime time.Duration,
	maxWaitTime time.Duration,
	blockDir string,
	currentStateDir string,
	startBlock uint64,
	endBlock uint64,
	configName string,
	configBytes []byte,
) {
	_ = "STUB: not implemented"
	return
}

// Set process group ID so we can kill all child processes

// 1. Start a reexecution test process using the Firewood state scheme

// 2. Allow the reexecution test to run for the specified wait duration

// 3. Forcefully terminate the process with SIGKILL to simulate a crash

// ExitCode() returns -1 when killed by signal

// 4. Open the VM database to read the last accepted block height from persisted state

// 5. Restart the reexecution test from the recovered height to verify state consistency

// openDB attempts to open a LevelDB database with retry logic and linear backoff.
// This is necessary after killing a process that held the database open, as the OS may
// need time to release file locks even after the process terminates.
//
// The backoff strategy increases by 500ms per attempt (500ms, 1s, 1.5s, 2s, ...).
func openDB(dbDir string, maxAttempts int) (database.Database, error) {
	_ = "STUB: not implemented"
	return *new(database.Database), nil
}

// createReexecutionCmd constructs a command to run the C-Chain reexecution test.
func createReexecutionCmd(
	blockDir string,
	currentStateDir string,
	startBlock uint64,
	endBlock uint64,
	configName string,
) *exec.Cmd {
	_ = "STUB: not implemented"
	return nil
}
