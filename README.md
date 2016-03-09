# Xorshift

Xorshift is a simple library that implements xorshift* and xorshift+ pseudo random number generators in GO.

This simple library in based on the work of Sebastiano Vigna (http://xorshift.di.unimi.it/).

[![Go Walker](https://img.shields.io/badge/Go%20Walker-API%20Documentation-green.svg?style=flat)](https://gowalker.org/github.com//vpxyz/xorshift)
[![GoDoc](https://godoc.org/github.com/vpxyz/xorshift?status.svg)](https://godoc.org/github.com/vpxyz/xorshift)
[![status](https://sourcegraph.com/api/repos/github.com/vpxyz/xorshift/.badges/status.svg)](https://sourcegraph.com/github.com/vpxyz/xorshift)

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
   	   tmpxs := xorshift.XorShift64Star{}
   	   tmpxs.Init(2343243232521)

       xs := xorshift.XorShift4096Star{}

       // you can use XorShift64Star for fill XorShift4096Star Seed
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

``` shellsession
    $ go test -bench=.
    PASS
    BenchmarkXorShift64Star-8       1000000000               2.95 ns/op
    BenchmarkXorshift128Plus-8      1000000000               2.41 ns/op
    BenchmarkXorshift1024Star-8     1000000000               2.74 ns/op
    BenchmarkXorshift4096Star-8     1000000000               2.91 ns/op
    BenchmarkRandSource-8           300000000                5.83 ns/o
    
```
