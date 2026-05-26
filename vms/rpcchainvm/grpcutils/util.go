// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package grpcutils

import (
	"net/http"
	"time"

	httppb "github.com/ava-labs/avalanchego/proto/pb/http"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

// GetGRPCErrorFromHTTPResponse takes an HandleSimpleHTTPResponse as input and returns a gRPC error.
func GetGRPCErrorFromHTTPResponse(resp *httppb.HandleSimpleHTTPResponse) error {
	_ = "STUB: not implemented"
	return nil
}

// GetHTTPResponseFromError takes an gRPC error as input and returns a gRPC
// HandleSimpleHTTPResponse.
func GetHTTPResponseFromError(err error) (*httppb.HandleSimpleHTTPResponse, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetHTTPHeader takes an http.Header as input and returns a slice of Header.
func GetHTTPHeader(hs http.Header) []*httppb.Element { _ = "STUB: not implemented"; return nil }

// SetHeaders sets headers to next
func SetHeaders(headers http.Header, next []*httppb.Element) { _ = "STUB: not implemented"; return }

// TimestampAsTime validates timestamppb timestamp and returns time.Time.
func TimestampAsTime(ts *tspb.Timestamp) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// TimestampFromTime converts time.Time to a timestamppb timestamp.
func TimestampFromTime(time time.Time) *tspb.Timestamp { _ = "STUB: not implemented"; return nil }

// EnsureValidResponseCode ensures that the response code is valid otherwise it returns 500.
func EnsureValidResponseCode(code int) int {
	_ = "STUB: not implemented"
	// Response code outside of this range is invalid and could panic.
	// ref. https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html
	return 0
}
