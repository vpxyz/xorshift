package xorshit

import (
	"log"
	"testing"
)

func TestXorshift64(t *testing.T) {
	xs := XorShift64Star{}

	xs.S = 2343243232521

	for i := 0; i < 10000; i++ {
		r := xs.Next()

		log.Printf("Iteration %v, next64 = %v\n", i, r)
	}

}

func TestXorshift128(t *testing.T) {
	xs := XorShift128Plus{}

	xs.S[0] = 43433241441424
	xs.S[1] = 3243241442214

	for i := 0; i < 10000; i++ {
		r := xs.Next()

		log.Printf("Iteration %v, next128 = %v\n", i, r)
	}
}

func TestXorshift1024(t *testing.T) {
	tmpxs := XorShift64Star{}
	tmpxs.S = 2343243232521

	xs := XorShift1024Star{}

	for i := 0; i < 16; i++ {
		xs.S[i] = tmpxs.Next()

	}

	for i := 0; i < 10000; i++ {
		r := xs.Next()

		log.Printf("Iteration %v, p = %v, next1024 = %v\n", i, xs.p, r)
	}
}

func TestXorshift4096(t *testing.T) {
	tmpxs := XorShift64Star{}
	tmpxs.S = 2343243232521

	xs := XorShift4096Star{}

	for i := 0; i < 64; i++ {
		xs.S[i] = tmpxs.Next()

	}

	for i := 0; i < 10000; i++ {
		r := xs.Next()

		log.Printf("Iteration %v, p = %v, next4096 = %v\n", i, xs.p, r)
	}
}
