// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package trace

import (
	"time"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const tracerProviderExportCreationTimeout = 5 * time.Second

type ExporterConfig struct {
	Type ExporterType `json:"type"`

	// Endpoint to send metrics to. If empty, the default endpoint will be used.
	Endpoint string `json:"endpoint"`

	// Headers to send with metrics
	Headers map[string]string `json:"headers"`

	// If true, don't use TLS
	Insecure bool `json:"insecure"`
}

func newExporter(config ExporterConfig) (sdktrace.SpanExporter, error) {
	_ = "STUB: not implemented"
	return *new(sdktrace.SpanExporter), nil
}
