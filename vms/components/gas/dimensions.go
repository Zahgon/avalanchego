// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gas

const (
	Bandwidth Dimension = iota
	DBRead
	DBWrite // includes deletes
	Compute

	NumDimensions = iota
)

type (
	Dimension  uint
	Dimensions [NumDimensions]uint64
)

// Add returns d + sum(os...).
//
// If overflow occurs, an error is returned.
func (d Dimensions) Add(os ...*Dimensions) (Dimensions, error) {
	_ = "STUB: not implemented"
	return *new(Dimensions), nil
}

// Sub returns d - sum(os...).
//
// If underflow occurs, an error is returned.
func (d Dimensions) Sub(os ...*Dimensions) (Dimensions, error) {
	_ = "STUB: not implemented"
	return *new(Dimensions), nil
}

// ToGas returns d · weights.
//
// If overflow occurs, an error is returned.
func (d Dimensions) ToGas(weights Dimensions) (Gas, error) {
	_ = "STUB: not implemented"
	return *new(Gas), nil
}
