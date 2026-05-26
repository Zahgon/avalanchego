// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package server

import (
	"net/http"

	"github.com/ava-labs/avalanchego/utils/set"
)

const wildcard = "*"

var _ http.Handler = (*allowedHostsHandler)(nil)

func filterInvalidHosts(
	handler http.Handler,
	allowed []string,
) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// wildcards match all hostnames, so just return the base handler

// allowedHostsHandler is an implementation of http.Handler that validates the
// http host header of incoming requests. This can prevent DNS rebinding attacks
// which do not utilize CORS-headers. Http request host headers are validated
// against a whitelist to determine whether the request should be dropped or
// not.
type allowedHostsHandler struct {
	handler http.Handler
	hosts   set.Set[string]
}

func (a *allowedHostsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	// if the host header is missing we can serve this request because dns
	// rebinding attacks rely on this header
	return
}

// either invalid (too many colons) or no port specified

// accept requests from ips

// a specific hostname - we need to check the whitelist to see if we should
// accept this r
