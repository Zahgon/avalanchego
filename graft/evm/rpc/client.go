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
	"sync/atomic"
	"time"
)

var (
	ErrBadResult                 = errors.New("bad result in JSON-RPC response")
	ErrClientQuit                = errors.New("client is closed")
	ErrNoResult                  = errors.New("JSON-RPC response has no result")
	ErrMissingBatchResponse      = errors.New("response batch did not contain a response to this call")
	ErrSubscriptionQueueOverflow = errors.New("subscription queue overflow")
	errClientReconnected         = errors.New("client reconnected")
	errDead                      = errors.New("connection lost")
)

// Timeouts
const (
	defaultDialTimeout = 10 * time.Second // used if context has no deadline
	subscribeTimeout   = 10 * time.Second // overall timeout eth_subscribe, rpc_modules calls
)

const (
	// Subscriptions are removed when the subscriber cannot keep up.
	//
	// This can be worked around by supplying a channel with sufficiently sized buffer,
	// but this can be inconvenient and hard to explain in the docs. Another issue with
	// buffered channels is that the buffer is static even though it might not be needed
	// most of the time.
	//
	// The approach taken here is to maintain a per-subscription linked list buffer
	// shrinks on demand. If the buffer reaches the size below, the subscription is
	// dropped.
	maxClientSubscriptionBuffer = 20000
)

// BatchElem is an element in a batch request.
type BatchElem struct {
	Method string
	Args   []interface{}
	// The result is unmarshaled into this field. Result must be set to a
	// non-nil pointer value of the desired type, otherwise the response will be
	// discarded.
	Result interface{}
	// Error is set if the server returns an error for this request, or if
	// unmarshaling into Result fails. It is not set for I/O errors.
	Error error
}

// Client represents a connection to an RPC server.
type Client struct {
	idgen    func() ID // for subscriptions
	isHTTP   bool      // isHTTP specifies if the client uses an HTTP connection
	services *serviceRegistry

	idCounter atomic.Uint32

	// This function, if non-nil, is called when the connection is lost.
	reconnectFunc reconnectFunc

	// config fields
	batchItemLimit       int
	batchResponseMaxSize int

	// writeConn is used for writing to the connection on the caller's goroutine. It should
	// only be accessed outside of dispatch, with the write lock held. The write lock is
	// taken by sending on reqInit and released by sending on reqSent.
	writeConn jsonWriter

	// for dispatch
	close       chan struct{}
	closing     chan struct{}    // closed when client is quitting
	didClose    chan struct{}    // closed when client quits
	reconnected chan ServerCodec // where write/reconnect sends the new connection
	readOp      chan readOp      // read messages
	readErr     chan error       // errors from read
	reqInit     chan *requestOp  // register response IDs, takes write lock
	reqSent     chan error       // signals write completion, releases write lock
	reqTimeout  chan *requestOp  // removes response IDs when call timeout expires
}

type reconnectFunc func(context.Context) (ServerCodec, error)

type clientContextKey struct{}

type clientConn struct {
	codec   ServerCodec
	handler *handler
}

func (c *Client) newClientConn(conn ServerCodec, apiMaxDuration, refillRate, maxStored time.Duration) *clientConn {
	_ = "STUB: not implemented"
	return nil
}

// When [apiMaxDuration] or [refillRate]/[maxStored] is 0 (as is the case for
// all client invocations of this function), it is ignored.

func (cc *clientConn) close(err error, inflightReq *requestOp) { _ = "STUB: not implemented"; return }

type readOp struct {
	msgs  []*jsonrpcMessage
	batch bool
}

// requestOp represents a pending request. This is used for both batch and non-batch
// requests.
type requestOp struct {
	ids         []json.RawMessage
	err         error
	resp        chan []*jsonrpcMessage // the response goes here
	sub         *ClientSubscription    // set for Subscribe requests.
	hadResponse bool                   // true when the request was responded to
}

