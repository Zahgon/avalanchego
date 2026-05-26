// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package json

type Float64 float64

func (f Float64) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *Float64) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
