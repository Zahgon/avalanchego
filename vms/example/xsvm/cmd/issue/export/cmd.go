// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package export

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/ava-labs/avalanchego/vms/example/xsvm/cmd/issue/status"
)

func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

func exportFunc(c *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func Export(ctx context.Context, config *Config) (*status.TxIssuance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
