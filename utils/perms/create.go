// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package perms

import (
	"os"
)

// Create a file at [filename] that has [perm] permissions.
func Create(filename string, perm os.FileMode) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The file currently has the wrong permissions, so update them.
