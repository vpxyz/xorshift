/*
Package xorshift1024star it's like other xorshift*, generators, but with a 2^1024 -1 period.
*/
package xorshift1024star

import (
	"github.com/vpxyz/xorshift/internal"
)

// XorShift1024Star holds the state required by XorShift1024Star generator.
type XorShift1024Star struct {
	// The state must be seeded with a nonzero value. Require 16 64-bit unsigned values.
	// The state must be seeded so that it is not everywhere zero. If you have a 64-bit seed,
	// we suggest to seed a xorshift64* generator and use its output to fill s .
	s [16]uint64
	p int
}

// NewSource return a new XorShift1024Star random number generator
func NewSource(seed int64) *XorShift1024Star {
	tmpxs := XorShift1024Star{}
	tmpxs.Seed(seed)
	return &tmpxs
}

// Seed use the provvided seed value to init XorShift4096Star internal state.
func (x *XorShift1024Star) Seed(seed int64) {
	tmpxs := internal.SplitMix64{}
	tmpxs.Seed(seed)

	for i := 0; i < len(x.s); i++ {
		x.s[i] = tmpxs.Uint64()

	}
	x.p = 0
}

// Uint64 returns the next pseudo random number generated, before start you must provvide seed.
func (x *XorShift1024Star) Uint64() uint64 {
	xpnew := (x.p + 1) & 15
	s0 := x.s[x.p]
	s1 := x.s[xpnew]

	s1 ^= s1 << 31 // a
	tmp := s1 ^ s0 ^ (s1 >> 11) ^ (s0 >> 30)

	// update the generator state
	x.s[xpnew] = tmp
	x.p = xpnew

	return tmp * uint64(1181783497276652981)
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (x *XorShift1024Star) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))
}

// Jump function for the generator. It is equivalent to 2^512 calls to  Uint64()
func (x *XorShift1024Star) Jump() {
	var t [16]uint64
	var b uint64

	for i := 0; i < len(internal.Jump1024); i++ {
		for b = 0; b < 64; b++ {
			if internal.Jump1024[i]&(uint64(1)<<b) != 0 {
				for j := 0; j < 16; j++ {
					t[j] ^= x.s[(j+x.p)&15]
				}
			}
			x.Uint64()
		}
	}

	for j := 0; j < 16; j++ {
		x.s[(j+x.p)&15] = t[j]
	}
}
