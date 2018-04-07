// Package splitmix64 generator is a fixed-increment version of Java 8's SplittableRandom generator.
package splitmix64

import (
	"github.com/vpxyz/xorshift/internal"
)

// SplitMix64 hold the state required by the SplitMix64 generator.
type SplitMix64 struct {
	is internal.SplitMix64
}

// NewSource return a new SplitMix64 random number generator
func NewSource(seed int64) *SplitMix64 { // rand.Source64 {
	tmpxs := SplitMix64{}
	tmpxs.Seed(seed)
	return &tmpxs
}

// Seed seed SplitMix64 random number generator with the given value.
func (x *SplitMix64) Seed(seed int64) {
	x.is.Seed(seed)
}

// Uint64 returns the next pseudo random number generated, before start you must provvide one 64 unsigned bit seed.
func (x *SplitMix64) Uint64() uint64 {
	return x.is.Uint64()
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (x *SplitMix64) Int63() int64 {
	return x.is.Int63()
}
