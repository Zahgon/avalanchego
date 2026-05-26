// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package signer

import "github.com/ava-labs/avalanchego/utils/crypto/bls"

var _ Signer = (*Empty)(nil)

type Empty struct{}

func (*Empty) Verify() error { _ = "STUB: not implemented"; return nil }

func (*Empty) Key() *bls.PublicKey { _ = "STUB: not implemented"; return nil }
