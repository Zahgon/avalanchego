// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package linearcodec

import (
	"reflect"
	"sync"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/utils/bimap"
	"github.com/ava-labs/avalanchego/utils/wrappers"
)

var (
	_ Codec              = (*linearCodec)(nil)
	_ codec.Codec        = (*linearCodec)(nil)
	_ codec.Registry     = (*linearCodec)(nil)
	_ codec.GeneralCodec = (*linearCodec)(nil)
)

// Codec marshals and unmarshals
type Codec interface {
	codec.Registry
	codec.Codec
	SkipRegistrations(int)
}

// Codec handles marshaling and unmarshaling of structs
type linearCodec struct {
	codec.Codec

	lock            sync.RWMutex
	nextTypeID      uint32
	registeredTypes *bimap.BiMap[uint32, reflect.Type]
}

// New returns a new, concurrency-safe codec; it allow to specify tagNames.
func New(tagNames []string) Codec { _ = "STUB: not implemented"; return *new(Codec) }

// NewDefault is a convenience constructor; it returns a new codec with default
// tagNames.
func NewDefault() Codec { _ = "STUB: not implemented"; return *new(Codec) }

// Skip some number of type IDs
func (c *linearCodec) SkipRegistrations(num int) { _ = "STUB: not implemented"; return }

// RegisterType is used to register types that may be unmarshaled into an interface
// [val] is a value of the type being registered
func (c *linearCodec) RegisterType(val interface{}) error { _ = "STUB: not implemented"; return nil }

func (*linearCodec) PrefixSize(reflect.Type) int {
	_ = "STUB: not implemented"
	// see PackPrefix implementation
	return 0
}

func (c *linearCodec) PackPrefix(p *wrappers.Packer, valueType reflect.Type) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the type ID of the value being marshaled

// Pack type ID so we know what to unmarshal this into

func (c *linearCodec) UnpackPrefix(p *wrappers.Packer, valueType reflect.Type) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// Get the type ID

// Get a type that implements the interface

// Ensure type actually does implement the interface

// instance of the proper type
