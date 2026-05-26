// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nat

import (
	"net/netip"
	"sync"
	"time"

	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/logging"
)

const (
	mapTimeout        = 30 * time.Minute
	maxRefreshRetries = 3
)

// Router describes the functionality that a network device must support to be
// able to open ports to an external IP.
type Router interface {
	// True iff this router supports NAT
	SupportsNAT() bool
	// Map external port [extPort] to internal port [intPort] for [duration]
	MapPort(intPort, extPort uint16, desc string, duration time.Duration) error
	// Undo a port mapping
	UnmapPort(intPort, extPort uint16) error
	// Return our external IP
	ExternalIP() (netip.Addr, error)
}

// GetRouter returns a router on the current network.
func GetRouter() Router { _ = "STUB: not implemented"; return *new(Router) }

// Mapper attempts to open a set of ports on a router
type Mapper struct {
	log    logging.Logger
	r      Router
	closer chan struct{}
	wg     sync.WaitGroup
}

// NewPortMapper returns an initialized mapper
func NewPortMapper(log logging.Logger, r Router) *Mapper { _ = "STUB: not implemented"; return nil }

// Map external port [extPort] (exposed to the internet) to internal port [intPort] (where our process is listening)
// and set [ip]. Does this every [updateTime]. [ip] may be nil.
func (m *Mapper) Map(
	intPort uint16,
	extPort uint16,
	desc string,
	ip *utils.Atomic[netip.AddrPort],
	updateTime time.Duration,
) {
	_ = "STUB: not implemented"
	return
}

// we attempt a port map, and log an Error if it fails.

// Retry port map up to maxRefreshRetries with a 1 second delay
func (m *Mapper) retryMapPort(intPort, extPort uint16, desc string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// log a message, sleep a second and retry.

// keepPortMapping runs in the background to keep a port mapped. It renews the mapping from [extPort]
// to [intPort]] every [updateTime]. Updates [ip] every [updateTime].
func (m *Mapper) keepPortMapping(
	intPort uint16,
	extPort uint16,
	desc string,
	ip *utils.Atomic[netip.AddrPort],
	updateTime time.Duration,
) {
	_ = "STUB: not implemented"
	return
}

func (m *Mapper) updateIP(ip *utils.Atomic[netip.AddrPort]) { _ = "STUB: not implemented"; return }

// UnmapAllPorts stops mapping all ports from this mapper and attempts to unmap
// them.
func (m *Mapper) UnmapAllPorts() { _ = "STUB: not implemented"; return }
