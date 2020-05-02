// Package xorshift define interfaces implemented by the generators
package xorshift

// The interfaces defined here, can be used to simplify your code if you want to switch from
// one generator to another.

// XorShift all sub packages implements this interface
type XorShift interface {
	Seed(seed int64)
	Uint64() uint64
}

// XorShiftExt optional functions, the xor... sub packages implements even this interface
type XorShiftExt interface {
	XorShift
	Jump()
}