func (op *requestOp) wait(ctx context.Context, c *Client) ([]*jsonrpcMessage, error) {
	_ = "STUB: not implemented"
	return nil,

		// Send the timeout to dispatch so it can remove the request IDs.
		nil
}

// Dial creates a new client for the given URL.
//
// The currently supported URL schemes are "http", "https", "ws" and "wss". If rawurl is a
// file name with no URL scheme, a local socket connection is established using UNIX
// domain sockets on supported platforms and named pipes on Windows.
//
// If you want to further configure the transport, use DialOptions instead of this
// function.
//
// For websocket connections, the origin is set to the local host name.
//
// The client reconnects automatically when the connection is lost.
func Dial(rawurl string) (*Client, error) { _ = "STUB: not implemented"; return nil, nil }

// DialContext creates a new RPC client, just like Dial.
//
// The context is used to cancel or time out the initial connection establishment. It does
// not affect subsequent interactions with the client.
func DialContext(ctx context.Context, rawurl string) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DialOptions creates a new RPC client for the given URL. You can supply any of the
// pre-defined client options to configure the underlying transport.
//
// The context is used to cancel or time out the initial connection establishment. It does
// not affect subsequent interactions with the client.
//
// The client reconnects automatically when the connection is lost.
func DialOptions(ctx context.Context, rawurl string, options ...ClientOption) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//case "stdio":
//reconnect = newClientTransportIO(os.Stdin, os.Stdout)
//case "":
//reconnect = newClientTransportIPC(rawurl)

