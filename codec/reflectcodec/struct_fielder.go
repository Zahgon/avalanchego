// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package reflectcodec

import (
	"reflect"
	"sync"
)

// TagValue is the value the tag must have to be serialized.
const TagValue = "true"

var _ StructFielder = (*structFielder)(nil)

// StructFielder handles discovery of serializable fields in a struct.
type StructFielder interface {
	// Returns the fields that have been marked as serializable in [t], which is
	// a struct type.
	// Returns an error if a field has tag "[tagName]: [TagValue]" but the field
	// is un-exported.
	// GetSerializedField(Foo) --> [1,5,8] means Foo.Field(1), Foo.Field(5),
	// Foo.Field(8) are to be serialized/deserialized.
	GetSerializedFields(t reflect.Type) ([]int, error)
}

func NewStructFielder(tagNames []string) StructFielder {
	_ = "STUB: not implemented"
	return *new(StructFielder)
}

type structFielder struct {
	lock sync.RWMutex

	// multiple tags per field can be specified. A field is serialized/deserialized
	// if it has at least one of the specified tags.
	tags []string

	// Key: a struct type
	// Value: Slice where each element is index in the struct type of a field
	// that is serialized/deserialized e.g. Foo --> [1,5,8] means Foo.Field(1),
	// etc. are to be serialized/deserialized. We assume this cache is pretty
	// small (a few hundred keys at most) and doesn't take up much memory.
	serializedFieldIndices map[reflect.Type][]int
}

func (s *structFielder) GetSerializedFields(t reflect.Type) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// use pre-computed result

// Go through all fields of this struct

// Multiple tags per fields can be specified.
// Serialize/Deserialize field if it has
// any tag with the right value

// Can only marshal exported fields

// cache result

func (s *structFielder) getCachedSerializedFields(t reflect.Type) ([]int, bool) {
	_ = "STUB: not implemented"
	return nil, false
}
