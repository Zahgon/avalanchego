// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package messagetest

import (
	"testing"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/graft/evm/message"
)

// ForEachCodec runs fn as a subtest for each supported codec (coreth and subnet-evm).
// The callback receives the codec manager and the corresponding leafs request type.
// The test case name is set automatically via t.Run.
func ForEachCodec(t *testing.T, fn func(codec.Manager, message.LeafsRequestType)) {
	_ = "STUB: not implemented"
	return
}
