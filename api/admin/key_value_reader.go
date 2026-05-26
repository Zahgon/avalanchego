// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package admin

import (
	"github.com/ava-labs/avalanchego/database"
)

var _ database.KeyValueReader = (*KeyValueReader)(nil)

type KeyValueReader struct {
	client *Client
}

func NewKeyValueReader(client *Client) *KeyValueReader { _ = "STUB: not implemented"; return nil }

func (r *KeyValueReader) Has(key []byte) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (r *KeyValueReader) Get(key []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
