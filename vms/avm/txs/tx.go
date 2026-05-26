// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p/gossip"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/avm/fxs"
	"github.com/ava-labs/avalanchego/vms/components/avax"
)

var _ gossip.Gossipable = (*Tx)(nil)

type UnsignedTx interface {
	snow.ContextInitializable

	SetBytes(unsignedBytes []byte)
	Bytes() []byte

	InputIDs() set.Set[ids.ID]

	NumCredentials() int
	// TODO: deprecate after x-chain linearization
	InputUTXOs() []*avax.UTXOID

	// Visit calls [visitor] with this transaction's concrete type
	Visit(visitor Visitor) error
}

// Tx is the core operation that can be performed. The tx uses the UTXO model.
// Specifically, a txs inputs will consume previous txs outputs. A tx will be
// valid if the inputs have the authority to consume the outputs they are
// attempting to consume and the inputs consume sufficient state to produce the
// outputs.
type Tx struct {
	Unsigned UnsignedTx          `serialize:"true" json:"unsignedTx"`
	Creds    []*fxs.FxCredential `serialize:"true" json:"credentials"` // The credentials of this transaction

	TxID  ids.ID `json:"id"`
	bytes []byte
}

func (t *Tx) Initialize(c codec.Manager) error { _ = "STUB: not implemented"; return nil }

func (t *Tx) SetBytes(unsignedBytes, signedBytes []byte) { _ = "STUB: not implemented"; return }

// ID returns the unique ID of this tx
func (t *Tx) ID() ids.ID {
	_ = "STUB: not implemented"

	// GossipID returns the unique ID that this tx should use for mempool gossip
	return *new(ids.ID)
}

func (t *Tx) GossipID() ids.ID {
	_ = "STUB: not implemented"

	// Bytes returns the binary representation of this tx
	return *new(ids.ID)
}

func (t *Tx) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (t *Tx) Size() int { _ = "STUB: not implemented"; return 0 }

// UTXOs returns the UTXOs transaction is producing.
func (t *Tx) UTXOs() []*avax.UTXO { _ = "STUB: not implemented"; return nil }

// The visit error is explicitly dropped here because no error is ever
// returned from the utxoGetter.

func (t *Tx) InputIDs() set.Set[ids.ID] { _ = "STUB: not implemented"; return nil }

func (t *Tx) SignSECP256K1Fx(c codec.Manager, signers [][]*secp256k1.PrivateKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tx) SignPropertyFx(c codec.Manager, signers [][]*secp256k1.PrivateKey) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Tx) SignNFTFx(c codec.Manager, signers [][]*secp256k1.PrivateKey) error {
	_ = "STUB: not implemented"
	return nil
}
