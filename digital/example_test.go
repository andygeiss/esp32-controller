package digital_test

import (
	"fmt"

	"github.com/andygeiss/esp32-controller/digital"
)

// Example switches an LED on: name the pin an output, then drive it high.
func Example() {
	pin := 2
	digital.PinMode(pin, digital.ModeOutput)
	digital.Write(pin, digital.High)

	fmt.Println(digital.GPIOValues[pin])
	// Output: 1
}
