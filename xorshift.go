/*
Package xorshift implements a simple library for pseudo random number generators based on xorshif* and xorshift+ .

xorshift* generators are obtained by scrambling the output of a Marsaglia xorshift generator with a 64-bit invertible multiplier.
xorshift+ generators are a 64-bit version of Saito and Matsumoto's XSadd generator.
This simple library in based on the work of Sebastiano Vigna (http://xorshift.di.unimi.it/).

The usage are very simple: just fill the seed with a nonzero value and call the Next() or SyncNext() function.

Example:

var xs XorShift64

xs.S = 43242434343434223

r := xs.Next()

*/
package xorshift

/*
XorShift64Star hold the state required by the XorShift64Star generators.
*/
type XorShift64Star struct {
	s uint64 // The state must be seeded with a nonzero value. Require a 64-bit unsigned values.
}

/*
XorShift128Plus holds the state required by XorShift128Plus generator.
*/
type XorShift128Plus struct {
	// The state must be seeded with a nonzero value. Require 2 64-bit unsigned values.
	// The state must be seeded so that it is not everywhere zero. If you have a 64-bit seed,
	// we suggest to seed a xorshift64* generator and use its output to fill S.
	s [2]uint64
}

/*
XorShift1024Star holds the state required by XorShift1024Star generator.
*/
type XorShift1024Star struct {
	// The state must be seeded with a nonzero value. Require 16 64-bit unsigned values.
	// The state must be seeded so that it is not everywhere zero. If you have a 64-bit seed,
	// we suggest to seed a xorshift64* generator and use its output to fill s .
	s [16]uint64
	p int
}

/*
XorShift4096Star holds the state required by XorShift4096Star generator.
*/
type XorShift4096Star struct {
	// The state must be seeded with a nonzero value. Require 64 64-bit unsigned values.
	// The state must be seeded so that it is not everywhere zero. If you have a 64-bit seed,
	// we suggest to seed a xorshift64* generator and use its output to fill s .
	s [64]uint64
	p int
}

/*
Next returns the next pseudo random number generated, before start you must provvide one 64 unsigned bit seed.
*/
func (x *XorShift64Star) Next() uint64 {
	s := x.s
	s ^= s >> 12
	s ^= s << 25
	s ^= s >> 27
	x.s = s

	return s * 2685821657736338717
}

/*
Init returns a new XorShift64Star source seeded with the given value.
*/
func (x *XorShift64Star) Init(seed uint64) {
	x.s = seed
}

/*
Next returns the next pseudo random number generated, before start you must provvide seed.
*/
func (x *XorShift128Plus) Next() uint64 {
	s1 := x.s[0]
	s0 := x.s[1]

	s1 ^= s1 << 23

	s1 = s1 ^ s0 ^ (s1 >> 18) ^ (s0 >> 5)

	// update the state of generator
	x.s[0] = s0
	x.s[1] = s1

	return s1 + s0 // b, c
}

/*
Init returns a new XorShift128Plus source seeded with a slice of 2 values.
*/
func (x *XorShift128Plus) Init(seed []uint64) {
	if len(seed) > 1 {
		x.s[0], x.s[1] = seed[0], seed[1]
		return
	}
	x.s[0] = seed[0]

}

/*
Next returns the next pseudo random number generated, before start you must provvide seed.
*/
func (x *XorShift1024Star) Next() uint64 {
	s0 := x.s[x.p]

	xpnew := (x.p + 1) & 15

	s1 := x.s[xpnew]

	s1 ^= s1 << 31 // a
	tmp := s1 ^ s0 ^ (s1 >> 11) ^ (s0 >> 30)

	// update the state of generator
	x.s[xpnew] = tmp
	x.p = xpnew

	return tmp * 1181783497276652981
}

/*
Init returns a new XorShift1024Star source seeded with a slice of 16 values.
*/
func (x *XorShift1024Star) Init(seed []uint64) {
	for i, v := range seed {
		if i < len(x.s) {
			x.s[i] = v
		}
	}
	x.p = 0
}

/*
Next returns the next pseudo random number generated, before start you must provvide seed.
*/
func (x *XorShift4096Star) Next() uint64 {
	xpnew := (x.p + 1) & 63
	s0 := x.s[x.p]
	s1 := x.s[xpnew]

	s1 ^= s1 << 25 // a
	s1 ^= s1 >> 3  // b
	s0 ^= s0 >> 49 // c

	tmp := s0 ^ s1

	// update the state of generator
	x.s[xpnew] = tmp
	x.p = xpnew

	return tmp * 8372773778140471301
}

/*
Init returns a new XorShift4096Star source seeded with a slicef of 64 values.
*/
func (x *XorShift4096Star) Init(seed []uint64) {
	for i, v := range seed {
		if i < len(x.s) {
			x.s[i] = v
		}
	}
	x.p = 0
}
