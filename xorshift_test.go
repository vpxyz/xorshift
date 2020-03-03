package xorshift

import (
	"math/rand"
	"testing"

	"github.com/vpxyz/xorshift/splitmix64"
	"github.com/vpxyz/xorshift/xoroshiro128plus"
	"github.com/vpxyz/xorshift/xoroshiro128starstar"
	"github.com/vpxyz/xorshift/xoroshiro256plus"
	"github.com/vpxyz/xorshift/xoroshiro256plusplus"
	"github.com/vpxyz/xorshift/xoroshiro256starstar"
	"github.com/vpxyz/xorshift/xoroshiro512plus"
	"github.com/vpxyz/xorshift/xoroshiro512starstar"
	"github.com/vpxyz/xorshift/xorshift1024star"
	"github.com/vpxyz/xorshift/xorshift1024starphi"
	"github.com/vpxyz/xorshift/xorshift128plus"
	"github.com/vpxyz/xorshift/xorshift4096star"
	"github.com/vpxyz/xorshift/xorshift64star"
)

const (
	SEED = 43433241441424
)

// benchmarks

func BenchmarkSplitMix64Source64(b *testing.B) {
	xs := splitmix64.NewSource(SEED)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = xs.Uint64()
	}
}

func BenchmarkSplitMix64asRand64(b *testing.B) {
	tmpxs := splitmix64.NewSource(SEED)
	b.ReportAllocs()
	r := rand.New(tmpxs)

	for i := 0; i < b.N; i++ {
		_ = r.ExpFloat64()
	}
}

func BenchmarkXorShift64StarSource64(b *testing.B) {
	xs := xorshift64star.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xs.Uint64()
	}
}

func BenchmarkXorShift64StarAsRand64(b *testing.B) {
	tmpxs := xorshift64star.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	r := rand.New(tmpxs)

	for i := 0; i < b.N; i++ {
		_ = r.ExpFloat64()
	}
}
func BenchmarkXorShift128PlusSource64(b *testing.B) {
	xs := xorshift128plus.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xs.Uint64()
	}
}

func BenchmarkXorShift128PlusAsRand64(b *testing.B) {
	tmpxs := xorshift128plus.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	r := rand.New(tmpxs)

	for i := 0; i < b.N; i++ {
		_ = r.ExpFloat64()
	}

}

func BenchmarkXoroShiro128PlusSource64(b *testing.B) {
	xs := xoroshiro128plus.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xs.Uint64()
	}
}

func BenchmarkXoroShiro128PlusAsRand64(b *testing.B) {
	tmpxs := xoroshiro128plus.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	r := rand.New(tmpxs)

	for i := 0; i < b.N; i++ {
		_ = r.ExpFloat64()
	}
}

func BenchmarkXoroShiro128StarStarSource64(b *testing.B) {
	xs := xoroshiro128starstar.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xs.Uint64()
	}
}

func BenchmarkXoroShiro128StarStarAsRand64(b *testing.B) {
	tmpxs := xoroshiro128starstar.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	r := rand.New(tmpxs)

	for i := 0; i < b.N; i++ {
		_ = r.ExpFloat64()
	}
}

func BenchmarkXoroShiro256PlusSource64(b *testing.B) {
	xs := xoroshiro256plus.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xs.Uint64()
	}
}

func BenchmarkXoroShiro256PlusAsRand64(b *testing.B) {
	tmpxs := xoroshiro256plus.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	r := rand.New(tmpxs)

	for i := 0; i < b.N; i++ {
		_ = r.ExpFloat64()
	}
}

func BenchmarkXoroShiro256PlusPlusSource64(b *testing.B) {
	xs := xoroshiro256plusplus.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xs.Uint64()
	}
}

func BenchmarkXoroShiro256PlusPlusAsRand64(b *testing.B) {
	tmpxs := xoroshiro256plusplus.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	r := rand.New(tmpxs)

	for i := 0; i < b.N; i++ {
		_ = r.ExpFloat64()
	}
}

func BenchmarkXoroShiro256StarStarSource64(b *testing.B) {
	xs := xoroshiro256starstar.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xs.Uint64()
	}
}

func BenchmarkXoroShiro256StarStarAsRand64(b *testing.B) {
	tmpxs := xoroshiro256starstar.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	r := rand.New(tmpxs)

	for i := 0; i < b.N; i++ {
		_ = r.ExpFloat64()
	}
}

func BenchmarkXoroShiro512PlusSource64(b *testing.B) {
	xs := xoroshiro512plus.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xs.Uint64()
	}
}

func BenchmarkXoroShiro512PlusAsRand64(b *testing.B) {
	tmpxs := xoroshiro512plus.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	r := rand.New(tmpxs)

	for i := 0; i < b.N; i++ {
		_ = r.ExpFloat64()
	}
}

func BenchmarkXoroShiro512StarStarSource64(b *testing.B) {
	xs := xoroshiro512starstar.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xs.Uint64()
	}
}

func BenchmarkXoroShiro512StarStarAsRand64(b *testing.B) {
	tmpxs := xoroshiro512starstar.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	r := rand.New(tmpxs)

	for i := 0; i < b.N; i++ {
		_ = r.ExpFloat64()
	}
}

func BenchmarkXorShift1024StarSource64(b *testing.B) {
	xs := xorshift1024star.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xs.Uint64()

	}
}

func BenchmarkXorShift1024StarAsRand64(b *testing.B) {
	tmpxs := xorshift1024star.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	r := rand.New(tmpxs)

	for i := 0; i < b.N; i++ {
		_ = r.ExpFloat64()
	}
}

func BenchmarkXorShift1024StarPhiSource64(b *testing.B) {
	xs := xorshift1024starphi.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xs.Uint64()

	}
}

func BenchmarkXorShift1024StarPhiAsRand64(b *testing.B) {
	tmpxs := xorshift1024starphi.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	r := rand.New(tmpxs)
	for i := 0; i < b.N; i++ {
		_ = r.ExpFloat64()
	}
}

func BenchmarkXorShift4096StarSource64(b *testing.B) {
	xs := xorshift4096star.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xs.Uint64()
	}
}

func BenchmarkXorShift4096StarAsRand64(b *testing.B) {
	tmpxs := xorshift4096star.NewSource(SEED)
	b.ReportAllocs()
	b.ResetTimer()
	r := rand.New(tmpxs)

	for i := 0; i < b.N; i++ {
		_ = r.ExpFloat64()
	}
}

func BenchmarkRandSource(b *testing.B) {
	b.ReportAllocs()
	s := rand.NewSource(SEED)
	for i := 0; i < b.N; i++ {
		_ = s.Int63()
	}
}

func BenchmarkRand(b *testing.B) {
	b.ReportAllocs()
	tmpxs := rand.NewSource(SEED)
	r := rand.New(tmpxs)
	for i := 0; i < b.N; i++ {
		_ = r.ExpFloat64()
	}
}
