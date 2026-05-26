// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gossip

import (
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/bloom"
)

func MarshalAppRequest(filter, salt []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseAppRequest(bytes []byte) (*bloom.ReadFilter, ids.ID, error) {
	_ = "STUB: not implemented"
	return nil, *new(ids.ID), nil
}

func MarshalAppResponse(gossip [][]byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseAppResponse(bytes []byte) ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func MarshalAppGossip(gossip [][]byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func ParseAppGossip(bytes []byte) ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }
