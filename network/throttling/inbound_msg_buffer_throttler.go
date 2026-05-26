// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package throttling

import (
	"context"
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/metric"
)

// See inbound_msg_throttler.go

func newInboundMsgBufferThrottler(
	registerer prometheus.Registerer,
	maxProcessingMsgsPerNode uint64,
) (*inboundMsgBufferThrottler, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Rate-limits inbound messages based on the number of
// messages from a given node that we're currently processing.
type inboundMsgBufferThrottler struct {
	lock    sync.Mutex
	metrics inboundMsgBufferThrottlerMetrics
	// Max number of messages currently processing from a
	// given node. We will stop reading messages from a
	// node until we're processing less than this many
	// messages from the node.
	// In this case, a message is "processing" if the corresponding
	// call to Acquire() has returned or is about to return,
	// but the corresponding call to Release() has not happened.
	// TODO: Different values for validators / non-validators?
	maxProcessingMsgsPerNode uint64
	// Node ID --> Number of messages from this node we're currently processing.
	// Must only be accessed when [lock] is held.
	nodeToNumProcessingMsgs map[ids.NodeID]uint64
	// Node ID --> Channel, when closed
	// causes a goroutine waiting in Acquire to return.
	// Must only be accessed when [lock] is held.
	awaitingAcquire map[ids.NodeID]chan struct{}
}

// Acquire returns when we've acquired space on the inbound message
// buffer so that we can read a message from [nodeID].
// The returned release function must be called (!) when done processing the message
// (or when we give up trying to read the message.)
//
// invariant: There should be a maximum of 1 blocking call to Acquire for a
// given nodeID. Callers must enforce this invariant.
func (t *inboundMsgBufferThrottler) Acquire(ctx context.Context, nodeID ids.NodeID) ReleaseFunc {
	_ = "STUB: not implemented"
	return *new(ReleaseFunc)
}

// We're currently processing the maximum number of
// messages from [nodeID]. Wait until we've finished
// processing some messages from [nodeID].
// [closeOnAcquireChan] will be closed inside Release()
// when we've acquired space on the inbound message buffer
// for this message.

// release marks that we've finished processing a message from [nodeID]
// and can release the space it took on the inbound message buffer.
func (t *inboundMsgBufferThrottler) release(nodeID ids.NodeID) { _ = "STUB: not implemented"; return }

// If we're waiting to acquire space on the inbound message
// buffer for messages from [nodeID], allow it to proceed
// (i.e. for its call to Acquire to return.)

type inboundMsgBufferThrottlerMetrics struct {
	acquireLatency  metric.Averager
	awaitingAcquire prometheus.Gauge
}

func (m *inboundMsgBufferThrottlerMetrics) initialize(reg prometheus.Registerer) error {
	_ = "STUB: not implemented"
	return nil
}
