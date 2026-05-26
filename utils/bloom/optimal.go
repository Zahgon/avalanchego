// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package bloom

import "math"

const ln2Squared = math.Ln2 * math.Ln2

// OptimalParameters calculates the optimal [numHashes] and [numEntries] that
// should be allocated for a bloom filter which will contain [count] and target
// [falsePositiveProbability].
func OptimalParameters(count int, falsePositiveProbability float64) (int, int) {
	_ = "STUB: not implemented"
	return 0, 0
}

// OptimalHashes calculates the number of hashes which will minimize the false
// positive probability of a bloom filter with [numEntries] after [count]
// additions.
//
// It is guaranteed to return a value in the range [minHashes, maxHashes].
//
// ref: https://en.wikipedia.org/wiki/Bloom_filter
func OptimalHashes(numEntries, count int) int { _ = "STUB: not implemented"; return 0 }

// Converting a floating-point value to an int produces an undefined value
// if the floating-point value cannot be represented as an int. To avoid
// this undefined behavior, we explicitly check against MaxInt here.
//
// ref: https://go.dev/ref/spec#Conversions

// OptimalEntries calculates the optimal number of entries to use when creating
// a new Bloom filter when targenting a size of [count] with
// [falsePositiveProbability] assuming that the optimal number of hashes is
// used.
//
// It is guaranteed to return a value in the range [minEntries, MaxInt].
//
// ref: https://en.wikipedia.org/wiki/Bloom_filter
func OptimalEntries(count int, falsePositiveProbability float64) int {
	_ = "STUB: not implemented"
	return 0
}

// Converting a floating-point value to an int produces an undefined value
// if the floating-point value cannot be represented as an int. To avoid
// this undefined behavior, we explicitly check against MaxInt here.
//
// ref: https://go.dev/ref/spec#Conversions

// EstimateCount estimates the number of additions a bloom filter with
// [numHashes] and [numEntries] must have to reach [falsePositiveProbability].
// This is derived by inversing a lower-bound on the probability of false
// positives. For values where numBits >> numHashes, the predicted probability
// is fairly accurate.
//
// It is guaranteed to return a value in the range [0, MaxInt].
//
// ref: https://tsapps.nist.gov/publication/get_pdf.cfm?pub_id=903775
func EstimateCount(numHashes, numEntries int, falsePositiveProbability float64) int {
	_ = "STUB: not implemented"
	return 0
}

// Converting a floating-point value to an int produces an undefined value
// if the floating-point value cannot be represented as an int. To avoid
// this undefined behavior, we explicitly check against MaxInt here.
//
// ref: https://go.dev/ref/spec#Conversions
