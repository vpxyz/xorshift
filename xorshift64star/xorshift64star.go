/*
Package xorshift64star it's like other xorshift*, generators, but with a 2^64 - 1 period.
*/
package xorshift64star

// XorShift64Star hold the state required by the XorShift64Star generator.
type XorShift64Star struct {
	s uint64 // The state must be seeded with a nonzero value. Require a 64-bit unsigned values.
}

// NewSource return a new XorShift64Star random number generator.
func NewSource(seed int64) *XorShift64Star {
	tmpxs := XorShift64Star{}
	tmpxs.s = uint64(seed)
	return &tmpxs
}

// Seed use the provvided seed value to init XorShift64Star internal state.
func (x *XorShift64Star) Seed(seed int64) {
	x.s = uint64(seed)
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (x *XorShift64Star) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))

}

// Uint64 returns the next pseudo random number generated, before start you must provvide one 64 unsigned bit seed.
func (x *XorShift64Star) Uint64() uint64 {
	r := x.s * uint64(2685821657736338717)
	x.s ^= x.s >> 12
	x.s ^= x.s << 25
	x.s ^= x.s >> 27

	return r
}
