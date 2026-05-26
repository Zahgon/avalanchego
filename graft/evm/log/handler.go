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
	"io"
	"sync"

	"golang.org/x/exp/slog"
)

type discardHandler struct{}

// DiscardHandler returns a no-op handler
func DiscardHandler() slog.Handler { _ = "STUB: not implemented"; return *new(slog.Handler) }

func (h *discardHandler) Handle(_ context.Context, r slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *discardHandler) Enabled(_ context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *discardHandler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *discardHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

type TerminalHandler struct {
	mu       sync.Mutex
	wr       io.Writer
	lvl      slog.Leveler
	useColor bool
	attrs    []slog.Attr
	// fieldPadding is a map with maximum field value lengths seen until now
	// to allow padding log contexts in a bit smarter way.
	fieldPadding map[string]int

	buf []byte

	// Prefix returns a string that is output before each log message.
	Prefix func(r slog.Record) string
}

// NewTerminalHandler returns a handler which formats log records at all levels optimized for human readability on
// a terminal with color-coded level output and terser human friendly timestamp.
// This format should only be used for interactive programs or while developing.
//
//	[LEVEL] [TIME] MESSAGE key=value key=value ...
//
// Example:
//
//	[DBUG] [May 16 20:58:45] remove route ns=haproxy addr=127.0.0.1:50002
func NewTerminalHandler(wr io.Writer, useColor bool) *TerminalHandler {
	_ = "STUB: not implemented"
	return nil
}

// NewTerminalHandlerWithLevel returns the same handler as NewTerminalHandler but only outputs
// records which are less than or equal to the specified verbosity level.
func NewTerminalHandlerWithLevel(wr io.Writer, lvl slog.Leveler, useColor bool) *TerminalHandler {
	_ = "STUB: not implemented"
	return nil
}

func (h *TerminalHandler) Handle(_ context.Context, r slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *TerminalHandler) Enabled(_ context.Context, level slog.Level) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *TerminalHandler) WithGroup(name string) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func (h *TerminalHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

// ResetFieldPadding zeroes the field-padding for all attribute pairs.
func (t *TerminalHandler) ResetFieldPadding() { _ = "STUB: not implemented"; return }

// JSONHandler returns a handler which prints records in JSON format.
func JSONHandler(wr io.Writer) slog.Handler { _ = "STUB: not implemented"; return *new(slog.Handler) }

func JSONHandlerWithLevel(wr io.Writer, level slog.Leveler) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

// LogfmtHandler returns a handler which prints records in logfmt format, an easy machine-parseable but human-readable
// format for key/value pairs.
//
// For more details see: http://godoc.org/github.com/kr/logfmt
func LogfmtHandler(wr io.Writer) slog.Handler { _ = "STUB: not implemented"; return *new(slog.Handler) }

// LogfmtHandlerWithLevel returns the same handler as LogfmtHandler but it only outputs
// records which are less than or equal to the specified verbosity level.
func LogfmtHandlerWithLevel(wr io.Writer, level slog.Leveler) slog.Handler {
	_ = "STUB: not implemented"
	return *new(slog.Handler)
}

func builtinReplaceLogfmt(_ []string, attr slog.Attr) slog.Attr {
	_ = "STUB: not implemented"
	return *new(slog.Attr)
}

func builtinReplaceJSON(_ []string, attr slog.Attr) slog.Attr {
	_ = "STUB: not implemented"
	return *new(slog.Attr)
}

func builtinReplace(_ []string, attr slog.Attr, logfmt bool) slog.Attr {
	_ = "STUB: not implemented"
	return *new(slog.Attr)
}
