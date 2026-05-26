// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// ACP181 implements the epoch logic specified here:
// https://github.com/avalanche-foundation/ACPs/blob/main/ACPs/181-p-chain-epoched-views/README.md
package acp181

import (
	"time"

	"github.com/ava-labs/avalanchego/upgrade"
	"github.com/ava-labs/avalanchego/vms/proposervm/block"
)

// NewEpoch returns a child block's epoch based on its parent.
func NewEpoch(
	upgrades upgrade.Config,
	parentPChainHeight uint64,
	parentEpoch block.Epoch,
	parentTimestamp time.Time,
	childTimestamp time.Time,
) block.Epoch {
	_ = "STUB: not implemented"
	return *new(block.Epoch)
}

// If the parent was not assigned an epoch, then the child is the first
// block of the initial epoch.

// If the parent was issued before the end of its epoch, then it did not
// seal the epoch.

// The parent sealed the epoch, so the child is the first block of the new
// epoch.
