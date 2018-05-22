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
       fmt.Printf("pseudo random = %v\n", xs.Uint64())
       
       // some generators has a Jump function
       // for XorShift2014Star is equivalent to 2^512 calls to Uint64()
       xs.Jump() 
       
       // because every generators implements Source64 interface, 
       // you can use it as drop-in replacement for rand.New()
       r := rand.New(tmpxs)

       for i := 0; i < b.N; i++ {
		       _ = r.ExpFloat64()
	   }
       
    }
```

## Benchmarks

On Fedora 28 with vanilla linux kernel 4.16.9, cpu i7-3840QM.

``` shellsession
    $ go test -bench=.
    goos: linux
    goarch: amd64
    pkg: github.com/vpxyz/xorshift
    BenchmarkSplitMix64Source64-8             	2000000000	         1.60 ns/op	       0 B/op	       0 allocs/op
    BenchmarkSplitMix64asRand64-8             	100000000	        12.5 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift64StarSource64-8         	1000000000	         2.93 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift64StarAsRand64-8         	100000000	        11.9 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift128PlusSource64-8        	1000000000	         2.40 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift128PlusAsRand64-8        	100000000	        12.4 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXoroShiro128PlusSource64-8       	1000000000	         2.71 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXoroShiro128PlusAsRand64-8       	100000000	        12.0 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXoroShiro128StarStarSource64-8   	500000000	         3.14 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXoroShiro128StarStarAsRand64-8   	100000000	        14.3 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXoroShiro256PlusSource64-8       	300000000	         4.44 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXoroShiro256PlusAsRand64-8       	100000000	        15.7 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXoroShiro256StarStarSource64-8   	300000000	         4.64 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXoroShiro256StarStarAsRand64-8   	100000000	        15.5 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXoroShiro512PlusSource64-8       	200000000	         6.61 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXoroShiro512PlusAsRand64-8       	100000000	        17.0 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXoroShiro512StarStarSource64-8   	200000000	         6.56 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXoroShiro512StarStarAsRand64-8   	100000000	        17.7 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift1024StarSource64-8       	1000000000	         2.42 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift1024StarAsRand64-8       	100000000	        13.7 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift1024StarPhiSource64-8    	1000000000	         2.41 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift1024StarPhiAsRand64-8    	100000000	        13.8 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift4096StarSource64-8       	1000000000	         2.31 ns/op	       0 B/op	       0 allocs/op
    BenchmarkXorShift4096StarAsRand64-8       	100000000	        13.6 ns/op	       0 B/op	       0 allocs/op
    BenchmarkRandSource-8                     	300000000	         4.89 ns/op	       0 B/op	       0 allocs/op
    BenchmarkRand-8                           	100000000	        14.3 ns/op	       0 B/op	       0 allocs/op
    PASS
    ok  	github.com/vpxyz/xorshift	50.267s
    
    
```
