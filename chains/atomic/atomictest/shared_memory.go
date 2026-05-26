// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomictest

import (
	"testing"

	"github.com/ava-labs/avalanchego/chains/atomic"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
)

// SharedMemoryTests is a list of all shared memory tests
var SharedMemoryTests = []func(t *testing.T, chainID0, chainID1 ids.ID, sm0, sm1 atomic.SharedMemory, db database.Database){
	TestSharedMemoryPutAndGet,
	TestSharedMemoryLargePutGetAndRemove,
	TestSharedMemoryIndexed,
	TestSharedMemoryLargeIndexed,
	TestSharedMemoryCantDuplicatePut,
	TestSharedMemoryCantDuplicateRemove,
	TestSharedMemoryCommitOnPut,
	TestSharedMemoryCommitOnRemove,
	TestSharedMemoryLargeBatchSize,
	TestPutAndRemoveBatch,
}

func TestSharedMemoryPutAndGet(t *testing.T, chainID0, chainID1 ids.ID, sm0, sm1 atomic.SharedMemory, _ database.Database) {
	_ = "STUB: not implemented"
	return
}

// TestSharedMemoryLargePutGetAndRemove tests to make sure that the interface
// can support large values.
func TestSharedMemoryLargePutGetAndRemove(t *testing.T, chainID0, chainID1 ids.ID, sm0, sm1 atomic.SharedMemory, _ database.Database) {
	_ = "STUB: not implemented"
	return
}

//#nosec G404

// 16 MiB
// 4 KiB
// 8 KiB

// #nosec G404

func TestSharedMemoryIndexed(t *testing.T, chainID0, chainID1 ids.ID, sm0, sm1 atomic.SharedMemory, _ database.Database) {
	_ = "STUB: not implemented"
	return
}

func TestSharedMemoryLargeIndexed(t *testing.T, chainID0, chainID1 ids.ID, sm0, sm1 atomic.SharedMemory, _ database.Database) {
	_ = "STUB: not implemented"
	return
}

// 8 MiB
// 1 KiB
// 3 KiB

// #nosec G404

func TestSharedMemoryCantDuplicatePut(t *testing.T, _, chainID1 ids.ID, sm0, _ atomic.SharedMemory, _ database.Database) {
	_ = "STUB: not implemented"
	return
}

// TODO: require error to be errDuplicatedOperation
//nolint:forbidigo // currently returns grpc errors too

// TODO: require error to be errDuplicatedOperation
//nolint:forbidigo // currently returns grpc errors too

func TestSharedMemoryCantDuplicateRemove(t *testing.T, _, chainID1 ids.ID, sm0, _ atomic.SharedMemory, _ database.Database) {
	_ = "STUB: not implemented"
	return
}

// TODO: require error to be errDuplicatedOperation
//nolint:forbidigo // currently returns grpc errors too

func TestSharedMemoryCommitOnPut(t *testing.T, _, chainID1 ids.ID, sm0, _ atomic.SharedMemory, db database.Database) {
	_ = "STUB: not implemented"
	return
}

func TestSharedMemoryCommitOnRemove(t *testing.T, _, chainID1 ids.ID, sm0, _ atomic.SharedMemory, db database.Database) {
	_ = "STUB: not implemented"
	return
}

// TestPutAndRemoveBatch tests to make sure multiple put and remove requests work properly
func TestPutAndRemoveBatch(t *testing.T, chainID0, _ ids.ID, _, sm1 atomic.SharedMemory, db database.Database) {
	_ = "STUB: not implemented"
	return
}

// TestSharedMemoryLargeBatchSize tests to make sure that the interface can
// support large batches.
func TestSharedMemoryLargeBatchSize(t *testing.T, _, chainID1 ids.ID, sm0, _ atomic.SharedMemory, db database.Database) {
	_ = "STUB: not implemented"
	return
}

//#nosec G404

// 8 MiB
// 4 KiB
// 8 KiB

// #nosec G404
