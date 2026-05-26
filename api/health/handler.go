// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package health

import (
	"net/http"

	"github.com/ava-labs/avalanchego/utils/logging"
)

// NewGetAndPostHandler returns a health handler that supports GET and jsonrpc
// POST requests.
func NewGetAndPostHandler(log logging.Logger, reporter Reporter) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

// If a GET request is sent, we respond with a 200 if the node is healthy or
// a 503 if the node isn't healthy.

// NewGetHandler return a health handler that supports GET requests reporting
// the result of the provided [reporter].
func NewGetHandler(reporter func(tags ...string) (map[string]Result, bool)) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Make sure the content type is set before writing the header.

// If a health check has failed, we should return a 503.

// The encoder will call write on the writer, which will write the
// header with a 200.
