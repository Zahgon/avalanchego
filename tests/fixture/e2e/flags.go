// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package e2e

import (
	"time"

	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet"
	"github.com/ava-labs/avalanchego/tests/fixture/tmpnet/flags"
)

type NetworkCmd int

const (
	EmptyNetworkCmd NetworkCmd = iota
	StartNetworkCmd
	StopNetworkCmd
	RestartNetworkCmd
	ReuseNetworkCmd
)

type FlagVars struct {
	startNetwork     bool
	startNetworkVars *flags.StartNetworkVars

	collectorVars *flags.CollectorVars

	checkMetricsCollected bool
	checkLogsCollected    bool

	networkDir     string
	reuseNetwork   bool
	stopNetwork    bool
	restartNetwork bool

	activateLatest bool
}

func (v *FlagVars) NetworkCmd() (NetworkCmd, error) {
	_ = "STUB: not implemented"
	return *new(NetworkCmd), nil
}

func (v *FlagVars) RootNetworkDir() string { _ = "STUB: not implemented"; return "" }

func (v *FlagVars) NetworkOwner() string { _ = "STUB: not implemented"; return "" }

func (v *FlagVars) NodeCount() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (v *FlagVars) NodeRuntimeConfig() (*tmpnet.NodeRuntimeConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *FlagVars) StartMetricsCollector() bool { _ = "STUB: not implemented"; return false }

func (v *FlagVars) StartLogsCollector() bool { _ = "STUB: not implemented"; return false }

func (v *FlagVars) CheckMetricsCollected() bool { _ = "STUB: not implemented"; return false }

func (v *FlagVars) CheckLogsCollected() bool { _ = "STUB: not implemented"; return false }

func (v *FlagVars) NetworkDir() string { _ = "STUB: not implemented"; return "" }

func (v *FlagVars) NetworkShutdownDelay() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Only return a non-zero value if we want to ensure the collectors have
// a chance to collect the metrics at the end of the test.

func (v *FlagVars) ActivateLatest() bool { _ = "STUB: not implemented"; return false }

type DefaultOption func(*DefaultOptions)

type DefaultOptions struct {
	owner     string
	nodeCount int
}

func newDefaultOptions(ops []DefaultOption) *DefaultOptions { _ = "STUB: not implemented"; return nil }

func (d *DefaultOptions) Owner() string { _ = "STUB: not implemented"; return "" }

func (d *DefaultOptions) NodeCount() int { _ = "STUB: not implemented"; return 0 }

func WithDefaultOwner(owner string) DefaultOption {
	_ = "STUB: not implemented"
	return *new(DefaultOption)
}

func WithDefaultNodeCount(nodeCount int) DefaultOption {
	_ = "STUB: not implemented"
	return *new(DefaultOption)
}

func RegisterFlags(ops ...DefaultOption) *FlagVars { _ = "STUB: not implemented"; return nil }

func SetCheckCollectionFlags(checkMetricsCollected *bool, checkLogsCollected *bool) {
	_ = "STUB: not implemented"
	return
}
