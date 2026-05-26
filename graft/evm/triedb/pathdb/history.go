// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2022 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>

package pathdb

import (
	"github.com/ava-labs/libevm/common"
	"github.com/ava-labs/libevm/trie/triestate"
)

// State history records the state changes involved in executing a block. The
// state can be reverted to the previous version by applying the associated
// history object (state reverse diff). State history objects are kept to
// guarantee that the system can perform state rollbacks in case of deep reorg.
//
// Each state transition will generate a state history object. Note that not
// every block has a corresponding state history object. If a block performs
// no state changes whatsoever, no state is created for it. Each state history
// will have a sequentially increasing number acting as its unique identifier.
//
// The state history is written to disk (ancient store) when the corresponding
// diff layer is merged into the disk layer. At the same time, system can prune
// the oldest histories according to config.
//
//                                                        Disk State
//                                                            ^
//                                                            |
//   +------------+     +---------+     +---------+     +---------+
//   | Init State |---->| State 1 |---->|   ...   |---->| State n |
//   +------------+     +---------+     +---------+     +---------+
//
//                     +-----------+      +------+     +-----------+
//                     | History 1 |----> | ...  |---->| History n |
//                     +-----------+      +------+     +-----------+
//
// # Rollback
//
// If the system wants to roll back to a previous state n, it needs to ensure
// all history objects from n+1 up to the current disk layer are existent. The
// history objects are applied to the state in reverse order, starting from the
// current disk layer.

const (
	accountIndexSize = common.AddressLength + 13 // The length of encoded account index
	slotIndexSize    = common.HashLength + 5     // The length of encoded slot index
	historyMetaSize  = 9 + 2*common.HashLength   // The length of fixed size part of meta object

	stateHistoryVersion = uint8(0) // initial version of state history structure.
)

// Each state history entry is consisted of five elements:
//
// # metadata
//  This object contains a few meta fields, such as the associated state root,
//  block number, version tag and so on. This object may contain an extra
//  accountHash list which means the storage changes belong to these accounts
//  are not complete due to large contract destruction. The incomplete history
//  can not be used for rollback and serving archive state request.
//
// # account index
//  This object contains some index information of account. For example, offset
//  and length indicate the location of the data belonging to the account. Besides,
//  storageOffset and storageSlots indicate the storage modification location
//  belonging to the account.
//
//  The size of each account index is *fixed*, and all indexes are sorted
//  lexicographically. Thus binary search can be performed to quickly locate a
//  specific account.
//
// # account data
//  Account data is a concatenated byte stream composed of all account data.
//  The account data can be solved by the offset and length info indicated
//  by corresponding account index.
//
//            fixed size
//         ^             ^
//        /               \
//        +-----------------+-----------------+----------------+-----------------+
//        | Account index 1 | Account index 2 |       ...      | Account index N |
//        +-----------------+-----------------+----------------+-----------------+
//        |
//        |     length
// offset |----------------+
//        v                v
//        +----------------+----------------+----------------+----------------+
//        | Account data 1 | Account data 2 |       ...      | Account data N |
//        +----------------+----------------+----------------+----------------+
//
// # storage index
//  This object is similar with account index. It's also fixed size and contains
//  the location info of storage slot data.
//
// # storage data
//  Storage data is a concatenated byte stream composed of all storage slot data.
//  The storage slot data can be solved by the location info indicated by
//  corresponding account index and storage slot index.
//
//                    fixed size
//                 ^             ^
//                /               \
//                +-----------------+-----------------+----------------+-----------------+
//                | Account index 1 | Account index 2 |       ...      | Account index N |
//                +-----------------+-----------------+----------------+-----------------+
//                |
//                |                    storage slots
// storage offset |-----------------------------------------------------+
//                v                                                     v
//                +-----------------+-----------------+-----------------+
//                | storage index 1 | storage index 2 | storage index 3 |
//                +-----------------+-----------------+-----------------+
//                |     length
//         offset |-------------+
//                v             v
//                +-------------+
//                | slot data 1 |
//                +-------------+

// accountIndex describes the metadata belonging to an account.
type accountIndex struct {
	address       common.Address // The address of account
	length        uint8          // The length of account data, size limited by 255
	offset        uint32         // The offset of item in account data table
	storageOffset uint32         // The offset of storage index in storage index table
	storageSlots  uint32         // The number of mutated storage slots belonging to the account
}

// encode packs account index into byte stream.
func (i *accountIndex) encode() []byte { _ = "STUB: not implemented"; return nil }

