package digital_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/esp32-controller/digital"
	"testing"
)

func TestDigitalWrite(t *testing.T) {
	pin := 1
	digital.GPIOValues[pin] = digital.Low
	digital.Write(pin, digital.High)
	val := digital.GPIOValues[pin]
	assert.That("val is set to high", t, val, digital.High)
}

func TestPinMode(t *testing.T) {
	pin := 1
	digital.GPIOModes[pin] = digital.ModeInput
	digital.PinMode(pin, digital.ModeOutput)
	mode := digital.GPIOModes[pin]
	assert.That("mode is set to output", t, mode, digital.ModeOutput)
}
