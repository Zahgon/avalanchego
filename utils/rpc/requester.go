// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"
)

var _ EndpointRequester = (*avalancheEndpointRequester)(nil)

type EndpointRequester interface {
	SendRequest(ctx context.Context, method string, params interface{}, reply interface{}, options ...Option) error
}

type avalancheEndpointRequester struct {
	uri string
}

func NewEndpointRequester(uri string) EndpointRequester {
	_ = "STUB: not implemented"
	return *new(EndpointRequester)
}

func (e *avalancheEndpointRequester) SendRequest(
	ctx context.Context,
	method string,
	params interface{},
	reply interface{},
	options ...Option,
) error {
	_ = "STUB: not implemented"
	return nil
}
