// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package utils

import (
	"sync"
)

// RandomBytes returns a slice of n random bytes
// Intended for use in testing
func RandomBytes(n int) []byte { _ = "STUB: not implemented"; return nil }

// Constant taken from the "math" package
const intSize = 32 << (^uint(0) >> 63) // 32 or 64

// BytesPool tracks buckets of available buffers to be allocated. Each bucket
// allocates buffers of the following length:
//
// 0
// 1
// 3
// 7
// 15
// 31
// 63
// 127
// ...
// MaxInt
//
// In order to allocate a buffer of length 19 (for example), we calculate the
// number of bits required to represent 19 (5). And therefore allocate a slice
// from bucket 5, which has length 31. This is the bucket which produces the
// smallest slices that are at least length 19.
//
// When replacing a buffer of length 19, we calculate the number of bits
// required to represent 20 (5). And therefore place the slice into bucket 4,
// which has length 15. This is the bucket which produces the largest slices
// that a length 19 slice can be used for.
type BytesPool [intSize]sync.Pool

func NewBytesPool() *BytesPool { _ = "STUB: not implemented"; return nil }

// uint is used here to avoid overflowing int during the shift

// Sync pool needs to return pointer-like values to avoid memory
// allocations.

// Get returns a non-nil pointer to a slice with the requested length.
//
// It is not guaranteed for the returned bytes to have been zeroed.
func (p *BytesPool) Get(length int) *[]byte { _ = "STUB: not implemented"; return nil }

// Round up

// Set the length to be the expected value

// Put takes ownership of a non-nil pointer to a slice of bytes.
//
// Note: this function takes ownership of the underlying array. So, the length
// of the provided slice is ignored and only its capacity is used.
func (p *BytesPool) Put(bytes *[]byte) { _ = "STUB: not implemented"; return }

// Round down
