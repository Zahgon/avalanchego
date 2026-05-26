// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package execute

import (
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/example/xsvm/genesis"
)

func Genesis(db database.KeyValueReaderWriterDeleter, chainID ids.ID, g *genesis.Genesis) error {
	_ = "STUB: not implemented"
	return nil
}
