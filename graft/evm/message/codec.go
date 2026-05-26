// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/utils/units"
)

const (
	Version        = uint16(0)
	maxMessageSize = 2*units.MiB - 64*units.KiB // Subtract 64 KiB from p2p network cap to leave room for encoding overhead from AvalancheGo
)

// CodecType identifies the codec variant to use for message serialization.
// Each variant has its own skip count and LeafsRequest wire format.
type CodecType int

const (
	// CorethCodecType is used for C-Chain state sync messages.
	CorethCodecType CodecType = iota
	// SubnetEVMCodecType is used for Subnet-EVM state sync messages.
	SubnetEVMCodecType
)

var (
	// CorethCodec is used for C-Chain state sync messages
	CorethCodec codec.Manager
	// SubnetEVMCodec is used for Subnet-EVM state sync messages
	SubnetEVMCodec codec.Manager
)

func init() {
	CorethCodec = newCodec(CorethCodecType)
	SubnetEVMCodec = newCodec(SubnetEVMCodecType)
}

func newCodec(codecType CodecType) codec.Manager {
	_ = "STUB: not implemented"
	// Codec-specific configuration: skip count and LeafsRequest type.
	return *new(codec.Manager)
}

// deprecated gossip types and sync summary type

// deprecated gossip types

// Skip registrations to maintain type ID compatibility after removing deprecated types.

// Skip 3 registrations for deprecated Warp types:
// MessageSignatureRequest, BlockSignatureRequest, SignatureResponse
// See https://github.com/ava-labs/coreth/pull/999
