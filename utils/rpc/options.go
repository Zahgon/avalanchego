// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"net/http"
	"net/url"
)

type Option func(*Options)

type Options struct {
	headers     http.Header
	queryParams url.Values
}

func NewOptions(ops []Option) *Options { _ = "STUB: not implemented"; return nil }

func (o *Options) applyOptions(ops []Option) { _ = "STUB: not implemented"; return }

func (o *Options) Headers() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (o *Options) QueryParams() url.Values { _ = "STUB: not implemented"; return *new(url.Values) }

func WithHeader(key, val string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithQueryParam(key, val string) Option { _ = "STUB: not implemented"; return *new(Option) }
