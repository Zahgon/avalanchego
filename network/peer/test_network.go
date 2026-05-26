// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/bloom"
	"github.com/ava-labs/avalanchego/utils/ips"
	"github.com/ava-labs/avalanchego/utils/set"
)

var TestNetwork Network = testNetwork{}

type testNetwork struct{}

func (testNetwork) Connected(ids.NodeID) { _ = "STUB: not implemented"; return }

func (testNetwork) AllowConnection(ids.NodeID) bool { _ = "STUB: not implemented"; return false }

func (testNetwork) Track([]*ips.ClaimedIPPort) error { _ = "STUB: not implemented"; return nil }

func (testNetwork) Disconnected(ids.NodeID) { _ = "STUB: not implemented"; return }

func (testNetwork) KnownPeers() ([]byte, []byte) { _ = "STUB: not implemented"; return nil, nil }

func (testNetwork) Peers(
	ids.NodeID,
	set.Set[ids.ID],
	bool,
	*bloom.ReadFilter,
	[]byte,
) []*ips.ClaimedIPPort {
	_ = "STUB: not implemented"
	return nil
}
