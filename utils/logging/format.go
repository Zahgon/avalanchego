// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package logging

import (
	"errors"

	"go.uber.org/zap/zapcore"
)

// Format modes available
const (
	Auto Format = iota
	Plain
	Colors
	JSON

	AutoString   = "auto"
	PlainString  = "plain"
	ColorsString = "colors"
	JSONString   = "json"

	FormatDescription = "The structure of log format. Defaults to 'auto' which formats terminal-like logs, when the output is a terminal. Otherwise, should be one of {auto, plain, colors, json}"

	termTimeFormat = "[01-02|15:04:05.000]"
)

var (
	formatJSON = []string{
		`"auto"`,
		`"plain"`,
		`"colors"`,
		`"json"`,
	}

	errUnknownFormat = errors.New("unknown format")

	defaultEncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	jsonEncoderConfig zapcore.EncoderConfig

	termTimeEncoder = zapcore.TimeEncoderOfLayout(termTimeFormat)
)

func init() {
	jsonEncoderConfig = defaultEncoderConfig
	jsonEncoderConfig.EncodeLevel = jsonLevelEncoder
	jsonEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	jsonEncoderConfig.EncodeDuration = zapcore.NanosDurationEncoder
}

// Highlight mode to apply to displayed logs
type Format int

// ToFormat chooses a highlighting mode
func ToFormat(h string, fd uintptr) (Format, error) {
	_ = "STUB: not implemented"
	return *new(Format), nil
}

func (f Format) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (f Format) WrapPrefix(prefix string) string { _ = "STUB: not implemented"; return "" }

func (f Format) ConsoleEncoder() zapcore.Encoder {
	_ = "STUB: not implemented"
	return *new(zapcore.Encoder)
}

func (f Format) FileEncoder() zapcore.Encoder {
	_ = "STUB: not implemented"
	return *new(zapcore.Encoder)
}

func newTermEncoderConfig(lvlEncoder zapcore.LevelEncoder) zapcore.EncoderConfig {
	_ = "STUB: not implemented"
	return *new(zapcore.EncoderConfig)
}

func levelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func jsonLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}

func ConsoleColorLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	_ = "STUB: not implemented"
	return
}
