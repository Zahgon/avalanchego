// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package json

type Uint64 uint64

func (u Uint64) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (u *Uint64) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
