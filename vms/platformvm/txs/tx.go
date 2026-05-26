// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"errors"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p/gossip"
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/utils/crypto/secp256k1"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/components/verify"
)

var (
	_ gossip.Gossipable = (*Tx)(nil)

	ErrNilSignedTx = errors.New("nil signed tx is not valid")

	errSignedTxNotInitialized = errors.New("signed tx was never initialized and is not valid")
)

// Tx is a signed transaction
type Tx struct {
	// The body of this transaction
	Unsigned UnsignedTx `serialize:"true" json:"unsignedTx"`

	// The credentials of this transaction
	Creds []verify.Verifiable `serialize:"true" json:"credentials"`

	TxID  ids.ID `json:"id"`
	bytes []byte
}

func NewSigned(
	unsigned UnsignedTx,
	c codec.Manager,
	signers [][]*secp256k1.PrivateKey,
) (*Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Tx) Initialize(c codec.Manager) error { _ = "STUB: not implemented"; return nil }

func (tx *Tx) SetBytes(unsignedBytes, signedBytes []byte) { _ = "STUB: not implemented"; return }

// Parse signed tx starting from its byte representation.
// Note: We explicitly pass the codec in Parse since we may need to parse
// P-Chain genesis txs whose length exceed the max length of txs.Codec.
func Parse(c codec.Manager, signedBytes []byte) (*Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Tx) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (tx *Tx) Size() int { _ = "STUB: not implemented"; return 0 }

func (tx *Tx) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (tx *Tx) GossipID() ids.ID {
	_ = "STUB: not implemented"

	// UTXOs returns the UTXOs transaction is producing.
	return *new(ids.ID)
}

func (tx *Tx) UTXOs() []*avax.UTXO { _ = "STUB: not implemented"; return nil }

// InputIDs returns the set of inputs this transaction consumes
func (tx *Tx) InputIDs() set.Set[ids.ID] { _ = "STUB: not implemented"; return nil }

func (tx *Tx) SyntacticVerify(ctx *snow.Context) error { _ = "STUB: not implemented"; return nil }

// Sign this transaction with the provided signers
// Note: We explicitly pass the codec in Sign since we may need to sign P-Chain
// genesis txs whose length exceed the max length of txs.Codec.
func (tx *Tx) Sign(c codec.Manager, signers [][]*secp256k1.PrivateKey) error {
	_ = "STUB: not implemented"
	return nil
}

// Attach credentials

// Sign hash

// Attach credential
