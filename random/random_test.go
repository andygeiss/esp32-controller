package random_test

import (
	"testing"

	"github.com/andygeiss/esp32-controller/random"
	"github.com/andygeiss/utils/assert"
)

func TestNum(t *testing.T) {
	num := random.Num(10)
	assert.That("num is max 10", t, num <= 10, true)
}

func TestNumXY(t *testing.T) {
	x := random.Num(100)
	y := random.Num(100)
	assert.That("x and y are not equal", t, x != y, true)
}

func TestNumBetween(t *testing.T) {
	num := random.NumBetween(100, 200)
	between100And200 := num >= 100 && num <= 200
	assert.That("num is between 100 and 200", t, between100And200, true)
}

func TestSeed(t *testing.T) {
	random.Seed(42)
	x := random.Num(100)
	y := random.Num(100)
	assert.That("x and y are not equal", t, x != y, true)
	assert.That("x is max 100", t, x <= 100, true)
	assert.That("y is max 100", t, y <= 100, true)
}
