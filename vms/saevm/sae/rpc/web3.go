// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"github.com/ava-labs/libevm/common/hexutil"
)

// web3API offers the `web3` RPCs.
type web3API struct {
	clientVersion string
}

func newWeb3API() *web3API { _ = "STUB: not implemented"; return nil }

func (w *web3API) ClientVersion() string { _ = "STUB: not implemented"; return "" }

func (*web3API) Sha3(input hexutil.Bytes) hexutil.Bytes {
	_ = "STUB: not implemented"
	return *new(hexutil.Bytes)
}
