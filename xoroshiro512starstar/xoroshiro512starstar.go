// Package xoroshiro512starstar (XOR/rotate/shift/rotate) all-purpose generator with internal 512 bits state.
package xoroshiro512starstar

import (
	"math/bits"

	"github.com/vpxyz/xorshift/internal"
)

// XoroShiro512StarStar holds the state required by XoroShiro512StarStar generator
type XoroShiro512StarStar struct {
	// The state must be seeded with a nonzero value. Require 2 64-bit unsigned values.
	// The state must be seeded so that it is not everywhere zero. The state are filled using
	// the SplitMix64 generator with the provvided seed.
	s [8]uint64
}

// NewSource return a new XoroShiro512StarStar random number generator
func NewSource(seed int64) *XoroShiro512StarStar {
	tmpxs := XoroShiro512StarStar{}
	tmpxs.Seed(seed)
	return &tmpxs
}

// Seed use the provvided seed value to init XoroShiro512StarStar internal state.
func (x *XoroShiro512StarStar) Seed(seed int64) {
	tmpxs := internal.SplitMix64{}
	tmpxs.Seed(seed)

	for i := 0; i < len(x.s); i++ {
		x.s[i] = tmpxs.Uint64()

	}
}

// Uint64 returns the next pseudo random number generated, before start you must provvide seed.
func (x *XoroShiro512StarStar) Uint64() uint64 {
	// Yeah, I know that I can use an array, but the Go compiler isn't smart as gcc, the generate code are slower.
	s0, s1, s2, s3, s4, s5, s6, s7 := x.s[0], x.s[1], x.s[2], x.s[3], x.s[4], x.s[5], x.s[6], x.s[7]

	x.s[0] = s0 ^ s6
	x.s[1] = s1 ^ s2 ^ s0
	x.s[2] = s2 ^ s0
	x.s[3] = s3 ^ s4
	x.s[4] = s4 ^ s5 ^ s1
	x.s[5] = s5 ^ s1
	x.s[6] = (s1 << 11) ^ s6 ^ ^s7 ^ s3
	x.s[7] = bits.RotateLeft64(s7^s3, 21)

	return bits.RotateLeft64(s1*5, 7) * 9
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (x *XoroShiro512StarStar) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))

}

// Jump it is equivalent to 2^256 calls to Uint64().
func (x *XoroShiro512StarStar) Jump() {
	var s [8]uint64
	var b uint64
	jump := []uint64{0x33ed89b6e7a353f9, 0x760083d7955323be, 0x2837f2fbb5f22fae, 0x4b8c5674d309511c,
		0xb11ac47a7ba28c25, 0xf1be7667092bcc1c, 0x53851efdb6df0aaf, 0x1ebbc8b23eaf25db}

	for i := 0; i < len(jump); i++ {
		for b = 0; b < 64; b++ {
			if jump[i]&(uint64(1)<<b) != 0 {
				s[7] ^= x.s[7]
				s[6] ^= x.s[6]
				s[5] ^= x.s[5]
				s[4] ^= x.s[4]
				s[3] ^= x.s[3]
				s[2] ^= x.s[2]
				s[1] ^= x.s[1]
				s[0] ^= x.s[0]
			}
			x.Uint64()
		}
	}
	x.s = s
}
