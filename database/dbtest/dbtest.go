// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package dbtest

import (
	"testing"

	"github.com/ava-labs/avalanchego/database"
)

// TestsBasic is a list of all basic database tests that require only
// a KeyValueReaderWriterDeleter.
var TestsBasic = map[string]func(t *testing.T, db database.KeyValueReaderWriterDeleter){
	"SimpleKeyValue":       TestSimpleKeyValue,
	"OverwriteKeyValue":    TestOverwriteKeyValue,
	"EmptyKey":             TestEmptyKey,
	"KeyEmptyValue":        TestKeyEmptyValue,
	"MemorySafetyDatabase": TestMemorySafetyDatabase,
	"ModifyValueAfterPut":  TestModifyValueAfterPut,
	"PutGetEmpty":          TestPutGetEmpty,
}

// Tests is a list of all database tests
var Tests = map[string]func(t *testing.T, db database.Database){
	"SimpleKeyValueClosed":             TestSimpleKeyValueClosed,
	"NewBatchClosed":                   TestNewBatchClosed,
	"BatchPut":                         TestBatchPut,
	"BatchDelete":                      TestBatchDelete,
	"BatchReset":                       TestBatchReset,
	"BatchReuse":                       TestBatchReuse,
	"BatchRewrite":                     TestBatchRewrite,
	"BatchReplay":                      TestBatchReplay,
	"BatchReplayPropagateError":        TestBatchReplayPropagateError,
	"BatchInner":                       TestBatchInner,
	"BatchLargeSize":                   TestBatchLargeSize,
	"IteratorSnapshot":                 TestIteratorSnapshot,
	"Iterator":                         TestIterator,
	"IteratorStart":                    TestIteratorStart,
	"IteratorPrefix":                   TestIteratorPrefix,
	"IteratorStartPrefix":              TestIteratorStartPrefix,
	"IteratorMemorySafety":             TestIteratorMemorySafety,
	"IteratorClosed":                   TestIteratorClosed,
	"IteratorError":                    TestIteratorError,
	"IteratorErrorAfterRelease":        TestIteratorErrorAfterRelease,
	"CompactNoPanic":                   TestCompactNoPanic,
	"MemorySafetyBatch":                TestMemorySafetyBatch,
	"AtomicClear":                      TestAtomicClear,
	"Clear":                            TestClear,
	"AtomicClearPrefix":                TestAtomicClearPrefix,
	"ClearPrefix":                      TestClearPrefix,
	"ModifyValueAfterBatchPut":         TestModifyValueAfterBatchPut,
	"ModifyValueAfterBatchPutReplay":   TestModifyValueAfterBatchPutReplay,
	"ConcurrentBatches":                TestConcurrentBatches,
	"ManySmallConcurrentKVPairBatches": TestManySmallConcurrentKVPairBatches,
}

func init() {
	// Add all basic database tests to the database tests
	for name, test := range TestsBasic {
		Tests[name] = func(t *testing.T, db database.Database) {
			test(t, db)
		}
	}
}

// TestSimpleKeyValue tests to make sure that simple Put + Get + Delete + Has
// calls return the expected values.
func TestSimpleKeyValue(t *testing.T, db database.KeyValueReaderWriterDeleter) {
	_ = "STUB: not implemented"
	return
}

func TestOverwriteKeyValue(t *testing.T, db database.KeyValueReaderWriterDeleter) {
	_ = "STUB: not implemented"
	return
}

func TestKeyEmptyValue(t *testing.T, db database.KeyValueReaderWriterDeleter) {
	_ = "STUB: not implemented"
	return
}

func TestEmptyKey(t *testing.T, db database.KeyValueReaderWriterDeleter) {
	_ = "STUB: not implemented"
	return
}

// Test that nil key can be retrieved by empty key

// Test that empty key can be retrieved by nil key

// TestSimpleKeyValueClosed tests to make sure that Put + Get + Delete + Has
// calls return the correct error when the database has been closed.
func TestSimpleKeyValueClosed(t *testing.T, db database.Database) {
	_ = "STUB: not implemented"
	return
}

// TestMemorySafetyDatabase ensures it is safe to modify a key after passing it
// to Database.Put and Database.Get.
func TestMemorySafetyDatabase(t *testing.T, db database.KeyValueReaderWriterDeleter) {
	_ = "STUB: not implemented"
	return
}

// Put key in the database directly

// Put key2 in the database by modifying key, which should be safe
// to modify after the Put call

// Get the value for [key]

// Modify [key]; make sure the value we got before hasn't changed

// Reset [key] to its original value and make sure it's correct

