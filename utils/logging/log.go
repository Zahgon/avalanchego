// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package logging

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Logger = (*log)(nil)

type log struct {
	wrappedCores   []WrappedCore
	internalLogger *zap.Logger
}

type WrappedCore struct {
	Core           zapcore.Core
	Writer         io.WriteCloser
	WriterDisabled bool
	AtomicLevel    zap.AtomicLevel
}

func NewWrappedCore(level Level, rw io.WriteCloser, encoder zapcore.Encoder) WrappedCore {
	_ = "STUB: not implemented"
	return *new(WrappedCore)
}

func newZapLogger(prefix string, wrappedCores ...WrappedCore) *zap.Logger {
	_ = "STUB: not implemented"
	return nil
}

// New returns a new logger set up according to [config]
func NewLogger(prefix string, wrappedCores ...WrappedCore) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

// TODO: return errors here
func (l *log) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// TODO: return errors here
func (l *log) Stop() { _ = "STUB: not implemented"; return }

// Enabled returns true if the given level is at or above this level.
func (l *log) Enabled(lvl Level) bool { _ = "STUB: not implemented"; return false }

// Should only be called from [Level] functions.
func (l *log) log(level Level, msg string, fields ...zap.Field) { _ = "STUB: not implemented"; return }

func (l *log) Fatal(msg string, fields ...zap.Field) { _ = "STUB: not implemented"; return }

func (l *log) Error(msg string, fields ...zap.Field) { _ = "STUB: not implemented"; return }

func (l *log) Warn(msg string, fields ...zap.Field) { _ = "STUB: not implemented"; return }

func (l *log) Info(msg string, fields ...zap.Field) { _ = "STUB: not implemented"; return }

func (l *log) Trace(msg string, fields ...zap.Field) { _ = "STUB: not implemented"; return }

func (l *log) Debug(msg string, fields ...zap.Field) { _ = "STUB: not implemented"; return }

func (l *log) Verbo(msg string, fields ...zap.Field) { _ = "STUB: not implemented"; return }

func (l *log) With(fields ...zap.Field) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (l *log) WithOptions(opts ...zap.Option) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

func (l *log) SetLevel(level Level) { _ = "STUB: not implemented"; return }

func (l *log) StopOnPanic() { _ = "STUB: not implemented"; return }

func (l *log) RecoverAndPanic(f func()) { _ = "STUB: not implemented"; return }

func (l *log) stopAndExit(exit func()) { _ = "STUB: not implemented"; return }

func (l *log) RecoverAndExit(f, exit func()) { _ = "STUB: not implemented"; return }
