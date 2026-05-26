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
// Copyright 2015 The go-ethereum Authors
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
	"io"
	"reflect"
	"sync"
	"time"
)

const (
	vsn                      = "2.0"
	serviceMethodSeparator   = "_"
	subscribeMethodSuffix    = "_subscribe"
	unsubscribeMethodSuffix  = "_unsubscribe"
	notificationMethodSuffix = "_subscription"

	defaultWriteTimeout = 10 * time.Second // used if context has no deadline
)

var null = json.RawMessage("null")

type subscriptionResult struct {
	ID     string          `json:"subscription"`
	Result json.RawMessage `json:"result,omitempty"`
}

type subscriptionResultEnc struct {
	ID     string `json:"subscription"`
	Result any    `json:"result"`
}

type jsonrpcSubscriptionNotification struct {
	Version string                `json:"jsonrpc"`
	Method  string                `json:"method"`
	Params  subscriptionResultEnc `json:"params"`
}

// A value of this type can a JSON-RPC request, notification, successful response or
// error response. Which one it is depends on the fields.
type jsonrpcMessage struct {
	Version string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Error   *jsonError      `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

func (msg *jsonrpcMessage) isNotification() bool { _ = "STUB: not implemented"; return false }

func (msg *jsonrpcMessage) isCall() bool { _ = "STUB: not implemented"; return false }

func (msg *jsonrpcMessage) isResponse() bool { _ = "STUB: not implemented"; return false }

func (msg *jsonrpcMessage) hasValidID() bool { _ = "STUB: not implemented"; return false }

func (msg *jsonrpcMessage) hasValidVersion() bool { _ = "STUB: not implemented"; return false }

func (msg *jsonrpcMessage) isSubscribe() bool { _ = "STUB: not implemented"; return false }

func (msg *jsonrpcMessage) isUnsubscribe() bool { _ = "STUB: not implemented"; return false }

func (msg *jsonrpcMessage) namespace() string { _ = "STUB: not implemented"; return "" }

func (msg *jsonrpcMessage) String() string { _ = "STUB: not implemented"; return "" }

func (msg *jsonrpcMessage) errorResponse(err error) *jsonrpcMessage {
	_ = "STUB: not implemented"
	return nil
}

func (msg *jsonrpcMessage) response(result interface{}) *jsonrpcMessage {
	_ = "STUB: not implemented"
	return nil
}

func errorMessage(err error) *jsonrpcMessage { _ = "STUB: not implemented"; return nil }

type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (err *jsonError) Error() string { _ = "STUB: not implemented"; return "" }

func (err *jsonError) ErrorCode() int { _ = "STUB: not implemented"; return 0 }

func (err *jsonError) ErrorData() interface{} {
	_ = "STUB: not implemented"

	// Conn is a subset of the methods of net.Conn which are sufficient for ServerCodec.
	return nil
}

type Conn interface {
	io.ReadWriteCloser
	SetWriteDeadline(time.Time) error
}

type deadlineCloser interface {
	io.Closer
	SetWriteDeadline(time.Time) error
}

// ConnRemoteAddr wraps the RemoteAddr operation, which returns a description
// of the peer address of a connection. If a Conn also implements ConnRemoteAddr, this
// description is used in log messages.
type ConnRemoteAddr interface {
	RemoteAddr() string
}

// jsonCodec reads and writes JSON-RPC messages to the underlying connection. It also has
// support for parsing arguments and serializing (result) objects.
type jsonCodec struct {
	remote  string
	closer  sync.Once        // close closed channel once
	closeCh chan interface{} // closed on Close
	decode  decodeFunc       // decoder to allow multiple transports
	encMu   sync.Mutex       // guards the encoder
	encode  encodeFunc       // encoder to allow multiple transports
	conn    deadlineCloser
}

type encodeFunc = func(v interface{}, isErrorResponse bool) error

type decodeFunc = func(v interface{}) error

// NewFuncCodec creates a codec which uses the given functions to read and write. If conn
// implements ConnRemoteAddr, log messages will use it to include the remote address of
// the connection.
func NewFuncCodec(conn deadlineCloser, encode encodeFunc, decode decodeFunc) ServerCodec {
	_ = "STUB: not implemented"
	return *new(ServerCodec)
}

// NewCodec creates a codec on the given connection. If conn implements ConnRemoteAddr, log
// messages will use it to include the remote address of the connection.
func NewCodec(conn Conn) ServerCodec { _ = "STUB: not implemented"; return *new(ServerCodec) }

func (c *jsonCodec) peerInfo() PeerInfo {
	_ = "STUB: not implemented"
	// This returns "ipc" because all other built-in transports have a separate codec type.
	return *new(PeerInfo)
}

func (c *jsonCodec) remoteAddr() string { _ = "STUB: not implemented"; return "" }

func (c *jsonCodec) readBatch() (messages []*jsonrpcMessage, batch bool, err error) {
	_ = "STUB: not implemented"
	// Decode the next JSON object in the input stream.
	// This verifies basic syntax, etc.
	return nil, false, nil
}

// Message is JSON 'null'. Replace with zero value so it
// will be treated like any other invalid message.

func (c *jsonCodec) writeJSON(ctx context.Context, val interface{}, isErrorResponse bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *jsonCodec) writeJSONSkipDeadline(ctx context.Context, v interface{}, isErrorResponse bool, skip bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *jsonCodec) close() { _ = "STUB: not implemented"; return }

// Closed returns a channel which will be closed when Close is called
func (c *jsonCodec) closed() <-chan interface{} {
	_ = "STUB: not implemented"

	// parseMessage parses raw bytes as a (batch of) JSON-RPC message(s). There are no error
	// checks in this function because the raw message has already been syntax-checked when it
	// is called. Any non-JSON-RPC messages in the input return the zero value of
	// jsonrpcMessage.
	return nil
}

func parseMessage(raw json.RawMessage) ([]*jsonrpcMessage, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// skip '['

// isBatch returns true when the first non-whitespace characters is '['
func isBatch(raw json.RawMessage) bool { _ = "STUB: not implemented"; return false }

// skip insignificant whitespace (http://www.ietf.org/rfc/rfc4627.txt)

// parsePositionalArguments tries to parse the given args to an array of values with the
// given types. It returns the parsed values or an error when the args could not be
// parsed. Missing optional arguments are returned as reflect.Zero values.
func parsePositionalArguments(rawArgs json.RawMessage, types []reflect.Type) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// "params" is optional and may be empty. Also allow "params":null even though it's
// not in the spec because our own client used to send it.

// Read argument array.

// Set any missing args to nil.

func parseArgumentArray(dec *json.Decoder, types []reflect.Type) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read end of args array.

// parseSubscriptionName extracts the subscription name from an encoded argument array.
func parseSubscriptionName(rawArgs json.RawMessage) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
