// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowball

import "github.com/ava-labs/avalanchego/ids"

var (
	SnowballFactory  Factory = snowballFactory{}
	SnowflakeFactory Factory = snowflakeFactory{}
)

type snowballFactory struct{}

func (snowballFactory) NewNnary(params Parameters, choice ids.ID) Nnary {
	_ = "STUB: not implemented"
	return *new(Nnary)
}

func (snowballFactory) NewUnary(params Parameters) Unary {
	_ = "STUB: not implemented"
	return *new(Unary)
}

type snowflakeFactory struct{}

func (snowflakeFactory) NewNnary(params Parameters, choice ids.ID) Nnary {
	_ = "STUB: not implemented"
	return *new(Nnary)
}

func (snowflakeFactory) NewUnary(params Parameters) Unary {
	_ = "STUB: not implemented"
	return *new(Unary)
}
