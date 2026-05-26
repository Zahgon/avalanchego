// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"context"
	"sync"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
)

const errThrottleTime = 100 * time.Millisecond

// Subscription is a function that blocks until either the given context is cancelled, or a message is returned.
// It is used to receive messages from a VM such as Pending transactions, state sync completion, etc.
// The function returns the message received, or an error if the context is cancelled.
type Subscription func(ctx context.Context) (Message, error)

type Notifier interface {
	Notify(context.Context, Message) error
}

// NotificationForwarder is a component that listens for notifications from a Subscription,
// and forwards them to a Notifier.
// When CheckForEvent is called mid-subscription, it retries the subscription.
// After Notify is called, it waits for CheckForEvent to be called before subscribing again.
type NotificationForwarder struct {
	Engine    Notifier
	Subscribe Subscription
	Log       logging.Logger

	lock          sync.Mutex
	executing     sync.WaitGroup
	execCtx       context.Context
	haltExecution context.CancelFunc
	abortContext  context.CancelFunc
}

func NewNotificationForwarder(
	engine Notifier,
	subscribe Subscription,
	log logging.Logger,
) *NotificationForwarder {
	_ = "STUB: not implemented"
	return nil
}

func (nf *NotificationForwarder) start() { _ = "STUB: not implemented"; return }

func (nf *NotificationForwarder) run() { _ = "STUB: not implemented"; return }

func (nf *NotificationForwarder) forwardNotification() { _ = "STUB: not implemented"; return }

// Wait to retry

// Wait to retry

// Wait for the context to be cancelled before proceeding to the next subscription,
// in order to subscribe after a block was accepted or a state sync was completed.

// CheckForEvent cancels any outstanding WaitForEvent calls and schedules a new WaitForEvent call.
func (nf *NotificationForwarder) CheckForEvent() { _ = "STUB: not implemented"; return }

func (nf *NotificationForwarder) cancelContext() { _ = "STUB: not implemented"; return }

func (nf *NotificationForwarder) setAndGetContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Close cancels any outstanding WaitForEvent calls and waits for them to return.
// After Close returns, no future WaitForEvent calls will be made by the notification forwarder.
func (nf *NotificationForwarder) Close() { _ = "STUB: not implemented"; return }
