// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/job"
)

var _ job.Job[ids.ID] = (*voter)(nil)

// Voter records chits received from [nodeID] once its dependencies are met.
type voter struct {
	e               *Engine
	nodeID          ids.NodeID
	requestID       uint32
	responseOptions []ids.ID
}

// The resolution results from the dependencies of the voter aren't explicitly
// used. The responseOptions are used to determine which block to apply the vote
// to. The dependencies are only used to optimistically delay the application of
// the vote until the blocks have been issued.
func (v *voter) Execute(ctx context.Context, _ []ids.ID, _ []ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// To prevent any potential deadlocks with undisclosed dependencies,
// votes must be bubbled to the nearest valid block
