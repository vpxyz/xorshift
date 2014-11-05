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


