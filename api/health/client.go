// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package health

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/utils/rpc"
)

type Client struct {
	Requester rpc.EndpointRequester
}

func NewClient(uri string) *Client { _ = "STUB: not implemented"; return nil }

// Readiness returns if the node has finished initialization
func (c *Client) Readiness(ctx context.Context, tags []string, options ...rpc.Option) (*APIReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Health returns a summation of the health of the node
func (c *Client) Health(ctx context.Context, tags []string, options ...rpc.Option) (*APIReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Liveness returns if the node is in need of a restart
func (c *Client) Liveness(ctx context.Context, tags []string, options ...rpc.Option) (*APIReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AwaitReady polls the node every [freq] until the node reports ready.
// Only returns an error if [ctx] returns an error.
func AwaitReady(ctx context.Context, c *Client, freq time.Duration, tags []string, options ...rpc.Option) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AwaitHealthy polls the node every [freq] until the node reports healthy.
// Only returns an error if [ctx] returns an error.
func AwaitHealthy(ctx context.Context, c *Client, freq time.Duration, tags []string, options ...rpc.Option) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AwaitAlive polls the node every [freq] until the node reports liveness.
// Only returns an error if [ctx] returns an error.
func AwaitAlive(ctx context.Context, c *Client, freq time.Duration, tags []string, options ...rpc.Option) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func await(
	ctx context.Context,
	freq time.Duration,
	check func(ctx context.Context, tags []string, options ...rpc.Option) (*APIReply, error),
	tags []string,
	options ...rpc.Option,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
