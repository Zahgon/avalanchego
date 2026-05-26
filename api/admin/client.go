// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package admin

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/rpc"
)

type Client struct {
	Requester rpc.EndpointRequester
}

func NewClient(uri string) *Client { _ = "STUB: not implemented"; return nil }

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

func (c *Client) Alias(ctx context.Context, endpoint, alias string, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) AliasChain(ctx context.Context, chain, alias string, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) GetChainAliases(ctx context.Context, chain string, options ...rpc.Option) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Stacktrace(ctx context.Context, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) LoadVMs(ctx context.Context, options ...rpc.Option) (map[ids.ID][]string, map[ids.ID]string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (c *Client) SetLoggerLevel(
	ctx context.Context,
	loggerName,
	logLevel,
	displayLevel string,
	options ...rpc.Option,
) (map[string]LogAndDisplayLevels, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) GetLoggerLevel(
	ctx context.Context,
	loggerName string,
	options ...rpc.Option,
) (map[string]LogAndDisplayLevels, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) GetConfig(ctx context.Context, options ...rpc.Option) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) DBGet(ctx context.Context, key []byte, options ...rpc.Option) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
