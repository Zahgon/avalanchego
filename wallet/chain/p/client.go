// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package p

import (
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
	"github.com/ava-labs/avalanchego/wallet/chain/p/wallet"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
)

var _ wallet.Client = (*Client)(nil)

func NewClient(
	c *platformvm.Client,
	b wallet.Backend,
) *Client {
	_ = "STUB: not implemented"
	return nil
}

type Client struct {
	client  *platformvm.Client
	backend wallet.Backend
}

func (c *Client) IssueTx(
	tx *txs.Tx,
	options ...common.Option,
) error {
	_ = "STUB: not implemented"
	return nil
}
