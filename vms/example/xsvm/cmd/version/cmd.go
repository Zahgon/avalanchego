// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"github.com/spf13/cobra"
)

const format = `%s:
  VMID:           %s
  Version:        %s
  Plugin Version: %d
`

func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

func versionFunc(*cobra.Command, []string) error { _ = "STUB: not implemented"; return nil }
