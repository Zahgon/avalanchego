// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package logging

import (
	"sync"

	"go.uber.org/zap"
)

var _ Factory = (*factory)(nil)

// Factory creates new instances of different types of Logger
type Factory interface {
	// Make creates a new logger with name [name]
	Make(name string) (Logger, error)

	// MakeChain creates a new logger to log the events of chain [chainID]
	MakeChain(chainID string) (Logger, error)

	// SetLogLevel sets log levels for all loggers in factory with given logger name, level pairs.
	SetLogLevel(name string, level Level) error

	// SetDisplayLevel sets log display levels for all loggers in factory with given logger name, level pairs.
	SetDisplayLevel(name string, level Level) error

	// GetLogLevel returns all log levels in factory as name, level pairs
	GetLogLevel(name string) (Level, error)

	// GetDisplayLevel returns all log display levels in factory as name, level pairs
	GetDisplayLevel(name string) (Level, error)

	// GetLoggerNames returns the names of all logs created by this factory
	GetLoggerNames() []string

	// Close stops and clears all of a Factory's instantiated loggers
	Close()
}

type logWrapper struct {
	logger       Logger
	displayLevel zap.AtomicLevel
	logLevel     zap.AtomicLevel
}

type factory struct {
	config Config
	lock   sync.RWMutex

	// For each logger created by this factory:
	// Logger name --> the logger.
	loggers map[string]logWrapper
}

// NewFactory returns a new instance of a Factory producing loggers configured with
// the values set in the [config] parameter
func NewFactory(config Config) Factory { _ = "STUB: not implemented"; return *new(Factory) }

// Assumes [f.lock] is held
func (f *factory) makeLogger(config Config) (Logger, error) {
	_ = "STUB: not implemented"
	return *new(Logger), nil
}

// megabytes
// days
// files

func (f *factory) Make(name string) (Logger, error) {
	_ = "STUB: not implemented"
	return *new(Logger), nil
}

func (f *factory) MakeChain(chainID string) (Logger, error) {
	_ = "STUB: not implemented"
	return *new(Logger), nil
}

func (f *factory) SetLogLevel(name string, level Level) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *factory) SetDisplayLevel(name string, level Level) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *factory) GetLogLevel(name string) (Level, error) {
	_ = "STUB: not implemented"
	return *new(Level), nil
}

func (f *factory) GetDisplayLevel(name string) (Level, error) {
	_ = "STUB: not implemented"
	return *new(Level), nil
}

func (f *factory) GetLoggerNames() []string { _ = "STUB: not implemented"; return nil }

func (f *factory) Close() { _ = "STUB: not implemented"; return }
