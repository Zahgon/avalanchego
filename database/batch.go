// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// For ease of implementation, our database's interface matches Ethereum's
// database implementation. This was to allow us to use Geth code as is for the
// EVM chain.

package database

// Batch is a write-only database that commits changes to its host database
// when Write is called. A batch cannot be used concurrently.
type Batch interface {
	KeyValueWriterDeleter

	// Size retrieves the amount of data queued up for writing, this includes
	// the keys, values, and deleted keys.
	Size() int

	// Write flushes any accumulated data to disk.
	Write() error

	// Reset resets the batch for reuse.
	Reset()

	// Replay replays the batch contents in the same order they were written
	// to the batch.
	Replay(w KeyValueWriterDeleter) error

	// Inner returns a Batch writing to the inner database, if one exists. If
	// this batch is already writing to the base DB, then itself should be
	// returned.
	Inner() Batch
}

// Batcher wraps the NewBatch method of a backing data store.
type Batcher interface {
	// NewBatch creates a write-only database that buffers changes to its host db
	// until a final write is called.
	NewBatch() Batch
}

type BatchOp struct {
	Key    []byte
	Value  []byte
	Delete bool
}

type BatchOps struct {
	Ops  []BatchOp
	size int
}

func (b *BatchOps) Put(key, value []byte) error { _ = "STUB: not implemented"; return nil }

func (b *BatchOps) Delete(key []byte) error { _ = "STUB: not implemented"; return nil }

func (b *BatchOps) Size() int { _ = "STUB: not implemented"; return 0 }

func (b *BatchOps) Reset() { _ = "STUB: not implemented"; return }

func (b *BatchOps) Replay(w KeyValueWriterDeleter) error { _ = "STUB: not implemented"; return nil }
