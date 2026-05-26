// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tracedvm

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
)

func (vm *blockVM) SetPreferenceWithContext(ctx context.Context, blkID ids.ID, blockCtx *block.Context) error {
	_ = "STUB: not implemented"
	return nil
}
