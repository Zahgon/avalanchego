// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package info

import (
	"context"
	"net/netip"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/utils/rpc"
	"github.com/ava-labs/avalanchego/vms/platformvm/signer"
)

type Client struct {
	Requester rpc.EndpointRequester
}

func NewClient(uri string) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) GetNodeVersion(ctx context.Context, options ...rpc.Option) (*GetNodeVersionReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) GetNodeID(ctx context.Context, options ...rpc.Option) (ids.NodeID, *signer.ProofOfPossession, error) {
	_ = "STUB: not implemented"
	return *new(ids.NodeID), nil, nil
}

func (c *Client) GetNodeIP(ctx context.Context, options ...rpc.Option) (netip.AddrPort, error) {
	_ = "STUB: not implemented"
	return *new(netip.AddrPort), nil
}

func (c *Client) GetNetworkID(ctx context.Context, options ...rpc.Option) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Client) GetNetworkName(ctx context.Context, options ...rpc.Option) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *Client) GetBlockchainID(ctx context.Context, alias string, options ...rpc.Option) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (c *Client) Peers(ctx context.Context, nodeIDs []ids.NodeID, options ...rpc.Option) ([]Peer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) IsBootstrapped(ctx context.Context, chainID string, options ...rpc.Option) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (c *Client) Upgrades(ctx context.Context, options ...rpc.Option) (*upgrade.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Uptime(ctx context.Context, options ...rpc.Option) (*UptimeResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) GetVMs(ctx context.Context, options ...rpc.Option) (map[ids.ID][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AwaitBootstrapped polls the node every [freq] to check if [chainID] has
// finished bootstrapping. Returns true once [chainID] reports that it has
// finished bootstrapping.
// Only returns an error if [ctx] returns an error.
func AwaitBootstrapped(ctx context.Context, c *Client, chainID string, freq time.Duration, options ...rpc.Option) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}
