// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package logging

import (
	"go.uber.org/zap"
)

type sanitizedString string

func (s sanitizedString) String() string { _ = "STUB: not implemented"; return "" }

// UserString constructs a field with the given key and the value stripped of
// newlines. The value is sanitized lazily.
func UserString(key, val string) zap.Field { _ = "STUB: not implemented"; return *new(zap.Field) }

type sanitizedStrings []string

func (s sanitizedStrings) String() string { _ = "STUB: not implemented"; return "" }

// UserStrings constructs a field with the given key and the values stripped of
// newlines. The values are sanitized lazily.
func UserStrings(key string, val []string) zap.Field {
	_ = "STUB: not implemented"
	return *new(zap.Field)
}
