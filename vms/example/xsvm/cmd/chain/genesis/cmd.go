// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"errors"

	"github.com/spf13/cobra"
)

var errUnknownEncoding = errors.New("unknown encoding")

func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

func genesisFunc(c *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }
