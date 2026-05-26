// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package execute

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/snow"

	smblock "github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	xsblock "github.com/ava-labs/avalanchego/vms/example/xsvm/block"
)

var errNoTxs = errors.New("no transactions")

func Block(
	ctx context.Context,
	chainContext *snow.Context,
	db database.KeyValueReaderWriterDeleter,
	skipVerify bool,
	blockContext *smblock.Context,
	blk *xsblock.Stateless,
) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: populate fees
