// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package reflectcodec

import (
	"errors"
	"reflect"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/utils/set"
	"github.com/ava-labs/avalanchego/utils/wrappers"
)

const (
	// DefaultTagName that enables serialization.
	DefaultTagName  = "serialize"
	initialSliceLen = 16
)

var (
	_ codec.Codec = (*genericCodec)(nil)

	errNeedPointer             = errors.New("argument to unmarshal must be a pointer")
	errRecursiveInterfaceTypes = errors.New("recursive interface types")
)

type TypeCodec interface {
	// UnpackPrefix unpacks the prefix of an interface from the given packer.
	// The prefix specifies the concrete type that the interface should be
	// deserialized into. This function returns a new instance of that concrete
	// type. The concrete type must implement the given type.
	UnpackPrefix(*wrappers.Packer, reflect.Type) (reflect.Value, error)

	// PackPrefix packs the prefix for the given type into the given packer.
	// This identifies the bytes that follow, which are the byte representation
	// of an interface, as having the given concrete type.
	// When deserializing the bytes, the prefix specifies which concrete type
	// to deserialize into.
	PackPrefix(*wrappers.Packer, reflect.Type) error

	// PrefixSize returns prefix length for the given type into the given
	// packer.
	PrefixSize(reflect.Type) int
}

// genericCodec handles marshaling and unmarshaling of structs with a generic
// implementation for interface encoding.
//
// A few notes:
//
//  1. We use "marshal" and "serialize" interchangeably, and "unmarshal" and
//     "deserialize" interchangeably
//  2. To include a field of a struct in the serialized form, add the tag
//     `{tagName}:"true"` to it. `{tagName}` defaults to `serialize`.
//  3. These typed members of a struct may be serialized:
//     bool, string, uint[8,16,32,64], int[8,16,32,64],
//     structs, slices, arrays, maps, interface.
//     structs, slices, maps and arrays can only be serialized if their constituent
//     values can be.
//  4. To marshal an interface, you must pass a pointer to the value
//  5. To unmarshal an interface, you must call
//     codec.RegisterType([instance of the type that fulfills the interface]).
//  6. Serialized fields must be exported
//  7. nil slices are marshaled as empty slices
type genericCodec struct {
	typer   TypeCodec
	fielder StructFielder
}

// New returns a new, concurrency-safe codec
func New(typer TypeCodec, tagNames []string) codec.Codec {
	_ = "STUB: not implemented"
	return *new(codec.Codec)
}

func (c *genericCodec) Size(value interface{}) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

/*=typeStack*/

// size returns the size of the value along with whether the value is constant
// sized.
func (c *genericCodec) size(
	value reflect.Value,
	typeStack set.Set[reflect.Type],
) (int, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// For fixed-size types we manually calculate lengths rather than
// processing each element separately to improve performance.

// For fixed-size types we manually calculate lengths rather than
// processing each element separately to improve performance.

// To marshal an interface, [value] must be a pointer to the interface
func (c *genericCodec) MarshalInto(value interface{}, p *wrappers.Packer) error {
	_ = "STUB: not implemented"
	return nil
}

/*=typeStack*/

// marshal writes the byte representation of [value] to [p]
//
// c.lock should be held for the duration of this function
func (c *genericCodec) marshal(
	value reflect.Value,
	p *wrappers.Packer,
	typeStack set.Set[reflect.Type],
) error {
	_ = "STUB: not implemented"
	return nil
}

// # elements in the slice/array. 0 if this slice is nil.

// pack # elements

// Returning here prevents execution of the (expensive) reflect
// calls below which check if the slice is []byte and, if it is,
// the call of value.Bytes()

// If this is a slice of bytes, manually pack the bytes rather
// than calling marshal on each byte. This improves performance.

// Process each element in the slice

// Process each element in the array

// Go through all fields of this struct that are serialized
// Serialize the field and write to byte array

// pack # elements

// pack key-value pairs sorted by increasing key

// pack key

// serialize and pack value

// UnmarshalFrom unmarshals [p.Bytes] into [dest], where [dest] must be a pointer or
// interface
func (c *genericCodec) UnmarshalFrom(p *wrappers.Packer, dest interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

/*=typeStack*/

// Unmarshal from p.Bytes into [value]. [value] must be addressable.
//
// c.lock should be held for the duration of this function
func (c *genericCodec) unmarshal(
	p *wrappers.Packer,
	value reflect.Value,
	typeStack set.Set[reflect.Type],
) error {
	_ = "STUB: not implemented"
	return nil
}

// If this is a slice of bytes, manually unpack the bytes rather
// than calling unmarshal on each byte. This improves performance.

// Unmarshal each element and append it into the slice.

// Get a slice to the underlying array value

// Unmarshal into the struct

// Get indices of fields that will be unmarshaled into

// Go through the fields and unmarshal into them

// Get the type this pointer points to

// Create a new pointer to a new value of the underlying type

// Fill the value

// Assign to the top-level struct's member

// Set [value] to be a new map of the appropriate type.

// Get the key's byte representation and check that the new key is
// actually bigger (according to bytes.Compare) than the previous
// key.
//
// We do this to enforce that key-value pairs are sorted by
// increasing key.

// Get the value

// Assign the key-value pair in the map
