// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/utils/set"
)

// Op is an opcode
type Op byte

// Types of messages that may be sent between nodes
// Note: If you add a new parseable Op below, you must add it to either
// [UnrequestedOps] or [FailedToResponseOps].
const (
	// Handshake:
	PingOp Op = iota
	PongOp
	HandshakeOp
	GetPeerListOp
	PeerListOp
	// State sync:
	GetStateSummaryFrontierOp
	GetStateSummaryFrontierFailedOp
	StateSummaryFrontierOp
	GetAcceptedStateSummaryOp
	GetAcceptedStateSummaryFailedOp
	AcceptedStateSummaryOp
	// Bootstrapping:
	GetAcceptedFrontierOp
	GetAcceptedFrontierFailedOp
	AcceptedFrontierOp
	GetAcceptedOp
	GetAcceptedFailedOp
	AcceptedOp
	GetAncestorsOp
	GetAncestorsFailedOp
	AncestorsOp
	// Consensus:
	GetOp
	GetFailedOp
	PutOp
	PushQueryOp
	PullQueryOp
	QueryFailedOp
	ChitsOp
	// Application:
	AppRequestOp
	AppErrorOp
	AppResponseOp
	AppGossipOp
	// Internal:
	ConnectedOp
	DisconnectedOp
	NotifyOp
	GossipRequestOp
	// Simplex
	SimplexOp
)

var (
	// UnrequestedOps are operations that are expected to be seen without having
	// requested them. For example, a peer may receive a Get request for a block
	// without having sent a message previously.
	UnrequestedOps = set.Of(
		GetAcceptedFrontierOp,
		GetAcceptedOp,
		GetAncestorsOp,
		GetOp,
		PushQueryOp,
		PullQueryOp,
		AppRequestOp,
		AppGossipOp,
		GetStateSummaryFrontierOp,
		GetAcceptedStateSummaryOp,
		SimplexOp,
	)
	// FailedToResponseOps maps response failure messages to their successful
	// counterparts.
	FailedToResponseOps = map[Op]Op{
		GetStateSummaryFrontierFailedOp: StateSummaryFrontierOp,
		GetAcceptedStateSummaryFailedOp: AcceptedStateSummaryOp,
		GetAcceptedFrontierFailedOp:     AcceptedFrontierOp,
		GetAcceptedFailedOp:             AcceptedOp,
		GetAncestorsFailedOp:            AncestorsOp,
		GetFailedOp:                     PutOp,
		QueryFailedOp:                   ChitsOp,
		AppErrorOp:                      AppResponseOp,
	}

	errUnknownMessageType = errors.New("unknown message type")
)

func (op Op) String() string {
	_ = "STUB: not implemented"

	// Handshake
	return ""
}

// State sync

// Bootstrapping

// Consensus

// Application

// Internal

// Simplex

func Unwrap(m *p2p.Message) (fmt.Stringer, error) {
	_ = "STUB: not implemented"
	return *new(fmt.Stringer), nil
}

// Handshake:

// State sync:

// Bootstrapping:

// Consensus:

// Application:

// Simplex

func ToOp(m *p2p.Message) (Op, error) { _ = "STUB: not implemented"; return *new(Op), nil }
