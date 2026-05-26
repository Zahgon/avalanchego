// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package client

import (
	"context"
	"errors"
	"time"

	"golang.org/x/exp/slog"

	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/atomic"
	"github.com/ava-labs/avalanchego/graft/coreth/plugin/evm/config"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/json"
	"github.com/ava-labs/avalanchego/utils/rpc"
)

var errInvalidAddr = errors.New("invalid hex address")

// Client for interacting with EVM [chain]
type Client struct {
	requester      rpc.EndpointRequester
	adminRequester rpc.EndpointRequester
}

// NewClient returns a Client for interacting with EVM [chain]
func NewClient(uri, chain string) *Client { _ = "STUB: not implemented"; return nil }

// NewCChainClient returns a Client for interacting with the C Chain
func NewCChainClient(uri string) *Client { _ = "STUB: not implemented"; return nil }

// IssueTx issues a transaction to a node and returns the TxID
func (c *Client) IssueTx(ctx context.Context, txBytes []byte, options ...rpc.Option) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

// GetAtomicTxStatusReply defines the GetAtomicTxStatus replies returned from the API
type GetAtomicTxStatusReply struct {
	Status      atomic.Status `json:"status"`
	BlockHeight *json.Uint64  `json:"blockHeight,omitempty"`
}

// GetAtomicTxStatus returns the status of [txID]
func (c *Client) GetAtomicTxStatus(ctx context.Context, txID ids.ID, options ...rpc.Option) (atomic.Status, error) {
	_ = "STUB: not implemented"
	return *new(atomic.Status), nil
}

// GetAtomicTx returns the byte representation of [txID]
func (c *Client) GetAtomicTx(ctx context.Context, txID ids.ID, options ...rpc.Option) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAtomicUTXOs returns the byte representation of the atomic UTXOs controlled by [addresses]
// from [sourceChain]
func (c *Client) GetAtomicUTXOs(ctx context.Context, addrs []ids.ShortID, sourceChain string, limit uint32, startAddress ids.ShortID, startUTXOID ids.ID, options ...rpc.Option) ([][]byte, ids.ShortID, ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ShortID), *new(ids.ID), nil
}

func (c *Client) StartCPUProfiler(ctx context.Context, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) StopCPUProfiler(ctx context.Context, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) MemoryProfile(ctx context.Context, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) LockProfile(ctx context.Context, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}

type SetLogLevelArgs struct {
	Level string `json:"level"`
}

// SetLogLevel dynamically sets the log level for the C Chain
func (c *Client) SetLogLevel(ctx context.Context, level slog.Level, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}

type ConfigReply struct {
	Config *config.Config `json:"config"`
}

// GetVMConfig returns the current config of the VM
func (c *Client) GetVMConfig(ctx context.Context, options ...rpc.Option) (*config.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AwaitTxAccepted polls GetAtomicTxStatus every freq until txID is accepted
// or ctx is cancelled.
func (c *Client) AwaitTxAccepted(ctx context.Context, txID ids.ID, freq time.Duration, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}
