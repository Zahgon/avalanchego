// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package wrappers

import (
	"errors"
	"math"
)

const (
	MaxStringLen = math.MaxUint16

	// ByteLen is the number of bytes per byte...
	ByteLen = 1
	// ShortLen is the number of bytes per short
	ShortLen = 2
	// IntLen is the number of bytes per int
	IntLen = 4
	// LongLen is the number of bytes per long
	LongLen = 8
	// BoolLen is the number of bytes per bool
	BoolLen = 1
)

func StringLen(str string) int {
	_ = "STUB: not implemented"
	// note: there is a max length for string ([MaxStringLen])
	// we defer to PackString checking whether str is within limits
	return 0
}

var (
	ErrInsufficientLength = errors.New("packer has insufficient length for input")
	errNegativeOffset     = errors.New("negative offset")
	errInvalidInput       = errors.New("input does not match expected format")
	errBadBool            = errors.New("unexpected value when unpacking bool")
	errOversized          = errors.New("size is larger than limit")
)

// Packer packs and unpacks a byte array from/to standard values
type Packer struct {
	Errs

	// The largest allowed size of expanding the byte array
	MaxSize int
	// The current byte array
	Bytes []byte
	// The offset that is being written to in the byte array
	Offset int
}

// PackByte append a byte to the byte array
func (p *Packer) PackByte(val byte) { _ = "STUB: not implemented"; return }

// UnpackByte unpack a byte from the byte array
func (p *Packer) UnpackByte() byte { _ = "STUB: not implemented"; return 0 }

// PackShort append a short to the byte array
func (p *Packer) PackShort(val uint16) { _ = "STUB: not implemented"; return }

// UnpackShort unpack a short from the byte array
func (p *Packer) UnpackShort() uint16 { _ = "STUB: not implemented"; return 0 }

// PackInt append an int to the byte array
func (p *Packer) PackInt(val uint32) { _ = "STUB: not implemented"; return }

// UnpackInt unpack an int from the byte array
func (p *Packer) UnpackInt() uint32 { _ = "STUB: not implemented"; return 0 }

// PackLong append a long to the byte array
func (p *Packer) PackLong(val uint64) { _ = "STUB: not implemented"; return }

// UnpackLong unpack a long from the byte array
func (p *Packer) UnpackLong() uint64 { _ = "STUB: not implemented"; return 0 }

// PackBool packs a bool into the byte array
func (p *Packer) PackBool(b bool) { _ = "STUB: not implemented"; return }

// UnpackBool unpacks a bool from the byte array
func (p *Packer) UnpackBool() bool { _ = "STUB: not implemented"; return false }

// PackFixedBytes append a byte slice, with no length descriptor to the byte
// array
func (p *Packer) PackFixedBytes(bytes []byte) { _ = "STUB: not implemented"; return }

// UnpackFixedBytes unpack a byte slice, with no length descriptor from the byte
// array
func (p *Packer) UnpackFixedBytes(size int) []byte { _ = "STUB: not implemented"; return nil }

// PackBytes append a byte slice to the byte array
func (p *Packer) PackBytes(bytes []byte) { _ = "STUB: not implemented"; return }

// UnpackBytes unpack a byte slice from the byte array
func (p *Packer) UnpackBytes() []byte { _ = "STUB: not implemented"; return nil }

// UnpackLimitedBytes unpacks a byte slice. If the size of the slice is greater
// than [limit], adds [errOversized] to the packer and returns nil.
func (p *Packer) UnpackLimitedBytes(limit uint32) []byte { _ = "STUB: not implemented"; return nil }

// PackStr append a string to the byte array
func (p *Packer) PackStr(str string) { _ = "STUB: not implemented"; return }

// UnpackStr unpacks a string from the byte array
func (p *Packer) UnpackStr() string { _ = "STUB: not implemented"; return "" }

// UnpackLimitedStr unpacks a string. If the size of the string is greater than
// [limit], adds [errOversized] to the packer and returns the empty string.
func (p *Packer) UnpackLimitedStr(limit uint16) string { _ = "STUB: not implemented"; return "" }

// checkSpace requires that there is at least [bytes] of write space left in the
// byte array. If this is not true, an error is added to the packer
func (p *Packer) checkSpace(bytes int) { _ = "STUB: not implemented"; return }

// expand ensures that there is [bytes] bytes left of space in the byte slice.
// If this is not allowed due to the maximum size, an error is added to the packer
// In order to understand this code, its important to understand the difference
// between a slice's length and its capacity.
func (p *Packer) expand(bytes int) { _ = "STUB: not implemented"; return }

// Need byte slice's length to be at least [neededSize]

// Byte slice has sufficient length already

// Lengthening the byte slice would cause it to grow too large

// Byte slice has sufficient capacity to lengthen it without mem alloc

// Add capacity/length to byte slice
