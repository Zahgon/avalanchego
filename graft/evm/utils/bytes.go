// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import "github.com/ava-labs/libevm/common"

// IncrOne increments bytes value by one
func IncrOne(bytes []byte) { _ = "STUB: not implemented"; return }

// HashSliceToBytes serializes a []common.Hash into a tightly packed byte array.
func HashSliceToBytes(hashes []common.Hash) []byte { _ = "STUB: not implemented"; return nil }

// BytesToHashSlice packs [b] into a slice of hash values with zero padding
// to the right if the length of b is not a multiple of 32.
func BytesToHashSlice(b []byte) []common.Hash { _ = "STUB: not implemented"; return nil }
