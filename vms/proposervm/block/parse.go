// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"github.com/ava-labs/avalanchego/ids"
)

type ParseResult struct {
	Block Block
	Err   error
}

// ParseBlocks parses the given raw blocks into tuples of (Block, error).
// Each ParseResult is returned in the same order as its corresponding bytes in the input.
func ParseBlocks(blks [][]byte, chainID ids.ID) []ParseResult {
	_ = "STUB: not implemented"
	return nil
}

// Parse a block and verify that the signature attached to the block is valid
// for the certificate provided in the block and that the block has a valid
// representation.
func Parse(bytes []byte, chainID ids.ID) (Block, error) {
	_ = "STUB: not implemented"
	return *new(Block), nil
}

// ParseWithoutVerification parses a block without verifying that the signature
// on the block is correct or has valid representation.
func ParseWithoutVerification(bytes []byte) (Block, error) {
	_ = "STUB: not implemented"
	return *new(Block), nil
}

func ParseHeader(bytes []byte) (Header, error) { _ = "STUB: not implemented"; return *new(Header), nil }
