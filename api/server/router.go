// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package server

import (
	"errors"
	"net/http"
	"sync"

	"github.com/gorilla/mux"

	"github.com/ava-labs/avalanchego/utils/set"
)

const HTTPHeaderRoute = "Avalanche-Api-Route"

var (
	errUnknownBaseURL  = errors.New("unknown base url")
	errUnknownEndpoint = errors.New("unknown endpoint")
	errAlreadyReserved = errors.New("route is either already aliased or already maps to a handle")
)

type router struct {
	lock   sync.RWMutex
	router *mux.Router

	routeLock      sync.Mutex
	reservedRoutes set.Set[string]     // Reserves routes so that there can't be alias that conflict
	aliases        map[string][]string // Maps a route to a set of reserved routes
	// headerRoutes contains routes based on http headers
	// aliasing is not currently supported
	headerRoutes map[string]http.Handler
	// legacy url-based routing
	routes map[string]map[string]http.Handler // Maps routes to a handler
}

func newRouter() *router { _ = "STUB: not implemented"; return nil }

func (r *router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_ = "STUB: not implemented"
	return
}

// If there is no routing header, fall-back to the legacy path-based
// routing

// Request specified the routing header key but did not provide a
// corresponding value

func (r *router) GetHandler(base, endpoint string) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func (r *router) AddHeaderRoute(route string, handler http.Handler) bool {
	_ = "STUB: not implemented"
	return false
}

func (r *router) AddRouter(base, endpoint string, handler http.Handler) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *router) addRouter(base, endpoint string, handler http.Handler) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *router) forceAddRouter(base, endpoint string, handler http.Handler) error {
	_ = "STUB: not implemented"
	return nil
}

// Name routes based on their URL for easy retrieval in the future

func (r *router) AddAlias(base string, aliases ...string) error {
	_ = "STUB: not implemented"
	return nil
}
