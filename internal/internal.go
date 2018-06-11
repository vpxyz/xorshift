package internal

var (
	// Jump128 "const" for Jump function
	Jump128 = []uint64{0x8a5cd789635d2dff, 0x121fd2155c472f96}

	// Jump1024 "const" for Jump function
	Jump1024 = []uint64{
		0x84242f96eca9c41d,
		0xa3c65b8776f96855, 0x5b34a39f070b5837, 0x4489affce4f31a1e,
		0x2ffeeb0a48316f40, 0xdc2d9891fe68c022, 0x3659132bb12fea70,
		0xaac17d8efa43cab8, 0xc4cb815590989b13, 0x5ee975283d71c93b,
		0x691548c86c1bd540, 0x7910c41d10a1e6a5, 0x0b5fc64563b3e2a8,
		0x047f7684e9fc949d, 0xb99181f2d8f685ca, 0x284600e3f30e38c3,
	}
)

// SplitMix64 hold the state required by the SplitMix64 generator.
type SplitMix64 struct {
	s uint64 // The state can be seeded with any value
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (x *SplitMix64) Int63() int64 {
	return int64(x.Uint64() & (1<<63 - 1))

}

// Uint64 returns the next pseudo random number generated, before start you must provvide one 64 unsigned bit seed.
func (x *SplitMix64) Uint64() uint64 {
	x.s = x.s + uint64(0x9E3779B97F4A7C15)
	z := x.s
	z = (z ^ (z >> 30)) * uint64(0xBF58476D1CE4E5B9)
	z = (z ^ (z >> 27)) * uint64(0x94D049BB133111EB)
	return z ^ (z >> 31)

}

// Seed seed SplitMix64 random number generator with the given value.
func (x *SplitMix64) Seed(seed int64) {
	x.s = uint64(seed)
}

// Rotl rotate bits to left
func Rotl(x uint64, k uint64) uint64 {
	return (x << k) | (x >> (64 - k))
}
