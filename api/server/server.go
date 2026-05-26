// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package server

import (
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/snow/engine/common"
	"github.com/ava-labs/avalanchego/trace"
	"github.com/ava-labs/avalanchego/utils/logging"
)

const (
	baseURL              = "/ext"
	maxConcurrentStreams = 64
)

var (
	_ PathAdder = readPathAdder{}
	_ Server    = (*server)(nil)
)

type PathAdder interface {
	// AddRoute registers a route to a handler.
	AddRoute(handler http.Handler, base, endpoint string) error

	// AddAliases registers aliases to the server
	AddAliases(endpoint string, aliases ...string) error
}

type PathAdderWithReadLock interface {
	// AddRouteWithReadLock registers a route to a handler assuming the http
	// read lock is currently held.
	AddRouteWithReadLock(handler http.Handler, base, endpoint string) error

	// AddAliasesWithReadLock registers aliases to the server assuming the http read
	// lock is currently held.
	AddAliasesWithReadLock(endpoint string, aliases ...string) error
}

// Server maintains the HTTP router
type Server interface {
	PathAdder
	PathAdderWithReadLock
	// Dispatch starts the API server
	Dispatch() error
	// RegisterChain registers the API endpoints associated with this chain.
	// That is, add <route, handler> pairs to server so that API calls can be
	// made to the VM.
	RegisterChain(chainName string, ctx *snow.ConsensusContext, vm common.VM)
	// Shutdown this server
	Shutdown() error
}

type HTTPConfig struct {
	ReadTimeout       time.Duration `json:"readTimeout"`
	ReadHeaderTimeout time.Duration `json:"readHeaderTimeout"`
	WriteTimeout      time.Duration `json:"writeHeaderTimeout"`
	IdleTimeout       time.Duration `json:"idleTimeout"`
}

type server struct {
	// log this server writes to
	log logging.Logger

	shutdownTimeout time.Duration

	tracingEnabled bool
	tracer         trace.Tracer

	metrics *metrics

	// Maps endpoints to handlers
	router *router

	srv *http.Server

	// Listener used to serve traffic
	listener net.Listener
}

// New returns an instance of a Server.
func New(
	log logging.Logger,
	listener net.Listener,
	allowedOrigins []string,
	shutdownTimeout time.Duration,
	nodeID ids.NodeID,
	tracingEnabled bool,
	tracer trace.Tracer,
	registerer prometheus.Registerer,
	httpConfig HTTPConfig,
	allowedHosts []string,
) (Server, error) {
	_ = "STUB: not implemented"
	return *new(Server), nil
}

func (s *server) Dispatch() error { _ = "STUB: not implemented"; return nil }

func (s *server) RegisterChain(chainName string, ctx *snow.ConsensusContext, vm common.VM) {
	_ = "STUB: not implemented"
	return
}

// all subroutes to a chain begin with "bc/<the chain's ID>"

// Register each endpoint

// Validate that the route being added is valid
// e.g. "/foo" and "" are ok but "\n" is not

func (s *server) addChainRoute(chainName string, handler http.Handler, ctx *snow.ConsensusContext, base, endpoint string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) wrapMiddleware(chainName string, handler http.Handler, ctx *snow.ConsensusContext) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Apply middleware to reject calls to the handler before the chain finishes bootstrapping

func (s *server) AddRoute(handler http.Handler, base, endpoint string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) AddRouteWithReadLock(handler http.Handler, base, endpoint string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) addRoute(handler http.Handler, base, endpoint string) error {
	_ = "STUB: not implemented"
	return nil
}

// Reject middleware wraps a handler. If the chain that the context describes is
// not done state-syncing/bootstrapping, writes back an error.
func rejectMiddleware(handler http.Handler, ctx *snow.ConsensusContext) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// If chain isn't done bootstrapping, ignore API calls

func (s *server) AddAliases(endpoint string, aliases ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *server) AddAliasesWithReadLock(endpoint string, aliases ...string) error {
	_ = "STUB: not implemented"
	// This is safe, as the read lock doesn't actually need to be held once the
	// http handler is called. However, it is unlocked later, so this function
	// must end with the lock held.
	return nil
}

func (s *server) Shutdown() error { _ = "STUB: not implemented"; return nil }

// Close the listener here in case Serve() was never called on the server.

// If shutdown times out, make sure the server is still shutdown.

type readPathAdder struct {
	pather PathAdderWithReadLock
}

func PathWriterFromWithReadLock(pather PathAdderWithReadLock) PathAdder {
	_ = "STUB: not implemented"
	return *new(PathAdder)
}

func (a readPathAdder) AddRoute(handler http.Handler, base, endpoint string) error {
	_ = "STUB: not implemented"
	return nil
}

func (a readPathAdder) AddAliases(endpoint string, aliases ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func wrapHandler(
	handler http.Handler,
	nodeID ids.NodeID,
	allowedOrigins []string,
	allowedHosts []string,
) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// Attach this node's ID as a header
