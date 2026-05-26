// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package throttling

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"
)

var (
	_ OutboundMsgThrottler = (*outboundMsgThrottler)(nil)
	_ OutboundMsgThrottler = (*noOutboundMsgThrottler)(nil)
)

// Rate-limits outgoing messages
type OutboundMsgThrottler interface {
	// Returns true if we can queue the message [msg] to be sent to node [nodeID].
	// Returns false if the message should be dropped (not sent to [nodeID]).
	// If this method returns true, Release([msg], [nodeID]) must be called (!) when
	// the message is sent (or when we give up trying to send the message, if applicable.)
	// If this method returns false, do not make a corresponding call to Release.
	Acquire(msg *message.OutboundMessage, nodeID ids.NodeID) bool

	// Mark that a message [msg] has been sent to [nodeID] or we have given up
	// sending the message. Must correspond to a previous call to
	// Acquire([msg], [nodeID]) that returned true.
	Release(msg *message.OutboundMessage, nodeID ids.NodeID)
}

type outboundMsgThrottler struct {
	commonMsgThrottler
	metrics outboundMsgThrottlerMetrics
}

func NewSybilOutboundMsgThrottler(
	log logging.Logger,
	registerer prometheus.Registerer,
	vdrs validators.Manager,
	config MsgByteThrottlerConfig,
) (OutboundMsgThrottler, error) {
	_ = "STUB: not implemented"
	return *new(OutboundMsgThrottler), nil
}

func (t *outboundMsgThrottler) Acquire(msg *message.OutboundMessage, nodeID ids.NodeID) bool {
	_ = "STUB: not implemented"
	// no need to acquire for this message
	return false
}

// Take as many bytes as we can from the at-large allocation.

// only give as many bytes as needed

// don't exceed per-node limit

// don't give more bytes than are in the allocation

// Take as many bytes as we can from [nodeID]'s validator allocation.
// Calculate [nodeID]'s validator allocation size based on its weight

// [vdrBytesAllowed] is the number of bytes this node
// may take from its validator allocation.

// We're already using all the bytes we can from the validator allocation

// Can't acquire enough bytes to queue this message to be sent

// Can acquire enough bytes to queue this message to be sent.
// Update the state.

// Mark that [nodeID] used [vdrBytesUsed] from its validator allocation

func (t *outboundMsgThrottler) Release(msg *message.OutboundMessage, nodeID ids.NodeID) {
	_ = "STUB: not implemented"
	// no need to release for this message
	return
}

// [vdrBytesToReturn] is the number of bytes from [msgSize]
// that will be given back to [nodeID]'s validator allocation.

// [atLargeBytesToReturn] is the number of bytes from [msgSize]
// that will be given to the at-large allocation.

// Mark that [nodeID] has released these bytes.

type outboundMsgThrottlerMetrics struct {
	acquireSuccesses      prometheus.Counter
	acquireFailures       prometheus.Counter
	remainingAtLargeBytes prometheus.Gauge
	remainingVdrBytes     prometheus.Gauge
	awaitingRelease       prometheus.Gauge
}

func (m *outboundMsgThrottlerMetrics) initialize(registerer prometheus.Registerer) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNoOutboundThrottler() OutboundMsgThrottler {
	_ = "STUB: not implemented"
	return *new(OutboundMsgThrottler)
}

// [Acquire] always returns true. [Release] does nothing.
type noOutboundMsgThrottler struct{}

func (*noOutboundMsgThrottler) Acquire(*message.OutboundMessage, ids.NodeID) bool {
	_ = "STUB: not implemented"
	return false
}

func (*noOutboundMsgThrottler) Release(*message.OutboundMessage, ids.NodeID) {
	_ = "STUB: not implemented"
	return
}
