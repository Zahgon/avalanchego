// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package simplex

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/ava-labs/simplex"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/proto/pb/p2p"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/snow/engine/snowman/block"
	"github.com/ava-labs/avalanchego/snow/validators"
	"github.com/ava-labs/avalanchego/utils/logging"

	simplexparams "github.com/ava-labs/avalanchego/snow/consensus/simplex"
)

var _ common.Engine = (*Engine)(nil)

var (
	errUnknownMessageType   = errors.New("unknown message type")
	errNilSimplexParameters = errors.New("simplex parameters cannot be nil")
)

type Engine struct {
	// list of NoOpsHandler for messages dropped by engine
	common.AllGetsServer
	common.StateSummaryFrontierHandler
	common.AcceptedStateSummaryHandler
	common.AcceptedFrontierHandler
	common.AcceptedHandler
	common.AncestorsHandler
	common.PutHandler
	common.QueryHandler
	common.ChitsHandler

	// Handler that passes application messages to the VM
	common.AppHandler
	validators.Connector
	vm block.ChainVM

	epoch              *simplex.Epoch
	blockDeserializer  *blockDeserializer
	quorumDeserializer *QCDeserializer
	logger             logging.Logger

	tickInterval time.Duration
	shutdown     chan struct{}
	shutdownOnce sync.Once
}

// NewEngine creates a new simplex engine. The VM must be initialized before
// calling this function.
func NewEngine(ctx context.Context, config *Config) (*Engine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newEngineWithSignerVerifier(ctx context.Context, config *Config, signer BLSSigner, verifier BLSVerifier) (*Engine, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Initialize the blockTracker with the last block fetched from Storage.

func (e *Engine) Start(_ context.Context, _ uint32) error { _ = "STUB: not implemented"; return nil }

// getTickInterval defines a reasonable tick interval for simplex to advance time.
func getTickInterval(params *simplexparams.Parameters) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// tick periodically advances the engine's time.
func (e *Engine) tick() { _ = "STUB: not implemented"; return }

func (e *Engine) Simplex(ctx context.Context, nodeID ids.NodeID, msg *p2p.Simplex) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) p2pToSimplexMessage(ctx context.Context, msg *p2p.Simplex) (*simplex.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Gossip is a no-op because there is no need for the Simplex engine
// to periodically pull/push messages from the network.
// This is all handled internally inside of Simplex consensus.
func (*Engine) Gossip(_ context.Context) error {
	_ = "STUB: not implemented"

	// Notify is a no-op because the Simplex engine does not need to be notified of any events from the VM.
	// This is because the Simplex instance listens to the VM events by directly calling `WaitForEvent` when needed.
	return nil
}

func (*Engine) Notify(_ context.Context, _ common.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *Engine) HealthCheck(ctx context.Context) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *Engine) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }
