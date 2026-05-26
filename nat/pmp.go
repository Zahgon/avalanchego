// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nat

import (
	"errors"
	"net/netip"
	"time"

	natpmp "github.com/jackpal/go-nat-pmp"
)

const (
	// pmpProtocol is intentionally lowercase and should not be confused with
	// upnpProtocol.
	// See:
	// - https://github.com/jackpal/go-nat-pmp/blob/v1.0.2/natpmp.go#L82
	pmpProtocol      = "tcp"
	pmpClientTimeout = 500 * time.Millisecond
)

var (
	_ Router = (*pmpRouter)(nil)

	errInvalidLifetime = errors.New("invalid mapping duration range")
)

// pmpRouter adapts the NAT-PMP protocol implementation so it conforms to the
// common interface.
type pmpRouter struct {
	client *natpmp.Client
}

func (*pmpRouter) SupportsNAT() bool { _ = "STUB: not implemented"; return false }

func (r *pmpRouter) MapPort(
	newInternalPort uint16,
	newExternalPort uint16,
	_ string,
	mappingDuration time.Duration,
) error {
	_ = "STUB: not implemented"
	return nil
}

// go-nat-pmp uses seconds to denote their lifetime

// Assumes the architecture is at least 32-bit

func (r *pmpRouter) UnmapPort(internalPort uint16, _ uint16) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *pmpRouter) ExternalIP() (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

func getPMPRouter() *pmpRouter { _ = "STUB: not implemented"; return nil }
