// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package extras

import (
	"github.com/ava-labs/avalanchego/graft/subnet-evm/precompile/precompileconfig"
)

type Precompiles map[string]precompileconfig.Config

// UnmarshalJSON parses the JSON-encoded data into the ChainConfigPrecompiles.
// ChainConfigPrecompiles is a map of precompile module keys to their
// configuration.
func (ccp *Precompiles) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
