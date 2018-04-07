/*
Package xorshift4096star it's like other xorshift*, generators, but with a 4096 bit internal state.
*/
package xorshift4096star

import (
	"github.com/vpxyz/xorshift/internal"
)

// XorShift4096Star holds the state required by XorShift4096Star generator.
type XorShift4096Star struct {
	// The state must be seeded with a nonzero value. Require 64 64-bit unsigned values.
	// The state must be seeded so that it is not everywhere zero. The state are filled using
	// the SplitMix64 generator with the provvided seed.
	s [64]uint64
	p int
}

// NewSource return a new XorShift4096Star random number generator.
func NewSource(seed int64) *XorShift4096Star {
	tmpxs := XorShift4096Star{}
	tmpxs.Seed(seed)
	return &tmpxs
}

// Seed returns a new XorShift4096Star source seeded with a slice of 16 values.
func (x *XorShift4096Star) Seed(seed int64) {
	tmpxs := internal.SplitMix64{}
	tmpxs.Seed(seed)

	for i := 0; i < len(x.s); i++ {
		x.s[i] = tmpxs.Uint64()

	}
	x.p = 0
}

// Uint64 returns the next pseudo random number generated, before start you must provvide seed.
func (x *XorShift4096Star) Uint64() uint64 {
	xpnew := (x.p + 1) & 63
	s0 := x.s[x.p]
	s1 := x.s[xpnew]

	s1 ^= s1 << 25 // a
	s1 ^= s1 >> 3  // b
	s0 ^= s0 >> 49 // c

	tmp := s0 ^ s1

	// update the generator state
	x.s[xpnew] = tmp
	x.p = xpnew

	return tmp * uint64(8372773778140471301)
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (x *XorShift4096Star) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))
}
