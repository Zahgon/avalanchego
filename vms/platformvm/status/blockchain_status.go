// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package status

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ava-labs/avalanchego/vms/components/verify"
)

// List of possible blockchain status values:
// - [UnknownChain] This node is not aware of the existence of this blockchain
// - [Created] This node is not currently validating this blockchain
// - [Preferred] This blockchain is currently in the preferred tip
// - [Validating] This node is currently validating this blockchain
// - [Syncing] This node is syncing up to the preferred block height
const (
	UnknownChain BlockchainStatus = iota
	Created
	Preferred
	Validating
	Syncing
)

var (
	errUnknownBlockchainStatus = errors.New("unknown blockchain status")

	_ json.Marshaler    = BlockchainStatus(0)
	_ verify.Verifiable = BlockchainStatus(0)
	_ fmt.Stringer      = BlockchainStatus(0)
)

type BlockchainStatus uint32

func (s BlockchainStatus) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *BlockchainStatus) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Verify that this is a valid status.
func (s BlockchainStatus) Verify() error { _ = "STUB: not implemented"; return nil }

func (s BlockchainStatus) String() string { _ = "STUB: not implemented"; return "" }
