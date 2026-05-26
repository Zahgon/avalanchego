// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"context"
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/choices"
	"github.com/ava-labs/avalanchego/utils/rpc"
)

var ErrRejected = errors.New("rejected")

type Client struct {
	Requester rpc.EndpointRequester
}

func NewClient(uri, chain string) *Client { _ = "STUB: not implemented"; return nil }

// GetBlock returns the block with the given id.
func (c *Client) GetBlock(ctx context.Context, blkID ids.ID, options ...rpc.Option) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBlockByHeight returns the block at the given height.
func (c *Client) GetBlockByHeight(ctx context.Context, height uint64, options ...rpc.Option) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetHeight returns the height of the last accepted block.
func (c *Client) GetHeight(ctx context.Context, options ...rpc.Option) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// IssueTx issues a transaction to a node and returns the TxID
func (c *Client) IssueTx(ctx context.Context, txBytes []byte, options ...rpc.Option) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

// GetTxStatus returns the status of [txID]
//
// Deprecated: GetTxStatus only returns Accepted or Unknown, GetTx should be
// used instead to determine if the tx was accepted.
func (c *Client) GetTxStatus(ctx context.Context, txID ids.ID, options ...rpc.Option) (choices.Status, error) {
	_ = "STUB: not implemented"
	return *new(choices.Status), nil
}

// GetTx returns the byte representation of txID.
func (c *Client) GetTx(ctx context.Context, txID ids.ID, options ...rpc.Option) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetUTXOs returns the byte representation of the UTXOs controlled by addrs.
func (c *Client) GetUTXOs(
	ctx context.Context,
	addrs []ids.ShortID,
	limit uint32,
	startAddress ids.ShortID,
	startUTXOID ids.ID,
	options ...rpc.Option,
) ([][]byte, ids.ShortID, ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ShortID), *new(ids.ID), nil
}

// GetAtomicUTXOs returns the byte representation of the atomic UTXOs controlled
// by addrs from sourceChain.
func (c *Client) GetAtomicUTXOs(
	ctx context.Context,
	addrs []ids.ShortID,
	sourceChain string,
	limit uint32,
	startAddress ids.ShortID,
	startUTXOID ids.ID,
	options ...rpc.Option,
) ([][]byte, ids.ShortID, ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ShortID), *new(ids.ID), nil
}

// GetAssetDescription returns a description of assetID.
func (c *Client) GetAssetDescription(ctx context.Context, assetID string, options ...rpc.Option) (*GetAssetDescriptionReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBalance returns the balance of assetID held by addr.
//
// If includePartial is set, balance includes partial owned (i.e. in a multisig)
// funds.
//
// Deprecated: GetUTXOs should be used instead.
func (c *Client) GetBalance(
	ctx context.Context,
	addr ids.ShortID,
	assetID string,
	includePartial bool,
	options ...rpc.Option,
) (*GetBalanceReply, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAllBalances returns all asset balances for addr.
//
// Deprecated: GetUTXOs should be used instead.
func (c *Client) GetAllBalances(
	ctx context.Context,
	addr ids.ShortID,
	includePartial bool,
	options ...rpc.Option,
) ([]Balance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetTxFee returns the cost to issue certain transactions.
func (c *Client) GetTxFee(ctx context.Context, options ...rpc.Option) (uint64, uint64, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func (c *Client) AwaitTxAccepted(ctx context.Context, txID ids.ID, freq time.Duration, options ...rpc.Option) error {
	_ = "STUB: not implemented"
	return nil
}
