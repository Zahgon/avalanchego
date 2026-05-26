// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package simplex

//go:generate go tool canoto $GOFILE

import (
	"context"
	"errors"
	"sync"

	"github.com/ava-labs/simplex"

	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/utils/tree"
)

var (
	_ simplex.BlockDeserializer = (*blockDeserializer)(nil)
	_ simplex.Block             = (*Block)(nil)
	_ simplex.VerifiedBlock     = (*Block)(nil)

	errDigestNotFound         = errors.New("digest not found in block tracker")
	errMismatchedPrevDigest   = errors.New("prev digest does not match block parent")
	errGenesisVerification    = errors.New("genesis block should not be verified")
	errFailedToParseMetadata  = errors.New("failed to parse protocol metadata")
	errFailedToParseBlacklist = errors.New("failed to parse blacklist")
)

type Block struct {
	digest simplex.Digest

	// metadata contains protocol metadata for the block
	metadata simplex.ProtocolMetadata

	// the parsed block
	vmBlock snowman.Block

	blockTracker *blockTracker

	blacklist simplex.Blacklist
}

func newBlock(metadata simplex.ProtocolMetadata, blacklist simplex.Blacklist, vmBlock snowman.Block, blockTracker *blockTracker) (*Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CanotoSimplexBlock is the Canoto representation of a block
type canotoSimplexBlock struct {
	Metadata   []byte `canoto:"bytes,1"`
	InnerBlock []byte `canoto:"bytes,2"`
	Blacklist  []byte `canoto:"bytes,3"`

	canotoData canotoData_canotoSimplexBlock
}

// BlockHeader returns the block header for the block.
func (b *Block) BlockHeader() simplex.BlockHeader {
	_ = "STUB: not implemented"
	return *new(simplex.BlockHeader)
}

// Bytes returns the serialized bytes of the block.
func (b *Block) Bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (b *Block) Blacklist() simplex.Blacklist {
	_ = "STUB: not implemented"
	return *

	// Verify verifies the block.
	new(simplex.Blacklist)
}

func (b *Block) Verify(ctx context.Context) (simplex.VerifiedBlock, error) {
	_ = "STUB: not implemented"
	// we should not verify the genesis block
	return *new(simplex.VerifiedBlock), nil
}

// verifyParentMatchesPrevBlock verifies that the previous block referenced in the current block's metadata
// matches the parent of the current block's vmBlock.
func (b *Block) verifyParentMatchesPrevBlock() error { _ = "STUB: not implemented"; return nil }

func computeDigest(bytes []byte) simplex.Digest {
	_ = "STUB: not implemented"
	return *new(simplex.Digest)
}

type blockDeserializer struct {
	parser       block.Parser
	blockTracker *blockTracker
}

func (d *blockDeserializer) DeserializeBlock(ctx context.Context, bytes []byte) (simplex.Block, error) {
	_ = "STUB: not implemented"
	return *new(simplex.Block), nil
}

// blockTracker is used to ensure that blocks are properly rejected, if competing blocks are accepted.
type blockTracker struct {
	lock sync.Mutex

	// tracks the simplex digests to the blocks that have been verified
	simplexDigestsToBlock map[simplex.Digest]*Block

	// handles block acceptance and rejection of inner blocks
	tree tree.Tree

	vm block.ChainVM
}

func newBlockTracker(vm block.ChainVM) *blockTracker { _ = "STUB: not implemented"; return nil }

// init sets the latest block in the tracker.
// This should only be called once, with the genesis or latest block.
func (bt *blockTracker) init(latestBlock *Block) {
	bt.simplexDigestsToBlock[latestBlock.digest] = latestBlock
}

func (bt *blockTracker) getBlockByDigest(digest simplex.Digest) (*Block, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// verifyAndTrackBlock verifies the block, sets it as the VM's preference, and tracks it in the block tracker.
// If the block is already verified, it does nothing.
func (bt *blockTracker) verifyAndTrackBlock(ctx context.Context, block *Block) error {
	_ = "STUB: not implemented"
	return nil
}

// check if the block is already verified

// track the block

// indexBlock calls accept on the block with the given digest, and reject on competing blocks.
func (bt *blockTracker) indexBlock(ctx context.Context, digest simplex.Digest) error {
	_ = "STUB: not implemented"
	return nil
}

// removes all digests with a lower seq

// notify the VM that we are accepting this block, and reject all competing blocks
