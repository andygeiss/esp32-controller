package random_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/esp32-controller/random"
	"testing"
)

func TestNum(t *testing.T) {
	num := random.Num(10)
	assert.That(t, num, 0)
}

func TestNumXY(t *testing.T) {
	x := random.Num(100)
	y := random.Num(100)
	assert.That(t, x, y)
}

func TestNumBetween(t *testing.T) {
	num := random.NumBetween(100, 200)
	between100And200 := num >= 100 && num <= 200
	assert.That(t, between100And200, true)
}

func TestSeed(t *testing.T) {
	random.Seed(42)
	x := random.Num(100)
	y := random.Num(100)
	assert.That(t, x, y)
}
