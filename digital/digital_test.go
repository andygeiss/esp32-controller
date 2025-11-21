package digital_test

import (
	"testing"

	"github.com/andygeiss/cloud-native-utils/assert"
	"github.com/andygeiss/esp32-controller/digital"
)

func TestDigitalWrite(t *testing.T) {
	pin := 1
	digital.GPIOValues[pin] = digital.Low
	digital.Write(pin, digital.High)
	val := digital.GPIOValues[pin]
	assert.That(t, "val is set to high", val, digital.High)
}

func TestPinMode(t *testing.T) {
	pin := 1
	digital.GPIOModes[pin] = digital.ModeInput
	digital.PinMode(pin, digital.ModeOutput)
	mode := digital.GPIOModes[pin]
	assert.That(t, "mode is set to output", mode, digital.ModeOutput)
}
