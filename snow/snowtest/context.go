// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowtest

import (
	"errors"
	"testing"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/constants"
)

var (
	PChainID    = constants.PlatformChainID
	XChainID    = ids.GenerateTestID()
	CChainID    = ids.GenerateTestID()
	AVAXAssetID = ids.GenerateTestID()

	errMissing = errors.New("missing")

	_ snow.Acceptor = noOpAcceptor{}
)

type noOpAcceptor struct{}

func (noOpAcceptor) Accept(*snow.ConsensusContext, ids.ID, []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func ConsensusContext(ctx *snow.Context) *snow.ConsensusContext {
	_ = "STUB: not implemented"
	return nil
}

func Context(tb testing.TB, chainID ids.ID) *snow.Context { _ = "STUB: not implemented"; return nil }
