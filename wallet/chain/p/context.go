// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package p

import (
	"context"

	"github.com/ava-labs/avalanchego/api/info"
	"github.com/ava-labs/avalanchego/vms/platformvm"
	"github.com/ava-labs/avalanchego/wallet/chain/p/builder"
)

// gasPriceMultiplier increases the gas price to support multiple transactions
// to be issued.
//
// TODO: Handle this better. Either here or in the mempool.
const gasPriceMultiplier = 2

func NewContextFromURI(ctx context.Context, uri string) (*builder.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewContextFromClients(
	ctx context.Context,
	infoClient *info.Client,
	chainClient *platformvm.Client,
) (*builder.Context, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
