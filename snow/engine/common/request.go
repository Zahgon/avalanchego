// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"github.com/ava-labs/avalanchego/ids"
)

type Request struct {
	NodeID    ids.NodeID
	RequestID uint32
}

func (r Request) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
