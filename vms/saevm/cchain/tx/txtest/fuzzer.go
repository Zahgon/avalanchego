// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package txtest provides test helpers for using [tx.Tx].
package txtest

import (
	"testing"

	"github.com/ava-labs/libevm/common"

	// Imported for [codec.Manager] comment resolution.
	_ "github.com/ava-labs/avalanchego/codec"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/vms/components/avax"
	"github.com/ava-labs/avalanchego/vms/saevm/cchain/tx"
)

// F works like [testing.F], but allows for the usage of [tx.Tx].
//
// This type should be used over [testing.F] when it is desired to consistently
// produce a parseable transaction.
type F struct {
	*testing.F

	// Addresses, if non-empty, biases addresses toward this alphabet.
	Addresses []common.Address
	// AssetIDs, if non-empty, biases assetIDs toward this alphabet.
	AssetIDs []ids.ID
}

// Add works like [testing.F.Add], but expects a [tx.Tx].
func (f *F) Add(tx *tx.Tx) { _ = "STUB: not implemented"; return }

// Fuzz works like [testing.F.Fuzz], but provides a [tx.Tx].
func (f *F) Fuzz(ff func(t *testing.T, tx *tx.Tx)) { _ = "STUB: not implemented"; return }

// It's possible for the fuzzer to generate a tx that exceeds the codec
// size limits.

// decoder turns a byte stream into structured data.
//
// The byte stream is consumed as structured data is produced; once exhausted,
// methods return the zero value of their result type.
type decoder struct {
	// data is the byte stream backing the decoder.
	data []byte
	// addresses, if non-empty, biases [decoder.address] toward this alphabet.
	addresses []common.Address
	// assetIDs, if non-empty, biases [decoder.assetID] toward this alphabet.
	assetIDs []ids.ID
}

// bytes always returns a slice of length n, even if the decoder is exhausted.
func (d *decoder) bytes(n int) []byte { _ = "STUB: not implemented"; return nil }

func (d *decoder) bool() bool           { _ = "STUB: not implemented"; return false }
func (d *decoder) uint32() uint32       { _ = "STUB: not implemented"; return 0 }
func (d *decoder) uint64() uint64       { _ = "STUB: not implemented"; return 0 }
func (d *decoder) addr() common.Address { _ = "STUB: not implemented"; return *new(common.Address) }
func (d *decoder) id() ids.ID           { _ = "STUB: not implemented"; return *new(ids.ID) }
func (d *decoder) shortID() ids.ShortID { _ = "STUB: not implemented"; return *new(ids.ShortID) }
func (d *decoder) signature() [65]byte  { _ = "STUB: not implemented"; return nil }

// intn returns a value in [0, x).
func (d *decoder) intn(x int) int { _ = "STUB: not implemented"; return 0 }

//#nosec G115 -- Overflow is impossible.

// element returns a random element biased towards values in s.
func element[T any](d *decoder, s []T, gen func(*decoder) T) T {
	_ = "STUB: not implemented"
	// [decoder.bool] must be read before checking whether or not s is empty so
	// that encoding can assume that a bool will be read.
	return *new(T)
}

// sliceOf generates a random slice of generated entries. The returned value is
// non-nil. The length is random, but is typically small.
func sliceOf[T any](d *decoder, gen func(*decoder) T) []T {
	_ = "STUB: not implemented"
	// The [codec.Manager] always generates non-nil slices. To avoid nil and
	// empty comparison issues, we also always generate non-nil slices.
	return nil
}

// [decoder.bool] returns false once the data is exhausted, so this loop
// will eventually terminate.

func (d *decoder) address() common.Address { _ = "STUB: not implemented"; return *new(common.Address) }

func (d *decoder) assetID() ids.ID { _ = "STUB: not implemented"; return *new(ids.ID) }

func (d *decoder) transferableInput() *avax.TransferableInput {
	_ = "STUB: not implemented"
	return nil
}

func (d *decoder) transferableOutput() *avax.TransferableOutput {
	_ = "STUB: not implemented"
	return nil
}

func (d *decoder) input() tx.Input { _ = "STUB: not implemented"; return *new(tx.Input) }

func (d *decoder) output() tx.Output { _ = "STUB: not implemented"; return *new(tx.Output) }

func (d *decoder) importTx() *tx.Import { _ = "STUB: not implemented"; return nil }

func (d *decoder) exportTx() *tx.Export { _ = "STUB: not implemented"; return nil }

func (d *decoder) unsigned() tx.Unsigned { _ = "STUB: not implemented"; return *new(tx.Unsigned) }

func (d *decoder) credential() tx.Credential { _ = "STUB: not implemented"; return *new(tx.Credential) }

func (d *decoder) tx() *tx.Tx { _ = "STUB: not implemented"; return nil }

// encoder writes a byte stream that [decoder] will decode into the same
// structure.
type encoder []byte

func (e *encoder) bytes(b []byte)        { _ = "STUB: not implemented"; return }
func (e *encoder) uint32(v uint32)       { _ = "STUB: not implemented"; return }
func (e *encoder) uint64(v uint64)       { _ = "STUB: not implemented"; return }
func (e *encoder) addr(v common.Address) { _ = "STUB: not implemented"; return }
func (e *encoder) id(v ids.ID)           { _ = "STUB: not implemented"; return }
func (e *encoder) shortID(v ids.ShortID) { _ = "STUB: not implemented"; return }
func (e *encoder) signature(v [65]byte)  { _ = "STUB: not implemented"; return }

func (e *encoder) bool(b bool) { _ = "STUB: not implemented"; return }

// address always picks the raw-bytes branch in [element] so the encoded value
// is independent of the alphabet.
func (e *encoder) address(v common.Address) { _ = "STUB: not implemented"; return }

// assetID always picks the raw-bytes branch in [element] so the encoded value
// is independent of the alphabet.
func (e *encoder) assetID(v ids.ID) { _ = "STUB: not implemented"; return }

func sliceTo[T any](e *encoder, items []T, gen func(*encoder, T)) {
	_ = "STUB: not implemented"
	return
}

func (e *encoder) transferableInput(in *avax.TransferableInput) { _ = "STUB: not implemented"; return }

func (e *encoder) transferableOutput(out *avax.TransferableOutput) {
	_ = "STUB: not implemented"
	return
}

func (e *encoder) input(i tx.Input) { _ = "STUB: not implemented"; return }

func (e *encoder) output(o tx.Output) { _ = "STUB: not implemented"; return }

func (e *encoder) importTx(t *tx.Import) { _ = "STUB: not implemented"; return }

func (e *encoder) exportTx(t *tx.Export) { _ = "STUB: not implemented"; return }

func (e *encoder) unsigned(u tx.Unsigned) { _ = "STUB: not implemented"; return }

func (e *encoder) credential(c tx.Credential) { _ = "STUB: not implemented"; return }

func (e *encoder) tx(tx *tx.Tx) { _ = "STUB: not implemented"; return }
