// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//go:build linux || darwin || unix

package tmpnet

import (
	"os/exec"
)

func configureDetachedProcess(cmd *exec.Cmd) { _ = "STUB: not implemented"; return }
