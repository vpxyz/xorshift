/*
Package xorshift implements a simple library for pseudo random number generators based on xorshift*, xorshift+, xoroshiro+ and splitmix64.

Xorshift* generators are obtained by scrambling the output of a Marsaglia xorshift generator with a 64-bit invertible multiplier.
Xorshift+ generators are a 64-bit version of Saito and Matsumoto's XSadd generator.
Xoroshiro128+ (XOR/rotate/shift/rotate) is the successor to xorshift128+.
Splitmix64 generator is a fixed-increment version of Java 8's SplittableRandom generator.

It's based on the work of Sebastiano Vigna (http://xoroshiro.di.unimi.it/).

All the generators implements rand.Source64 interface and can be used
as a drop-in replacement for rand.New() parameter.

Some generators have a Jump() function that is equivalent to call the generator many times.

NOTE: Not concurrency-safe! You can wrap generator with a monitor goroutine, for e.g.

*/
package xorshift
