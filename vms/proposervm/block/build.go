// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"crypto"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/staking"
)

func BuildUnsigned(
	parentID ids.ID,
	timestamp time.Time,
	pChainHeight uint64,
	epoch Epoch,
	blockBytes []byte,
) (SignedBlock, error) {
	_ = "STUB: not implemented"
	return *new(SignedBlock), nil
}

func Build(
	parentID ids.ID,
	timestamp time.Time,
	pChainHeight uint64,
	epoch Epoch,
	cert *staking.Certificate,
	blockBytes []byte,
	chainID ids.ID,
	key crypto.Signer,
) (SignedBlock, error) {
	_ = "STUB: not implemented"
	return *new(SignedBlock), nil
}

// The serialized form of the block is the unsignedBytes followed by the
// signature, which is prefixed by a uint32. Because we are marshalling the
// block with an empty signature, we only need to strip off the length
// prefix to get the unsigned bytes.

// Marshal the final block with signature

// Set the metadata

func BuildHeader(
	chainID ids.ID,
	parentID ids.ID,
	bodyID ids.ID,
) (Header, error) {
	_ = "STUB: not implemented"
	return *new(Header), nil
}

// BuildOption the option block
// [parentID] is the ID of this option's wrapper parent block
// [innerBytes] is the byte representation of a child option block
func BuildOption(
	parentID ids.ID,
	innerBytes []byte,
) (Block, error) {
	_ = "STUB: not implemented"
	return *new(Block), nil
}
