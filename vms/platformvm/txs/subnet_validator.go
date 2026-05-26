// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"github.com/ava-labs/avalanchego/ids"
)

// SubnetValidator validates a subnet on the Avalanche network.
type SubnetValidator struct {
	Validator `serialize:"true"`

	// ID of the subnet this validator is validating
	Subnet ids.ID `serialize:"true" json:"subnetID"`
}

// SubnetID is the ID of the subnet this validator is validating
func (v *SubnetValidator) SubnetID() ids.ID {
	_ = "STUB: not implemented"

	// Verify this validator is valid
	return *new(ids.ID)
}

func (v *SubnetValidator) Verify() error { _ = "STUB: not implemented"; return nil }
