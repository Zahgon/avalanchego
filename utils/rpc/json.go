// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"
	"io"
	"net/url"
)

// CleanlyCloseBody avoids sending unnecessary RST_STREAM and PING frames by
// ensuring the whole body is read before being closed.
// See https://blog.cloudflare.com/go-and-enhance-your-calm/#reading-bodies-in-go-can-be-unintuitive
func CleanlyCloseBody(body io.ReadCloser) { _ = "STUB: not implemented"; return }

func SendJSONRequest(
	ctx context.Context,
	uri *url.URL,
	method string,
	params interface{},
	reply interface{},
	options ...Option,
) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:bodyclose // body is closed via CleanlyCloseBody

// Return an error for any non successful status code
