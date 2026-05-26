// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package json

type Float32 float32

func (f Float32) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (f *Float32) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
