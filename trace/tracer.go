// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package trace

import (
	"io"
	"time"

	"go.opentelemetry.io/otel/trace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	tracerExportTimeout = 10 * time.Second
	// [tracerProviderShutdownTimeout] is longer than [tracerExportTimeout] so
	// in-flight exports can finish before the tracer provider shuts down.
	tracerProviderShutdownTimeout = 15 * time.Second
)

type Config struct {
	ExporterConfig `json:"exporterConfig"`

	// The fraction of traces to sample.
	// If >= 1 always samples.
	// If <= 0 never samples.
	TraceSampleRate float64 `json:"traceSampleRate"`

	AppName string `json:"appName"`
	Version string `json:"version"`
}

type Tracer interface {
	trace.Tracer
	io.Closer
}

type tracer struct {
	trace.Tracer

	tp *sdktrace.TracerProvider
}

func (t *tracer) Close() error { _ = "STUB: not implemented"; return nil }

func New(config Config) (Tracer, error) { _ = "STUB: not implemented"; return *new(Tracer), nil }
