// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// Package codectest provides a test suite for testing codec implementations.
package codectest

import (
	"testing"

	codecpkg "github.com/ava-labs/avalanchego/codec"
)

// A NamedTest couples a test in the suite with a human-readable name.
type NamedTest struct {
	Name string
	Test func(testing.TB, codecpkg.GeneralCodec)
}

// Run runs the test on the GeneralCodec.
func (tt *NamedTest) Run(t *testing.T, c codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

// RunAll runs all [Tests], constructing a new GeneralCodec for each.
func RunAll(t *testing.T, ctor func() codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

// RunAllMultipleTags runs all [MultipleTagsTests], constructing a new GeneralCodec for each.
func RunAllMultipleTags(t *testing.T, ctor func() codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

var (
	Tests = []NamedTest{
		{"Struct", TestStruct},
		{"Register Struct Twice", TestRegisterStructTwice},
		{"UInt32", TestUInt32},
		{"UIntPtr", TestUIntPtr},
		{"Slice", TestSlice},
		{"Max-Size Slice", TestMaxSizeSlice},
		{"Bool", TestBool},
		{"Array", TestArray},
		{"Big Array", TestBigArray},
		{"Pointer To Struct", TestPointerToStruct},
		{"Slice Of Struct", TestSliceOfStruct},
		{"Interface", TestInterface},
		{"Slice Of Interface", TestSliceOfInterface},
		{"Array Of Interface", TestArrayOfInterface},
		{"Pointer To Interface", TestPointerToInterface},
		{"String", TestString},
		{"Nil Slice", TestNilSlice},
		{"Serialize Unexported Field", TestSerializeUnexportedField},
		{"Serialize Of NoSerialize Field", TestSerializeOfNoSerializeField},
		{"Nil Slice Serialization", TestNilSliceSerialization},
		{"Empty Slice Serialization", TestEmptySliceSerialization},
		{"Slice With Empty Serialization", TestSliceWithEmptySerialization},
		{"Slice With Empty Serialization Error", TestSliceWithEmptySerializationError},
		{"Map With Empty Serialization", TestMapWithEmptySerialization},
		{"Map With Empty Serialization Error", TestMapWithEmptySerializationError},
		{"Slice Too Large", TestSliceTooLarge},
		{"Negative Numbers", TestNegativeNumbers},
		{"Too Large Unmarshal", TestTooLargeUnmarshal},
		{"Unmarshal Invalid Interface", TestUnmarshalInvalidInterface},
		{"Extra Space", TestExtraSpace},
		{"Slice Length Overflow", TestSliceLengthOverflow},
		{"Map", TestMap},
		{"Can Marshal Large Slices", TestCanMarshalLargeSlices},
		{"Implements UnmarshalFrom", TestImplementsUnmarshalFrom},
	}

	MultipleTagsTests = []NamedTest{
		{"Multiple Tags", TestMultipleTags},
	}
)

// The below structs and interfaces exist
// for the sake of testing

var (
	_ Foo = (*MyInnerStruct)(nil)
	_ Foo = (*MyInnerStruct2)(nil)
)

type Foo interface {
	Foo() int
}

type MyInnerStruct struct {
	Str string `serialize:"true"`
}

func (*MyInnerStruct) Foo() int { _ = "STUB: not implemented"; return 0 }

type MyInnerStruct2 struct {
	Bool bool `serialize:"true"`
}

func (*MyInnerStruct2) Foo() int {
	_ = "STUB: not implemented"

	// MyInnerStruct3 embeds Foo, an interface,
	// so it has to implement TypeID and ConcreteInstance
	return 0
}

type MyInnerStruct3 struct {
	Str string        `serialize:"true"`
	M1  MyInnerStruct `serialize:"true"`
	F   Foo           `serialize:"true"`
}

type myStruct struct {
	InnerStruct  MyInnerStruct               `serialize:"true"`
	InnerStruct2 *MyInnerStruct              `serialize:"true"`
	Member1      int64                       `serialize:"true"`
	Member2      uint16                      `serialize:"true"`
	MyArray2     [5]string                   `serialize:"true"`
	MyArray3     [3]MyInnerStruct            `serialize:"true"`
	MyArray4     [2]*MyInnerStruct2          `serialize:"true"`
	MySlice      []byte                      `serialize:"true"`
	MySlice2     []string                    `serialize:"true"`
	MySlice3     []MyInnerStruct             `serialize:"true"`
	MySlice4     []*MyInnerStruct2           `serialize:"true"`
	MyArray      [4]byte                     `serialize:"true"`
	MyInterface  Foo                         `serialize:"true"`
	MySlice5     []Foo                       `serialize:"true"`
	InnerStruct3 MyInnerStruct3              `serialize:"true"`
	MyPointer    *Foo                        `serialize:"true"`
	MyMap1       map[string]string           `serialize:"true"`
	MyMap2       map[int32][]MyInnerStruct3  `serialize:"true"`
	MyMap3       map[MyInnerStruct2][]int32  `serialize:"true"`
	MyMap4       map[int32]*int32            `serialize:"true"`
	MyMap5       map[int32]int32             `serialize:"true"`
	MyMap6       map[[5]int32]int32          `serialize:"true"`
	MyMap7       map[interface{}]interface{} `serialize:"true"`
	Uint8        uint8                       `serialize:"true"`
	Int8         int8                        `serialize:"true"`
	Uint16       uint16                      `serialize:"true"`
	Int16        int16                       `serialize:"true"`
	Uint32       uint32                      `serialize:"true"`
	Int32        int32                       `serialize:"true"`
	Uint64       uint64                      `serialize:"true"`
	Int64        int64                       `serialize:"true"`
	Bool         bool                        `serialize:"true"`
	String       string                      `serialize:"true"`
}

// Test marshaling/unmarshaling a complicated struct
func TestStruct(t testing.TB, codec codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

// Register the types that may be unmarshaled into interfaces

// In myStructInstance MyMap4 is nil and in myStructUnmarshaled MyMap4 is an
// empty map

func TestRegisterStructTwice(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

func TestUInt32(t testing.TB, codec codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

func TestUIntPtr(t testing.TB, codec codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

func TestSlice(t testing.TB, codec codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

// Test marshalling/unmarshalling largest possible slice
func TestMaxSizeSlice(t testing.TB, codec codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

// Test marshalling a bool
func TestBool(t testing.TB, codec codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

// Test marshalling an array
func TestArray(t testing.TB, codec codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

// Test marshalling a really big array
func TestBigArray(t testing.TB, codec codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

// Test marshalling a pointer to a struct
func TestPointerToStruct(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// Test marshalling a slice of structs
func TestSliceOfStruct(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// Test marshalling an interface
func TestInterface(t testing.TB, codec codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

// Test marshalling a slice of interfaces
func TestSliceOfInterface(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// Test marshalling an array of interfaces
func TestArrayOfInterface(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// Test marshalling a pointer to an interface
func TestPointerToInterface(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// Test marshalling a string
func TestString(t testing.TB, codec codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

// Ensure a nil slice is unmarshaled to slice with length 0
func TestNilSlice(t testing.TB, codec codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

// Ensure that trying to serialize a struct with an unexported member
// that has `serialize:"true"` returns error
func TestSerializeUnexportedField(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

//nolint:revive,unused

func TestSerializeOfNoSerializeField(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// Test marshalling of nil slice
func TestNilSliceSerialization(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// 0 for codec version, then nil slice marshaled as 0 length slice

// Test marshaling a slice that has 0 elements (but isn't nil)
func TestEmptySliceSerialization(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// 0 for codec version (uint16) and 0 for size (uint32)

// Test marshaling empty slice of zero length structs
func TestSliceWithEmptySerialization(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// codec version (0x00, 0x00) then (0x00, 0x00, 0x00, 0x00) for numElts

func TestSliceWithEmptySerializationError(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// codec version (0x00, 0x00) then (0x00, 0x00, 0x00, 0x01) for numElts

// Test marshaling empty map of zero length structs
func TestMapWithEmptySerialization(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// codec version (0x00, 0x00) then (0x00, 0x00, 0x00, 0x00) for numElts

func TestMapWithEmptySerializationError(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// codec version (0x00, 0x00) then (0x00, 0x00, 0x00, 0x01) for numElts

func TestSliceTooLarge(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// Ensure serializing structs with negative number members works
func TestNegativeNumbers(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// Ensure deserializing structs with too many bytes errors correctly
func TestTooLargeUnmarshal(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

type outerInterface interface {
	ToInt() int
}

type outer struct {
	Interface outerInterface `serialize:"true"`
}

type innerInterface struct{}

func (*innerInterface) ToInt() int { _ = "STUB: not implemented"; return 0 }

type innerNoInterface struct{}

// Ensure deserializing structs into the wrong interface errors gracefully
func TestUnmarshalInvalidInterface(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// Test unmarshaling something with extra data
func TestExtraSpace(t testing.TB, codec codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

// codec version 0x0000 then 0x01 for b then 0x02 as extra data.

// Ensure deserializing slices whose lengths exceed MaxInt32 error correctly
func TestSliceLengthOverflow(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// Codec Version:

// Slice Length:

type MultipleVersionsStruct struct {
	BothTags    string `tag1:"true"  tag2:"true"`
	SingleTag1  string `tag1:"true"`
	SingleTag2  string `             tag2:"true"`
	EitherTags1 string `tag1:"false" tag2:"true"`
	EitherTags2 string `tag1:"true"  tag2:"false"`
	NoTags      string `tag1:"false" tag2:"false"`
}

func TestMultipleTags(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	// received codec is expected to have both v1 and v2 registered as tags
	return
}

func TestMap(t testing.TB, codec codecpkg.GeneralCodec) { _ = "STUB: not implemented"; return }

// data1 and data2 should have the same byte representation even though
// their key-value pairs were defined in a different order.

// Make sure Size returns the correct size for the marshalled data

func TestCanMarshalLargeSlices(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

func FuzzStructUnmarshal(codec codecpkg.GeneralCodec, f *testing.F) {
	_ = "STUB: not implemented"
	return
}

// Register the types that may be unmarshaled into interfaces

func TestImplementsUnmarshalFrom(t testing.TB, codec codecpkg.GeneralCodec) {
	_ = "STUB: not implemented"
	return
}

// pack 3 extra bytes prefix

// pack 3 extra bytes suffix
