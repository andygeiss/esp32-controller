// Package random draws pseudo-random numbers, the way an Arduino sketch does.
//
// [Num] and [NumBetween] draw a number; [Seed] restarts the sequence at a known
// point. An unseeded program draws the same numbers on every run — Arduino's
// random() behaves the same way — so call Seed with a value that varies when the
// numbers have to differ between runs.
//
// The generator is package state and is not safe to use from more than one
// goroutine.
package random

import "math/rand/v2"

// generator starts from a fixed point so an unseeded program repeats itself,
// which is what Arduino's random() does until randomSeed() is called.
var generator = rand.New(rand.NewPCG(0, 0))

// Num returns a number from 0 up to, but not including, max.
// It panics when max is not positive.
//
// See https://www.arduino.cc/reference/en/language/functions/random-numbers/random/
func Num(max int) int {
	return generator.IntN(max)
}

// NumBetween returns a number from min up to, but not including, max.
// It panics when max is not greater than min.
//
// See https://www.arduino.cc/reference/en/language/functions/random-numbers/random/
func NumBetween(min, max int) int {
	return min + generator.IntN(max-min)
}

// Seed restarts the sequence at a known point. Seeding twice with the same value
// draws the same numbers again.
//
// See https://www.arduino.cc/reference/en/language/functions/random-numbers/randomseed/
func Seed(seed int64) {
	generator = rand.New(rand.NewPCG(uint64(seed), uint64(seed)))
}
