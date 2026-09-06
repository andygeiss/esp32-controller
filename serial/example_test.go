package serial_test

import (
	"github.com/andygeiss/esp32-controller/serial"
)

// Example opens the port and sends a line of text.
func Example() {
	serial.Begin(serial.BaudRate115200)
	serial.Println("ready")
	// Output: ready
}
