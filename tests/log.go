// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tests

import (
	"github.com/ava-labs/avalanchego/utils/logging"
)

func NewDefaultLogger(prefix string) logging.Logger {
	_ = "STUB: not implemented"
	return *new(logging.Logger)
}

// This should never happen since auto is a valid log format

// TODO(marun) Does/should the logging package have a function like this?
func LoggerForFormat(prefix string, rawLogFormat string) (logging.Logger, error) {
	_ = "STUB: not implemented"
	return *new(logging.Logger), nil
}

// TODO(marun) Make the log level configurable
