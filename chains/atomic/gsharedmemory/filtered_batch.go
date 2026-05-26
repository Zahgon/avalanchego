// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gsharedmemory

import (
	"github.com/ava-labs/avalanchego/utils/set"

	sharedmemorypb "github.com/ava-labs/avalanchego/proto/pb/sharedmemory"
)

type filteredBatch struct {
	writes  map[string][]byte
	deletes set.Set[string]
}

func (b *filteredBatch) Put(key []byte, value []byte) error { _ = "STUB: not implemented"; return nil }

func (b *filteredBatch) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (b *filteredBatch) PutRequests() []*sharedmemorypb.BatchPut {
	_ = "STUB: not implemented"
	return nil
}

func (b *filteredBatch) DeleteRequests() []*sharedmemorypb.BatchDelete {
	_ = "STUB: not implemented"
	return nil
}
