// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package json

import (
	"errors"
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

const (
	// Null is the string representation of a null value
	Null = "null"
)

var (
	errUppercaseMethod = errors.New("method must start with a non-uppercase letter")
	errInvalidArg      = errors.New("couldn't unmarshal an argument. Ensure arguments are valid and properly formatted. See documentation for example calls")
)

// NewCodec returns a new json codec that will convert the first character of
// the method to uppercase
func NewCodec() rpc.Codec { _ = "STUB: not implemented"; return *new(rpc.Codec) }

type lowercase struct{ *json2.Codec }

func (lc lowercase) NewRequest(r *http.Request) rpc.CodecRequest {
	_ = "STUB: not implemented"
	return *new(rpc.CodecRequest)
}

type request struct{ *json2.CodecRequest }

func (r *request) Method() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (r *request) ReadRequest(args interface{}) error { _ = "STUB: not implemented"; return nil }
