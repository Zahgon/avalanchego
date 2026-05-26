// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package logging

import (
	"io"

	"go.uber.org/zap"
)

var (
	// Discard is a mock WriterCloser that drops all writes and close requests
	Discard io.WriteCloser = discard{}

	_ Logger = NoLog{}
)

type NoLog struct{}

func (NoLog) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (NoLog) Fatal(string, ...zap.Field) { _ = "STUB: not implemented"; return }

func (NoLog) Error(string, ...zap.Field) { _ = "STUB: not implemented"; return }

func (NoLog) Warn(string, ...zap.Field) { _ = "STUB: not implemented"; return }

func (NoLog) Info(string, ...zap.Field) { _ = "STUB: not implemented"; return }

func (NoLog) Trace(string, ...zap.Field) { _ = "STUB: not implemented"; return }

func (NoLog) Debug(string, ...zap.Field) { _ = "STUB: not implemented"; return }

func (NoLog) Verbo(string, ...zap.Field) { _ = "STUB: not implemented"; return }

func (n NoLog) With(...zap.Field) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (n NoLog) WithOptions(...zap.Option) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (NoLog) SetLevel(Level) { _ = "STUB: not implemented"; return }

func (NoLog) Enabled(Level) bool { _ = "STUB: not implemented"; return false }

func (NoLog) StopOnPanic() { _ = "STUB: not implemented"; return }

func (NoLog) RecoverAndPanic(f func()) { _ = "STUB: not implemented"; return }

func (NoLog) RecoverAndExit(f, exit func()) { _ = "STUB: not implemented"; return }

func (NoLog) Stop() { _ = "STUB: not implemented"; return }

type NoWarn struct{ NoLog }

func (NoWarn) Fatal(string, ...zap.Field) { _ = "STUB: not implemented"; return }

func (NoWarn) Error(string, ...zap.Field) { _ = "STUB: not implemented"; return }

func (NoWarn) Warn(string, ...zap.Field) { _ = "STUB: not implemented"; return }

type discard struct{}

func (discard) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (discard) Close() error { _ = "STUB: not implemented"; return nil }
