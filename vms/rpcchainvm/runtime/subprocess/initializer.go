// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package subprocess

import (
	"context"
	"sync"

	"github.com/ava-labs/avalanchego/vms/rpcchainvm/runtime"
)

var _ runtime.Initializer = (*initializer)(nil)

// Subprocess VM Runtime initializer.
type initializer struct {
	path string

	once sync.Once
	// Address of the RPC Chain VM server
	vmAddr string
	// Error, if one occurred, during Initialization
	err error
	// Initialized is closed once Initialize is called
	initialized chan struct{}
}

func newInitializer(path string) *initializer { _ = "STUB: not implemented"; return nil }

func (i *initializer) Initialize(_ context.Context, protocolVersion uint, vmAddr string) error {
	_ = "STUB: not implemented"
	return nil
}
