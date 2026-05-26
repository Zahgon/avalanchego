// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package archivedb

import "github.com/ava-labs/avalanchego/database"

var _ database.KeyValueReader = (*Reader)(nil)

type Reader struct {
	db     *Database
	height uint64
}

func (r *Reader) Has(key []byte) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (r *Reader) Get(key []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// GetEntry retrieves the value of the provided key, the height it was last
// modified at, and a boolean to indicate if the last modification was an
// insertion. If the key has never been modified, ErrNotFound will be returned.
func (r *Reader) GetEntry(key []byte) ([]byte, uint64, bool, error) {
	_ = "STUB: not implemented"
	return nil, 0, false, nil
}

// There is no available key with the requested prefix
