// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package saetest

import (
	"context"
	"testing"

	"go.uber.org/zap"

	"github.com/ava-labs/avalanchego/utils/logging"
)

// logger is the common wrapper around [LogRecorder] and [tbLogger] handlers,
// plumbing all levels into the handler.
type logger struct {
	level   logging.Level
	handler interface {
		log(logging.Level, string, ...zap.Field)
	}
	with []zap.Field
	// Some methods will panic, in which case they need to be implemented. This
	// is better than embedding a [logging.NoLog], which could silently drop
	// important entries.
	logging.Logger
}

var _ logging.Logger = (*logger)(nil)

func (l *logger) With(fields ...zap.Field) logging.Logger {
	_ = "STUB: not implemented"
	return *new(logging.Logger)
}

func (l *logger) log(lvl logging.Level, msg string, fields ...zap.Field) {
	_ = "STUB: not implemented"
	return
}

func (l *logger) Debug(msg string, fs ...zap.Field) { _ = "STUB: not implemented"; return }
func (l *logger) Trace(msg string, fs ...zap.Field) { _ = "STUB: not implemented"; return }
func (l *logger) Info(msg string, fs ...zap.Field)  { _ = "STUB: not implemented"; return }
func (l *logger) Warn(msg string, fs ...zap.Field)  { _ = "STUB: not implemented"; return }
func (l *logger) Error(msg string, fs ...zap.Field) { _ = "STUB: not implemented"; return }
func (l *logger) Fatal(msg string, fs ...zap.Field) { _ = "STUB: not implemented"; return }

// NewLogRecorder constructs a new [LogRecorder] at the specified level.
func NewLogRecorder(level logging.Level) *LogRecorder { _ = "STUB: not implemented"; return nil }

// yes, the recursion is gross, but that's composition for you ¯\_(ツ)_/¯

// A LogRecorder is a [logging.Logger] that stores all logs as [LogRecord]
// entries for inspection.
type LogRecorder struct {
	*logger
	Records []*LogRecord
}

// A LogRecord is a single entry in a [LogRecorder].
type LogRecord struct {
	Level  logging.Level
	Msg    string
	Fields []zap.Field
}

func (l *LogRecorder) log(lvl logging.Level, msg string, fields ...zap.Field) {
	_ = "STUB: not implemented"
	return
}

// Filter returns the recorded logs for which `fn` returns true.
func (l *LogRecorder) Filter(fn func(*LogRecord) bool) []*LogRecord {
	_ = "STUB: not implemented"
	return nil
}

// At returns all recorded logs at the specified [logging.Level].
func (l *LogRecorder) At(lvl logging.Level) []*LogRecord { _ = "STUB: not implemented"; return nil }

// AtLeast returns all recorded logs at or above the specified [logging.Level].
func (l *LogRecorder) AtLeast(lvl logging.Level) []*LogRecord {
	_ = "STUB: not implemented"
	return nil
}

// NewTBLogger constructs a logger that propagates logs to [testing.TB]. WARNING
// and ERROR logs are sent to [testing.TB.Errorf] while FATAL is sent to
// [testing.TB.Fatalf]. All other logs are sent to [testing.TB.Logf]. Although
// the level can be configured, it is silently capped at [logging.Warn].
//
//nolint:thelper // The outputs include the logging site while the TB site is most useful if here
func NewTBLogger(tb testing.TB, level logging.Level) *TBLogger {
	_ = "STUB: not implemented"
	return nil
}

// TODO(arr4n) remove the recursion here and in [LogRecorder]

// TBLogger is a [logging.Logger] that propagates logs to [testing.TB].
type TBLogger struct {
	*logger
	tb      testing.TB
	onError []context.CancelFunc
}

// CancelOnError pipes `ctx` to and from [context.WithCancel], calling the
// [context.CancelFunc] after logs >= [logging.Error], and during [testing.TB]
// cleanup.
func (l *TBLogger) CancelOnError(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (l *TBLogger) log(lvl logging.Level, msg string, fields ...zap.Field) {
	_ = "STUB: not implemented"
	return
}

// because @ARR4N says warnings in tests are errors
