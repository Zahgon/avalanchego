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

package log

import (
	"context"
	"math"

	"golang.org/x/exp/slog"
)

const errorKey = "LOG_ERROR"

const (
	legacyLevelCrit = iota
	legacyLevelError
	legacyLevelWarn
	legacyLevelInfo
	legacyLevelDebug
	legacyLevelTrace
)

const (
	levelMaxVerbosity slog.Level = math.MinInt
	LevelTrace        slog.Level = -8
	LevelDebug                   = slog.LevelDebug
	LevelInfo                    = slog.LevelInfo
	LevelWarn                    = slog.LevelWarn
	LevelError                   = slog.LevelError
	LevelCrit         slog.Level = 12

	// for backward-compatibility
	LvlTrace = LevelTrace
	LvlInfo  = LevelInfo
	LvlDebug = LevelDebug
)

// LvlFromString returns the appropriate Lvl from a string name.
// Useful for parsing command line args and configuration files.
func LvlFromString(lvlString string) (slog.Level, error) {
	_ = "STUB: not implemented"
	return *new(slog.Level), nil
}

// convert from old Geth verbosity level constants
// to levels defined by slog
func FromLegacyLevel(lvl int) slog.Level { _ = "STUB: not implemented"; return *new(slog.Level) }

// TODO: should we allow use of custom levels or force them to match existing max/min if they fall outside the range as I am doing here?

// LevelAlignedString returns a 5-character string containing the name of a Lvl.
func LevelAlignedString(l slog.Level) string { _ = "STUB: not implemented"; return "" }

// LevelString returns a string containing the name of a Lvl.
func LevelString(l slog.Level) string { _ = "STUB: not implemented"; return "" }

// A Logger writes key/value pairs to a Handler
type Logger interface {
	// With returns a new Logger that has this logger's attributes plus the given attributes
	With(ctx ...interface{}) Logger

	// With returns a new Logger that has this logger's attributes plus the given attributes. Identical to 'With'.
	New(ctx ...interface{}) Logger

	// Log logs a message at the specified level with context key/value pairs
	Log(level slog.Level, msg string, ctx ...interface{})

	// Trace log a message at the trace level with context key/value pairs
	Trace(msg string, ctx ...interface{})

	// Debug logs a message at the debug level with context key/value pairs
	Debug(msg string, ctx ...interface{})

	// Info logs a message at the info level with context key/value pairs
	Info(msg string, ctx ...interface{})

	// Warn logs a message at the warn level with context key/value pairs
	Warn(msg string, ctx ...interface{})

	// Error logs a message at the error level with context key/value pairs
	Error(msg string, ctx ...interface{})

	// Crit logs a message at the crit level with context key/value pairs, and exits
	Crit(msg string, ctx ...interface{})

	// Write logs a message at the specified level
	Write(level slog.Level, msg string, attrs ...any)

	// Enabled reports whether l emits log records at the given context and level.
	Enabled(ctx context.Context, level slog.Level) bool
}

type logger struct {
	inner *slog.Logger
}

// NewLogger returns a logger with the specified handler set
func NewLogger(h slog.Handler) Logger { _ = "STUB: not implemented"; return *new(Logger) }

// write logs a message at the specified level:
func (l *logger) Write(level slog.Level, msg string, attrs ...any) {
	_ = "STUB: not implemented"
	return
}

func (l *logger) Log(level slog.Level, msg string, attrs ...any) { _ = "STUB: not implemented"; return }

func (l *logger) With(ctx ...interface{}) Logger { _ = "STUB: not implemented"; return *new(Logger) }

func (l *logger) New(ctx ...interface{}) Logger {
	_ = "STUB: not implemented"
	return *

	// Enabled reports whether l emits log records at the given context and level.
	new(Logger)
}

func (l *logger) Enabled(ctx context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (l *logger) Trace(msg string, ctx ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Debug(msg string, ctx ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Info(msg string, ctx ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Warn(msg string, ctx ...any) { _ = "STUB: not implemented"; return }

func (l *logger) Error(msg string, ctx ...interface{}) { _ = "STUB: not implemented"; return }

func (l *logger) Crit(msg string, ctx ...interface{}) { _ = "STUB: not implemented"; return }
