# Xorshift

Xorshift is a simple library that implements xorshift*, xorshift+, xoroshiro+ and splitmix64 pseudo random number generators in GO.

It's based on the work of Sebastiano Vigna (http://xoroshiro.di.unimi.it/).

As suggested by Vigna, generators with size of internal state greater then uint64, are filled using SplitMix64.

[![Go Walker](https://img.shields.io/badge/Go%20Walker-API%20Documentation-green.svg?style=flat)](https://gowalker.org/github.com//vpxyz/xorshift)
[![GoDoc](https://godoc.org/github.com/vpxyz/xorshift?status.svg)](https://godoc.org/github.com/vpxyz/xorshift)
[![status](https://sourcegraph.com/api/repos/github.com/vpxyz/xorshift/.badges/status.svg)](https://sourcegraph.com/github.com/vpxyz/xorshift)
[![Go Report Card](https://goreportcard.com/badge/github.com/vpxyz/xorshift)](https://goreportcard.com/report/github.com/vpxyz/xorshift)

*NOTE*: Not concurrency-safe! You can wrap generator with a monitor goroutine, for e.g.

## Install

This package is "go-gettable", just do:

    go get github.com/vpxyz/xorshift...

go >= 1.9 are required

## Example




``` go
    package main
    
    import (
       "github.com/vpxyz/xorshift/xorshift1024star"
       "fmt"
    )

    func main() {
    
       xs := xorshift1024star.NewSource(2343243232521)

       // use the generator
       fmt.Printf("pseudo random number = %v\n", xs.Uint64())
       
       // some generators has a Jump function
       // for XorShift2014Star is equivalent to 2^512 calls to Uint64()
       xs.Jump() 
       
       // because every generators implements Source64 interface, 
       // you can use it as drop-in replacement for rand.New()
       r := rand.New(xs)

       for i := 0; i < 10; i++ {
		       fmt.Printf("pseudo random number using Source64 interface = %v\n", r.ExpFloat64())
	   }
       
    }
```

## Benchmarks

On Fedora 31 with vanilla linux kernel 5.5.7, cpu i7-3840QM.

``` shellsession
    $ go test -bench=.
      goos: linux
      goarch: amd64
      pkg: github.com/vpxyz/xorshift
      BenchmarkSplitMix64Source64-8                   727509061                1.59 ns/op            0 B/op          0 allocs/op
      BenchmarkSplitMix64asRand64-8                   139142323                8.75 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift64StarSource64-8               402109586                2.92 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift64StarAsRand64-8               150146994                7.97 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift128PlusSource64-8              489134055                2.40 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift128PlusAsRand64-8              142476627                8.36 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro128PlusSource64-8             450397389                2.64 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro128PlusAsRand64-8             148400413                8.08 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro128StarStarSource64-8         445290165                2.64 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro128StarStarAsRand64-8         141690339                8.41 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro256PlusSource64-8             447325856                2.63 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro256PlusAsRand64-8             124359404                9.49 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro256PlusPlusSource64-8         306806460                3.90 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro256PlusPlusAsRand64-8         102424314               11.7 ns/op             0 B/op          0 allocs/op
      BenchmarkXoroShiro256StarStarSource64-8         303311925                3.90 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro256StarStarAsRand64-8         101823260               11.6 ns/op             0 B/op          0 allocs/op
      BenchmarkXoroShiro512PlusSource64-8             195214028                6.13 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro512PlusAsRand64-8             87043470                13.1 ns/op             0 B/op          0 allocs/op
      BenchmarkXoroShiro512StarStarSource64-8         188351828                6.28 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro512StarStarAsRand64-8         89132463                13.5 ns/op             0 B/op          0 allocs/op
      BenchmarkXorShift1024StarSource64-8             489940494                2.39 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift1024StarAsRand64-8             125186522                9.58 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift1024StarPhiSource64-8          489869648                2.40 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift1024StarPhiAsRand64-8          124680717                9.71 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift4096StarSource64-8             516819006                2.30 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift4096StarAsRand64-8             126090891                9.53 ns/op            0 B/op          0 allocs/op
      BenchmarkRandSource-8                           253307964                4.71 ns/op            0 B/op          0 allocs/op
      BenchmarkRand-8                                 120000142               10.0 ns/op             0 B/op          0 allocs/op
      PASS
      ok      github.com/vpxyz/xorshift       49.473s

    
```
