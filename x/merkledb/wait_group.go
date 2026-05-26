// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package merkledb

import "sync"

// waitGroup is a small wrapper of a sync.WaitGroup that avoids performing a
// memory allocation when Add is never called.
type waitGroup struct {
	wg *sync.WaitGroup
}

func (wg *waitGroup) Add(delta int) { _ = "STUB: not implemented"; return }

func (wg *waitGroup) Wait() { _ = "STUB: not implemented"; return }
