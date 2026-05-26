// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package compression

import (
	"errors"
)

var errUnknownCompressionType = errors.New("unknown compression type")

type Type byte

const (
	TypeNone Type = iota + 1
	TypeZstd
)

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

func TypeFromString(s string) (Type, error) { _ = "STUB: not implemented"; return *new(Type), nil }

func (t Type) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
