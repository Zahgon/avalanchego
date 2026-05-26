// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/vms/rpcchainvm/grpcutils"
)

const defaultRuntimeDialTimeout = 5 * time.Second

// The address of the Runtime server is expected to be passed via ENV `runtime.EngineAddressKey`.
// This address is used by the Runtime client to send Initialize RPC to server.
//
// Serve starts the RPC Chain VM server and performs a handshake with the VM runtime service.
func Serve(ctx context.Context, vm block.ChainVM, opts ...grpcutils.ServerOption) error {
	_ = "STUB: not implemented"
	return nil
}

// We drop all signals until our parent process has notified us
// that we are shutting down. Once we are in the shutdown
// workflow, we will gracefully exit upon receiving a SIGTERM.

// address of Runtime server from ENV

// start RPC Chain VM server

// Returns an RPC Chain VM server serving health and VM services.
func newVMServer(vm block.ChainVM, allowShutdown *utils.Atomic[bool], opts ...grpcutils.ServerOption) *grpc.Server {
	_ = "STUB: not implemented"
	return nil
}
