// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package fee

import (
	"errors"

	"github.com/ava-labs/avalanchego/vms/components/gas"
	"github.com/ava-labs/avalanchego/vms/platformvm/txs"
)

var (
	_ Calculator = (*dynamicCalculator)(nil)

	ErrCalculatingComplexity = errors.New("error calculating complexity")
	ErrCalculatingGas        = errors.New("error calculating gas")
	ErrCalculatingCost       = errors.New("error calculating cost")
)

func NewDynamicCalculator(
	weights gas.Dimensions,
	price gas.Price,
) Calculator {
	_ = "STUB: not implemented"
	return *new(Calculator)
}

type dynamicCalculator struct {
	weights gas.Dimensions
	price   gas.Price
}

func (c *dynamicCalculator) CalculateFee(tx txs.UnsignedTx) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
