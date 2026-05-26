// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package importtx

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/ava-labs/avalanchego/vms/example/xsvm/cmd/issue/status"
)

func Command() *cobra.Command { _ = "STUB: not implemented"; return nil }

func importFunc(c *cobra.Command, args []string) error { _ = "STUB: not implemented"; return nil }

func Import(ctx context.Context, config *Config) (*status.TxIssuance, error) {
	_ = "STUB: not implemented"

	// Note: here we assume the unsigned message is correct from the last
	//       URI in sourceURIs. In practice this shouldn't be done.
	return nil, nil
}

// Note: assumes that sourceURIs are all of the validators of the subnet
//       and that they do not share public keys.

// Note: the public key should not be fetched from the node in practice.
//       The public key should be fetched from the P-chain directly.
