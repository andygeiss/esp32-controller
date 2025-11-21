package random_test

import (
	"testing"

	"github.com/andygeiss/cloud-native-utils/assert"
	"github.com/andygeiss/esp32-controller/random"
)

func TestNum(t *testing.T) {
	num := random.Num(10)
	assert.That(t, "num is max 10", num <= 10, true)
}

func TestNumXY(t *testing.T) {
	x := random.Num(100)
	y := random.Num(100)
	assert.That(t, "x and y are not equal", x != y, true)
}

func TestNumBetween(t *testing.T) {
	num := random.NumBetween(100, 200)
	between100And200 := num >= 100 && num <= 200
	assert.That(t, "num is between 100 and 200", between100And200, true)
}

func TestSeed(t *testing.T) {
	random.Seed(42)
	x := random.Num(100)
	y := random.Num(100)
	assert.That(t, "x and y are not equal", x != y, true)
	assert.That(t, "x is max 100", x <= 100, true)
	assert.That(t, "y is max 100", y <= 100, true)
}
