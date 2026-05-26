// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package json

type Uint32 uint32

func (u Uint32) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (u *Uint32) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
