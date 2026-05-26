// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"crypto/tls"
	"errors"
	"net"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/staking"
)

var (
	errNoCert = errors.New("tls handshake finished with no peer certificate")

	_ Upgrader = (*tlsServerUpgrader)(nil)
	_ Upgrader = (*tlsClientUpgrader)(nil)
)

type Upgrader interface {
	// Must be thread safe
	Upgrade(net.Conn) (ids.NodeID, net.Conn, *staking.Certificate, error)
}

type tlsServerUpgrader struct {
	config       *tls.Config
	invalidCerts prometheus.Counter
}

func NewTLSServerUpgrader(config *tls.Config, invalidCerts prometheus.Counter) Upgrader {
	_ = "STUB: not implemented"
	return *new(Upgrader)
}

func (t *tlsServerUpgrader) Upgrade(conn net.Conn) (ids.NodeID, net.Conn, *staking.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(ids.NodeID), *new(net.Conn), nil, nil
}

type tlsClientUpgrader struct {
	config       *tls.Config
	invalidCerts prometheus.Counter
}

func NewTLSClientUpgrader(config *tls.Config, invalidCerts prometheus.Counter) Upgrader {
	_ = "STUB: not implemented"
	return *new(Upgrader)
}

func (t *tlsClientUpgrader) Upgrade(conn net.Conn) (ids.NodeID, net.Conn, *staking.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(ids.NodeID), *new(net.Conn), nil, nil
}

func connToIDAndCert(conn *tls.Conn, invalidCerts prometheus.Counter) (ids.NodeID, net.Conn, *staking.Certificate, error) {
	_ = "STUB: not implemented"
	return *new(ids.NodeID), *new(net.Conn), nil, nil
}
