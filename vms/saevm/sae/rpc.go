// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sae

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/core/types"
	"github.com/ava-labs/libevm/ethdb"
	"github.com/ava-labs/libevm/event"

	"github.com/ava-labs/avalanchego/network/p2p"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/vms/saevm/blocks"
	"github.com/ava-labs/avalanchego/vms/saevm/hook"
	"github.com/ava-labs/avalanchego/vms/saevm/saexec"
	"github.com/ava-labs/avalanchego/vms/saevm/txgossip"

	saerpc "github.com/ava-labs/avalanchego/vms/saevm/sae/rpc"
	saetypes "github.com/ava-labs/avalanchego/vms/saevm/types"
)

// GethRPCBackends returns the backing infrastructure for geth's implementations
// of the JSON-RPC namespaces supported by the VM.
func (vm *VM) GethRPCBackends() saerpc.GethBackends {
	_ = "STUB: not implemented"
	return *new(saerpc.GethBackends)
}

func (vm *VM) chain() saerpc.Chain { _ = "STUB: not implemented"; return *new(saerpc.Chain) }

type chain struct {
	*VM
	*saexec.Executor
}

func (c chain) Logger() logging.Logger { _ = "STUB: not implemented"; return *new(logging.Logger) }
func (c chain) Hooks() hook.Points     { _ = "STUB: not implemented"; return *new(hook.Points) }
func (c chain) DB() ethdb.Database     { _ = "STUB: not implemented"; return *new(ethdb.Database) }
func (c chain) XDB() saetypes.ExecutionResults {
	_ = "STUB: not implemented"
	return *new(saetypes.ExecutionResults)
}
func (c chain) Mempool() *txgossip.Set      { _ = "STUB: not implemented"; return nil }
func (c chain) Peers() *p2p.Peers           { _ = "STUB: not implemented"; return nil }
func (c chain) LastAccepted() *blocks.Block { _ = "STUB: not implemented"; return nil }
func (c chain) LastSettled() *blocks.Block  { _ = "STUB: not implemented"; return nil }

func (c chain) ConsensusCriticalBlock(h common.Hash) (*blocks.Block, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c chain) NewBlock(eth *types.Block, parent, lastSettled *blocks.Block) (*blocks.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c chain) SubscribeAcceptedBlocks(ch chan<- *blocks.Block) event.Subscription {
	_ = "STUB: not implemented"
	return *new(event.Subscription)
}
