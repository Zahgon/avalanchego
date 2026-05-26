// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package storage

// FileExists checks if a file exists before we
// try using it to prevent further errors.
func FileExists(filePath string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ReadFileWithName reads a single file with name fileNameWithoutExt without specifying any extension.
// it errors when there are more than 1 file with the given fileName
func ReadFileWithName(parentDir string, fileNameNoExt string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// all possible extensions

// no file found, return nothing

// FolderExists checks if a folder exists before we
// try using it to prevent further errors.
func FolderExists(filePath string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func DirSize(path string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
