// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"math/big"
	"time"
)

func TimeToNewUint64(time time.Time) *uint64 { _ = "STUB: not implemented"; return nil }

func Uint64ToTime(val *uint64) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Uint64PtrEqual returns true if x and y pointers are equivalent ie. both nil or both
// contain the same value.
func Uint64PtrEqual(x, y *uint64) bool { _ = "STUB: not implemented"; return false }

// BigEqual returns true if a is equal to b. If a and b are nil, it returns
// true.
func BigEqual(a, b *big.Int) bool { _ = "STUB: not implemented"; return false }

// BigEqualUint64 returns true if a is equal to b. If a is nil or not a uint64,
// it returns false.
func BigEqualUint64(a *big.Int, b uint64) bool { _ = "STUB: not implemented"; return false }

// BigLessOrEqualUint64 returns true if a is less than or equal to b. If a is
// nil or not a uint64, it returns false.
func BigLessOrEqualUint64(a *big.Int, b uint64) bool { _ = "STUB: not implemented"; return false }
