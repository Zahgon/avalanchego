// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package log

import (
	"context"
	"io"

	"golang.org/x/exp/slog"

	ethlog "github.com/ava-labs/libevm/log"
)

type Logger struct {
	ethlog.Logger

	logLevel *slog.LevelVar
}

// InitLogger initializes logger with alias and sets the log level and format with the original [os.StdErr] interface
// along with the context logger.
func InitLogger(alias string, level string, jsonFormat bool, writer io.Writer) (Logger, error) {
	_ = "STUB: not implemented"
	return *new(Logger), nil
}

// Create handler

// SetLogLevel sets the log level of initialized log handler.
func (l *Logger) SetLogLevel(level string) error {
	_ = "STUB: not implemented"
	// Set log level
	return nil
}

// locationTrims are trimmed for display to avoid unwieldy log lines.
var locationTrims = []string{
	"coreth",
}

func trimPrefixes(s string) string { _ = "STUB: not implemented"; return "" }

func getSource(r slog.Record) (string, int) { _ = "STUB: not implemented"; return "", 0 }

type addContext struct {
	slog.Handler

	logger string
}

func (a *addContext) Handle(ctx context.Context, r slog.Record) error {
	_ = "STUB: not implemented"
	return nil
}
