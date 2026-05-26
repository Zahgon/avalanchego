// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package api

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/rpc"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/block"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/genesis"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/tx"
	"github.com/ava-labs/avalanchego/vms/platformvm/warp"
)

const DefaultPollingInterval = 50 * time.Millisecond

func NewClient(uri, chain string) *Client { _ = "STUB: not implemented"; return nil }

type Client struct {
	Req rpc.EndpointRequester
}

func (c *Client) Network(
	ctx context.Context,
	options ...rpc.Option,
) (uint32, ids.ID, ids.ID, error) {
	_ = "STUB: not implemented"
	return 0, *new(ids.ID), *new(ids.ID), nil
}

func (c *Client) Genesis(
	ctx context.Context,
	options ...rpc.Option,
) (*genesis.Genesis, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Nonce(
	ctx context.Context,
	address ids.ShortID,
	options ...rpc.Option,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Client) Balance(
	ctx context.Context,
	address ids.ShortID,
	assetID ids.ID,
	options ...rpc.Option,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Client) Loan(
	ctx context.Context,
	chainID ids.ID,
	options ...rpc.Option,
) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Client) IssueTx(
	ctx context.Context,
	newTx *tx.Tx,
	options ...rpc.Option,
) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (c *Client) LastAccepted(
	ctx context.Context,
	options ...rpc.Option,
) (ids.ID, *block.Stateless, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil, nil
}

func (c *Client) Block(
	ctx context.Context,
	blkID ids.ID,
	options ...rpc.Option,
) (*block.Stateless, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Message(
	ctx context.Context,
	txID ids.ID,
	options ...rpc.Option,
) (*warp.UnsignedMessage, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func AwaitTxAccepted(
	ctx context.Context,
	c *Client,
	address ids.ShortID,
	nonce uint64,
	freq time.Duration,
	options ...rpc.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

// The nonce increasing indicates the acceptance of a transaction
// issued with the specified nonce.
