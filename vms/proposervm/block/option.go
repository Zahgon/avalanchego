// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"github.com/ava-labs/avalanchego/ids"
)

type option struct {
	PrntID     ids.ID `serialize:"true"`
	InnerBytes []byte `serialize:"true"`

	id    ids.ID
	bytes []byte
}

func (b *option) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (b *option) ParentID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (b *option) Block() []byte { _ = "STUB: not implemented"; return nil }

func (b *option) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *option) initialize(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (*option) verify(ids.ID) error { _ = "STUB: not implemented"; return nil }
