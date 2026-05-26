// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bloom

func Add(f *Filter, key, salt []byte) bool { _ = "STUB: not implemented"; return false }

func Contains(c Checker, key, salt []byte) bool { _ = "STUB: not implemented"; return false }

type Checker interface {
	Contains(hash uint64) bool
}

func Hash(key, salt []byte) uint64 {
	_ = "STUB: not implemented"

	// sha256.Write never returns errors
	return 0
}
