// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package extrastest

import (
	"github.com/ava-labs/avalanchego/graft/coreth/params/extras"
	"github.com/ava-labs/avalanchego/upgrade/upgradetest"
)

func ForkToRules(fork upgradetest.Fork) *extras.Rules { _ = "STUB: not implemented"; return nil }

func ForkToAvalancheRules(fork upgradetest.Fork) extras.AvalancheRules {
	_ = "STUB: not implemented"
	return *new(extras.AvalancheRules)
}
