// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package registry

import (
	"errors"

	"github.com/ava-labs/avalanchego/api/metrics"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/filesystem"
	"github.com/ava-labs/avalanchego/utils/resource"
	"github.com/ava-labs/avalanchego/vms"
	"github.com/ava-labs/avalanchego/vms/rpcchainvm/runtime"
)

var (
	_ VMGetter = (*vmGetter)(nil)

	errInvalidVMID = errors.New("invalid vmID")
)

// VMGetter defines functionality to get the plugins on the node.
type VMGetter interface {
	// Get fetches the VMs that are registered and the VMs that are not
	// registered but available to be installed on the node.
	Get() (
		registeredVMs map[ids.ID]vms.Factory,
		unregisteredVMs map[ids.ID]vms.Factory,
		err error,
	)
}

// VMGetterConfig defines settings for VMGetter
type VMGetterConfig struct {
	FileReader      filesystem.Reader
	Manager         *vms.Manager
	PluginDirectory string
	CPUTracker      resource.ProcessTracker
	RuntimeTracker  runtime.Tracker
	MetricsGatherer metrics.MultiGatherer
}

type vmGetter struct {
	config VMGetterConfig
}

// NewVMGetter returns a new instance of a VMGetter
func NewVMGetter(config VMGetterConfig) VMGetter { _ = "STUB: not implemented"; return *new(VMGetter) }

func (getter *vmGetter) Get() (map[ids.ID]vms.Factory, map[ids.ID]vms.Factory, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Strip any extension from the file. This is to support windows .exe
// files.

// Skip hidden files.

// there is no alias with plugin name, try to use full vmID.

// If we already have the VM registered, we shouldn't attempt to
// register it again.

// If the error isn't "not found", then we should report the error.
