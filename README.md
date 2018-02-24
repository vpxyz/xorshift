# Xorshift

Xorshift is a simple library that implements xorshift*, xorshift+, xoroshiro+ and splitmix64 pseudo random number generators in GO.

It's based on the work of Sebastiano Vigna (http://xoroshiro.di.unimi.it/).

[![Go Walker](https://img.shields.io/badge/Go%20Walker-API%20Documentation-green.svg?style=flat)](https://gowalker.org/github.com//vpxyz/xorshift)
[![GoDoc](https://godoc.org/github.com/vpxyz/xorshift?status.svg)](https://godoc.org/github.com/vpxyz/xorshift)
[![status](https://sourcegraph.com/api/repos/github.com/vpxyz/xorshift/.badges/status.svg)](https://sourcegraph.com/github.com/vpxyz/xorshift)
[![Go Report Card](https://goreportcard.com/badge/github.com/vpxyz/xorshift)](https://goreportcard.com/report/github.com/vpxyz/xorshift)

*NOTE*: Not concurrency-safe! You can wrap generator with a monitor goroutine, for e.g.

## Install

This package is "go-gettable", just do:

    go get github.com/vpxyz/xorshift

## Example

I suggest to use SplitMix64 for fill seed.


``` go
    package main
    
    import (
       "github.com/vpxyz/xorshift"
       "fmt"
    )

    func main() {
   	   tmpxs := xorshift.SplitMix64{}
   	   tmpxs.Init(2343243232521)

       xs := xorshift.XorShift4096Star{}

       // you can use SplitMix64 for fill Seed
       seed := make([]uint64, 64)

	   for i := 0; i < 64; i++ {
	       seed[i] = tmpxs.Next()

       }

       xs.Init(seed)
       
       // use the generator
       fmt.Printf("pseudo random = %v\n", xs.Next())
       
    }
```

## Benchmarks

On Fedora 27 with vanilla linux kernel 4.15.4, cpu i7-3840QM.

``` shellsession
    $ go test -bench=.
    PASS
    goos: linux
    goarch: amd64
    pkg: github.com/vpxyz/xorshift
    BenchmarkSplitMix64-8            	2000000000	         1.61 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift64Star-8        	1000000000	         2.95 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift128Plus-8       	1000000000	         2.41 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXoroShiro128Plus-8      	1000000000	         2.64 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift1024Star-8      	1000000000	         2.34 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift1024StarPhi-8   	1000000000	         2.34 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift4096Star-8      	1000000000	         2.47 ns/op	       0 B/op	       0 allocs/op
    BenchmarkRandSource-8            	300000000	         4.88 ns/op	       0 B/op	       0 allocs/op

    
```
