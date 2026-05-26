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
// Copyright 2019 The go-ethereum Authors
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
	"reflect"
	"sync"
)

var (
	contextType      = reflect.TypeOf((*context.Context)(nil)).Elem()
	errorType        = reflect.TypeOf((*error)(nil)).Elem()
	subscriptionType = reflect.TypeOf(Subscription{})
	stringType       = reflect.TypeOf("")
)

type serviceRegistry struct {
	mu       sync.Mutex
	services map[string]service
}

// service represents a registered object.
type service struct {
	name          string               // name for service
	callbacks     map[string]*callback // registered handlers
	subscriptions map[string]*callback // available subscriptions/notifications
}

// callback is a method callback which was registered in the server
type callback struct {
	fn          reflect.Value  // the function
	rcvr        reflect.Value  // receiver object of method, set if fn is method
	argTypes    []reflect.Type // input argument types
	hasCtx      bool           // method's first argument is a context (not included in argTypes)
	errPos      int            // err return idx, of -1 when method cannot return error
	isSubscribe bool           // true if this is a subscription callback
}

func (r *serviceRegistry) registerName(name string, rcvr interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// callback returns the callback corresponding to the given RPC method name.
func (r *serviceRegistry) callback(method string) *callback { _ = "STUB: not implemented"; return nil }

// subscription returns a subscription callback in the given service.
func (r *serviceRegistry) subscription(service, name string) *callback {
	_ = "STUB: not implemented"
	return nil
}

// suitableCallbacks iterates over the methods of the given type. It determines if a method
// satisfies the criteria for a RPC callback or a subscription callback and adds it to the
// collection of callbacks. See server documentation for a summary of these criteria.
func suitableCallbacks(receiver reflect.Value) map[string]*callback {
	_ = "STUB: not implemented"
	return nil
}

// method not exported

// function invalid

// newCallback turns fn (a function) into a callback object. It returns nil if the function
// is unsuitable as an RPC callback.
func newCallback(receiver, fn reflect.Value) *callback { _ = "STUB: not implemented"; return nil }

// Determine parameter types. They must all be exported or builtin types.

// Verify return types. The function must return at most one error
// and/or one other non-error value.

// If an error is returned, it must be the last returned value.

// makeArgTypes composes the argTypes list.
func (c *callback) makeArgTypes() { _ = "STUB: not implemented"; return }

// Skip receiver and context.Context parameter (if present).

// Add all remaining parameters.

// call invokes the callback.
func (c *callback) call(ctx context.Context, method string, args []reflect.Value) (res interface{}, errRes error) {
	_ = "STUB: not implemented"
	// Create the argument slice.
	return nil, nil
}

// Catch panic while running the callback.

// Run the callback.

// Method has returned non-nil error value.

// Does t satisfy the error interface?
func isErrorType(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

// Is t Subscription or *Subscription?
func isSubscriptionType(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

// isPubSub tests whether the given method has as as first argument a context.Context and
// returns the pair (Subscription, error).
func isPubSub(methodType reflect.Type) bool {
	_ = "STUB: not implemented"
	// numIn(0) is the receiver type
	return false
}

// formatName converts to first character of name to lowercase.
func formatName(name string) string { _ = "STUB: not implemented"; return "" }
