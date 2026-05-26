// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package galiasreader

import (
	"github.com/ava-labs/avalanchego/ids"

	aliasreaderpb "github.com/ava-labs/avalanchego/proto/pb/aliasreader"
)

var _ ids.AliaserReader = (*Client)(nil)

// Client implements alias lookups that talk over RPC.
type Client struct {
	client aliasreaderpb.AliasReaderClient
}

// NewClient returns an alias lookup instance connected to a remote alias lookup
// instance
func NewClient(client aliasreaderpb.AliasReaderClient) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Lookup(alias string) (ids.ID, error) {
	_ = "STUB: not implemented"
	return *new(ids.ID), nil
}

func (c *Client) PrimaryAlias(id ids.ID) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c *Client) Aliases(id ids.ID) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
