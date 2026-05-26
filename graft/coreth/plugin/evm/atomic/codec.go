// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomic

import (
	"errors"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/codec/linearcodec"
	"github.com/ava-labs/avalanchego/utils/wrappers"
	"github.com/ava-labs/avalanchego/vms/secp256k1fx"
)

const CodecVersion = uint16(0)

var (
	// Codec does serialization and deserialization
	Codec codec.Manager

	errMissingAtomicTxs = errors.New("cannot build a block with non-empty extra data and zero atomic transactions")
)

func init() {
	Codec = codec.NewDefaultManager()

	var (
		lc   = linearcodec.NewDefault()
		errs = wrappers.Errs{}
	)
	errs.Add(
		lc.RegisterType(&UnsignedImportTx{}),
		lc.RegisterType(&UnsignedExportTx{}),
	)
	lc.SkipRegistrations(3)
	errs.Add(
		lc.RegisterType(&secp256k1fx.TransferInput{}),
	)
	lc.SkipRegistrations(1)
	errs.Add(
		lc.RegisterType(&secp256k1fx.TransferOutput{}),
	)
	lc.SkipRegistrations(1)
	errs.Add(
		lc.RegisterType(&secp256k1fx.Credential{}),
		lc.RegisterType(&secp256k1fx.Input{}),
		lc.RegisterType(&secp256k1fx.OutputOwners{}),
		Codec.RegisterCodec(CodecVersion, lc),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
}

// extractAtomicTxs returns the atomic transactions in [atomicTxBytes] if
// they exist.
// if [batch] is true, it attempts to unmarshal [atomicTxBytes] as a slice of
// transactions (post-ApricotPhase5), and if it is false, then it unmarshals
// it as a single atomic transaction.
func ExtractAtomicTxs(atomicTxBytes []byte, batch bool, codec codec.Manager) ([]*Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// [ExtractAtomicTx] extracts a singular atomic transaction from [atomicTxBytes]
// and returns a slice of atomic transactions for compatibility with the type returned post
// ApricotPhase5.
// Note: this function assumes [atomicTxBytes] is non-empty.
func ExtractAtomicTx(atomicTxBytes []byte, codec codec.Manager) (*Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// [ExtractAtomicTxsBatch] extracts a slice of atomic transactions from [atomicTxBytes].
// Note: this function assumes [atomicTxBytes] is non-empty.
func ExtractAtomicTxsBatch(atomicTxBytes []byte, codec codec.Manager) ([]*Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Do not allow non-empty extra data field to contain zero atomic transactions. This would allow
// people to construct a block that contains useless data.
