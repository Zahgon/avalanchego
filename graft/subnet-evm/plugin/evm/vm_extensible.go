// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"errors"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/database/versiondb"
	"github.com/ava-labs/avalanchego/graft/evm/sync/engine"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/core"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/params"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/plugin/evm/config"
	"github.com/ava-labs/avalanchego/graft/subnet-evm/plugin/evm/extension"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p"
)

var _ extension.InnerVM = (*VM)(nil)

var (
	errVMAlreadyInitialized      = errors.New("vm already initialized")
	errExtensionConfigAlreadySet = errors.New("extension config already set")
)

func (vm *VM) SetExtensionConfig(config *extension.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// All these methods below assumes that VM is already initialized

func (vm *VM) GetExtendedBlock(ctx context.Context, blkID ids.ID) (extension.ExtendedBlock, error) {
	_ = "STUB: not implemented"
	// Since each internal handler used by [vm.State] always returns a block
	// with non-nil ethBlock value, GetBlockInternal should never return a
	// (*Block) with a nil ethBlock value.
	return *new(extension.ExtendedBlock), nil
}

func (vm *VM) LastAcceptedExtendedBlock() extension.ExtendedBlock {
	_ = "STUB: not implemented"
	return *new(extension.ExtendedBlock)
}

// ChainConfig returns the chain config for the VM
// Even though this is available through Blockchain().Config(),
// ChainConfig() here will be available before the blockchain is initialized.
func (vm *VM) ChainConfig() *params.ChainConfig { _ = "STUB: not implemented"; return nil }

func (vm *VM) Blockchain() *core.BlockChain { _ = "STUB: not implemented"; return nil }

func (vm *VM) Config() config.Config { _ = "STUB: not implemented"; return *new(config.Config) }

func (vm *VM) MetricRegistry() *prometheus.Registry { _ = "STUB: not implemented"; return nil }

func (vm *VM) Validators() *p2p.Validators { _ = "STUB: not implemented"; return nil }

func (vm *VM) VersionDB() *versiondb.Database { _ = "STUB: not implemented"; return nil }

func (vm *VM) SyncerClient() engine.Client { _ = "STUB: not implemented"; return *new(engine.Client) }
