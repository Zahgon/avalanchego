// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package dbtest

import (
	"testing"

	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/utils/units"
)

var (
	// Benchmarks is a list of all database benchmarks
	Benchmarks = map[string]func(b *testing.B, db database.Database, keys, values [][]byte){
		"Get":            BenchmarkGet,
		"Put":            BenchmarkPut,
		"Delete":         BenchmarkDelete,
		"BatchPut":       BenchmarkBatchPut,
		"BatchDelete":    BenchmarkBatchDelete,
		"BatchWrite":     BenchmarkBatchWrite,
		"ParallelGet":    BenchmarkParallelGet,
		"ParallelPut":    BenchmarkParallelPut,
		"ParallelDelete": BenchmarkParallelDelete,
	}
	// BenchmarkSizes to use with each benchmark
	BenchmarkSizes = [][]int{
		// count, keySize, valueSize
		{1024, 32, 32},
		{1024, 256, 256},
		{1024, 2 * units.KiB, 2 * units.KiB},
	}
)

// Writes size data into the db in order to setup reads in subsequent tests.
func SetupBenchmark(b *testing.B, count int, keySize, valueSize int) ([][]byte, [][]byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// #nosec G404

// #nosec G404

// BenchmarkGet measures the time it takes to get an operation from a database.
func BenchmarkGet(b *testing.B, db database.Database, keys, values [][]byte) {
	_ = "STUB: not implemented"
	return
}

// Reads b.N values from the db

// BenchmarkPut measures the time it takes to write an operation to a database.
func BenchmarkPut(b *testing.B, db database.Database, keys, values [][]byte) {
	_ = "STUB: not implemented"
	return
}

// Writes b.N values to the db

// BenchmarkDelete measures the time it takes to delete a (k, v) from a database.
func BenchmarkDelete(b *testing.B, db database.Database, keys, values [][]byte) {
	_ = "STUB: not implemented"
	return
}

// Writes random values of size _size_ to the database

// Deletes b.N values from the db

// BenchmarkBatchPut measures the time it takes to batch put.
func BenchmarkBatchPut(b *testing.B, db database.Database, keys, values [][]byte) {
	_ = "STUB: not implemented"
	return
}

// BenchmarkBatchDelete measures the time it takes to batch delete.
func BenchmarkBatchDelete(b *testing.B, db database.Database, keys, _ [][]byte) {
	_ = "STUB: not implemented"
	return
}

// BenchmarkBatchWrite measures the time it takes to batch write.
func BenchmarkBatchWrite(b *testing.B, db database.Database, keys, values [][]byte) {
	_ = "STUB: not implemented"
	return
}

// BenchmarkParallelGet measures the time it takes to read in parallel.
func BenchmarkParallelGet(b *testing.B, db database.Database, keys, values [][]byte) {
	_ = "STUB: not implemented"
	return
}

// BenchmarkParallelPut measures the time it takes to write to the db in parallel.
func BenchmarkParallelPut(b *testing.B, db database.Database, keys, values [][]byte) {
	_ = "STUB: not implemented"
	return
}

// Write N values to the db

// BenchmarkParallelDelete measures the time it takes to delete a (k, v) from the db.
func BenchmarkParallelDelete(b *testing.B, db database.Database, keys, values [][]byte) {
	_ = "STUB: not implemented"
	return
}

// Deletes b.N values from the db
