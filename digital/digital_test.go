package digital_test

import (
	"testing"

	"github.com/andygeiss/esp32-controller/digital"
)

func TestIsPinValid(t *testing.T) {
	tests := []struct {
		name string
		pin  int
		want bool
	}{
		{"a pin in the middle is valid", 1, true},
		// GPIO 0 is a real pin, and the sketch's pinMode(0, ...) acts on it.
		{"the lowest pin is valid", 0, true},
		{"the highest pin is valid", digital.PinsMax, true},
		{"one past the highest is not", digital.PinsMax + 1, false},
		{"a negative pin is not", -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := digital.IsPinValid(tt.pin); got != tt.want {
				t.Errorf("IsPinValid(%d) = %v, want %v", tt.pin, got, tt.want)
			}
		})
	}
}

func TestPinMode(t *testing.T) {
	pin := 1
	digital.GPIOModes[pin] = digital.ModeInput
	digital.PinMode(pin, digital.ModeOutput)
	if got := digital.GPIOModes[pin]; got != digital.ModeOutput {
		t.Errorf("GPIOModes[%d] = %v, want %v", pin, got, digital.ModeOutput)
	}
}

func TestPinModeIgnoresAnInvalidPin(t *testing.T) {
	pin := digital.PinsMax + 1
	delete(digital.GPIOModes, pin)
	digital.PinMode(pin, digital.ModeOutput)
	if _, ok := digital.GPIOModes[pin]; ok {
		t.Errorf("GPIOModes[%d] was set, want it left alone", pin)
	}
}

func TestWrite(t *testing.T) {
	pin := 1
	digital.GPIOValues[pin] = digital.Low
	digital.Write(pin, digital.High)
	if got := digital.GPIOValues[pin]; got != digital.High {
		t.Errorf("GPIOValues[%d] = %v, want %v", pin, got, digital.High)
	}
}

func TestWriteIgnoresAnInvalidPin(t *testing.T) {
	pin := -1
	delete(digital.GPIOValues, pin)
	digital.Write(pin, digital.High)
	if _, ok := digital.GPIOValues[pin]; ok {
		t.Errorf("GPIOValues[%d] was set, want it left alone", pin)
	}
}
