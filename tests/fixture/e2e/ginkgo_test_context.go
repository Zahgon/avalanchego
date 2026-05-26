// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package e2e

import (
	"context"
	"time"

	"go.uber.org/zap/zapcore"

	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary/common"
)

var _ tests.TestContext = (*GinkgoTestContext)(nil)

type ginkgoWriteCloser struct{}

func (*ginkgoWriteCloser) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	// Add a leading space to better differentiate from other ginkgo output
	return 0, nil
}

func (*ginkgoWriteCloser) Close() error {
	_ = "STUB: not implemented"

	// Define a simple encoder config appropriate for logging with ginkgo
	return nil
}

var ginkgoEncoderConfig = zapcore.EncoderConfig{
	// Time, name and caller are omitted for consistency with previous output.
	TimeKey:        "",
	LevelKey:       "level",
	NameKey:        "",
	CallerKey:      "",
	MessageKey:     "msg",
	StacktraceKey:  "stacktrace",
	EncodeLevel:    logging.ConsoleColorLevelEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
}

// NewGinkgoLogger returns a logger with limited output
func newGinkgoLogger(cfg zapcore.Encoder) logging.Logger {
	_ = "STUB: not implemented"
	return *new(logging.Logger)
}

type GinkgoTestContext struct {
	logger logging.Logger
}

// NewEventHandlerTestContext provides a logger with full output to
// account for the limited context otherwise provided in an event
// handler e.g. SynchronizedBeforeSuite.
func NewEventHandlerTestContext() *GinkgoTestContext { _ = "STUB: not implemented"; return nil }

// NewTestContext provides a logger with limited output to account for
// the context already provided by ginkgo for test logging.
func NewTestContext() *GinkgoTestContext { _ = "STUB: not implemented"; return nil }

func (*GinkgoTestContext) Errorf(format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (*GinkgoTestContext) FailNow() { _ = "STUB: not implemented"; return }

func (tc *GinkgoTestContext) Log() logging.Logger {
	_ = "STUB: not implemented"
	return *new(logging.Logger)
}

func (*GinkgoTestContext) Cleanup() {
	_ = "STUB: not implemented"
	// No-op - ginkgo does this automatically
	return
}

func (*GinkgoTestContext) DeferCleanup(cleanup func()) { _ = "STUB: not implemented"; return }

func (*GinkgoTestContext) By(text string, callback ...func()) { _ = "STUB: not implemented"; return }

// Helper simplifying use of a timed context by canceling the context on ginkgo teardown.
func (tc *GinkgoTestContext) ContextWithTimeout(duration time.Duration) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Helper simplifying use of a timed context configured with the default timeout.
func (tc *GinkgoTestContext) DefaultContext() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Helper simplifying use via an option of a timed context configured with the default timeout.
func (tc *GinkgoTestContext) WithDefaultContext() common.Option {
	_ = "STUB: not implemented"
	return *new(common.Option)
}

func (*GinkgoTestContext) GetDefaultContextParent() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// Re-implementation of testify/require.Eventually that is compatible with ginkgo. testify's
// version calls the condition function with a goroutine and ginkgo assertions don't work
// properly in goroutines.
func (*GinkgoTestContext) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msg string) {
	_ = "STUB: not implemented"
	return
}
