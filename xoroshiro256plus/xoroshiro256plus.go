// Package xoroshiro256plus (XOR/rotate/shift/rotate) with 256 bits internal state, fast generator for floating-point numbers.
package xoroshiro256plus

import (
	"github.com/vpxyz/xorshift/internal"
)

// XoroShiro256Plus holds the state required by XoroShiro256Plus generator
type XoroShiro256Plus struct {
	// The state must be seeded with a nonzero value. Require 2 64-bit unsigned values.
	// The state must be seeded so that it is not everywhere zero. The state are filled using
	// the SplitMix64 generator with the provvided seed.
	s [4]uint64
}

// NewSource return a new XoroShiro256Plus random number generator
func NewSource(seed int64) *XoroShiro256Plus {
	tmpxs := XoroShiro256Plus{}
	tmpxs.Seed(seed)
	return &tmpxs
}

// Seed use the provvided seed value to init XoroShiro256Plus internal state.
func (x *XoroShiro256Plus) Seed(seed int64) {
	tmpxs := internal.SplitMix64{}
	tmpxs.Seed(seed)

	for i := 0; i < len(x.s); i++ {
		x.s[i] = tmpxs.Uint64()

	}
}

// Uint64 returns the next pseudo random number generated, before start you must provvide seed.
func (x *XoroShiro256Plus) Uint64() uint64 {
	// Yeah, I know that I can use an array, but the Go compiler isn't smart as gcc, the generate code are slower.
	s0, s1, s2, s3 := x.s[0], x.s[1], x.s[2], x.s[3]

	x.s[0] = s0 ^ s3 ^ s1
	x.s[1] = s1 ^ s2 ^ s0
	x.s[2] = s2 ^ s0 ^ (s1 << 17)
	x.s[3] = internal.Rotl(s3^s1, 45)

	return s0 + s3
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (x *XoroShiro256Plus) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))

}

// Jump it is equivalent to 2^128 calls to Uint64().
func (x *XoroShiro256Plus) Jump() {
	var s0, s1, s2, s3 uint64
	var b uint64
	jump := []uint64{0x180ec6d33cfd0aba, 0xd5a61266f0c9392c, 0xa9582618e03fc9aa, 0x39abdc4529b1661c}

	for i := 0; i < len(jump); i++ {
		for b = 0; b < 64; b++ {
			if jump[i]&(uint64(1)<<b) != 0 {
				s3 ^= x.s[3]
				s2 ^= x.s[2]
				s1 ^= x.s[1]
				s0 ^= x.s[0]
			}
			x.Uint64()
		}
	}

	x.s[3] = s3
	x.s[2] = s2
	x.s[1] = s1
	x.s[0] = s0
}
