// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package c

import (
	"context"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms/avm"
)

const Alias = "C"

type Context struct {
	NetworkID    uint32
	BlockchainID ids.ID
	AVAXAssetID  ids.ID
}

func NewContextFromURI(ctx context.Context, uri string) (*Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewContextFromClients(
	ctx context.Context,
	infoClient *info.Client,
	xChainClient *avm.Client,
) (*Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newSnowContext(c *Context) (*snow.Context, error) { _ = "STUB: not implemented"; return nil, nil }
