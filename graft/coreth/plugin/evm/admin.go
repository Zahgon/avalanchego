// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"net/http"

	"github.com/ava-labs/avalanchego/api"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/client"
	"github.com/ava-labs/avalanchego/utils/profiler"
)

// Admin is the API service for admin API calls
type Admin struct {
	vm       *VM
	profiler profiler.Profiler
}

func NewAdminService(vm *VM, performanceDir string) *Admin { _ = "STUB: not implemented"; return nil }

// StartCPUProfiler starts a cpu profile writing to the specified file
func (p *Admin) StartCPUProfiler(*http.Request, *struct{}, *api.EmptyReply) error {
	_ = "STUB: not implemented"
	return nil
}

// StopCPUProfiler stops the cpu profile
func (p *Admin) StopCPUProfiler(*http.Request, *struct{}, *api.EmptyReply) error {
	_ = "STUB: not implemented"
	return nil
}

// MemoryProfile runs a memory profile writing to the specified file
func (p *Admin) MemoryProfile(*http.Request, *struct{}, *api.EmptyReply) error {
	_ = "STUB: not implemented"
	return nil
}

// LockProfile runs a mutex profile writing to the specified file
func (p *Admin) LockProfile(*http.Request, *struct{}, *api.EmptyReply) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Admin) SetLogLevel(_ *http.Request, args *client.SetLogLevelArgs, _ *api.EmptyReply) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Admin) GetVMConfig(_ *http.Request, _ *struct{}, reply *client.ConfigReply) error {
	_ = "STUB: not implemented"
	return nil
}
