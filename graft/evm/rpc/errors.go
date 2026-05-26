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

// HTTPError is returned by client operations when the HTTP status code of the
// response is not a 2xx status.
type HTTPError struct {
	StatusCode int
	Status     string
	Body       []byte
}

func (err HTTPError) Error() string { _ = "STUB: not implemented"; return "" }

// Error wraps RPC errors, which contain an error code in addition to the message.
type Error interface {
	Error() string  // returns the message
	ErrorCode() int // returns the code
}

// A DataError contains some data in addition to the error message.
type DataError interface {
	Error() string          // returns the message
	ErrorData() interface{} // returns the error data
}

// Error types defined below are the built-in JSON-RPC errors.

var (
	_ Error = new(methodNotFoundError)
	_ Error = new(subscriptionNotFoundError)
	_ Error = new(parseError)
	_ Error = new(invalidRequestError)
	_ Error = new(invalidMessageError)
	_ Error = new(invalidParamsError)
	_ Error = new(internalServerError)
)

const (
	errcodeDefault          = -32000
	errcodeTimeout          = -32002
	errcodeResponseTooLarge = -32003
	errcodePanic            = -32603
	errcodeMarshalError     = -32603

	legacyErrcodeNotificationsUnsupported = -32001
)

const (
	errMsgTimeout          = "request timed out"
	errMsgResponseTooLarge = "response too large"
	errMsgBatchTooLarge    = "batch too large"
)

type methodNotFoundError struct{ method string }

func (e *methodNotFoundError) ErrorCode() int { _ = "STUB: not implemented"; return 0 }

func (e *methodNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

type notificationsUnsupportedError struct{}

func (e notificationsUnsupportedError) Error() string { _ = "STUB: not implemented"; return "" }

func (e notificationsUnsupportedError) ErrorCode() int {
	_ = "STUB: not implemented"

	// Is checks for equivalence to another error. Here we define that all errors with code
	// -32601 (method not found) are equivalent to notificationsUnsupportedError. This is
	// done to enable the following pattern:
	//
	//	sub, err := client.Subscribe(...)
	//	if errors.Is(err, rpc.ErrNotificationsUnsupported) {
	//		// server doesn't support subscriptions
	//	}
	return 0
}

func (e notificationsUnsupportedError) Is(other error) bool {
	_ = "STUB: not implemented"
	return false
}

type subscriptionNotFoundError struct{ namespace, subscription string }

func (e *subscriptionNotFoundError) ErrorCode() int { _ = "STUB: not implemented"; return 0 }

func (e *subscriptionNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

// Invalid JSON was received by the server.
type parseError struct{ message string }

func (e *parseError) ErrorCode() int { _ = "STUB: not implemented"; return 0 }

func (e *parseError) Error() string {
	_ = "STUB: not implemented"

	// received message isn't a valid request
	return ""
}

type invalidRequestError struct{ message string }

func (e *invalidRequestError) ErrorCode() int { _ = "STUB: not implemented"; return 0 }

func (e *invalidRequestError) Error() string {
	_ = "STUB: not implemented"

	// received message is invalid
	return ""
}

type invalidMessageError struct{ message string }

func (e *invalidMessageError) ErrorCode() int { _ = "STUB: not implemented"; return 0 }

func (e *invalidMessageError) Error() string {
	_ = "STUB: not implemented"

	// unable to decode supplied params, or an invalid number of parameters
	return ""
}

type invalidParamsError struct{ message string }

func (e *invalidParamsError) ErrorCode() int { _ = "STUB: not implemented"; return 0 }

func (e *invalidParamsError) Error() string {
	_ = "STUB: not implemented"

	// internalServerError is used for server errors during request processing.
	return ""
}

type internalServerError struct {
	code    int
	message string
}

func (e *internalServerError) ErrorCode() int { _ = "STUB: not implemented"; return 0 }

func (e *internalServerError) Error() string { _ = "STUB: not implemented"; return "" }
