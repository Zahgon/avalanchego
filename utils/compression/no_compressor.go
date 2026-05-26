// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package compression

var _ Compressor = (*noCompressor)(nil)

type noCompressor struct{}

func (*noCompressor) Compress(msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (*noCompressor) Decompress(msg []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewNoCompressor() Compressor { _ = "STUB: not implemented"; return *new(Compressor) }
