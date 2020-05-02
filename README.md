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


If you need to switch between different implementation or pass around the generators, you can use the interface defined in the xorshift package.

```go
    package main
    import (
       "github.com/vpxyz/xorshift"
       "github.com/vpxyz/xorshift/xorshift1024star"
       "github.com/vpxyz/xorshift/xoroshiro256plus"
    )
   func useRandom(x XorShiftExt) {
       x.Jump()
       _ = x.Uint64()
   }
    
    func main() {
    
       xs := xorshift1024star.NewSource(2343243232521)
       useRandom(xs)
       
       xss :=  xorshift256plus.NewSource(2343243232521)
       useRandom(xss)
    }
```

## Benchmarks

On Fedora 31 with vanilla linux kernel 5.6.9, cpu i7-3840QM and go 1.14

``` shellsession
    $ go test -bench=.
      goos: linux
      goarch: amd64
      pkg: github.com/vpxyz/xorshift
      BenchmarkSplitMix64Source64-8                   717461050                1.61 ns/op            0 B/op          0 allocs/op
      BenchmarkSplitMix64asRand64-8                   138133461                8.66 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift64StarSource64-8               406665097                2.93 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift64StarAsRand64-8               150135656                8.03 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift128PlusSource64-8              488120306                2.42 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift128PlusAsRand64-8              143114881                8.42 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro128PlusSource64-8             436295998                2.66 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro128PlusAsRand64-8             145748610                8.10 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro128StarStarSource64-8         444608847                2.63 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro128StarStarAsRand64-8         140955052                8.44 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro256PlusSource64-8             445946502                2.64 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro256PlusAsRand64-8             125309582                9.55 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro256PlusPlusSource64-8         304200223                3.86 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro256PlusPlusAsRand64-8         97761489                11.7 ns/op             0 B/op          0 allocs/op
      BenchmarkXoroShiro256StarStarSource64-8         295416097                3.93 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro256StarStarAsRand64-8         103112442               11.7 ns/op             0 B/op          0 allocs/op
      BenchmarkXoroShiro512PlusSource64-8             193433206                6.15 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro512PlusAsRand64-8             91186728                13.2 ns/op             0 B/op          0 allocs/op
      BenchmarkXoroShiro512StarStarSource64-8         189024180                6.34 ns/op            0 B/op          0 allocs/op
      BenchmarkXoroShiro512StarStarAsRand64-8         77784883                13.6 ns/op             0 B/op          0 allocs/op
      BenchmarkXorShift1024StarSource64-8             490018650                2.40 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift1024StarAsRand64-8             123210558                9.61 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift1024StarPhiSource64-8          491318622                2.39 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift1024StarPhiAsRand64-8          124519191                9.67 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift4096StarSource64-8             503598228                2.32 ns/op            0 B/op          0 allocs/op
      BenchmarkXorShift4096StarAsRand64-8             124942081                9.58 ns/op            0 B/op          0 allocs/op
      BenchmarkRandSource-8                           244402022                4.93 ns/op            0 B/op          0 allocs/op
      BenchmarkRand-8                                 116951886               10.2 ns/op             0 B/op          0 allocs/op
      PASS
      ok      github.com/vpxyz/xorshift       48.133s
    
```
