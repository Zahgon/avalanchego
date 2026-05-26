// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package codec

import (
	"errors"
	"sync"

	"github.com/ava-labs/avalanchego/utils/units"
	"github.com/ava-labs/avalanchego/utils/wrappers"
)

const (
	VersionSize = wrappers.ShortLen

	// default max size, in bytes, of something being marshaled by Marshal()
	defaultMaxSize = 256 * units.KiB

	// initial capacity of byte slice that values are marshaled into.
	// Larger value --> need less memory allocations but possibly have allocated but unused memory
	// Smaller value --> need more memory allocations but more efficient use of allocated memory
	initialSliceCap = 128
)

var (
	ErrUnknownVersion    = errors.New("unknown codec version")
	ErrMarshalNil        = errors.New("can't marshal nil pointer or interface")
	ErrUnmarshalNil      = errors.New("can't unmarshal nil")
	ErrUnmarshalTooBig   = errors.New("byte array exceeds maximum length")
	ErrCantPackVersion   = errors.New("couldn't pack codec version")
	ErrCantUnpackVersion = errors.New("couldn't unpack codec version")
	ErrDuplicatedVersion = errors.New("duplicated codec version")
	ErrExtraSpace        = errors.New("trailing buffer space")
)

var _ Manager = (*manager)(nil)

// Manager describes the functionality for managing codec versions.
type Manager interface {
	// Associate the given codec with the given version ID
	RegisterCodec(version uint16, codec Codec) error

	// Size returns the size, in bytes, of [value] when it's marshaled
	// using the codec with the given version.
	// RegisterCodec must have been called with that version.
	// If [value] is nil, returns [ErrMarshalNil]
	Size(version uint16, value interface{}) (int, error)

	// Marshal the given value using the codec with the given version.
	// RegisterCodec must have been called with that version.
	Marshal(version uint16, source interface{}) (destination []byte, err error)

	// Unmarshal the given bytes into the given destination. [destination] must
	// be a pointer or an interface. Returns the version of the codec that
	// produces the given bytes.
	Unmarshal(source []byte, destination interface{}) (version uint16, err error)
}

// NewManager returns a new codec manager.
func NewManager(maxSize int) Manager { _ = "STUB: not implemented"; return *new(Manager) }

// NewDefaultManager returns a new codec manager.
func NewDefaultManager() Manager { _ = "STUB: not implemented"; return *new(Manager) }

type manager struct {
	lock    sync.RWMutex
	maxSize int
	codecs  map[uint16]Codec
}

// RegisterCodec is used to register a new codec version that can be used to
// (un)marshal with.
func (m *manager) RegisterCodec(version uint16, codec Codec) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *manager) Size(version uint16, value interface{}) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// can't marshal nil

// Add [VersionSize] for the codec version

// To marshal an interface, [value] must be a pointer to the interface.
func (m *manager) Marshal(version uint16, value interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// can't marshal nil

// Should never happen

// Unmarshal unmarshals [bytes] into [dest], where [dest] must be a pointer or
// interface.
func (m *manager) Unmarshal(bytes []byte, dest interface{}) (uint16, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Make sure the codec version is correct
