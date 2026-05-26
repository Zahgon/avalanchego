// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ghttp

import (
	"net/http"

	"github.com/ava-labs/avalanchego/utils/logging"

	httppb "github.com/ava-labs/avalanchego/proto/pb/http"
)

var _ http.Handler = (*Client)(nil)

// Client is an http.Handler that talks over RPC.
type Client struct {
	client httppb.HTTPClient
	log    logging.Logger
}

// NewClient returns an HTTP handler database instance connected to a remote
// HTTP handler instance
func NewClient(client httppb.HTTPClient, log logging.Logger) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	// rfc2616#section-14.42: The Upgrade general-header allows the client
	// to specify a communication protocols it supports and would like to
	// use. Upgrade (e.g. websockets) is a more expensive transaction and
	// if not required use the less expensive HTTPSimple.
	//
	// Http/2 explicitly does not allow the use of the Upgrade header.
	// (ref: https://httpwg.org/specs/rfc9113.html#informational-responses)
	return
}

// Wrap [w] with a lock to ensure that it is accessed in a thread-safe manner.

// Start responsewriter gRPC service.

// serveHTTPSimple converts an http request to a gRPC HTTPRequest and returns the
// response to the client. Protocol upgrade requests (websockets) are not supported
// and should use ServeHTTP.
func (c *Client) serveHTTPSimple(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// Some errors will actually contain a valid resp, just need to unpack it

// getHTTPSimpleRequest takes an http request as input and returns a gRPC HandleSimpleHTTPRequest.
func getHTTPSimpleRequest(w http.ResponseWriter, r *http.Request) (*httppb.HandleSimpleHTTPRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// convertWriteResponse converts a gRPC HandleSimpleHTTPResponse to an HTTP response.
func convertWriteResponse(w http.ResponseWriter, resp *httppb.HandleSimpleHTTPResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// isUpgradeRequest returns true if the upgrade key exists in header and value is non empty.
func isUpgradeRequest(req *http.Request) bool { _ = "STUB: not implemented"; return false }

func isHTTP2Request(req *http.Request) bool { _ = "STUB: not implemented"; return false }
