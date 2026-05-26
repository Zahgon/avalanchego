// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import "github.com/ava-labs/avalanchego/ids"

type Header interface {
	ChainID() ids.ID
	ParentID() ids.ID
	BodyID() ids.ID
	Bytes() []byte
}

type statelessHeader struct {
	Chain  ids.ID `serialize:"true"`
	Parent ids.ID `serialize:"true"`
	Body   ids.ID `serialize:"true"`

	bytes []byte
}

func (h *statelessHeader) ChainID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (h *statelessHeader) ParentID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (h *statelessHeader) BodyID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (h *statelessHeader) Bytes() []byte { _ = "STUB: not implemented"; return nil }
