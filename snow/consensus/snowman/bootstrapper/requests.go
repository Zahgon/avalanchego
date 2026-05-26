// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bootstrapper

import (
	"context"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/set"
)

type requests struct {
	maxOutstanding int

	pendingSend set.Set[ids.NodeID]
	outstanding set.Set[ids.NodeID]
}

func (r *requests) GetPeers(context.Context) set.Set[ids.NodeID] {
	_ = "STUB: not implemented"
	return nil
}

func (r *requests) recordResponse(nodeID ids.NodeID) bool { _ = "STUB: not implemented"; return false }

func (r *requests) finished() bool { _ = "STUB: not implemented"; return false }
