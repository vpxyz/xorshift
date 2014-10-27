/*
Package xorshift implements a simple library for pseudo random number generators based on xorshif* and xorshift+ .

xorshift* generators are obtained by scrambling the output of a Marsaglia xorshift generator with a 64-bit invertible multiplier.
xorshift+ generators are a 64-bit version of Saito and Matsumoto's XSadd generator.
This simple library in based on the work of Sebastiano Vigna (http://xorshift.di.unimi.it/).

The usage are very simple: just fill the seed with a nonzero value and call the Next() function.

Example:

var xs XorShift64

xs.S = 43242434343434223

r := xs.Next()

*/
package xorshit

/*
XorShift64Star hold the state required by the XorShift64Star generators.
*/
type XorShift64Star struct {
	S uint64 // The state must be seeded with a nonzero value. Require a 64-bit unsigned values.
}

/*
XorShift128Plus holds the state required by XorShift128Plus generator.
*/
type XorShift128Plus struct {
	// The state must be seeded with a nonzero value. Require 2 64-bit unsigned values. The state must be seeded so that it is not everywhere zero. If you have a 64-bit seed,  we suggest to seed a xorshift64* generator and use its output to fill S.
	S [2]uint64
}

/*
XorShift1024Star holds the state required by XorShift1024Star generator.
*/
type XorShift1024Star struct {
	// The state must be seeded with a nonzero value. Require 16 64-bit unsigned values. The state must be seeded so that it is not everywhere zero. If you have a 64-bit seed,  we suggest to seed a xorshift64* generator and use its output to fill s .
	S [16]uint64
	p int
}

/*
XorShift4096Star holds the state required by XorShift4096Star generator.
*/
type XorShift4096Star struct {
	// The state must be seeded with a nonzero value. Require 64 64-bit unsigned values. The state must be seeded so that it is not everywhere zero. If you have a 64-bit seed,  we suggest to seed a xorshift64* generator and use its output to fill s .
	S [64]uint64
	p int
}

/*
Next return the next pseudo random number generated, before start you must provvide one 64 unsigned bit seed.
*/
func (x *XorShift64Star) Next() uint64 {
	x.S ^= x.S >> 12
	x.S ^= x.S << 25
	x.S ^= x.S >> 27

	return x.S * 2685821657736338717
}

/*
Next return the next pseudo random number generated, before start you must provvide seed.
*/
func (x *XorShift128Plus) Next() uint64 {
	s1 := x.S[0]
	s0 := x.S[1]

	s1 ^= s1 << 23

	// update the state of generator
	x.S[0] = s0
	x.S[1] = (s1 ^ s0 ^ (s1 >> 17) ^ (s0 >> 26))

	return x.S[1] + s0 // b, c
}

/*
Next return the next pseudo random number generated, before start you must provvide seed.
*/
func (x *XorShift1024Star) Next() uint64 {
	s0 := x.S[x.p]

	xpnew := (x.p + 1) & 15

	s1 := x.S[xpnew]

	s1 ^= s1 << 31 // a
	s1 ^= s1 >> 11 // b
	s0 ^= s0 >> 30 // c

	// update the state of generator
	x.S[xpnew] = (s0 ^ s1)
	x.p = xpnew

	return x.S[xpnew] * 1181783497276652981
}

/*
Next return the next pseudo random number generated, before start you must provvide seed.
*/
func (x *XorShift4096Star) Next() uint64 {
	xpnew := (x.p + 1) & 63
	s0 := x.S[x.p]
	s1 := x.S[xpnew]

	s1 ^= s1 << 25 // a
	s1 ^= s1 >> 3  // b
	s0 ^= s0 >> 49 // c

	// update the state of generator
	x.p = xpnew
	x.S[xpnew] = s0 ^ s1

	return x.S[xpnew] * 8372773778140471301
}
