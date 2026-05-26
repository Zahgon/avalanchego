// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"errors"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/staking"
)

var (
	_ SignedBlock = (*statelessBlock)(nil)
	_ SignedBlock = (*statelessGraniteBlock)(nil)

	errUnexpectedSignature = errors.New("signature provided when none was expected")
	errInvalidCertificate  = errors.New("invalid certificate")
	errZeroEpoch           = errors.New("epoch must be provided after granite")
)

type Block interface {
	ID() ids.ID
	ParentID() ids.ID
	Block() []byte
	Bytes() []byte

	initialize(bytes []byte) error
	verify(chainID ids.ID) error
}

type SignedBlock interface {
	Block

	PChainHeight() uint64
	PChainEpoch() Epoch
	Timestamp() time.Time

	// Proposer returns the ID of the node that proposed this block. If no node
	// signed this block, [ids.EmptyNodeID] will be returned.
	Proposer() ids.NodeID
}

type statelessUnsignedBlock struct {
	ParentID     ids.ID `serialize:"true" json:"parentID"`
	Timestamp    int64  `serialize:"true" json:"timestamp"`
	PChainHeight uint64 `serialize:"true" json:"pChainHeight"`
	Certificate  []byte `serialize:"true" json:"certificate"`
	Block        []byte `serialize:"true" json:"block"`
}

type statelessUnsignedGraniteBlock struct {
	StatelessBlock statelessUnsignedBlock `serialize:"true" json:"statelessBlock"`
	Epoch          Epoch                  `serialize:"true" json:"epoch"`
}

type Epoch struct {
	PChainHeight uint64 `serialize:"true" json:"pChainHeight"`
	Number       uint64 `serialize:"true" json:"number"`
	StartTime    int64  `serialize:"true" json:"startTime"`
}

type statelessBlockMetadata struct {
	id        ids.ID
	timestamp time.Time
	cert      *staking.Certificate
	proposer  ids.NodeID
	bytes     []byte
}

func (m *statelessBlockMetadata) initialize(
	b *statelessUnsignedBlock,
	sig []byte,
	bytes []byte,
) error {
	_ = "STUB: not implemented"

	// The serialized form of the block is the unsignedBytes followed by the
	// signature, which is prefixed by a uint32. So, we need to strip off the
	// signature as well as it's length prefix to get the unsigned bytes.
	return nil
}

func (m *statelessBlockMetadata) verify(
	b *statelessUnsignedBlock,
	sig []byte,
	chainID ids.ID,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *statelessBlockMetadata) ID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (m *statelessBlockMetadata) Timestamp() time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (m *statelessBlockMetadata) Proposer() ids.NodeID {
	_ = "STUB: not implemented"
	return *new(ids.NodeID)
}

func (m *statelessBlockMetadata) Bytes() []byte { _ = "STUB: not implemented"; return nil }

type statelessBlock struct {
	statelessBlockMetadata

	StatelessBlock statelessUnsignedBlock `serialize:"true" json:"statelessBlock"`
	Signature      []byte                 `serialize:"true" json:"signature"`
}

type statelessGraniteBlock struct {
	statelessBlockMetadata

	StatelessGraniteBlock statelessUnsignedGraniteBlock `serialize:"true" json:"statelessGraniteBlock"`
	Signature             []byte                        `serialize:"true" json:"signature"`
}

func (b *statelessBlock) ParentID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (b *statelessBlock) Block() []byte { _ = "STUB: not implemented"; return nil }

func (b *statelessBlock) initialize(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (b *statelessBlock) verify(chainID ids.ID) error { _ = "STUB: not implemented"; return nil }

func (b *statelessBlock) PChainHeight() uint64 { _ = "STUB: not implemented"; return 0 }

func (*statelessBlock) PChainEpoch() Epoch { _ = "STUB: not implemented"; return *new(Epoch) }

func (b *statelessGraniteBlock) ParentID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (b *statelessGraniteBlock) Block() []byte { _ = "STUB: not implemented"; return nil }

func (b *statelessGraniteBlock) PChainHeight() uint64 { _ = "STUB: not implemented"; return 0 }

func (b *statelessGraniteBlock) PChainEpoch() Epoch { _ = "STUB: not implemented"; return *new(Epoch) }

func (b *statelessGraniteBlock) initialize(bytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *statelessGraniteBlock) verify(chainID ids.ID) error { _ = "STUB: not implemented"; return nil }
