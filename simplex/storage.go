// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package simplex

//go:generate go tool canoto $GOFILE

import (
	"context"
	"errors"
	"sync/atomic"

	"github.com/ava-labs/simplex"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/utils/logging"
)

var (
	_               simplex.Storage = (*Storage)(nil)
	genesisMetadata                 = simplex.ProtocolMetadata{
		Version: 0,
		Epoch:   0,
		Round:   0,
		Seq:     0,
	}

	errUnexpectedSeq    = errors.New("unexpected sequence number")
	errInvalidQC        = errors.New("invalid quorum certificate")
	errMismatchedDigest = errors.New("mismatched digest in finalization")

	finalizationPrefix = []byte("f")
	blacklistPrefix    = []byte("b")
)

type Storage struct {
	// numBlocks represents the number of blocks indexed in storage, also known as the height of the chain
	numBlocks atomic.Uint64

	// db is the underlying database used to store finalizations.
	db database.KeyValueReaderWriter

	// genesisBlock is the genesis block data. It is stored as the first block in the storage.
	genesisBlock *Block

	// lastIndexed is the last indexed block digest.
	lastIndexedDigest simplex.Digest

	// deserializer is used to deserialize quorum certificates from bytes.
	deserializer *QCDeserializer

	// blockTracker is used to manage blocks that have been indexed.
	blockTracker *blockTracker

	vm block.ChainVM

	log logging.Logger
}

// newStorage creates a new prefixed database to store
// finalizations according to their sequence numbers.
// The VM is assumed to be initialized before calling this function.
func newStorage(ctx context.Context, config *Config, qcDeserializer *QCDeserializer, blockTracker *blockTracker) (*Storage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// set the last accepted digest by retrieving the last accepted simplex block

func (s *Storage) NumBlocks() uint64 { _ = "STUB: not implemented"; return 0 }

// Retrieve returns the block and finalization at [seq].
// If [seq] is not found, returns simplex.ErrBlockNotFound.
func (s *Storage) Retrieve(seq uint64) (simplex.VerifiedBlock, simplex.Finalization, error) {
	_ = "STUB: not implemented"
	// The genesis block doesn't have a finalization, so we need to handle it specifically.
	return *new(simplex.VerifiedBlock), *new(simplex.Finalization), nil
}

// Index indexes the finalization in the storage.
// It stores the finalization bytes and increments numBlocks.
func (s *Storage) Index(ctx context.Context, block simplex.VerifiedBlock, finalization simplex.Finalization) error {
	_ = "STUB: not implemented"
	return nil
}

// only increment numBlocks after successful indexing

func finalizationKey(seq uint64) []byte { _ = "STUB: not implemented"; return nil }

func blacklistKey(seq uint64) []byte { _ = "STUB: not implemented"; return nil }

// getGenesisBlock returns the genesis block wrapped as a Block instance.
func getGenesisBlock(ctx context.Context, config *Config, blockTracker *blockTracker) (*Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// set the digest

// retrieveFinalization retrieves the finalization at [seq].
// If the finalization is not found, it returns false.
func (s *Storage) retrieveFinalization(seq uint64) (simplex.Finalization, error) {
	_ = "STUB: not implemented"
	return *new(simplex.Finalization), nil
}

func (s *Storage) retrieveBlacklist(seq uint64) (simplex.Blacklist, error) {
	_ = "STUB: not implemented"
	return *new(simplex.Blacklist), nil
}

func getBlock(ctx context.Context, vm block.ChainVM, height uint64) (snowman.Block, error) {
	_ = "STUB: not implemented"
	return *new(snowman.Block), nil
}

// finalizationToBytes serializes the simplex.Finalization into bytes.
func finalizationToBytes(finalization simplex.Finalization) []byte {
	_ = "STUB: not implemented"
	return nil
}

type canotoFinalization struct {
	Finalization []byte `canoto:"bytes,1"`
	QC           []byte `canoto:"bytes,2"`

	canotoData canotoData_canotoFinalization
}

// finalizationFromBytes deserialized the bytes into a simplex.Finalization.
func (c *canotoFinalization) toFinalization(d *QCDeserializer) (simplex.Finalization, error) {
	_ = "STUB: not implemented"
	return *new(simplex.Finalization), nil
}
