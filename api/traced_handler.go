// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package api

import (
	"net/http"

	"github.com/ava-labs/avalanchego/trace"
)

var _ http.Handler = (*tracedHandler)(nil)

type tracedHandler struct {
	h            http.Handler
	serveHTTPTag string
	tracer       trace.Tracer
}

func TraceHandler(h http.Handler, name string, tracer trace.Tracer) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (h *tracedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
