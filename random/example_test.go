package random_test

import (
	"fmt"

	"github.com/andygeiss/esp32-controller/random"
)

// ExampleSeed shows what a seed buys you: the same seed draws the same numbers,
// so a run can be repeated.
func ExampleSeed() {
	random.Seed(42)
	first := random.Num(100)
	random.Seed(42)
	second := random.Num(100)

	fmt.Println(first == second)
	// Output: true
}

// ExampleNumBetween draws a number from 10 up to, but not including, 100.
func ExampleNumBetween() {
	n := random.NumBetween(10, 100)

	fmt.Println(n >= 10 && n < 100)
	// Output: true
}
