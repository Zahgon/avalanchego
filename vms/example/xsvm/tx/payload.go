// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package tx

import "github.com/ava-labs/avalanchego/ids"

type Payload struct {
	// Sender + Nonce provides replay protection
	Sender   ids.ShortID `serialize:"true" json:"sender"`
	Nonce    uint64      `serialize:"true" json:"nonce"`
	IsReturn bool        `serialize:"true" json:"isReturn"`
	Amount   uint64      `serialize:"true" json:"amount"`
	To       ids.ShortID `serialize:"true" json:"to"`

	bytes []byte
}

func (p *Payload) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func NewPayload(
	sender ids.ShortID,
	nonce uint64,
	isReturn bool,
	amount uint64,
	to ids.ShortID,
) (*Payload, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParsePayload(bytes []byte) (*Payload, error) { _ = "STUB: not implemented"; return nil, nil }
