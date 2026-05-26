// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package handler

import (
	"context"
	"sync"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/snow/networking/tracker"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/buffer"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
)

var _ MessageQueue = (*messageQueue)(nil)

// Message defines individual messages that have been parsed from the network
// and are now pending execution from the chain.
type Message struct {
	// The original message from the peer
	*message.InboundMessage
	// The desired engine type to execute this message. If not specified,
	// the current executing engine type is used.
	EngineType p2p.EngineType
}

type MessageQueue interface {
	// Add a message.
	//
	// If called after [Shutdown], the message will immediately be marked as
	// having been handled.
	Push(context.Context, Message)

	// Remove and return a message and its context.
	//
	// If there are no available messages, this function will block until a
	// message becomes available or the queue is [Shutdown].
	Pop() (context.Context, Message, bool)

	// Returns the number of messages currently on the queue
	Len() int

	// Shutdown and empty the queue.
	Shutdown()
}

// TODO: Use a better data structure for this.
// We can do something better than pushing to the back of a queue. A multi-level
// queue?
type messageQueue struct {
	// Useful for faking time in tests
	clock   mockable.Clock
	metrics messageQueueMetrics

	log      logging.Logger
	subnetID ids.ID
	// Validator set for the chain associated with this
	vdrs validators.Manager
	// Tracks CPU utilization of each node
	cpuTracker tracker.Tracker

	cond   *sync.Cond
	closed bool
	// Node ID --> Messages this node has in [msgs]
	nodeToUnprocessedMsgs map[ids.NodeID]int
	// Unprocessed messages
	msgAndCtxs buffer.Deque[*msgAndContext]
}

func NewMessageQueue(
	log logging.Logger,
	subnetID ids.ID,
	vdrs validators.Manager,
	cpuTracker tracker.Tracker,
	metricsNamespace string,
	reg prometheus.Registerer,
) (MessageQueue, error) {
	_ = "STUB: not implemented"
	return *new(MessageQueue), nil
}

/*=initSize*/

func (m *messageQueue) Push(ctx context.Context, msg Message) { _ = "STUB: not implemented"; return }

// Add the message to the queue

// Update metrics

// Signal a waiting thread

// FIFO, but skip over messages whose senders whose messages have caused us to
// use excessive CPU recently.
func (m *messageQueue) Pop() (context.Context, Message, bool) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(Message), false
}

// note that n > 0

// See if it's OK to process [msg] next
// i should never == n but handle anyway as a fail-safe

// [msg.nodeID] is causing excessive CPU usage.
// Push [msg] to back of [m.msgs] and handle it later.

func (m *messageQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (m *messageQueue) Shutdown() { _ = "STUB: not implemented"; return }

// Remove all the current messages from the queue

// Update metrics

// Mark the queue as closed

// canPop will return true for at least one message in [m.msgs]
func (m *messageQueue) canPop(msg *message.InboundMessage) bool {
	_ = "STUB: not implemented"
	// Always pop connected and disconnected messages.
	return false
}

// If the deadline to handle [msg] has passed, always pop it.
// It will be dropped immediately.

// Every node has some allowed CPU allocation depending on
// the number of nodes with unprocessed messages.

// The sum of validator weights should never overflow, but if they do,
// we treat portionWeight as 0.

// The sum of validator weights should never be 0, but handle that case
// for completeness here to avoid divide by 0.

// Validators are allowed to use more CPU. More weight --> more CPU use allowed.

type msgAndContext struct {
	msg Message
	ctx context.Context
}
