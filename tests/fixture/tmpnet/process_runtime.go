// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"context"
	"errors"
	"net/netip"
	"os"
	"time"

	"github.com/ava-labs/avalanchego/config"
	"github.com/ava-labs/avalanchego/config/node"
	"github.com/ava-labs/avalanchego/utils/logging"
)

const (
	AvalancheGoPathEnvName = "AVALANCHEGO_PATH"

	defaultNodeInitTimeout = 10 * time.Second
)

var (
	AvalancheGoPluginDirEnvName = config.EnvVarName(config.EnvPrefix, config.PluginDirKey)

	errNodeAlreadyRunning = errors.New("failed to start node: node is already running")
	errNotRunning         = errors.New("node is not running")
)

type ProcessRuntimeConfig struct {
	AvalancheGoPath   string `json:"avalancheGoPath,omitempty"`
	PluginDir         string `json:"pluginDir,omitempty"`
	ReuseDynamicPorts bool   `json:"reuseDynamicPorts,omitempty"`
}

// Defines local-specific node configuration. Supports setting default
// and node-specific values.
type ProcessRuntime struct {
	node *Node

	// PID of the node process
	pid int
}

func (p *ProcessRuntime) getRuntimeConfig() *ProcessRuntimeConfig {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProcessRuntime) setProcessContext(processContext node.ProcessContext) {
	_ = "STUB: not implemented"
	return
}

func (p *ProcessRuntime) readState(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// The absence of the process context file indicates the node is not running

// Start waits for the process context to be written which
// indicates that the node will be accepting connections on
// its staking port. The network will start faster with this
// synchronization due to the avoidance of exponential backoff
// if a node tries to connect to a beacon that is not ready.
func (p *ProcessRuntime) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Avoid attempting to start an already running node.

// Attempt to check for rpc version compatibility

// Ensure a stale process context file is removed so that the
// creation of a new file can indicate node start.

// All arguments are provided in the flags file

// Ensure process is detached from the parent process so that an error in the parent will not affect the child

// Watch the node's main.log file in the background for FATAL log entries that indicate
// a configuration error preventing startup. Such a log entry will be provided to the
// cancelWithCause function so that waitForProcessContext can exit early with an error
// that includes the log entry.

// A node writes a process context file on start. If the file is not
// found in a reasonable amount of time, the node is unlikely to have
// started successfully.

// Configure collection of metrics and logs

func (p *ProcessRuntime) writeFlags() error { _ = "STUB: not implemented"; return nil }

// Use dynamic port allocation

// Binding to localhost on macos avoids having a permission
// dialog pop up for every node that tries to bind to
// non-localhost interfaces

// Only configure the plugin dir with a non-empty value to ensure the use of the
// default value (`[datadir]/plugins`) when no plugin dir is configured.

// Ensure a non-empty plugin directory exists or the node will fail to start.

// Signals the node process to stop.
func (p *ProcessRuntime) InitiateStop(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Already stopped

// Waits for the node process to stop.
func (p *ProcessRuntime) WaitForStopped(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Restarts the node
func (p *ProcessRuntime) Restart(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Attempt to save the API port currently being used so the
// restarted node can reuse it. This may result in the node
// failing to start if the operating system allocates the port
// to a different process between node stop and start.

func (p *ProcessRuntime) IsHealthy(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	// Check that the node process is running as a precondition for
	// checking health. getProcess will also ensure that the node's
	// API URI is current.
	return false, nil
}

func (p *ProcessRuntime) getProcessContextPath() string { _ = "STUB: not implemented"; return "" }

func (p *ProcessRuntime) waitForProcessContext(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// Retrieve the node process if it is running. As part of determining
// process liveness, the node's process context will be refreshed if
// live or cleared if not running.
func (p *ProcessRuntime) getProcess() (*os.Process, error) {
	_ = "STUB: not implemented"
	// This context is not used but a non-nil value must be supplied to satisfy the linter
	return nil, nil
}

// Read the process context to ensure freshness. The node may have
// stopped or been restarted since last read.

// Process is not running

// getProcess retrieves the process if it is running.
func getProcess(pid int) (*os.Process, error) { _ = "STUB: not implemented"; return nil, nil }

// Sending 0 will not actually send a signal but will perform
// error checking.

// Process is running

// Process is not running

// Write monitoring configuration enabling collection of metrics and logs from the node.
func (p *ProcessRuntime) writeMonitoringConfig() error {
	_ = "STUB: not implemented"
	// Ensure labeling that uniquely identifies the node and its network
	return nil
}

// Return the path for this node's prometheus configuration.
func (p *ProcessRuntime) getMonitoringConfigPath(name string) (string, error) {
	_ = "STUB: not implemented"
	// Ensure a unique filename to allow config files to be added and removed
	// by multiple nodes without conflict.
	return "", nil
}

// Ensure the removal of the monitoring configuration files for this node.
func (p *ProcessRuntime) removeMonitoringConfig() error { _ = "STUB: not implemented"; return nil }

// Write the configuration for a type of monitoring (e.g. prometheus, promtail).
func (p *ProcessRuntime) writeMonitoringConfigFile(name string, config []ConfigMap) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAccessibleURI returns the URI that can be used to access the node's API.
func (p *ProcessRuntime) GetAccessibleURI() string { _ = "STUB: not implemented"; return "" }

func (p *ProcessRuntime) GetAccessibleStakingAddress(_ context.Context) (netip.AddrPort, func(), error) {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort), nil, nil
}

// Saves the currently allocated API port to the node's configuration
// for use across restarts.
func (p *ProcessRuntime) saveAPIPort() error { _ = "STUB: not implemented"; return nil }

// Without an API URI there is nothing to save

// watchLogFileForFatal waits for the specified file path to exist and then checks each of
// its lines for the string 'FATAL' until such a line is observed or the provided context
// is canceled. If line containing 'FATAL' is encountered, it will be provided as an error
// to the provided cancelWithCause function.
//
// Errors encountered while looking for FATAL log entries are considered potential rather
// than positive indications of failure and are printed to the provided writer instead of
// being provided to the cancelWithCause function.
func watchLogFileForFatal(ctx context.Context, cancelWithCause context.CancelCauseFunc, log logging.Logger, path string) {
	_ = "STUB: not implemented"
	return
}

// Wait for the file to exist

// File does not exist yet - wait and try again

// Open the file

// Scan for lines in the file containing 'FATAL'

// Read a line from the file

// If end of file is reached, wait and try again
