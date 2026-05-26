// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package json

type Uint16 uint16

func (u Uint16) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (u *Uint16) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
