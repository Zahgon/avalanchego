// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package dialer

import (
	"context"
	"net"
	"net/netip"
	"time"

	"github.com/ava-labs/avalanchego/network/throttling"
	"github.com/ava-labs/avalanchego/utils/logging"
)

var _ Dialer = (*dialer)(nil)

// Dialer attempts to create a connection with the provided IP/port pair
type Dialer interface {
	// If [ctx] is canceled, gives up trying to connect to [ip]
	// and returns an error.
	Dial(ctx context.Context, ip netip.AddrPort) (net.Conn, error)
}

type dialer struct {
	dialer    net.Dialer
	log       logging.Logger
	network   string
	throttler throttling.DialThrottler
}

type Config struct {
	ThrottleRps       uint32        `json:"throttleRps"`
	ConnectionTimeout time.Duration `json:"connectionTimeout"`
}

// NewDialer returns a new Dialer that calls net.Dial with the provided network.
// [network] is the network passed into Dial. Should probably be "TCP".
// [dialerConfig.connectionTimeout] gives the timeout when dialing an IP.
// [dialerConfig.throttleRps] gives the max number of outgoing connection attempts/second.
// If [dialerConfig.throttleRps] == 0, outgoing connections aren't rate-limited.
func NewDialer(network string, dialerConfig Config, log logging.Logger) Dialer {
	_ = "STUB: not implemented"
	return *new(Dialer)
}

func (d *dialer) Dial(ctx context.Context, ip netip.AddrPort) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
