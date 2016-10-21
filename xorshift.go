/*
Package xorshift implements a simple library for pseudo random number generators based on xorshif* and xorshift+ .

xorshift* generators are obtained by scrambling the output of a Marsaglia xorshift generator with a 64-bit invertible multiplier.
xorshift+ generators are a 64-bit version of Saito and Matsumoto's XSadd generator.
This simple library in based on the work of Sebastiano Vigna (http://xorshift.di.unimi.it/).

The usage are very simple: just fill the seed with a nonzero value and call the Next() or SyncNext() function.

NOTE:Not concurrency-safe! You must wrap into monitor goroutine or a mutex.

Example:

var xs XorShift64

xs.S = 43242434343434223

r := xs.Next()

*/
package xorshift

var (
	// "const" for Jump function
	jump128 = []uint64{0x8a5cd789635d2dff, 0x121fd2155c472f96}

	jump1024 = []uint64{
		0x84242f96eca9c41d,
		0xa3c65b8776f96855, 0x5b34a39f070b5837, 0x4489affce4f31a1e,
		0x2ffeeb0a48316f40, 0xdc2d9891fe68c022, 0x3659132bb12fea70,
		0xaac17d8efa43cab8, 0xc4cb815590989b13, 0x5ee975283d71c93b,
		0x691548c86c1bd540, 0x7910c41d10a1e6a5, 0x0b5fc64563b3e2a8,
		0x047f7684e9fc949d, 0xb99181f2d8f685ca, 0x284600e3f30e38c3,
	}
)

// XorShift64Star hold the state required by the XorShift64Star generators.
type XorShift64Star struct {
	s uint64 // The state must be seeded with a nonzero value. Require a 64-bit unsigned values.
}

// XorShift128Plus holds the state required by XorShift128Plus generator.
type XorShift128Plus struct {
	// The state must be seeded with a nonzero value. Require 2 64-bit unsigned values.
	// The state must be seeded so that it is not everywhere zero. If you have a 64-bit seed,
	// we suggest to seed a xorshift64* generator and use its output to fill S.
	s [2]uint64
}

// XoroShiro128Plus holds the state required by XoroShiro128Plus generator
type XoroShiro128Plus struct {
	// The state must be seeded with a nonzero value. Require 2 64-bit unsigned values.
	// The state must be seeded so that it is not everywhere zero. If you have a 64-bit seed,
	// we suggest to seed a xorshift64* generator and use its output to fill S.
	s [2]uint64
}

// XorShift1024Star holds the state required by XorShift1024Star generator.
type XorShift1024Star struct {
	// The state must be seeded with a nonzero value. Require 16 64-bit unsigned values.
	// The state must be seeded so that it is not everywhere zero. If you have a 64-bit seed,
	// we suggest to seed a xorshift64* generator and use its output to fill s .
	s [16]uint64
	p int
}

// XorShift4096Star holds the state required by XorShift4096Star generator.
type XorShift4096Star struct {
	// The state must be seeded with a nonzero value. Require 64 64-bit unsigned values.
	// The state must be seeded so that it is not everywhere zero. If you have a 64-bit seed,
	// we suggest to seed a xorshift64* generator and use its output to fill s .
	s [64]uint64
	p int
}

// Next returns the next pseudo random number generated, before start you must provvide one 64 unsigned bit seed.
func (x *XorShift64Star) Next() uint64 {
	r := x.s * uint64(2685821657736338717)
	x.s ^= x.s >> 12
	x.s ^= x.s << 25
	x.s ^= x.s >> 27

	return r
}

// Init returns a new XorShift64Star source seeded with the given value.
func (x *XorShift64Star) Init(seed uint64) {
	x.s = seed
}

// Next returns the next pseudo random number generated, before start you must provvide seed.
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

// Jump it is equivalent to 2^64 calls to Next();
func (x *XorShift128Plus) Jump() {
	var s0, s1 uint64 = 0, 0
	var b uint64

	for i := 0; i < len(jump128); i++ {
		for b = 0; b < 64; b++ {
			if jump128[i]&uint64(1)<<b != 0 {
				s0 ^= x.s[0]
				s1 ^= x.s[1]
			}
			x.Next()
		}
	}

	x.s[0] = s0
	x.s[1] = s1
}

// Init returns a new XorShift128Plus source seeded with a slice of 2 values.
func (x *XorShift128Plus) Init(seed []uint64) {
	if len(seed) > 1 {
		x.s[0], x.s[1] = seed[0], seed[1]
		return
	}
	x.s[0] = seed[0]

}

// Next returns the next pseudo random number generated, before start you must provvide seed.
func (x *XoroShiro128Plus) Next() uint64 {
	s0, s1 := x.s[0], x.s[1]
	r := s0 + s1

	s1 ^= s0

	// update the state of generator
	x.s[0] = ((s0 << 55) | (s0 >> (64 - 55))) ^ s1 ^ (s1 << 14) // a,b
	x.s[1] = ((s1 << 36) | (s1 >> (64 - 36)))

	return r
}

// Jump it is equivalent to 2^64 calls to Next();
func (x *XoroShiro128Plus) Jump() {
	var s0, s1 uint64 = 0, 0
	var b uint64

	for i := 0; i < len(jump128); i++ {
		for b = 0; b < 64; b++ {
			if jump128[i]&uint64(1)<<b != 0 {
				s0 ^= x.s[0]
				s1 ^= x.s[1]
			}
			x.Next()
		}
	}

	x.s[0] = s0
	x.s[1] = s1
}

// Init returns a new XoroShiro128Plus source seeded with a slice of 2 values.
func (x *XoroShiro128Plus) Init(seed []uint64) {
	if len(seed) > 1 {
		x.s[0], x.s[1] = seed[0], seed[1]
		return
	}
	x.s[0] = seed[0]

}

// Next returns the next pseudo random number generated, before start you must provvide seed.
func (x *XorShift1024Star) Next() uint64 {
	s0 := x.s[x.p]

	xpnew := (x.p + 1) & 15

	s1 := x.s[xpnew]

	s1 ^= s1 << 31 // a
	tmp := s1 ^ s0 ^ (s1 >> 11) ^ (s0 >> 30)

	// update the state of generator
	x.s[xpnew] = tmp
	x.p = xpnew

	return tmp * uint64(1181783497276652981)
}

// Jump function for the generator. It is equivalent to 2^512 calls to next()
func (x *XorShift1024Star) Jump() {
	var t [16]uint64
	var b uint64

	for i := 0; i < len(jump1024); i++ {
		for b = 0; b < 64; b++ {
			if jump1024[i]&uint64(1)<<b != 0 {
				for j := 0; j < 16; j++ {
					t[j] ^= x.s[(j+x.p)&15]
				}
			}
			x.Next()
		}
	}

	for j := 0; j < 16; j++ {
		x.s[(j+x.p)&15] = t[j]
	}
}

// Init returns a new XorShift1024Star source seeded with a slice of 16 values.
func (x *XorShift1024Star) Init(seed []uint64) {
	for i, v := range seed {
		if i < len(x.s) {
			x.s[i] = v
		}
	}
	x.p = 0
}

// Next returns the next pseudo random number generated, before start you must provvide seed.
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

	return tmp * uint64(8372773778140471301)
}

// Init returns a new XorShift4096Star source seeded with a slicef of 64 values.
func (x *XorShift4096Star) Init(seed []uint64) {
	for i, v := range seed {
		if i < len(x.s) {
			x.s[i] = v
		}
	}
	x.p = 0
}
