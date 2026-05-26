// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ids

// NumBits is the number of bits this patricia tree manages
const NumBits = 256

// BitsPerByte is the number of bits per byte
const BitsPerByte = 8

// EqualSubset takes in two indices and two ids and returns if the ids are
// equal from bit start to bit end (non-inclusive). Bit indices are defined as:
// [7 6 5 4 3 2 1 0] [15 14 13 12 11 10 9 8] ... [255 254 253 252 251 250 249 248]
// Where index 7 is the MSB of byte 0.
func EqualSubset(start, stop int, id1, id2 ID) bool { _ = "STUB: not implemented"; return false }

// If there is a series of bytes between the first byte and the last byte, they must be equal

// Index in the byte that the first bit is at
// Index in the byte that the last bit is at

// 111...0... The number of 0s is equal to startBit
// 000...1... The number of 1s is equal to stopBit+1

// If we are looking at the same byte, both masks need to be applied

// The index here could be startIndex or stopIndex, as they are equal

// FirstDifferenceSubset takes in two indices and two ids and returns the index
// of the first difference between the ids inside bit start to bit end
// (non-inclusive). Bit indices are defined above
func FirstDifferenceSubset(start, stop int, id1, id2 ID) (int, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// Index in the byte that the first bit is at
// Index in the byte that the last bit is at

// 111...0... The number of 0s is equal to startBit
// 000...1... The number of 1s is equal to stopBit+1

// If we are looking at the same byte, both masks need to be applied

// The index here could be startIndex or stopIndex, as they are equal

// Check the first byte, may have some bits masked

// Check all the interior bits

// Check the last byte, may have some bits masked

// No difference was found
