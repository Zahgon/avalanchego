// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package connectclient

import (
	"net/http"

	"connectrpc.com/connect"
)

var _ connect.Interceptor = (*SetRouteHeaderInterceptor)(nil)

// SetRouteHeaderInterceptor sets the api routing header for connect-rpc
// requests
type SetRouteHeaderInterceptor struct {
	Route []string
}

func (s SetRouteHeaderInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	_ = "STUB: not implemented"
	return *new(connect.UnaryFunc)
}

func (s SetRouteHeaderInterceptor) WrapStreamingClient(next connect.StreamingClientFunc) connect.StreamingClientFunc {
	_ = "STUB: not implemented"
	return *new(connect.StreamingClientFunc)
}

func (SetRouteHeaderInterceptor) WrapStreamingHandler(next connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	_ = "STUB: not implemented"
	return *new(connect.StreamingHandlerFunc)
}

func New() *http.Client { _ = "STUB: not implemented"; return nil }

// Skip TLS to use h2c
