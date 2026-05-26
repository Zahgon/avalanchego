// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nat

import (
	"net"
	"net/netip"
	"time"

	"github.com/huin/goupnp"
)

const (
	// upnpProtocol is intentionally uppercase and should not be confused with
	// pmpProtocol.
	// See:
	// - https://github.com/huin/goupnp/blob/v1.0.3/dcps/internetgateway1/internetgateway1.go#L2361
	// - https://github.com/huin/goupnp/blob/v1.0.3/dcps/internetgateway1/internetgateway1.go#L3618
	// - https://github.com/huin/goupnp/blob/v1.0.3/dcps/internetgateway2/internetgateway2.go#L3919
	upnpProtocol       = "TCP"
	soapRequestTimeout = 10 * time.Second
)

var _ Router = (*upnpRouter)(nil)

// upnpClient is the interface used by goupnp for their client implementations
type upnpClient interface {
	// attempts to map connection using the provided protocol from the external
	// port to the internal port for the lease duration.
	AddPortMapping(
		newRemoteHost string,
		newExternalPort uint16,
		newProtocol string,
		newInternalPort uint16,
		newInternalClient string,
		newEnabled bool,
		newPortMappingDescription string,
		newLeaseDuration uint32) error

	// attempt to remove any mapping from the external port.
	DeletePortMapping(
		newRemoteHost string,
		newExternalPort uint16,
		newProtocol string) error

	// attempts to return the external IP address, formatted as a string.
	GetExternalIPAddress() (ip string, err error)

	// returns if there is rsip available, nat enabled, or an unexpected error.
	GetNATRSIPStatus() (newRSIPAvailable bool, natEnabled bool, err error)

	// attempts to get port mapping information give a external port and protocol
	GetSpecificPortMappingEntry(
		NewRemoteHost string,
		NewExternalPort uint16,
		NewProtocol string,
	) (
		NewInternalPort uint16,
		NewInternalClient string,
		NewEnabled bool,
		NewPortMappingDescription string,
		NewLeaseDuration uint32,
		err error,
	)
}

type upnpRouter struct {
	dev    *goupnp.RootDevice
	client upnpClient
}

func (*upnpRouter) SupportsNAT() bool { _ = "STUB: not implemented"; return false }

func (r *upnpRouter) localIP() (net.IP, error) {
	_ = "STUB: not implemented"
	// attempt to get an address on the router
	return *new(net.IP), nil
}

// attempt to find one of my IPs that matches router's record

func (r *upnpRouter) ExternalIP() (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

func (r *upnpRouter) MapPort(
	intPort,
	extPort uint16,
	desc string,
	duration time.Duration,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *upnpRouter) UnmapPort(_, extPort uint16) error { _ = "STUB: not implemented"; return nil }

// create UPnP SOAP service client with URN
func getUPnPClient(client goupnp.ServiceClient) upnpClient {
	_ = "STUB: not implemented"
	return *new(upnpClient)
}

// discover() tries to find gateway device
func discover(target string) *upnpRouter { _ = "STUB: not implemented"; return nil }

// getUPnPRouter searches for internet gateway using both Device Control Protocol
// and returns the first one it can find. It returns nil if no UPnP gateway is found
func getUPnPRouter() *upnpRouter { _ = "STUB: not implemented"; return nil }
