// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package perms

import (
	"os"
)

// ChmodR sets the permissions of all directories and optionally files to [perm]
// permissions.
func ChmodR(dir string, dirOnly bool, perm os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}
