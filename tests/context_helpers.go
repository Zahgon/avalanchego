// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
)

// A long default timeout used to timeout failed operations but unlikely to induce
// flaking due to unexpected resource contention.
const DefaultTimeout = 2 * time.Minute

// Helper simplifying use of a timed context by canceling the context with the test context.
func ContextWithTimeout(tc TestContext, duration time.Duration) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Helper simplifying use of a timed context configured with the default timeout.
func DefaultContext(tc TestContext) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Helper simplifying use via an option of a timed context configured with the default timeout.
func WithDefaultContext(tc TestContext) common.Option {
	_ = "STUB: not implemented"
	return *new(common.Option)
}
