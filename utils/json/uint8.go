// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package json

type Uint8 uint8

func (u Uint8) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (u *Uint8) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
