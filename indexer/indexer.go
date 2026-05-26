// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package indexer

import (
	"io"
	"sync"

	"github.com/ava-labs/avalanchego/api/server"
	"github.com/ava-labs/avalanchego/chains"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/utils/logging"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
)

const (
	indexNamePrefix         = "index-"
	txPrefix                = 0x01
	vtxPrefix               = 0x02
	blockPrefix             = 0x03
	isIncompletePrefix      = 0x04
	previouslyIndexedPrefix = 0x05
)

var (
	_ Indexer = (*indexer)(nil)

	hasRunKey = []byte{0x07}
)

// Config for an indexer
type Config struct {
	DB                   database.Database
	Log                  logging.Logger
	IndexingEnabled      bool
	AllowIncompleteIndex bool
	BlockAcceptorGroup   snow.AcceptorGroup
	TxAcceptorGroup      snow.AcceptorGroup
	VertexAcceptorGroup  snow.AcceptorGroup
	APIServer            server.PathAdder
	ShutdownF            func()
}

// Indexer causes accepted containers for a given chain
// to be indexed by their ID and by the order in which
// they were accepted by this node.
// Indexer is threadsafe.
type Indexer interface {
	chains.Registrant
	// Close will do nothing and return nil after the first call
	io.Closer
}

// NewIndexer returns a new Indexer and registers a new endpoint on the given API server.
func NewIndexer(config Config) (Indexer, error) {
	_ = "STUB: not implemented"
	return *new(Indexer), nil
}

type indexer struct {
	clock  mockable.Clock
	lock   sync.RWMutex
	log    logging.Logger
	db     database.Database
	closed bool

	// Called in a goroutine on shutdown
	shutdownF func()

	// true if this is not the first run using this database
	hasRunBefore bool

	// Used to add API endpoint for new indices
	pathAdder server.PathAdder

	// If true, allow running in such a way that could allow the creation
	// of an index which could be missing accepted containers.
	allowIncompleteIndex bool

	// If false, don't create index for a chain when RegisterChain is called
	indexingEnabled bool

	// Chain ID --> index of blocks of that chain (if applicable)
	blockIndices map[ids.ID]*index
	// Chain ID --> index of vertices of that chain (if applicable)
	vtxIndices map[ids.ID]*index
	// Chain ID --> index of txs of that chain (if applicable)
	txIndices map[ids.ID]*index

	// Notifies of newly accepted blocks
	blockAcceptorGroup snow.AcceptorGroup
	// Notifies of newly accepted transactions
	txAcceptorGroup snow.AcceptorGroup
	// Notifies of newly accepted vertices
	vertexAcceptorGroup snow.AcceptorGroup
}

// Assumes [ctx.Lock] is not held
func (i *indexer) RegisterChain(chainName string, ctx *snow.ConsensusContext, vm common.VM) {
	_ = "STUB: not implemented"
	return
}

// If the index is incomplete, make sure that's OK. Otherwise, cause node to die.

// See if this chain was indexed in a previous run

// Indexing is disabled

// We indexed this chain in a previous run but not in this run.
// This would create an incomplete index, which is not allowed, so exit.

// Creating an incomplete index is allowed. Mark index as incomplete.

// Mark that in this run, this chain was indexed

func (i *indexer) registerChainHelper(
	chainID ids.ID,
	prefixEnd byte,
	name, endpoint string,
	acceptorGroup snow.AcceptorGroup,
) (*index, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Register index to learn about new accepted vertices

// Create an API endpoint for this index

// Close this indexer. Stops indexing all chains.
// Closes [i.db]. Assumes Close is only called after
// the node is done making decisions.
// Calling Close after it has been called does nothing.
func (i *indexer) Close() error { _ = "STUB: not implemented"; return nil }

func (i *indexer) close() error { _ = "STUB: not implemented"; return nil }

func (i *indexer) markIncomplete(chainID ids.ID) error { _ = "STUB: not implemented"; return nil }

// Returns true if this chain is incomplete
func (i *indexer) isIncomplete(chainID ids.ID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (i *indexer) markPreviouslyIndexed(chainID ids.ID) error {
	_ = "STUB: not implemented"
	return nil
}

// Returns true if this chain is incomplete
func (i *indexer) previouslyIndexed(chainID ids.ID) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Mark that the node has run at least once
func (i *indexer) markHasRun() error { _ = "STUB: not implemented"; return nil }

// Returns true if the node has run before
func (i *indexer) hasRun() (bool, error) { _ = "STUB: not implemented"; return false, nil }
