# Xorshift

Xorshift is a simple library that implements xorshift*, xorshift+ and splitmix64 pseudo random number generators in GO.

This simple library in based on the work of Sebastiano Vigna (http://xorshift.di.unimi.it/).

[![Go Walker](https://img.shields.io/badge/Go%20Walker-API%20Documentation-green.svg?style=flat)](https://gowalker.org/github.com//vpxyz/xorshift)
[![GoDoc](https://godoc.org/github.com/vpxyz/xorshift?status.svg)](https://godoc.org/github.com/vpxyz/xorshift)
[![status](https://sourcegraph.com/api/repos/github.com/vpxyz/xorshift/.badges/status.svg)](https://sourcegraph.com/github.com/vpxyz/xorshift)
[![Go Report Card](https://goreportcard.com/badge/github.com/vpxyz/xorshift)](https://goreportcard.com/report/github.com/vpxyz/xorshift)

*NOTE*: Not concurrency-safe! You must wrap into monitor goroutine or a mutex.

## Install

This package is "go-gettable", just do:

    go get github.com/vpxyz/xorshift

## Example

Ok, that's all:

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

       // you can use SplitMix64 for fill XorShift4096Star Seed
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

On Fedora 26 with vanilla linux kernel 4.13.2, cpu i7-3840QM.

``` shellsession
    $ go test -bench=.
    PASS
    goos: linux
    goarch: amd64
    pkg: github.com/vpxyz/xorshift
    BenchmarkSplitMix64-8           2000000000               1.64 ns/op
    BenchmarkXorShift64Star-8       500000000                2.99 ns/op
    BenchmarkXorShift128Plus-8      1000000000               2.45 ns/op
    BenchmarkXoroShiro128Plus-8     1000000000               2.69 ns/op
    BenchmarkXorShift1024Star-8     1000000000               2.48 ns/op
    BenchmarkXorShift4096Star-8     1000000000               2.46 ns/op
    BenchmarkRandSource-8           300000000                4.78 ns/op
    
```
