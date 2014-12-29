package xorshit

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

	xs.S = SEED0

	for i := 0; i < 10000; i++ {
		r := xs.Next()

		log.Printf("Iteration %v, next64 = %v\n", i, r)
	}

}

func TestXorshift128Plus(t *testing.T) {
	xs := XorShift128Plus{}

	xs.S[0] = SEED0
	xs.S[1] = SEED1

	for i := 0; i < 10000; i++ {
		r := xs.Next()

		log.Printf("Iteration %v, next128 = %v\n", i, r)
	}
}

func TestXorshift1024Star(t *testing.T) {
	tmpxs := XorShift64Star{}
	tmpxs.S = SEED0

	xs := XorShift1024Star{}

	for i := 0; i < len(xs.S); i++ {
		xs.S[i] = tmpxs.Next()

	}

	for i := 0; i < 10000; i++ {
		r := xs.Next()

		log.Printf("Iteration %v, p = %v, next1024 = %v\n", i, xs.p, r)
	}
}

func TestXorshift4096Star(t *testing.T) {
	tmpxs := XorShift64Star{}
	tmpxs.S = SEED0

	xs := XorShift4096Star{}

	for i := 0; i < len(xs.S); i++ {
		xs.S[i] = tmpxs.Next()

	}

	for i := 0; i < 10000; i++ {
		r := xs.Next()

		log.Printf("Iteration %v, p = %v, next4096 = %v\n", i, xs.p, r)
	}
}

// benchmarks

func BenchmarkXorShift64Star(b *testing.B) {
	tmpxs := XorShift64Star{}
	tmpxs.S = SEED0

	for i := 0; i < b.N; i++ {
		_ = tmpxs.Next()
	}
}

func BenchmarkXorshift128Plus(b *testing.B) {
	xs := XorShift128Plus{}

	xs.S[0] = SEED0
	xs.S[1] = SEED1

	for i := 0; i < b.N; i++ {
		_ = xs.Next()
	}
}

func BenchmarkXorshift1024Star(b *testing.B) {
	tmpxs := XorShift64Star{}
	tmpxs.S = SEED0

	xs := XorShift1024Star{}

	for i := 0; i < len(xs.S); i++ {
		xs.S[i] = tmpxs.Next()

	}

	for i := 0; i < b.N; i++ {
		_ = xs.Next()

	}
}

func BenchmarkXorshift4096Star(b *testing.B) {
	tmpxs := XorShift64Star{}
	tmpxs.S = SEED0

	xs := XorShift4096Star{}

	for i := 0; i < len(xs.S); i++ {
		xs.S[i] = tmpxs.Next()

	}

	for i := 0; i < b.N; i++ {
		_ = xs.Next()
	}
}

func BenchmarkSyncXorShift64Star(b *testing.B) {
	tmpxs := XorShift64Star{}
	tmpxs.S = SEED0

	for i := 0; i < b.N; i++ {
		_ = tmpxs.SyncNext()
	}
}

func BenchmarkSyncXorshift128Plus(b *testing.B) {
	xs := XorShift128Plus{}

	xs.S[0] = SEED0
	xs.S[1] = SEED1

	for i := 0; i < b.N; i++ {
		_ = xs.SyncNext()
	}
}

func BenchmarkSyncXorshift1024Star(b *testing.B) {
	tmpxs := XorShift64Star{}
	tmpxs.S = SEED0

	xs := XorShift1024Star{}

	for i := 0; i < len(xs.S); i++ {
		xs.S[i] = tmpxs.SyncNext()

	}

	for i := 0; i < b.N; i++ {
		_ = xs.SyncNext()

	}
}

func BenchmarkSyncXorshift4096Star(b *testing.B) {
	tmpxs := XorShift64Star{}
	tmpxs.S = SEED0

	xs := XorShift4096Star{}

	for i := 0; i < len(xs.S); i++ {
		xs.S[i] = tmpxs.SyncNext()

	}

	for i := 0; i < b.N; i++ {
		_ = xs.SyncNext()
	}
}

func BenchmarkRandSource(b *testing.B) {
	s := rand.NewSource(SEED0)
	for i := 0; i < b.N; i++ {
		_ = s.Int63()
	}
}
