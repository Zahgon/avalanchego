// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"context"
	"time"

	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
)

var _ TestContext = (*SimpleTestContext)(nil)

const failNowMessage = "SimpleTestContext.FailNow called"

type ErrorfHandler func(format string, args ...any)

type PanicHandler func(any)

type SimpleTestContext struct {
	defaultContextParent context.Context
	log                  logging.Logger

	cleanupFuncs  []func()
	cleanupCalled bool

	errorfHandler ErrorfHandler
	panicHandler  PanicHandler
}

func NewTestContext(log logging.Logger) *SimpleTestContext { _ = "STUB: not implemented"; return nil }

func NewTestContextWithArgs(
	ctx context.Context,
	log logging.Logger,
	errorfHandler ErrorfHandler,
	panicHandler PanicHandler,
) *SimpleTestContext {
	_ = "STUB: not implemented"
	return nil
}

func (tc *SimpleTestContext) Errorf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (*SimpleTestContext) FailNow() { _ = "STUB: not implemented"; return }

// RecoverAndExit is intended to be deferred by the caller to ensure
// cleanup functions are called before exit or re-panic (in the event
// of an unexpected panic or a panic during cleanup).
func (tc *SimpleTestContext) RecoverAndExit() {
	_ = "STUB: not implemented"
	// Only exit non-zero if a cleanup caused a panic
	return
}

// Retain the panic data to raise after cleanup

// Ensure a non-zero exit due to an assertion failure

// Re-throw an unexpected (non-assertion) panic

// Recover is intended to be deferred in a function executing a test whose
// assertions may result in panics. Such a panic is intended to be recovered to
// allow cleanup functions to be called before execution continues.
func (tc *SimpleTestContext) Recover() {
	_ = "STUB: not implemented"
	/* rethrow */ return
}

// RecoverAndRethrow is intended to be deferred in a function executing a test
// whose assertions may result in panics.  Such a panic is intended to be recovered
// to allow cleanup functions to be called before the panic is rethrown.
func (tc *SimpleTestContext) RecoverAndRethrow() {
	_ = "STUB: not implemented"
	/* rethrow */ return
}

// Recover is intended to be deferred in a function executing a test
// whose assertions may result in panics. Such a panic is intended to
// be recovered to allow cleanup functions to be called. A panic can
// be optionally rethrown by setting `rethrow` to true.
func (tc *SimpleTestContext) recover(rethrow bool) {
	_ = "STUB: not implemented"
	// Recover from test failure
	return
}

// Ensure cleanup functions are called

// cleanup ensures that the registered cleanup functions have been
// called. Cleanup functions will be called at most once. Returns a
// boolean indication of whether a panic results from executing one or
// more cleanup functions i.e. to trigger a non-zero exit.
func (tc *SimpleTestContext) cleanup() bool { _ = "STUB: not implemented"; return false }

// Ensure a failed cleanup doesn't prevent subsequent cleanup functions from running

func (tc *SimpleTestContext) DeferCleanup(cleanup func()) { _ = "STUB: not implemented"; return }

func (tc *SimpleTestContext) By(msg string, callback ...func()) { _ = "STUB: not implemented"; return }

func (tc *SimpleTestContext) Log() logging.Logger {
	_ = "STUB: not implemented"

	// Helper simplifying use of a timed context by canceling the context on ginkgo teardown.
	return *new(logging.Logger)
}

func (tc *SimpleTestContext) ContextWithTimeout(duration time.Duration) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Helper simplifying use of a timed context configured with the default timeout.
func (tc *SimpleTestContext) DefaultContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Helper simplifying use via an option of a timed context configured with the default timeout.
func (tc *SimpleTestContext) WithDefaultContext() common.Option {
	_ = "STUB: not implemented"
	return *new(common.Option)
}

func (tc *SimpleTestContext) GetDefaultContextParent() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (tc *SimpleTestContext) SetDefaultContextParent(parent context.Context) {
	_ = "STUB: not implemented"
	return
}

func (tc *SimpleTestContext) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msg string) {
	_ = "STUB: not implemented"
	return
}
