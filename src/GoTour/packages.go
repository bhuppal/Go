package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Using rand.Intn example
	// pseudo-random sequence number
	// By default it depends on some seed number
	// Intn returns, as an int, a non-negative pseudo-random number in [0,n]. It panics if n <= 0.
	fmt.Println("My favorite number is ", rand.Intn(50))

	//type Rand
	// func New(src Source) *Rand -
	// New returns a new Rand that uses random values from src to generate other random values

	// time package
	// func(t Time) UnixNano() int64
	// UnixNano returns t as a Unix time, the number of nanoseconds elapsed since January 1, 1970 UTC.
	// The result is undefined if the Unix time in nanoseconds cannot be represented by an int64
	// (a date before the year 1678 or after 2262). Note that this means the result of calling UnixNano
	// on the zero Time is undefined. The result does not depend on the location associated with t.
	customSource := rand.NewSource(time.Now().UnixNano())
	customRand := rand.New(customSource)
	fmt.Println(customRand.Intn(100))

	// Using rand.Seed
	// Seed uses the provided seed value to initialize the default Source to a deterministic state.
	// If Seed is not called, the generator behaves as if seeded by Seed(1).
	// Seed values that have the same remainder when divided by 2³¹-1 generate the same pseudo-random sequence.
	// Seed, unlike the Rand.Seed method, is safe for concurrent use.
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Using seed function:", rand.Intn(200))
}