// decode unpacks account index from byte stream.
func (i *accountIndex) decode(blob []byte) { _ = "STUB: not implemented"; return }

// slotIndex describes the metadata belonging to a storage slot.
type slotIndex struct {
	hash   common.Hash // The hash of slot key
	length uint8       // The length of storage slot, up to 32 bytes defined in protocol
	offset uint32      // The offset of item in storage slot data table
}

// encode packs slot index into byte stream.
func (i *slotIndex) encode() []byte { _ = "STUB: not implemented"; return nil }

// decode unpack slot index from the byte stream.
func (i *slotIndex) decode(blob []byte) { _ = "STUB: not implemented"; return }

// meta describes the meta data of state history object.
type meta struct {
	version    uint8            // version tag of history object
	parent     common.Hash      // prev-state root before the state transition
	root       common.Hash      // post-state root after the state transition
	block      uint64           // associated block number
	incomplete []common.Address // list of address whose storage set is incomplete
}

// encode packs the meta object into byte stream.
func (m *meta) encode() []byte { _ = "STUB: not implemented"; return nil }

// decode unpacks the meta object from byte stream.
func (m *meta) decode(blob []byte) error { _ = "STUB: not implemented"; return nil }

// history represents a set of state changes belong to a block along with
// the metadata including the state roots involved in the state transition.
// State history objects in disk are linked with each other by a unique id
// (8-bytes integer), the oldest state history object can be pruned on demand
// in order to control the storage size.
type history struct {
	meta        *meta                                     // Meta data of history
	accounts    map[common.Address][]byte                 // Account data keyed by its address hash
	accountList []common.Address                          // Sorted account hash list
	storages    map[common.Address]map[common.Hash][]byte // Storage data keyed by its address hash and slot hash
	storageList map[common.Address][]common.Hash          // Sorted slot hash list
}

// newHistory constructs the state history object with provided state change set.
func newHistory(root common.Hash, parent common.Hash, block uint64, states *triestate.Set) *history {
	_ = "STUB: not implemented"
	return nil
}

// encode serializes the state history and returns four byte streams represent
// concatenated account/storage data, account/storage indexes respectively.
func (h *history) encode() ([]byte, []byte, []byte, []byte) {
	_ = "STUB: not implemented"
	return nil,

		// the number of processed slots
		nil, nil, nil
}

// the buffer for concatenated account data
// the buffer for concatenated storage data
// the buffer for concatenated account index
// the buffer for concatenated storage index

// Encode storage slots in order

// Fill up the storage meta in account index

// decoder wraps the byte streams for decoding with extra meta fields.
type decoder struct {
	accountData    []byte // the buffer for concatenated account data
	storageData    []byte // the buffer for concatenated storage data
	accountIndexes []byte // the buffer for concatenated account index
	storageIndexes []byte // the buffer for concatenated storage index

	lastAccount       *common.Address // the address of last resolved account
	lastAccountRead   uint32          // the read-cursor position of account data
	lastSlotIndexRead uint32          // the read-cursor position of storage slot index
	lastSlotDataRead  uint32          // the read-cursor position of storage slot data
}

// verify validates the provided byte streams for decoding state history. A few
// checks will be performed to quickly detect data corruption. The byte stream
// is regarded as corrupted if:
//
// - account indexes buffer is empty(empty state set is invalid)
// - account indexes/storage indexer buffer is not aligned
//
// note, these situations are allowed:
//
// - empty account data: all accounts were not present
// - empty storage set: no slots are modified
func (r *decoder) verify() error { _ = "STUB: not implemented"; return nil }

// readAccount parses the account from the byte stream with specified position.
func (r *decoder) readAccount(pos int) (accountIndex, []byte, error) {
	_ = "STUB: not implemented"
	// Decode account index from the index byte stream.
	return *new(accountIndex), nil, nil
}

// Perform validation before parsing account data, ensure
// - account is sorted in order in byte stream
// - account data is strictly encoded with no gap inside
// - account data is not out-of-slice
// zero address is possible

// readStorage parses the storage slots from the byte stream with specified account.
func (r *decoder) readStorage(accIndex accountIndex) ([]common.Hash, map[common.Hash][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Perform validation before parsing storage slot data, ensure
// - slot index is not out-of-slice
// - slot data is not out-of-slice
// - slot is sorted in order in byte stream
// - slot indexes is strictly encoded with no gap inside
// - slot data is strictly encoded with no gap inside

// decode deserializes the account and storage data from the provided byte stream.
func (h *history) decode(accountData, storageData, accountIndexes, storageIndexes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Resolve account first

// Resolve storage slots
