// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package client

import (
	"context"

	"golang.org/x/exp/slog"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/plugin/evm/config"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/rpc"
)

// Interface compliance
var _ Client = (*client)(nil)

type CurrentValidator struct {
	ValidationID     ids.ID     `json:"validationID"`
	NodeID           ids.NodeID `json:"nodeID"`
	Weight           uint64     `json:"weight"`
	StartTimestamp   uint64     `json:"startTimestamp"`
	IsActive         bool       `json:"isActive"`
	IsL1Validator    bool       `json:"isL1Validator"`
	IsConnected      bool       `json:"isConnected"`
	UptimePercentage float32    `json:"uptimePercentage"`
	UptimeSeconds    uint64     `json:"uptimeSeconds"`
}

// Client interface for interacting with EVM [chain]
type Client interface {
	StartCPUProfiler(ctx context.Context, options ...rpc.Option) error
	StopCPUProfiler(ctx context.Context, options ...rpc.Option) error
	MemoryProfile(ctx context.Context, options ...rpc.Option) error
	LockProfile(ctx context.Context, options ...rpc.Option) error
	SetLogLevel(ctx context.Context, level slog.Level, options ...rpc.Option) error
	GetVMConfig(ctx context.Context, options ...rpc.Option) (*config.Config, error)
	GetCurrentValidators(ctx context.Context, nodeIDs []ids.NodeID, options ...rpc.Option) ([]CurrentValidator, error)
}

// Client implementation for interacting with EVM [chain]
type client struct {
	adminRequester      rpc.EndpointRequester
	validatorsRequester rpc.EndpointRequester
}

// NewClient returns a Client for interacting with EVM [chain]
func NewClient(uri, chain string) Client { _ = "STUB: not implemented"; return *new(Client) }

// NewClientWithURL returns a Client for interacting with EVM [chain]
func NewClientWithURL(url string) Client { _ = "STUB: not implemented"; return *new(Client) }

func (c *client) StartCPUProfiler(ctx context.Context, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) StopCPUProfiler(ctx context.Context, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) MemoryProfile(ctx context.Context, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *client) LockProfile(ctx context.Context, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}

type SetLogLevelArgs struct {
	Level string `json:"level"`
}

// SetLogLevel dynamically sets the log level for the C Chain
func (c *client) SetLogLevel(ctx context.Context, level slog.Level, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}

type ConfigReply struct {
	Config *config.Config `json:"config"`
}

// GetVMConfig returns the current config of the VM
func (c *client) GetVMConfig(ctx context.Context, options ...rpc.Option) (*config.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type GetCurrentValidatorsRequest struct {
	NodeIDs []ids.NodeID `json:"nodeIDs"`
}

type GetCurrentValidatorsResponse struct {
	Validators []CurrentValidator `json:"validators"`
}

// GetCurrentValidators returns the current validators
func (c *client) GetCurrentValidators(ctx context.Context, nodeIDs []ids.NodeID, options ...rpc.Option) ([]CurrentValidator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
