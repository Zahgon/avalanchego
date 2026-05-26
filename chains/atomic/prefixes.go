// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomic

import (
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
)

var (
	inboundSmallerValuePrefix = []byte{0}
	inboundSmallerIndexPrefix = []byte{1}
	inboundLargerValuePrefix  = []byte{2}
	inboundLargerIndexPrefix  = []byte{3}

	// note that inbound and outbound have their smaller and larger values
	// swapped

	// inbound specifies the prefixes to use for inbound shared memory
	// ie. reading and deleting a message received from another chain.
	inbound = prefixes{
		smallerValuePrefix: inboundSmallerValuePrefix,
		smallerIndexPrefix: inboundSmallerIndexPrefix,
		largerValuePrefix:  inboundLargerValuePrefix,
		largerIndexPrefix:  inboundLargerIndexPrefix,
	}
	// outbound specifies the prefixes to use for outbound shared memory
	// ie. writing a message to another chain.
	outbound = prefixes{
		smallerValuePrefix: inboundLargerValuePrefix,
		smallerIndexPrefix: inboundLargerIndexPrefix,
		largerValuePrefix:  inboundSmallerValuePrefix,
		largerIndexPrefix:  inboundSmallerIndexPrefix,
	}
)

type prefixes struct {
	smallerValuePrefix []byte
	smallerIndexPrefix []byte
	largerValuePrefix  []byte
	largerIndexPrefix  []byte
}

func (p *prefixes) getValueDB(myChainID, peerChainID ids.ID, db database.Database) database.Database {
	_ = "STUB: not implemented"
	return *new(database.Database)
}

func (p *prefixes) getValueAndIndexDB(myChainID, peerChainID ids.ID, db database.Database) (database.Database, database.Database) {
	_ = "STUB: not implemented"
	return *new(database.Database), *new(database.Database)
}
