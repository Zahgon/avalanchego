// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package stacktrace

import (
	"os"
	"runtime"
)

// If the environment variable STACK_TRACE_ERRORS=1 is set, errors
// passing through the functions defined in this package will have a
// stack trace added to them. The following equivalents to stdlib error
// functions are provided:
//
// - `fmt.Errorf` -> `stacktrace.Errorf`
// - `errors.New` -> `stacktrace.New`
//
// Additionally, a stack trace can be added to an existing error with
// `stacktrace.Wrap(err)`.

var stackTraceErrors bool

func init() {
	if os.Getenv("STACK_TRACE_ERRORS") == "1" {
		stackTraceErrors = true
	}
}

type StackTraceError struct {
	StackTrace []runtime.Frame
	Cause      error
}

func (e StackTraceError) Error() string { _ = "STUB: not implemented"; return "" }

func (e StackTraceError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func New(msg string) error { _ = "STUB: not implemented"; return nil }

// Errorf adds a stack trace to the last argument provided if it is an
// error and stack traces are enabled.
func Errorf(format string, args ...any) error { _ = "STUB: not implemented"; return nil }

// Assume the last argument is an error requiring a stack trace if it is of type error

// If there's already a StackTraceError, preserve its stack but update the cause

// No stack trace exists, capture one now

func Wrap(err error) error { _ = "STUB: not implemented"; return nil }

// wrap adds a stack trace to err if stack traces are enabled and it
// doesn't already have one.
func wrap(err error) error { _ = "STUB: not implemented"; return nil }

// If there's already a StackTraceError in the chain, just return it

// Need to capture a stack trace

// skip wrap, New/Wrap/Errorf, and runtime.Callers
