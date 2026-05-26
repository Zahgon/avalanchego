// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package customtypes

import (
	"io"
	"math/big"

	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/common/hexutil"
	"github.com/ava-labs/libevm/rlp"

	"github.com/ava-labs/avalanchego/vms/evm/acp226"

	ethtypes "github.com/ava-labs/libevm/core/types"
)

// GetHeaderExtra returns the [HeaderExtra] from the given [Header].
func GetHeaderExtra(h *ethtypes.Header) *HeaderExtra { _ = "STUB: not implemented"; return nil }

// SetHeaderExtra sets the given [HeaderExtra] on the [Header].
func SetHeaderExtra(h *ethtypes.Header, extra *HeaderExtra) { _ = "STUB: not implemented"; return }

// WithHeaderExtra sets the given [HeaderExtra] on the [Header]
// and returns the [Header] for chaining.
func WithHeaderExtra(h *ethtypes.Header, extra *HeaderExtra) *ethtypes.Header {
	_ = "STUB: not implemented"
	return nil
}

// HeaderExtra is a struct that contains extra fields used by Avalanche
// in the block header.
// This type uses [HeaderSerializable] to encode and decode the extra fields
// along with the upstream type for compatibility with existing network blocks.
type HeaderExtra struct {
	ExtDataHash      common.Hash
	ExtDataGasUsed   *big.Int
	BlockGasCost     *big.Int
	TimeMilliseconds *uint64
	MinDelayExcess   *acp226.DelayExcess
}

// HeaderTimeMilliseconds returns the header timestamp in milliseconds.
// If the header has the Granite field TimeMilliseconds set in extras, it is used.
// Otherwise, it falls back to seconds-based Time multiplied by 1000.
func HeaderTimeMilliseconds(h *ethtypes.Header) uint64 { _ = "STUB: not implemented"; return 0 }

// EncodeRLP RLP encodes the given [ethtypes.Header] and [HeaderExtra] together
// to the `writer`. It does merge both structs into a single [HeaderSerializable].
func (h *HeaderExtra) EncodeRLP(eth *ethtypes.Header, writer io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

// DecodeRLP RLP decodes from the [*rlp.Stream] and writes the output to both the
// [ethtypes.Header] passed as argument and to the receiver [HeaderExtra].
func (h *HeaderExtra) DecodeRLP(eth *ethtypes.Header, stream *rlp.Stream) error {
	_ = "STUB: not implemented"
	return nil
}

// EncodeJSON JSON encodes the given [ethtypes.Header] and [HeaderExtra] together
// to the `writer`. It does merge both structs into a single [HeaderSerializable].
func (h *HeaderExtra) EncodeJSON(eth *ethtypes.Header) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DecodeJSON JSON decodes from the `input` bytes and writes the output to both the
// [ethtypes.Header] passed as argument and to the receiver [HeaderExtra].
func (h *HeaderExtra) DecodeJSON(eth *ethtypes.Header, input []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *HeaderExtra) PostCopy(dst *ethtypes.Header) { _ = "STUB: not implemented"; return }

func (h *HeaderSerializable) updateFromEth(eth *ethtypes.Header) { _ = "STUB: not implemented"; return }

func (h *HeaderSerializable) updateToEth(eth *ethtypes.Header) { _ = "STUB: not implemented"; return }

func (h *HeaderSerializable) updateFromExtras(extras *HeaderExtra) {
	_ = "STUB: not implemented"
	return
}

func (h *HeaderSerializable) updateToExtras(extras *HeaderExtra) { _ = "STUB: not implemented"; return }

// NOTE: both generators currently do not support type aliases.
// We are using custom versions of these programs for now to support type aliases,
// see https://github.com/ava-labs/avalanchego/graft/coreth/pull/746#discussion_r1969673252
//go:generate go tool gencodec -type HeaderSerializable -field-override headerMarshaling -out gen_header_serializable_json.go
//go:generate go tool rlpgen -type HeaderSerializable -out gen_header_serializable_rlp.go

// HeaderSerializable defines the header of a block in the Ethereum blockchain,
// as it is to be serialized into RLP and JSON. Note it must be exported so that
// rlpgen can generate the serialization code from it.
//
//nolint:tagalign
type HeaderSerializable struct {
	ParentHash  common.Hash         `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash         `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address      `json:"miner"            gencodec:"required"`
	Root        common.Hash         `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash         `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash         `json:"receiptsRoot"     gencodec:"required"`
	Bloom       ethtypes.Bloom      `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int            `json:"difficulty"       gencodec:"required"`
	Number      *big.Int            `json:"number"           gencodec:"required"`
	GasLimit    uint64              `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64              `json:"gasUsed"          gencodec:"required"`
	Time        uint64              `json:"timestamp"        gencodec:"required"`
	Extra       []byte              `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash         `json:"mixHash"`
	Nonce       ethtypes.BlockNonce `json:"nonce"`
	ExtDataHash common.Hash         `json:"extDataHash"      gencodec:"required"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`

	// ExtDataGasUsed was added by Apricot Phase 4 and is ignored in legacy
	// headers.
	//
	// It is not a uint64 like GasLimit or GasUsed because it is not possible to
	// correctly encode this field optionally with uint64.
	ExtDataGasUsed *big.Int `json:"extDataGasUsed" rlp:"optional"`

	// BlockGasCost was added by Apricot Phase 4 and is ignored in legacy
	// headers.
	BlockGasCost *big.Int `json:"blockGasCost" rlp:"optional"`

	// BlobGasUsed was added by EIP-4844 and is ignored in legacy headers.
	BlobGasUsed *uint64 `json:"blobGasUsed" rlp:"optional"`

	// ExcessBlobGas was added by EIP-4844 and is ignored in legacy headers.
	ExcessBlobGas *uint64 `json:"excessBlobGas" rlp:"optional"`

	// ParentBeaconRoot was added by EIP-4788 and is ignored in legacy headers.
	ParentBeaconRoot *common.Hash `json:"parentBeaconBlockRoot" rlp:"optional"`

	// TimeMilliseconds was added by Granite and is ignored in legacy headers.
	TimeMilliseconds *uint64 `json:"timestampMilliseconds" rlp:"optional"`

	// MinDelayExcess was added by Granite and is ignored in legacy headers.
	// We use *uint64 type here to avoid rlpgen generating incorrect code
	MinDelayExcess *uint64 `json:"minDelayExcess" rlp:"optional"`
}

// field type overrides for gencodec
type headerMarshaling struct {
	Difficulty       *hexutil.Big
	Number           *hexutil.Big
	GasLimit         hexutil.Uint64
	GasUsed          hexutil.Uint64
	Time             hexutil.Uint64
	Extra            hexutil.Bytes
	BaseFee          *hexutil.Big
	ExtDataGasUsed   *hexutil.Big
	BlockGasCost     *hexutil.Big
	Hash             common.Hash `json:"hash"` // adds call to Hash() in MarshalJSON
	BlobGasUsed      *hexutil.Uint64
	ExcessBlobGas    *hexutil.Uint64
	TimeMilliseconds *hexutil.Uint64
	MinDelayExcess   *hexutil.Uint64
}

// Hash returns the block hash of the header, which is simply the keccak256 hash of its
// RLP encoding.
// This function MUST be exported and is used in [HeaderSerializable.EncodeJSON] which is
// generated to the file gen_header_json.go.
func (h *HeaderSerializable) Hash() common.Hash {
	_ = "STUB: not implemented"
	return *new(common.Hash)
}
