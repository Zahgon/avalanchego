// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tmpnet

import (
	"context"
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/api/health"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
)

const (
	DefaultNodeTickerInterval = 50 * time.Millisecond
)

var ErrUnrecoverableNodeHealthCheck = errors.New("failed to query node health")

func CheckNodeHealth(ctx context.Context, uri string) (*health.APIReply, error) {
	_ = "STUB: not implemented"
	// Check that the node is reporting healthy
	return nil, nil
}

// Connection refused - potentially recoverable

// Connection refused - potentially recoverable

// Assume `503 Service Unavailable` is the result of the ingress
// for the node not being ready.
// TODO(marun) Update Client.Health() to return a typed error

// Assume all other errors are not recoverable

// NodeURI associates a node ID with its API URI.
type NodeURI struct {
	NodeID ids.NodeID
	URI    string
}

// GetNodeURIs returns the accessible URIs of the provided nodes that are running and not ephemeral.
func GetNodeURIs(nodes []*Node) []NodeURI { _ = "STUB: not implemented"; return nil }

// FilterAvailableNodes filters the provided nodes by whether they are running and not ephemeral.
func FilterAvailableNodes(nodes []*Node) []*Node { _ = "STUB: not implemented"; return nil }

// Avoid returning URIs for nodes whose lifespan is indeterminate

// Only running nodes have URIs

// GetNodeWebsocketURIs returns a list of websocket URIs for the given nodes and
// blockchain ID, in the form "ws://<node-uri>/ext/bc/<blockchain-id>/ws".
// Ephemeral and stopped nodes are ignored.
func GetNodeWebsocketURIs(
	nodes []*Node,
	blockchainID string,
) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use websocket to be able to stream events

// Marshal to json with default prefix and indent.
func DefaultJSONMarshal(v interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Helper simplifying creation of a set of private keys
func NewPrivateKeys(keyCount int) ([]*secp256k1.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NodesToIDs(nodes ...*Node) []ids.NodeID { _ = "STUB: not implemented"; return nil }

func GetEnvWithDefault(envVar, defaultVal string) string { _ = "STUB: not implemented"; return "" }
