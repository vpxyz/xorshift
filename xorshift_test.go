package xorshift

import (
	"log"
	"math/rand"
	"testing"
)

const (
	SEED0 = 43433241441424
	SEED1 = 3243241442214
)

func TestXorshift64Star(t *testing.T) {
	xs := XorShift64Star{}

	xs.Init(SEED0)

	log.Print("Xorshift64Star:\n")
	for i := 0; i < 10000; i++ {
		r := xs.Next()

		log.Printf("%10d, %20d\n", i, r)
	}

}

func TestXorshift128Plus(t *testing.T) {
	xs := XorShift128Plus{}

	xs.Init([]uint64{SEED0, SEED1})

	log.Print("Xorshift128Plus:\n")
	for i := 0; i < 10000; i++ {
		r := xs.Next()
		log.Printf("%10d, %20d\n", i, r)

	}
}

func TestXorshift1024Star(t *testing.T) {
	tmpxs := XorShift64Star{}
	tmpxs.s = SEED0

	xs := XorShift1024Star{}

	seed := make([]uint64, 16)

	for i := 0; i < 16; i++ {
		seed[i] = tmpxs.Next()

	}

	xs.Init(seed)

	log.Print("Xorshift1024Star:\n")
	for i := 0; i < 10000; i++ {
		r := xs.Next()

		log.Printf("%10d, %20d\n", i, r)
	}
}

func TestXorshift4096Star(t *testing.T) {
	tmpxs := XorShift64Star{}
	tmpxs.s = SEED0

	xs := XorShift4096Star{}

	seed := make([]uint64, 64)

	for i := 0; i < 64; i++ {
		seed[i] = tmpxs.Next()

	}

	xs.Init(seed)

	log.Print("Xorshift4096Star:\n")
	for i := 0; i < 10000; i++ {
		r := xs.Next()

		log.Printf("%10d, %20d\n", i, r)

	}
}

// benchmarks

func BenchmarkXorShift64Star(b *testing.B) {
	tmpxs := XorShift64Star{}

	tmpxs.Init(SEED0)

	for i := 0; i < b.N; i++ {
		_ = tmpxs.Next()
	}
}

func BenchmarkXorshift128Plus(b *testing.B) {
	xs := XorShift128Plus{}

	xs.Init([]uint64{SEED0, SEED1})

	for i := 0; i < b.N; i++ {
		_ = xs.Next()
	}
}

func BenchmarkXorshift1024Star(b *testing.B) {
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

func BenchmarkXorshift4096Star(b *testing.B) {
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
