// Package digital reads and writes the ESP32's general-purpose IO pins.
//
// [PinMode] sets a pin to input or output, and [Write] drives an output pin high
// or low. Both ignore a pin number [IsPinValid] rejects, because an Arduino
// sketch has nowhere to report an error to.
//
// Pin state lives in [GPIOModes] and [GPIOValues], the way a sketch keeps it in
// globals. Nothing here is safe to use from more than one goroutine.
package digital

const (
	// High is the level of a pin that is switched on.
	High = 1
	// Low is the level of a pin that is switched off.
	Low = 0
	// ModeInput makes a pin read a voltage.
	ModeInput = 0
	// ModeOutput makes a pin drive a voltage.
	ModeOutput = 1
	// PinsMax is the highest pin number this package acts on.
	PinsMax = 48
)

var (
	// GPIOModes holds the mode [PinMode] last set, per pin.
	GPIOModes = make(map[int]int, PinsMax+1)
	// GPIOValues holds the value [Write] last set, per pin.
	GPIOValues = make(map[int]int, PinsMax+1)
)

// IsPinValid reports whether pin is a pin number this package acts on: GPIO 0
// up to and including [PinsMax].
func IsPinValid(pin int) bool {
	return pin >= 0 && pin <= PinsMax
}

// PinMode sets pin to read a voltage (ModeInput) or drive one (ModeOutput).
// It does nothing when IsPinValid rejects pin.
//
// See https://www.arduino.cc/reference/en/language/functions/digital-io/pinmode/
func PinMode(pin, mode int) {
	if IsPinValid(pin) {
		GPIOModes[pin] = mode
	}
}

// Write sets pin to High or Low. It does nothing when IsPinValid rejects pin.
//
// See https://www.arduino.cc/reference/en/language/functions/digital-io/digitalwrite/
func Write(pin, value int) {
	if IsPinValid(pin) {
		GPIOValues[pin] = value
	}
}
