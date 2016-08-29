package xorshift

import (
	"math/rand"
	"testing"
)

const (
	SEED0 = 43433241441424
	SEED1 = 3243241442214
)

// benchmarks

func BenchmarkXorShift64Star(b *testing.B) {
	tmpxs := XorShift64Star{}

	tmpxs.Init(SEED0)

	for i := 0; i < b.N; i++ {
		_ = tmpxs.Next()
	}
}

func BenchmarkXorShift128Plus(b *testing.B) {
	xs := XorShift128Plus{}

	xs.Init([]uint64{SEED0, SEED1})

	for i := 0; i < b.N; i++ {
		_ = xs.Next()
	}
}

func BenchmarkXoroShiro128Plus(b *testing.B) {
	xs := XoroShiro128Plus{}

	xs.Init([]uint64{SEED0, SEED1})

	for i := 0; i < b.N; i++ {
		_ = xs.Next()
	}
}

func BenchmarkXorShift1024Star(b *testing.B) {
	tmpxs := XorShift64Star{}
	tmpxs.s = SEED0

	xs := XorShift1024Star{}

	seed := make([]uint64, 16)

	for i := 0; i < 16; i++ {
		seed[i] = tmpxs.Next()

	}

	xs.Init(seed)

	for i := 0; i < b.N; i++ {
		_ = xs.Next()

	}
}

func BenchmarkXorShift4096Star(b *testing.B) {
	tmpxs := XorShift64Star{}
	tmpxs.s = SEED0

	xs := XorShift4096Star{}

	seed := make([]uint64, 64)

	for i := 0; i < 64; i++ {
		seed[i] = tmpxs.Next()

	}

	xs.Init(seed)

	for i := 0; i < b.N; i++ {
		_ = xs.Next()
	}
}

func BenchmarkRandSource(b *testing.B) {
	s := rand.NewSource(SEED0)
	for i := 0; i < b.N; i++ {
		_ = s.Int63()
	}
}
