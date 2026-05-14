// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package sae

import (
	"github.com/ava-labs/libevm/core"
	"github.com/ava-labs/libevm/event"
	"github.com/ava-labs/libevm/libevm"
)

// LastExecutedState returns a [libevm.StateReader] backed by the post-execution
// state of the last-executed block.
func (vm *VM) LastExecutedState() (libevm.StateReader, error) {
	return vm.exec.StateDB(vm.exec.LastExecuted().PostExecutionStateRoot())
}

// SubscribeChainHeadEvent returns a new subscription for each
// [core.ChainHeadEvent] emitted after a block has been executed.
func (vm *VM) SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription {
	return vm.exec.SubscribeChainHeadEvent(ch)
}
