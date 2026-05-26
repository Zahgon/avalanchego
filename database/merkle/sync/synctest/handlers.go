// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package synctest

import (
	"github.com/ava-labs/avalanchego/network/p2p"
)

// NewCounterHandler returns a TestHandler that wraps the given innerHandler.
// The returned handler will call fn when the AppRequest method has been called
// index times. fn must be thread-safe only if called externally.
func NewCounterHandler(innerHandler p2p.Handler, fn func(), index int) *p2p.TestHandler {
	_ = "STUB: not implemented"
	return nil
}
