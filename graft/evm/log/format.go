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
	"bytes"
	"math/big"
	"time"

	"github.com/holiman/uint256"
	"golang.org/x/exp/slog"
)

const (
	timeFormat        = "2006-01-02T15:04:05-0700"
	floatFormat       = 'f'
	termMsgJust       = 40
	termCtxMaxPadding = 40
)

// 40 spaces
var spaces = []byte("                                        ")

// TerminalStringer is an analogous interface to the stdlib stringer, allowing
// own types to have custom shortened serialization formats when printed to the
// screen.
type TerminalStringer interface {
	TerminalString() string
}

func (h *TerminalHandler) format(buf []byte, r slog.Record, usecolor bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Note the timestamp is moved before the log level compared to upstream

// Start color

// Prefix is added compared to upstream

// try to justify the log output for short messages
//length := utf8.RuneCountInString(msg)

// print the attributes

func (h *TerminalHandler) formatAttributes(buf *bytes.Buffer, r slog.Record, color string) {
	_ = "STUB: not implemented"
	// tmp is a temporary buffer we use, until bytes.Buffer.AvailableBuffer() (1.21)
	// can be used.
	return
}

//buf.Write(appendEscapeString(buf.AvailableBuffer(), attr.Key))

//buf.Write(appendEscapeString(buf.AvailableBuffer(), attr.Key))

//val := FormatSlogValue(attr.Value, true, buf.AvailableBuffer())

// FormatSlogValue formats a slog.Value for serialization to terminal.
func FormatSlogValue(v slog.Value, tmp []byte) (result []byte) {
	_ = "STUB: not implemented"
	return nil
}

// All int-types (int8, int16 etc) wind up here

// All uint-types (uint8, uint16 etc) wind up here

// Performance optimization: No need for escaping since the provided
// timeFormat doesn't have any escape characters, and escaping is
// expensive.

// Need to be before fmt.Stringer-clause

// Need to be before fmt.Stringer-clause

// We can use the 'tmp' as a scratch-buffer, to first format the
// value, and in a second step do escaping.

// appendInt64 formats n with thousand separators and writes into buffer dst.
func appendInt64(dst []byte, n int64) []byte { _ = "STUB: not implemented"; return nil }

// appendUint64 formats n with thousand separators and writes into buffer dst.
func appendUint64(dst []byte, n uint64, neg bool) []byte {
	_ = "STUB: not implemented"
	// Small numbers are fine as is
	return nil
}

// Large numbers should be split

// FormatLogfmtUint64 formats n with thousand separators.
func FormatLogfmtUint64(n uint64) string { _ = "STUB: not implemented"; return "" }

// appendBigInt formats n with thousand separators and writes to dst.
func appendBigInt(dst []byte, n *big.Int) []byte { _ = "STUB: not implemented"; return nil }

// appendU256 formats n with thousand separators.
func appendU256(dst []byte, n *uint256.Int) []byte { _ = "STUB: not implemented"; return nil }

// appendEscapeString writes the string s to the given writer, with
// escaping/quoting if needed.
func appendEscapeString(dst []byte, s string) []byte { _ = "STUB: not implemented"; return nil }

// If it contains spaces or equal-sign, we need to quote it.

// We need to escape it, if it contains
// - character " (0x22) and lower (except space)
// - characters above ~ (0x7E), plus equal-sign

// No escaping needed, but we might have to place within quote-marks, in case
// it contained a space

// escapeMessage checks if the provided string needs escaping/quoting, similarly
// to escapeString. The difference is that this method is more lenient: it allows
// for spaces and linebreaks to occur without needing quoting.
func escapeMessage(s string) string { _ = "STUB: not implemented"; return "" }

// Allow CR/LF/TAB. This is to make multi-line messages work.

// We quote everything below <space> (0x20) and above~ (0x7E),
// plus equal-sign

// writeTimeTermFormat writes on the format "01-02|15:04:05.000"
func writeTimeTermFormat(buf *bytes.Buffer, t time.Time) { _ = "STUB: not implemented"; return }

// writePosIntWidth writes non-negative integer i to the buffer, padded on the left
// by zeroes to the given width. Use a width of 0 to omit padding.
// Adapted from golang.org/x/exp/slog/internal/buffer/buffer.go
func writePosIntWidth(b *bytes.Buffer, i, width int) {
	_ = "STUB: not implemented"
	// Cheap integer to fixed-width decimal ASCII.
	// Copied from log/log.go.
	return
}

// Assemble decimal in reverse order.

// i < 10
