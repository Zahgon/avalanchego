// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package x

import (
	"context"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/vms/avm"
	"github.com/ava-labs/avalanchego/wallet/chain/x/builder"
)

func NewContextFromURI(ctx context.Context, uri string) (*builder.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewContextFromClients(
	ctx context.Context,
	infoClient *info.Client,
	xChainClient *avm.Client,
) (*builder.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