// TestNewBatchClosed tests to make sure that calling NewBatch on a closed
// database returns a batch that errors correctly.
func TestNewBatchClosed(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// TestBatchPut tests to make sure that batched writes work as expected.
func TestBatchPut(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// TestBatchDelete tests to make sure that batched deletes work as expected.
func TestBatchDelete(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// TestMemorySafetyBatch ensures it is safe to modify a key after passing it
// to Batch.Put.
func TestMemorySafetyBatch(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// Put a key in the batch

// Modify the key

// Make sure the original key was written to the database

// Make sure the new key wasn't written to the database

// TestBatchReset tests to make sure that a batch drops un-written operations
// when it is reset.
func TestBatchReset(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// TestBatchReuse tests to make sure that a batch can be reused once it is
// reset.
func TestBatchReuse(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// TestBatchRewrite tests to make sure that write can be called multiple times
// on a batch and the values will be updated correctly.
func TestBatchRewrite(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// TestBatchReplay tests to make sure that batches will correctly replay their
// contents.
func TestBatchReplay(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// TestBatchReplayPropagateError tests to make sure that batches will correctly
// propagate any returned error during Replay.
func TestBatchReplayPropagateError(t *testing.T, db database.Database) {
	_ = "STUB: not implemented"
	return
}

// TestBatchInner tests to make sure that inner can be used to write to the
// database.
func TestBatchInner(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// TestBatchLargeSize tests to make sure that the batch can support a large
// amount of entries.
func TestBatchLargeSize(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// 8 KiB

// TestIteratorSnapshot tests to make sure the database iterates over a snapshot
// of the database at the time of the iterator creation.
func TestIteratorSnapshot(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// TestIterator tests to make sure the database iterates over the database
// contents lexicographically.
func TestIterator(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// TestIteratorStart tests to make sure the iterator can be configured to
// start mid way through the database.
func TestIteratorStart(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// TestIteratorPrefix tests to make sure the iterator can be configured to skip
// keys missing the provided prefix.
func TestIteratorPrefix(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// TestIteratorStartPrefix tests to make sure that the iterator can start mid
// way through the database while skipping a prefix.
func TestIteratorStartPrefix(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// TestIteratorMemorySafety tests to make sure that keys can values are able to
// be modified from the returned iterator.
func TestIteratorMemorySafety(t *testing.T, db database.Database) {
	_ = "STUB: not implemented"
	return
}

// TestIteratorClosed tests to make sure that an iterator that was created with
// a closed database will report a closed error correctly.
func TestIteratorClosed(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// TestIteratorError tests to make sure that an iterator on a database will report
// itself as being exhausted and return [database.ErrClosed] to indicate that the iteration
// was not successful.
// Additionally tests that an iterator that has already called Next() can still serve
// its current value after the underlying DB was closed.
func TestIteratorError(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// Call Next() and ensure that if the database is closed, the iterator
// can still report the current contents.

// Subsequent calls to the iterator should return false and report an error

// TestIteratorErrorAfterRelease tests to make sure that an iterator that was
// released still reports the error correctly.
func TestIteratorErrorAfterRelease(t *testing.T, db database.Database) {
	_ = "STUB: not implemented"
	return
}

// TestCompactNoPanic tests to make sure compact never panics.
func TestCompactNoPanic(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// Test compacting with nil bounds

// Test compacting when start > end

// Test compacting when start > largest key

func TestAtomicClear(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

func TestClear(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// testClear tests to make sure the deletion helper works as expected.
func testClear(t *testing.T, db database.Database, clearF func(database.Database) error) {
	_ = "STUB: not implemented"
	return
}

func TestAtomicClearPrefix(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

func TestClearPrefix(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

// testClearPrefix tests to make sure prefix deletion works as expected.
func testClearPrefix(t *testing.T, db database.Database, clearF func(database.Database, []byte) error) {
	_ = "STUB: not implemented"
	return
}

func TestModifyValueAfterPut(t *testing.T, db database.KeyValueReaderWriterDeleter) {
	_ = "STUB: not implemented"
	return
}

// Modify the value that was Put into the database
// to see if the database copied the value correctly.

func TestModifyValueAfterBatchPut(t *testing.T, db database.Database) {
	_ = "STUB: not implemented"
	return
}

// Modify the value that was Put into the Batch and then Write the
// batch to the database.

// Verify that the value written to the database contains matches the original
// value of the byte slice when Put was called.

func TestModifyValueAfterBatchPutReplay(t *testing.T, db database.Database) {
	_ = "STUB: not implemented"
	return
}

// Modify the value that was Put into the Batch and then Write the
// batch to the database.

// Create a new batch and replay the batch onto this one before writing it to the DB.

// Verify that the value written to the database contains matches the original
// value of the byte slice when Put was called.

func TestConcurrentBatches(t *testing.T, db database.Database) { _ = "STUB: not implemented"; return }

func TestManySmallConcurrentKVPairBatches(t *testing.T, db database.Database) {
	_ = "STUB: not implemented"
	return
}

func runConcurrentBatches(
	db database.Database,
	numBatches,
	keysPerBatch,
	keySize,
	valueSize int,
) error {
	_ = "STUB: not implemented"
	return nil
}

func TestPutGetEmpty(t *testing.T, db database.KeyValueReaderWriterDeleter) {
	_ = "STUB: not implemented"
	return
}

// May be nil or empty byte slice.

// May be nil or empty byte slice.

func FuzzKeyValue(f *testing.F, db database.KeyValueReaderWriterDeleter) {
	_ = "STUB: not implemented"
	return
}

func FuzzNewIteratorWithPrefix(f *testing.F, db database.Database) {
	_ = "STUB: not implemented"
	return
}

// #nosec G404

// Put a bunch of key-values

// #nosec G404

// #nosec G404

// Consistently treat zero length values as nil
// so that we can compare [expected] and [got] with
// require.Equal, which treats nil and empty byte
// as being unequal, whereas the database treats
// them as being equal.

// Assert the iterator returns the expected key-values.

// Clear the database for the next fuzz iteration.

func FuzzNewIteratorWithStartAndPrefix(f *testing.F, db database.Database) {
	_ = "STUB: not implemented"
	return
}

// #nosec G404

// Put a bunch of key-values

// #nosec G404

// #nosec G404

// Consistently treat zero length values as nil
// so that we can compare [expected] and [got] with
// require.Equal, which treats nil and empty byte
// as being unequal, whereas the database treats
// them as being equal.

// Assert the iterator returns the expected key-values.

// Clear the database for the next fuzz iteration.
