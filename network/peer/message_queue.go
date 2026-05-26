// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"context"
	"sync"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/message"
	"github.com/ava-labs/avalanchego/network/throttling"
	"github.com/ava-labs/avalanchego/utils/buffer"
	"github.com/ava-labs/avalanchego/utils/logging"
)

const initialQueueSize = 64

var (
	_ MessageQueue = (*throttledMessageQueue)(nil)
	_ MessageQueue = (*blockingMessageQueue)(nil)
)

type SendFailedCallback interface {
	SendFailed(*message.OutboundMessage)
}

type SendFailedFunc func(*message.OutboundMessage)

func (f SendFailedFunc) SendFailed(msg *message.OutboundMessage) { _ = "STUB: not implemented"; return }

type MessageQueue interface {
	// Push attempts to add the message to the queue. If the context is
	// canceled, then pushing the message will return `false` and the message
	// will not be added to the queue.
	Push(ctx context.Context, msg *message.OutboundMessage) bool

	// Pop blocks until a message is available and then returns the message. If
	// the queue is closed, then `false` is returned.
	Pop() (*message.OutboundMessage, bool)
	// PopNow attempts to return a message without blocking. If a message is not
	// available or the queue is closed, then `false` is returned.
	PopNow() (*message.OutboundMessage, bool)

	// Close empties the queue and prevents further messages from being pushed
	// onto it. After calling close once, future calls to close will do nothing.
	Close()
}

type throttledMessageQueue struct {
	onFailed SendFailedCallback
	// [id] of the peer we're sending messages to
	id                   ids.NodeID
	log                  logging.Logger
	outboundMsgThrottler throttling.OutboundMsgThrottler

	// Signalled when a message is added to the queue and when Close() is
	// called.
	cond *sync.Cond

	// closed flags whether the send queue has been closed.
	// [cond.L] must be held while accessing [closed].
	closed bool

	// queue of the messages
	// [cond.L] must be held while accessing [queue].
	queue buffer.Deque[*message.OutboundMessage]
}

func NewThrottledMessageQueue(
	onFailed SendFailedCallback,
	id ids.NodeID,
	log logging.Logger,
	outboundMsgThrottler throttling.OutboundMsgThrottler,
) MessageQueue {
	_ = "STUB: not implemented"
	return *new(MessageQueue)
}

func (q *throttledMessageQueue) Push(ctx context.Context, msg *message.OutboundMessage) bool {
	_ = "STUB: not implemented"
	return false
}

// Acquire space on the outbound message queue, or drop [msg] if we can't.

// Invariant: must call q.outboundMsgThrottler.Release(msg, q.id) when [msg]
// is popped or, if this queue closes before [msg] is popped, when this
// queue closes.

func (q *throttledMessageQueue) Pop() (*message.OutboundMessage, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// There is a message

// Wait until there is a message

func (q *throttledMessageQueue) PopNow() (*message.OutboundMessage, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// There isn't a message

func (q *throttledMessageQueue) pop() *message.OutboundMessage {
	_ = "STUB: not implemented"
	return nil
}

func (q *throttledMessageQueue) Close() { _ = "STUB: not implemented"; return }

type blockingMessageQueue struct {
	onFailed SendFailedCallback
	log      logging.Logger

	closeOnce   sync.Once
	closingLock sync.RWMutex
	closing     chan struct{}

	// queue of the messages
	queue chan *message.OutboundMessage
}

func NewBlockingMessageQueue(
	onFailed SendFailedCallback,
	log logging.Logger,
	bufferSize int,
) MessageQueue {
	_ = "STUB: not implemented"
	return *new(MessageQueue)
}

func (q *blockingMessageQueue) Push(ctx context.Context, msg *message.OutboundMessage) bool {
	_ = "STUB: not implemented"
	return false
}

func (q *blockingMessageQueue) Pop() (*message.OutboundMessage, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (q *blockingMessageQueue) PopNow() (*message.OutboundMessage, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (q *blockingMessageQueue) Close() { _ = "STUB: not implemented"; return }
