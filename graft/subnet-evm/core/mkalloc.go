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
// Copyright 2017 The go-ethereum Authors
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
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

//go:build none

/*
The mkalloc tool creates the genesis allocation constants in genesis_alloc.go
It outputs a const declaration that contains an RLP-encoded list of (address, balance) tuples.

	go run mkalloc.go genesis.json
*/
package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/ava-labs/avalanchego/graft/subnet-evm/core"
	"github.com/ava-labs/libevm/common"
)

type allocItem struct {
	Addr    *big.Int
	Balance *big.Int
	Misc    *allocItemMisc `rlp:"optional"`
}

type allocItemMisc struct {
	Nonce uint64
	Code  []byte
	Slots []allocItemStorageItem
}

type allocItemStorageItem struct {
	Key common.Hash
	Val common.Hash
}

func makelist(g *core.Genesis) []allocItem { _ = "STUB: not implemented"; return nil }

func makealloc(g *core.Genesis) string { _ = "STUB: not implemented"; return "" }

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: mkalloc genesis.json")
		os.Exit(1)
	}

	g := new(core.Genesis)
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(file).Decode(g); err != nil {
		panic(err)
	}
	fmt.Println("const allocData =", makealloc(g))
}
