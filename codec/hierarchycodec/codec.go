// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package hierarchycodec

import (
	"reflect"
	"sync"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/utils/bimap"
	"github.com/ava-labs/avalanchego/utils/wrappers"
)

var (
	_ Codec              = (*hierarchyCodec)(nil)
	_ codec.Codec        = (*hierarchyCodec)(nil)
	_ codec.Registry     = (*hierarchyCodec)(nil)
	_ codec.GeneralCodec = (*hierarchyCodec)(nil)
)

// Codec marshals and unmarshals
type Codec interface {
	codec.Registry
	codec.Codec
	SkipRegistrations(int)
	NextGroup()
}

type typeID struct {
	groupID uint16
	typeID  uint16
}

// Codec handles marshaling and unmarshaling of structs
type hierarchyCodec struct {
	codec.Codec

	lock            sync.RWMutex
	currentGroupID  uint16
	nextTypeID      uint16
	registeredTypes *bimap.BiMap[typeID, reflect.Type]
}

// New returns a new, concurrency-safe codec
func New(tagNames []string) Codec { _ = "STUB: not implemented"; return *new(Codec) }

// NewDefault returns a new codec with reasonable default values
func NewDefault() Codec { _ = "STUB: not implemented"; return *new(Codec) }

// SkipRegistrations some number of type IDs
func (c *hierarchyCodec) SkipRegistrations(num int) { _ = "STUB: not implemented"; return }

// NextGroup moves to the next group registry
func (c *hierarchyCodec) NextGroup() { _ = "STUB: not implemented"; return }

// RegisterType is used to register types that may be unmarshaled into an interface
// [val] is a value of the type being registered
func (c *hierarchyCodec) RegisterType(val interface{}) error { _ = "STUB: not implemented"; return nil }

func (*hierarchyCodec) PrefixSize(reflect.Type) int {
	_ = "STUB: not implemented"
	// see PackPrefix implementation
	return 0
}

func (c *hierarchyCodec) PackPrefix(p *wrappers.Packer, valueType reflect.Type) error {
	_ = "STUB: not implemented"
	return nil
}

// Get the type ID of the value being marshaled

// Pack type ID so we know what to unmarshal this into

func (c *hierarchyCodec) UnpackPrefix(p *wrappers.Packer, valueType reflect.Type) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// Get the group ID
// Get the type ID

// Get a type that implements the interface

// Ensure type actually does implement the interface

// instance of the proper type
