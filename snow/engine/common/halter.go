// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

var _ Haltable = (*Halter)(nil)

type Haltable interface {
	Halt()
	Halted() bool
}

type Halter struct {
	halted uint32
}

func (h *Halter) Halt() { _ = "STUB: not implemented"; return }

func (h *Halter) Halted() bool { _ = "STUB: not implemented"; return false }
