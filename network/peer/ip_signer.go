// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"crypto"
	"net/netip"
	"sync"

	"github.com/ava-labs/avalanchego/utils"
	"github.com/ava-labs/avalanchego/utils/crypto/bls"
	"github.com/ava-labs/avalanchego/utils/timer/mockable"
)

// IPSigner will return a signedIP for the current value of our dynamic IP.
type IPSigner struct {
	ip        *utils.Atomic[netip.AddrPort]
	clock     mockable.Clock
	tlsSigner crypto.Signer
	blsSigner bls.Signer

	// Must be held while accessing [signedIP]
	signedIPLock sync.RWMutex
	// Note that the values in [*signedIP] are constants and can be inspected
	// without holding [signedIPLock].
	signedIP *SignedIP
}

func NewIPSigner(
	ip *utils.Atomic[netip.AddrPort],
	tlsSigner crypto.Signer,
	blsSigner bls.Signer,
) *IPSigner {
	_ = "STUB: not implemented"
	return nil
}

// GetSignedIP returns the signedIP of the current value of the provided
// dynamicIP. If the dynamicIP hasn't changed since the prior call to
// GetSignedIP, then the same [SignedIP] will be returned.
//
// It's safe for multiple goroutines to concurrently call GetSignedIP.
func (s *IPSigner) GetSignedIP() (*SignedIP, error) {
	_ = "STUB: not implemented"
	// Optimistically, the IP should already be signed. By grabbing a read lock
	// here we enable full concurrency of new connections.
	return nil, nil
}

// If our current IP hasn't been signed yet - then we should sign it.

// It's possible that multiple threads read [n.signedIP] as incorrect at the
// same time, we should verify that we are the first thread to attempt to
// update it.

// We should now sign our new IP at the current timestamp.
