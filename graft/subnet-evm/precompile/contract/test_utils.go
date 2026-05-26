// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package contract

import (
	"github.com/ava-labs/libevm/common"
)

// PackOrderedHashesWithSelector packs the function selector and ordered list of hashes into [dst]
// byte slice.
// assumes that [dst] has sufficient room for [functionSelector] and [hashes].
// Kept for testing backwards compatibility.
func PackOrderedHashesWithSelector(dst []byte, functionSelector []byte, hashes []common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// PackOrderedHashes packs the ordered list of [hashes] into the [dst] byte buffer.
// assumes that [dst] has sufficient space to pack [hashes] or else this function will panic.
// Kept for testing backwards compatibility.
func PackOrderedHashes(dst []byte, hashes []common.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// PackedHash returns packed the byte slice with common.HashLength from [packed]
// at the given [index].
// Assumes that [packed] is composed entirely of packed 32 byte segments.
// Kept for testing backwards compatibility.
func PackedHash(packed []byte, index int) []byte { _ = "STUB: not implemented"; return nil }
