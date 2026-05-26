// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txgossip

import (
	"context"

	"github.com/ava-labs/libevm/core/types"

	"github.com/ava-labs/avalanchego/network/p2p/gossip"
)

// SendTx implements the respective method of [ethapi.Backend], accepting
// transactions submitted via the `eth_sendTransaction` RPC method. Unlike
// [gossip.BloomSet.Add], transactions added via this method will not be added
// to the Bloom filter, but will be push-gossiped; see
// [Set.RegisterPushGossiper].
func (s *Set) SendTx(ctx context.Context, ethTx *types.Transaction) error {
	_ = "STUB: not implemented"
	// protect the import for [comment] rendering
	return nil
}

// Note: the Bloom filter is only necessary for pull gossip, which is only
// done between validator nodes. Validators are encouraged to NOT expose RPC
// APIs (i.e. this method) and SHOULD therefore only add transactions via
// [gossip.BloomSet.Add]. Hence there is a MECE separation between
// pull+Bloom vs API+push under recommended usage patterns.

/*local*/

// RegisterPushGossiper registers [gossip.PushGossiper.Add] as a callback for
// every transaction received via [Set.SendTx].
//
// NOTE: it is not safe to call this method concurrently with [Set.SendTx];
// registration is expected to occur immediately after construction of the
// [Set].
func (s *Set) RegisterPushGossiper(push *gossip.PushGossiper[Transaction]) {
	_ = "STUB: not implemented"
	return
}
