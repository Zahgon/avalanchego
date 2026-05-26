// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bind

import (
	"github.com/ava-labs/libevm/accounts/abi"
)

type (
	// These types are exported for use in bind/precompilebind
	TmplContract = tmplContract
	TmplMethod   = tmplMethod
	TmplStruct   = tmplStruct
)

// BindHook is a callback function that can be used to customize the binding.
type BindHook func(lang Lang, pkg string, types []string, contracts map[string]*tmplContract, structs map[string]*tmplStruct) (data any, templateSource string, err error)

func IsKeyWord(arg string) bool { _ = "STUB: not implemented"; return false }

var bindTypeNew = map[Lang]func(kind abi.Type, structs map[string]*tmplStruct) string{
	LangGo: bindTypeNewGo,
}

// bindTypeNewGo converts new types to Go ones.
func bindTypeNewGo(kind abi.Type, structs map[string]*tmplStruct) string {
	_ = "STUB: not implemented"
	return ""
}

func mkList(args ...any) []any { _ = "STUB: not implemented"; return nil }

func add(a, b int) int { _ = "STUB: not implemented"; return 0 }
