// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package logging

import (
	"errors"

	"go.uber.org/zap/zapcore"
)

type Level zapcore.Level

const (
	Verbo Level = iota - 9
	Debug
	Trace
	Info
	Warn
	Error
	Fatal
	Off

	fatalStr   = "FATAL"
	errorStr   = "ERROR"
	warnStr    = "WARN"
	infoStr    = "INFO"
	traceStr   = "TRACE"
	debugStr   = "DEBUG"
	verboStr   = "VERBO"
	offStr     = "OFF"
	unknownStr = "UNKNO"

	fatalLowStr   = "fatal"
	errorLowStr   = "error"
	warnLowStr    = "warn"
	infoLowStr    = "info"
	traceLowStr   = "trace"
	debugLowStr   = "debug"
	verboLowStr   = "verbo"
	offLowStr     = "off"
	unknownLowStr = "unkno"
)

var ErrUnknownLevel = errors.New("unknown log level")

// Inverse of Level.String()
func ToLevel(l string) (Level, error) { _ = "STUB: not implemented"; return *new(Level), nil }

func (l Level) String() string { _ = "STUB: not implemented"; return "" }

// This should never happen

func (l Level) LowerString() string { _ = "STUB: not implemented"; return "" }

// This should never happen

func (l Level) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (l *Level) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
