# Xorshift

Xorshift is a simple library that implements xorshift* and xorshift+ pseudo random number generators in GO.

This simple library in based on the work of Sebastiano Vigna (http://xorshift.di.unimi.it/).

## Install

This package is "go-gettable", just do:

    go get github.com/vpxyz/xorshift

## Example

Ok, that's all:

    import (
       "github.com/vpxyx/xorshift"
       "fmt"
    )

    func main() {
   	   tmpxs := XorShift64Star{}
   	   tmpxs.S = 2343243232521

       xs := XorShift4096Star{}

       // you can use XorShift64Star for fill XorShift4096Star Seed
       for i := 0; i < len(xs.S); i++ {
          xs.S[i] = tmpxs.Next()
       }

       // use the generator
       fmt.Printf("pseudo random = %v\n", xs.Next())
       
    }


## Benchmarks

    BenchmarkXorShift64Star         300000000                4.68 ns/op
    BenchmarkXorshift128Plus        300000000                4.86 ns/op
    BenchmarkXorshift1024Star       200000000                7.03 ns/op
    BenchmarkXorshift4096Star       200000000                7.04 ns/op

    BenchmarkSyncXorShift64Star     10000000               175 ns/op
    BenchmarkSyncXorshift128Plus    10000000               180 ns/op
    BenchmarkSyncXorshift1024Star   10000000               184 ns/op
    BenchmarkSyncXorshift4096Star   10000000               183 ns/op

    BenchmarkRandSource             200000000                8.98 ns/op

