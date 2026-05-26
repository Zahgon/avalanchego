// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package warp

import (
	"context"
	"errors"

	"github.com/ava-labs/avalanchego/cache"
	"github.com/ava-labs/avalanchego/cache/lru"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/network/p2p/acp118"
	"github.com/ava-labs/avalanchego/snow/consensus/snowman"
	"github.com/ava-labs/avalanchego/vms/evm/uptimetracker"

	avalancheWarp "github.com/ava-labs/avalanchego/vms/platformvm/warp"
)

var (
	_                         Backend = (*backend)(nil)
	ErrValidateBlock                  = errors.New("failed to validate block message")
	ErrVerifyWarpMessage              = errors.New("failed to verify warp message")
	errParsingOffChainMessage         = errors.New("failed to parse off-chain message")

	messageCacheSize = 500
)

type BlockClient interface {
	GetAcceptedBlock(ctx context.Context, blockID ids.ID) (snowman.Block, error)
}

// Backend tracks signature-eligible warp messages and provides an interface to fetch them.
// The backend is also used to query for warp message signatures by the signature request handler.
type Backend interface {
	// AddMessage signs [unsignedMessage] and adds it to the warp backend database
	AddMessage(unsignedMessage *avalancheWarp.UnsignedMessage) error

	// GetMessageSignature validates the message and returns the signature of the requested message.
	GetMessageSignature(ctx context.Context, message *avalancheWarp.UnsignedMessage) ([]byte, error)

	// GetBlockSignature returns the signature of a hash payload containing blockID if it's the ID of an accepted block.
	GetBlockSignature(ctx context.Context, blockID ids.ID) ([]byte, error)

	// GetMessage retrieves the [unsignedMessage] from the warp backend database if available
	GetMessage(messageHash ids.ID) (*avalancheWarp.UnsignedMessage, error)

	acp118.Verifier
}

// backend implements Backend, keeps track of warp messages, and generates message signatures.
type backend struct {
	networkID                 uint32
	sourceChainID             ids.ID
	db                        database.Database
	warpSigner                avalancheWarp.Signer
	blockClient               BlockClient
	uptimeTracker             *uptimetracker.UptimeTracker
	signatureCache            cache.Cacher[ids.ID, []byte]
	messageCache              *lru.Cache[ids.ID, *avalancheWarp.UnsignedMessage]
	offchainAddressedCallMsgs map[ids.ID]*avalancheWarp.UnsignedMessage
	stats                     *verifierStats
}

// NewBackend creates a new Backend, and initializes the signature cache and message tracking database.
func NewBackend(
	networkID uint32,
	sourceChainID ids.ID,
	warpSigner avalancheWarp.Signer,
	blockClient BlockClient,
	uptimeTracker *uptimetracker.UptimeTracker,
	db database.Database,
	signatureCache cache.Cacher[ids.ID, []byte],
	offchainMessages [][]byte,
) (Backend, error) {
	_ = "STUB: not implemented"
	return *new(Backend), nil
}

func (b *backend) initOffChainMessages(offchainMessages [][]byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (b *backend) AddMessage(unsignedMessage *avalancheWarp.UnsignedMessage) error {
	_ = "STUB: not implemented"
	return nil
}

// In the case when a node restarts, and possibly changes its bls key, the cache gets emptied but the database does not.
// So to avoid having incorrect signatures saved in the database after a bls key change, we save the full message in the database.
// Whereas for the cache, after the node restart, the cache would be emptied so we can directly save the signatures.

func (b *backend) GetMessageSignature(ctx context.Context, unsignedMessage *avalancheWarp.UnsignedMessage) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *backend) GetBlockSignature(ctx context.Context, blockID ids.ID) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *backend) GetMessage(messageID ids.ID) (*avalancheWarp.UnsignedMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *backend) signMessage(unsignedMessage *avalancheWarp.UnsignedMessage) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
