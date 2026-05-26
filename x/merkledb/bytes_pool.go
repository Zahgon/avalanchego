// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import "sync"

type bytesPool struct {
	slots     chan struct{}
	bytesLock sync.Mutex
	bytes     [][]byte
}

func newBytesPool(numSlots int) *bytesPool { _ = "STUB: not implemented"; return nil }

func (p *bytesPool) Acquire() []byte { _ = "STUB: not implemented"; return nil }

func (p *bytesPool) TryAcquire() ([]byte, bool) { _ = "STUB: not implemented"; return nil, false }

func (p *bytesPool) pop() []byte { _ = "STUB: not implemented"; return nil }

func (p *bytesPool) Release(b []byte) {
	_ = "STUB: not implemented"
	// Before waking anyone waiting on a slot, return the bytes.
	return
}