// ClientFromContext retrieves the client from the context, if any. This can be used to perform
// 'reverse calls' in a handler method.
func ClientFromContext(ctx context.Context) (*Client, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func newClient(initctx context.Context, cfg *clientConfig, connect reconnectFunc) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func initClient(conn ServerCodec, services *serviceRegistry, cfg *clientConfig, apiMaxDuration, refillRate, maxStored time.Duration) *Client {
	_ = "STUB: not implemented"
	return nil
}

// Set defaults.

// Launch the main loop.

// RegisterName creates a service for the given receiver type under the given name. When no
// methods on the given receiver match the criteria to be either a RPC method or a
// subscription an error is returned. Otherwise a new service is created and added to the
// service collection this client provides to the server.
func (c *Client) RegisterName(name string, receiver interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) nextID() json.RawMessage { _ = "STUB: not implemented"; return *new(json.RawMessage) }

// SupportedModules calls the rpc_modules method, retrieving the list of
// APIs that are available on the server.
func (c *Client) SupportedModules() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Close closes the client, aborting any in-flight requests.
func (c *Client) Close() { _ = "STUB: not implemented"; return }

// SetHeader adds a custom HTTP header to the client's requests.
// This method only works for clients using HTTP, it doesn't have
// any effect for clients using another transport.
func (c *Client) SetHeader(key, value string) { _ = "STUB: not implemented"; return }

// Call performs a JSON-RPC call with the given arguments and unmarshals into
// result if no error occurred.
//
// The result must be a pointer so that package json can unmarshal into it. You
// can also pass nil, in which case the result is ignored.
func (c *Client) Call(result interface{}, method string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// CallContext performs a JSON-RPC call with the given arguments. If the context is
// canceled before the call has successfully returned, CallContext returns immediately.
//
// The result must be a pointer so that package json can unmarshal into it. You
// can also pass nil, in which case the result is ignored.
func (c *Client) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// dispatch has accepted the request and will close the channel when it quits.

// BatchCall sends all given requests as a single batch and waits for the server
// to return a response for all of them.
//
// In contrast to Call, BatchCall only returns I/O errors. Any error specific to
// a request is reported through the Error field of the corresponding BatchElem.
//
// Note that batch calls may not be executed atomically on the server side.
func (c *Client) BatchCall(b []BatchElem) error { _ = "STUB: not implemented"; return nil }

// BatchCallContext sends all given requests as a single batch and waits for the server
// to return a response for all of them. The wait duration is bounded by the
// context's deadline.
//
// In contrast to CallContext, BatchCallContext only returns errors that have occurred
// while sending the request. Any error specific to a request is reported through the
// Error field of the corresponding BatchElem.
//
// Note that batch calls may not be executed atomically on the server side.
func (c *Client) BatchCallContext(ctx context.Context, b []BatchElem) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait for all responses to come back.

// Ignore null responses. These can happen for batches sent via HTTP.

// Find the element corresponding to this response.

// Assign result and error.

// Check that all expected responses have been received.

// Notify sends a notification, i.e. a method call that doesn't expect a response.
func (c *Client) Notify(ctx context.Context, method string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// EthSubscribe registers a subscription under the "eth" namespace.
func (c *Client) EthSubscribe(ctx context.Context, channel interface{}, args ...interface{}) (*ClientSubscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ShhSubscribe registers a subscription under the "shh" namespace.
// Deprecated: use Subscribe(ctx, "shh", ...).
func (c *Client) ShhSubscribe(ctx context.Context, channel interface{}, args ...interface{}) (*ClientSubscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Subscribe calls the "<namespace>_subscribe" method with the given arguments,
// registering a subscription. Server notifications for the subscription are
// sent to the given channel. The element type of the channel must match the
// expected type of content returned by the subscription.
//
// The context argument cancels the RPC request that sets up the subscription but has no
// effect on the subscription after Subscribe has returned.
//
// Slow subscribers will be dropped eventually. Client buffers up to 20000 notifications
// before considering the subscriber dead. The subscription Err channel will receive
// ErrSubscriptionQueueOverflow. Use a sufficiently large buffer on the channel or ensure
// that the channel usually has at least one reader to prevent this issue.
func (c *Client) Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (*ClientSubscription, error) {
	_ = "STUB: not implemented"
	// Check type of channel first.
	return nil, nil
}

// Send the subscription request.
// The arrival and validity of the response is signaled on sub.quit.

// SupportsSubscriptions reports whether subscriptions are supported by the client
// transport. When this returns false, Subscribe and related methods will return
// ErrNotificationsUnsupported.
func (c *Client) SupportsSubscriptions() bool { _ = "STUB: not implemented"; return false }

func (c *Client) newMessage(method string, paramsIn ...interface{}) (*jsonrpcMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// prevent sending "params":null

// send registers op with the dispatch loop, then sends msg on the connection.
// if sending fails, op is deregistered.
func (c *Client) send(ctx context.Context, op *requestOp, msg interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// This can happen if the client is overloaded or unable to keep up with
// subscription notifications.

func (c *Client) write(ctx context.Context, msg interface{}, retry bool) error {
	_ = "STUB: not implemented"
	return nil

	// The previous write failed. Try to establish a new connection.
}

func (c *Client) reconnect(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// dispatch is the main loop of the client.
// It sends read messages to waiting calls to Call and BatchCall
// and subscription notifications to registered subscriptions.
func (c *Client) dispatch(codec ServerCodec, apiMaxDuration, refillRate, maxStored time.Duration) {
	_ = "STUB: not implemented"
	return
}

// tracks last send operation
// nil while the send lock is held

// Spawn the initial read loop.

// Read path:

// Reconnect:

// Wait for the previous read loop to exit. This is a rare case which
// happens if this loop isn't notified in time after the connection breaks.
// In those cases the caller will notice first and reconnect. Closing the
// handler terminates all waiting requests (closing op.resp) except for
// lastOp, which will be transferred to the new handler.

// Re-register the in-flight request on the new handler
// because that's where it will be sent.

// Send path:

// Stop listening for further requests until the current one has been sent.

// Remove response handlers for the last send. When the read loop
// goes down, it will signal all other current operations.

// Let the next request in.

// drainRead drops read messages until an error occurs.
func (c *Client) drainRead() { _ = "STUB: not implemented"; return }

// read decodes RPC messages from a codec, feeding them into dispatch.
func (c *Client) read(codec ServerCodec) { _ = "STUB: not implemented"; return }
