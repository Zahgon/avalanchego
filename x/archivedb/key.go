// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package archivedb

import (
	"errors"
)

var (
	ErrParsingKeyLength   = errors.New("failed reading key length")
	ErrIncorrectKeyLength = errors.New("incorrect key length")

	heightKey = newDBKeyFromMetadata([]byte{})
)

// The requirements of a database key are:
//
// 1. A given user key must have a unique database key prefix. This guarantees
// that user keys can not overlap on disk.
// 2. Inside of a database key prefix, the database keys must be sorted by
// decreasing height.
// 3. User keys must never overlap with any metadata keys.

// newDBKeyFromUser converts a user key and height into a database formatted
// key.
//
// To meet the requirements of a database key, the prefix is defined by
// concatenating the length of the user key and the user key. The suffix of the
// database key is the negation of the big endian encoded height. This suffix
// guarantees the keys are sorted correctly.
//
//	Example (Asumming heights are 1 byte):
//	 |  User key  |  Stored as  |
//	 |------------|-------------|
//	 |   foo:10   |  3:foo:245  |
//	 |   foo:20   |  3:foo:235  |
//
// Returns:
// - The database key
// - The database key prefix, which is independent of the height
func newDBKeyFromUser(key []byte, height uint64) ([]byte, []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseDBKeyFromUser takes a database formatted key and returns the user key
// along with its height.
//
// Note: An error should only be returned from this function if the database has
// been corrupted.
func parseDBKeyFromUser(dbKey []byte) ([]byte, uint64, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// newDBKeyFromMetadata converts a metadata key into a database formatted key.
//
// To meet the requirements of a database key, the key is defined by
// concatenating the length of the metadata key + 1 and the metadata key.
//
//	Example:
//	 |  Metadata key  |  Stored as  |
//	 |----------------|-------------|
//	 |       foo      |    4:foo    |
//	 |       fo       |    3:fo     |
func newDBKeyFromMetadata(key []byte) []byte { _ = "STUB: not implemented"; return nil }
