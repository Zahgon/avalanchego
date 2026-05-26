// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package filesystem

import (
	"io/fs"
)

var _ Reader = reader{}

// Reader is an interface for reading the filesystem.
type Reader interface {
	// ReadDir reads a given directory.
	// Returns the files in the directory.
	ReadDir(string) ([]fs.DirEntry, error)
}

type reader struct{}

// NewReader returns an instance of Reader
func NewReader() Reader {
	_ = "STUB: not implemented"

	// This is just a wrapper around os.ReadDir to make testing easier.
	return *new(Reader)
}

func (reader) ReadDir(dirname string) ([]fs.DirEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
