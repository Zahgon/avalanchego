// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package throttling

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/linked"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/metric"
)

// See inbound_msg_throttler.go

func newInboundMsgByteThrottler(
	log logging.Logger,
	registerer prometheus.Registerer,
	vdrs validators.Manager,
	config MsgByteThrottlerConfig,
) (*inboundMsgByteThrottler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Information about a message waiting to be read.
type msgMetadata struct {
	// Need this many more bytes before Acquire returns
	bytesNeeded uint64
	// The number of bytes that were attempted to be acquired
	msgSize uint64
	// The sender of this incoming message
	nodeID ids.NodeID
	// Closed when the message can be read.
	closeOnAcquireChan chan struct{}
}

// It gives more space to validators with more stake.
// Messages are guaranteed to make progress toward
// acquiring enough bytes to be read.
type inboundMsgByteThrottler struct {
	commonMsgThrottler
	metrics   inboundMsgByteThrottlerMetrics
	nextMsgID uint64
	// Node ID --> Msg ID for a message this node is waiting to acquire
	nodeToWaitingMsgID map[ids.NodeID]uint64
	// Msg ID --> *msgMetadata
	waitingToAcquire *linked.Hashmap[uint64, *msgMetadata]
	// Invariant: The node is only waiting on a single message at a time
	//
	// Invariant: waitingToAcquire.Get(nodeToWaitingMsgIDs[nodeID])
	// is the info about the message [nodeID] that has been blocking
	// on reading.
	//
	// Invariant: len(nodeToWaitingMsgIDs) >= 1
	// implies waitingToAcquire.Len() >= 1, and vice versa.
}

// Returns when we can read a message of size [msgSize] from node [nodeID].
// The returned ReleaseFunc must be called (!) when done with the message
// or when we give up trying to read the message, if applicable.
func (t *inboundMsgByteThrottler) Acquire(ctx context.Context, msgSize uint64, nodeID ids.NodeID) ReleaseFunc {
	_ = "STUB: not implemented"
	return *new(ReleaseFunc)
}

// If there is already a message waiting, log the error and return

// Take as many bytes as we can from the at-large allocation.

// only give as many bytes as needed

// don't exceed per-node limit

// don't give more bytes than are in the allocation

// If we acquired enough bytes, return

// Take as many bytes as we can from [nodeID]'s validator allocation.
// Calculate [nodeID]'s validator allocation size based on its weight

// [vdrBytesAllowed] is the number of bytes this node
// may take from its validator allocation.

// We're already using all the bytes we can from the validator allocation

// Mark that [nodeID] used [vdrBytesUsed] from its validator allocation

// If we acquired enough bytes, return

// We still haven't acquired enough bytes to read the message.
// Wait until more bytes are released.

// [closeOnAcquireChan] is closed when [msgSize] bytes have
// been acquired and the message can be read.

// Must correspond to a previous call of Acquire([msgSize], [nodeID])
func (t *inboundMsgByteThrottler) release(metadata *msgMetadata, nodeID ids.NodeID) {
	_ = "STUB: not implemented"
	return
}

// [vdrBytesToReturn] is the number of bytes from [msgSize]
// that will be given back to [nodeID]'s validator allocation
// or messages from [nodeID] currently waiting to acquire bytes.

// [atLargeBytesToReturn] is the number of bytes from [msgSize]
// that will be given to the at-large allocation or a message
// from any node currently waiting to acquire bytes.

// Mark that [nodeID] has released these bytes.

// Iterates over messages waiting to acquire bytes from oldest
// (waiting the longest) to newest. Try to give bytes to the
// oldest message, then next oldest, etc. until there are no
// waiting messages or we exhaust the bytes.

// From the at-large allocation, take the maximum number of bytes
// without exceeding the per-node limit on taking from at-large pool.

// don't give [msg] too many bytes

// don't exceed per-node limit

// don't give more bytes than are in the allocation

// Mark that we gave [atLargeBytesGiven] to [msg]

// [msg] has acquired enough bytes to be read.
// Unblock the corresponding thread in Acquire

// Mark that this message is no longer waiting to acquire bytes

// Get the message from [nodeID], if any, waiting to acquire

// Give [msg] all the bytes we can

// Unblock the corresponding thread in Acquire

// This should never happen

// We gave back all the bytes we could to waiting messages from [nodeID]
// but some are still left.

type inboundMsgByteThrottlerMetrics struct {
	acquireLatency        metric.Averager
	remainingAtLargeBytes prometheus.Gauge
	remainingVdrBytes     prometheus.Gauge
	awaitingAcquire       prometheus.Gauge
	awaitingRelease       prometheus.Gauge
}

func (m *inboundMsgByteThrottlerMetrics) initialize(reg prometheus.Registerer) error {
	_ = "STUB: not implemented"
	return nil
}
