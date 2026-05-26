// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package uptimetracker

import (
	"fmt"
	"math"
	"time"

	"github.com/ava-labs/avalanchego/codec"
	"github.com/ava-labs/avalanchego/codec/linearcodec"
	"github.com/ava-labs/avalanchego/database"
	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/snow/uptime"
	"github.com/ava-labs/avalanchego/utils/set"
)

const codecVersion uint16 = 0

var (
	codecManager codec.Manager
	_            uptime.State = (*state)(nil)
)

func init() {
	codecManager = codec.NewManager(math.MaxInt32)
	c := linearcodec.NewDefault()

	if err := c.RegisterType(validator{}); err != nil {
		panic(fmt.Errorf("failed to register type: %w", err))
	}

	if err := codecManager.RegisterCodec(codecVersion, c); err != nil {
		panic(fmt.Errorf("failed to register codec: %w", err))
	}
}

type validator struct {
	UpDuration    time.Duration `serialize:"true"`
	LastUpdated   uint64        `serialize:"true"`
	NodeID        ids.NodeID    `serialize:"true"`
	Weight        uint64        `serialize:"true"`
	StartTime     uint64        `serialize:"true"`
	IsActive      bool          `serialize:"true"`
	IsL1Validator bool          `serialize:"true"`

	validationID ids.ID
}

// state holds the on-disk and cached representation of the validator set
type state struct {
	db database.Database

	validators             map[ids.ID]*validator
	nodeIDsToValidationIDs map[ids.NodeID]ids.ID
	updatedValidators      set.Set[ids.ID]
	deletedValidators      set.Set[ids.ID]
}

func newState(db database.Database) (*state, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *state) GetUptime(nodeID ids.NodeID) (time.Duration, time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Time), nil
}

func (s *state) SetUptime(
	nodeID ids.NodeID,
	upDuration time.Duration,
	lastUpdated time.Time,
) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *state) GetStartTime(nodeID ids.NodeID) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// addNewValidator adds a new validator to the state and marks it for
// persistence to the database. This should be used when adding validators
// during runtime operations.
func (s *state) addNewValidator(vdr *validator) { _ = "STUB: not implemented"; return }

// updateValidator sets the isActive state of the validator with the given
// validationID -- this function assumes that a validator with the given
// validationID exists.
func (s *state) updateValidator(validationID ids.ID, isActive bool) bool {
	_ = "STUB: not implemented"
	return false
}

// deleteValidator deletes the validator with the given validationID -- this
// function assumes that a validator with the given validationID exists.
func (s *state) deleteValidator(validationID ids.ID) { _ = "STUB: not implemented"; return }

func (s *state) writeModifications() error { _ = "STUB: not implemented"; return nil }

// We have written all pending updates

// addValidatorToMemory adds a validator to the in-memory data structures only.
// This is used during initialization when loading from the database and does not
// mark the validator for persistence.
func (s *state) addValidatorToMemory(validationID ids.ID, validator *validator) {
	_ = "STUB: not implemented"
	return
}

func (s *state) getValidatorByNodeID(nodeID ids.NodeID) (*validator, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// we are guaranteed to have this validator

func (s *state) hasValidationID(validationID ids.ID) bool { _ = "STUB: not implemented"; return false }

func (s *state) getNodeID(validationID ids.ID) (ids.NodeID, bool) {
	_ = "STUB: not implemented"
	return *new(ids.NodeID), false
}

func (s *state) getNodeIDs() []ids.NodeID { _ = "STUB: not implemented"; return nil }
