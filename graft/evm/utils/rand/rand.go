// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rand

// Credit to Brandur Leach (@Brandur) for this implementation.
// https://brandur.org/fragments/crypto-rand-float64

// Intn is a shortcut for generating a random integer between 0 and
// n using crypto/rand.
func Intn(n int64) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// SecureFloat64 is a shortcut for generating a random float between 0 and
// 1 using crypto/rand.
func SecureFloat64() (float64, error) { _ = "STUB: not implemented"; return 0, nil }
