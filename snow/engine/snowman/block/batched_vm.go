// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"context"
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/utils/logging"
)

var ErrRemoteVMNotImplemented = errors.New("vm does not implement RemoteVM interface")

// BatchedChainVM extends the minimal functionalities exposed by ChainVM for VMs
// communicating over network (gRPC in our case). This allows more efficient
// operations since calls over network can be duly batched
type BatchedChainVM interface {
	GetAncestors(
		ctx context.Context,
		blkID ids.ID, // first requested block
		maxBlocksNum int, // max number of blocks to be retrieved
		maxBlocksSize int, // max cumulated byte size of retrieved blocks
		maxBlocksRetrivalTime time.Duration, // max duration of retrival operation
	) ([][]byte, error)

	BatchedParseBlock(ctx context.Context, blks [][]byte) ([]snowman.Block, error)
}

func GetAncestors(
	ctx context.Context,
	log logging.Logger,
	vm Getter, // fetch blocks
	blkID ids.ID, // first requested block
	maxBlocksNum int, // max number of blocks to be retrieved
	maxBlocksSize int, // max cumulated byte size of retrieved blocks
	maxBlocksRetrivalTime time.Duration, // max duration of retrival operation
) ([][]byte, error) {
	_ = "STUB: not implemented"
	// Try and batch GetBlock requests
	return nil, nil
}

// RemoteVM did not work, try local logic

// Special case ErrNotFound as an empty response: this signals
// the client to avoid contacting this node for further ancestors
// as they may have been pruned or unavailable due to state-sync.

// First elt is byte repr. of [blk], then its parent, then grandparent, etc.

// length, in bytes, of all elements of ancestors

// After state sync we may not have the full chain

// Ensure response size isn't too large. Include wrappers.IntLen because
// the size of the message is included with each container, and the size
// is repr. by an int.

// Reached maximum response size

func BatchedParseBlock(
	ctx context.Context,
	vm Parser,
	blks [][]byte,
) ([]snowman.Block, error) {
	_ = "STUB: not implemented"
	// Try and batch ParseBlock requests
	return nil, nil
}

// We couldn't batch the ParseBlock requests, try to parse them one at a
// time.
