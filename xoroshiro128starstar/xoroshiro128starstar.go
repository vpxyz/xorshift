// Package xoroshiro128starstar (XOR/rotate/shift/rotate) all-purpose generator with 128 bits internal state.
package xoroshiro128starstar

import (
	"github.com/vpxyz/xorshift/internal"
)

// XoroShiro128StarStar holds the state required by XoroShiro128StarStar generator
type XoroShiro128StarStar struct {
	// The state must be seeded with a nonzero value. Require 2 64-bit unsigned values.
	// The state must be seeded so that it is not everywhere zero. The state are filled using
	// the SplitMix64 generator with the provvided seed.
	s [2]uint64
}

// NewSource return a new XoroShiro128StarStar random number generator
func NewSource(seed int64) *XoroShiro128StarStar {
	tmpxs := XoroShiro128StarStar{}
	tmpxs.Seed(seed)
	return &tmpxs
}

// Seed use the provvided seed value to init XoroShiro128StarStar internal state.
func (x *XoroShiro128StarStar) Seed(seed int64) {
	tmpxs := internal.SplitMix64{}
	tmpxs.Seed(seed)

	for i := 0; i < len(x.s); i++ {
		x.s[i] = tmpxs.Uint64()

	}
}

// Uint64 returns the next pseudo random number generated, before start you must provvide seed.
func (x *XoroShiro128StarStar) Uint64() uint64 {
	s0, s1 := x.s[0], x.s[1]
	r := internal.Rotl(s0*5, 7) * 9

	s1 ^= s0

	// update the generator state
	x.s[0] = internal.Rotl(s0, 24) ^ s1 ^ (s1 << 16) // a, b
	x.s[1] = internal.Rotl(s1, 37)                   // c

	return r
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (x *XoroShiro128StarStar) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))

}

// Jump it is equivalent to 2^64 calls to Uint64().
func (x *XoroShiro128StarStar) Jump() {
	var s0, s1 uint64
	var b uint64
	jump := []uint64{0xdf900294d8f554a5, 0x170865df4b3201fc}

	for i := 0; i < len(jump); i++ {
		for b = 0; b < 64; b++ {
			if jump[i]&(uint64(1)<<b) != 0 {
				s1 ^= x.s[1]
				s0 ^= x.s[0]
			}
			x.Uint64()
		}
	}

	x.s[1] = s1
	x.s[0] = s0
}
