// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package filesystem

import "io/fs"

var _ fs.DirEntry = MockFile{}

// MockFile is an implementation of fs.File for unit testing.
type MockFile struct {
	MockName    string
	MockIsDir   bool
	MockType    fs.FileMode
	MockInfo    fs.FileInfo
	MockInfoErr error
}

func (m MockFile) Name() string { _ = "STUB: not implemented"; return "" }

func (m MockFile) IsDir() bool { _ = "STUB: not implemented"; return false }

func (m MockFile) Type() fs.FileMode { _ = "STUB: not implemented"; return *new(fs.FileMode) }

func (m MockFile) Info() (fs.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(fs.FileInfo), nil
}
