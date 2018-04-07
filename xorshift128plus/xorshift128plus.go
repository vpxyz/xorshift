/*
Package xorshift128plus is 64-bit version of Saito and Matsumoto's XSadd generator.
It use a 128 bit internal state.
*/
package xorshift128plus

import (
	"github.com/vpxyz/xorshift/internal"
)

// XorShift128Plus holds the state required by XorShift128Plus generator.
type XorShift128Plus struct {
	// The state must be seeded with a nonzero value. Require 2 64-bit unsigned values.
	// The state must be seeded so that it is not everywhere zero. The state are filled using
	// the SplitMix64 generator with the provvided seed.
	s [2]uint64
}

// NewSource return a new SplitMix64 random number generator.
func NewSource(seed int64) *XorShift128Plus {
	tmpxs := XorShift128Plus{}
	tmpxs.Seed(seed)
	return &tmpxs
}

// Seed use the provvided seed value to init XorShift128Plus internal state.
func (x *XorShift128Plus) Seed(seed int64) {
	tmpxs := internal.SplitMix64{}
	tmpxs.Seed(seed)

	for i := 0; i < len(x.s); i++ {
		x.s[i] = tmpxs.Uint64()

	}
}

// Uint64 returns the next pseudo random number generated, before start you must provvide seed.
func (x *XorShift128Plus) Uint64() uint64 {
	s1 := x.s[0]
	s0 := x.s[1]

	s1 ^= s1 << 23

	s1 = s1 ^ s0 ^ (s1 >> 18) ^ (s0 >> 5)

	// update the generator state
	x.s[0] = s0
	x.s[1] = s1

	return s1 + s0 // b, c
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (x *XorShift128Plus) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))
}

// Jump it is equivalent to 2^64 calls to Uint64().
func (x *XorShift128Plus) Jump() {
	var s0, s1 uint64 = 0, 0
	var b uint64

	for i := 0; i < len(internal.Jump128); i++ {
		for b = 0; b < 64; b++ {
			if internal.Jump128[i]&uint64(1)<<b != 0 {
				s0 ^= x.s[0]
				s1 ^= x.s[1]
			}
			x.Uint64()
		}
	}

	x.s[0] = s0
	x.s[1] = s1
}
