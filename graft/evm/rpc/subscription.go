// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"sync"
)

var (
	// ErrNotificationsUnsupported is returned by the client when the connection doesn't
	// support notifications. You can use this error value to check for subscription
	// support like this:
	//
	//	sub, err := client.EthSubscribe(ctx, channel, "newHeads", true)
	//	if errors.Is(err, rpc.ErrNotificationsUnsupported) {
	//		// Server does not support subscriptions, fall back to polling.
	//	}
	//
	ErrNotificationsUnsupported = notificationsUnsupportedError{}

	// ErrSubscriptionNotFound is returned when the notification for the given id is not found
	ErrSubscriptionNotFound = errors.New("subscription not found")
)

var globalGen = randomIDGenerator()

// ID defines a pseudo random number that is used to identify RPC subscriptions.
type ID string

// NewID returns a new, random ID.
func NewID() ID {
	_ = "STUB: not implemented"

	// randomIDGenerator returns a function generates a random IDs.
	return *new(ID)
}

func randomIDGenerator() func() ID { _ = "STUB: not implemented"; return nil }

func encodeID(b []byte) ID { _ = "STUB: not implemented"; return *new(ID) }

// ID's are RPC quantities, no leading zero's and 0 is 0x0.

type notifierKey struct{}

// NotifierFromContext returns the Notifier value stored in ctx, if any.
func NotifierFromContext(ctx context.Context) (*Notifier, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Notifier is tied to a RPC connection that supports subscriptions.
// Server callbacks use the notifier to send notifications.
type Notifier struct {
	h         *handler
	namespace string

	mu           sync.Mutex
	sub          *Subscription
	buffer       []any
	callReturned bool
	activated    bool
}

// CreateSubscription returns a new subscription that is coupled to the
// RPC connection. By default subscriptions are inactive and notifications
// are dropped until the subscription is marked as active. This is done
// by the RPC server after the subscription ID is send to the client.
func (n *Notifier) CreateSubscription() *Subscription { _ = "STUB: not implemented"; return nil }

// Notify sends a notification to the client with the given data as payload.
// If an error occurs the RPC connection is closed and the error is returned.
func (n *Notifier) Notify(id ID, data any) error { _ = "STUB: not implemented"; return nil }

// Closed returns a channel that is closed when the RPC connection is closed.
// Deprecated: use subscription error channel
func (n *Notifier) Closed() <-chan interface{} { _ = "STUB: not implemented"; return nil }

// takeSubscription returns the subscription (if one has been created). No subscription can
// be created after this call.
func (n *Notifier) takeSubscription() *Subscription { _ = "STUB: not implemented"; return nil }

// activate is called after the subscription ID was sent to client. Notifications are
// buffered before activation. This prevents notifications being sent to the client before
// the subscription ID is sent to the client.
func (n *Notifier) activate() error { _ = "STUB: not implemented"; return nil }

func (n *Notifier) send(sub *Subscription, data any) error { _ = "STUB: not implemented"; return nil }

// A Subscription is created by a notifier and tied to that notifier. The client can use
// this subscription to wait for an unsubscribe request for the client, see Err().
type Subscription struct {
	ID        ID
	namespace string
	err       chan error // closed on unsubscribe
}

// Err returns a channel that is closed when the client send an unsubscribe request.
func (s *Subscription) Err() <-chan error {
	_ = "STUB: not implemented"

	// MarshalJSON marshals a subscription as its ID.
	return nil
}

func (s *Subscription) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// ClientSubscription is a subscription established through the Client's Subscribe or
		// EthSubscribe methods.
		nil
}

type ClientSubscription struct {
	client    *Client
	etype     reflect.Type
	channel   reflect.Value
	namespace string
	subid     string

	// The in channel receives notification values from client dispatcher.
	in chan json.RawMessage

	// The error channel receives the error from the forwarding loop.
	// It is closed by Unsubscribe.
	err     chan error
	errOnce sync.Once

	// Closing of the subscription is requested by sending on 'quit'. This is handled by
	// the forwarding loop, which closes 'forwardDone' when it has stopped sending to
	// sub.channel. Finally, 'unsubDone' is closed after unsubscribing on the server side.
	quit        chan error
	forwardDone chan struct{}
	unsubDone   chan struct{}
}

// This is the sentinel value sent on sub.quit when Unsubscribe is called.
var errUnsubscribed = errors.New("unsubscribed")

func newClientSubscription(c *Client, namespace string, channel reflect.Value) *ClientSubscription {
	_ = "STUB: not implemented"
	return nil
}

// Err returns the subscription error channel. The intended use of Err is to schedule
// resubscription when the client connection is closed unexpectedly.
//
// The error channel receives a value when the subscription has ended due to an error. The
// received error is nil if Close has been called on the underlying client and no other
// error has occurred.
//
// The error channel is closed when Unsubscribe is called on the subscription.
func (sub *ClientSubscription) Err() <-chan error {
	_ = "STUB: not implemented"

	// Unsubscribe unsubscribes the notification and closes the error channel.
	// It can safely be called more than once.
	return nil
}

func (sub *ClientSubscription) Unsubscribe() { _ = "STUB: not implemented"; return }

// deliver is called by the client's message dispatcher to send a notification value.
func (sub *ClientSubscription) deliver(result json.RawMessage) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

// close is called by the client's message dispatcher when the connection is closed.
func (sub *ClientSubscription) close(err error) { _ = "STUB: not implemented"; return }

// run is the forwarding loop of the subscription. It runs in its own goroutine and
// is launched by the client's handler after the subscription has been created.
func (sub *ClientSubscription) run() { _ = "STUB: not implemented"; return }

// The client's dispatch loop won't be able to execute the unsubscribe call if it is
// blocked in sub.deliver() or sub.close(). Closing forwardDone unblocks them.

// Call the unsubscribe method on the server.

// Send the error.

// ErrClientQuit gets here when Client.Close is called. This is reported as a
// nil error because it's not an error, but we can't close sub.err here.

// forward is the forwarding loop. It takes in RPC notifications and sends them
// on the subscription channel.
func (sub *ClientSubscription) forward() (unsubscribeServer bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Idle, omit send case.

// Non-empty buffer, send the first queued item.

// <-sub.quit

// Exiting because Unsubscribe was called, unsubscribe on server.

// <-sub.in

// sub.channel<-
// Don't hold onto the value.

func (sub *ClientSubscription) unmarshal(result json.RawMessage) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (sub *ClientSubscription) requestUnsubscribe() error { _ = "STUB: not implemented"; return nil }
