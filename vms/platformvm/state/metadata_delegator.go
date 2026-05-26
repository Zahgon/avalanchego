// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
)

type delegatorMetadata struct {
	PotentialReward uint64 `v1:"true"`
	StakerStartTime uint64 `v1:"true"`

	txID ids.ID
}

func parseDelegatorMetadata(bytes []byte, metadata *delegatorMetadata) error {
	_ = "STUB: not implemented"
	return nil
}

// only potential reward was stored

func writeDelegatorMetadata(db database.KeyValueWriter, metadata *delegatorMetadata, codecVersion uint16) error {
	_ = "STUB: not implemented"
	// The "0" codec is skipped for [delegatorMetadata]. This is to ensure the
	// [validatorMetadata] codec version is the same as the [delegatorMetadata]
	// codec version.
	//
	// TODO: Cleanup post-Durango activation.
	return nil
}
