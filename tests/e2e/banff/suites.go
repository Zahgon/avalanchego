// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Implements tests for the banff network upgrade.
package banff

import (
	"github.com/onsi/ginkgo/v2"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/tests"
	"github.com/ava-labs/avalanchego/tests/fixture/e2e"
	"github.com/ava-labs/avalanchego/wallet/subnet/primary"
)

var _ = ginkgo.Describe("[Banff]", func() {
	ginkgo.It("can send custom assets X->P and P->X", func() {
		e2e.ExecuteAPITest(TestCustomAssetTransfer)
	})
})

func TestCustomAssetTransfer(
	tc tests.TestContext,
	wallet primary.Wallet,
	ownerAddress ids.ShortID,
) {
	_ = "STUB: not implemented"
	return
}

// Get the P-chain and the X-chain wallets

// Pull out useful constants to use when issuing transactions.
