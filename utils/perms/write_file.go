// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package perms

import (
	"os"
)

// WriteFile writes [data] to [filename] and ensures that [filename] has [perm]
// permissions. Will write atomically on linux/macos and fall back to non-atomic
// ioutil.WriteFile on windows.
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

// The file doesn't exist, so try to write it.

// The file currently has the wrong permissions, so update them.

// The file has the right permissions, so truncate any data and write the
// file.
