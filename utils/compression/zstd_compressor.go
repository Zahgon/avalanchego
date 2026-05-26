// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package compression

import (
	"errors"
)

var (
	_ Compressor = (*zstdCompressor)(nil)

	ErrInvalidMaxSizeCompressor = errors.New("invalid compressor max size")
	ErrDecompressedMsgTooLarge  = errors.New("decompressed msg too large")
	ErrMsgTooLarge              = errors.New("msg too large to be compressed")
)

func NewZstdCompressor(maxSize int64) (Compressor, error) {
	_ = "STUB: not implemented"
	return *new(Compressor), nil
}

func NewZstdCompressorWithLevel(maxSize int64, level int) (Compressor, error) {
	_ = "STUB: not implemented"
	return *new(Compressor), nil
}

// "Decompress" creates "io.LimitReader" with max size + 1:
// if the max size + 1 overflows, "io.LimitReader" reads nothing
// returning 0 byte for the decompress call
// require max size < math.MaxInt64 to prevent int64 overflows

type zstdCompressor struct {
	maxSize int64
	level   int
}

func (z *zstdCompressor) Compress(msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (z *zstdCompressor) Decompress(msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We allow [io.LimitReader] to read up to [z.maxSize + 1] bytes, so that if
// the decompressed payload is greater than the maximum size, this function
// will return the appropriate error instead of an incomplete byte slice.
