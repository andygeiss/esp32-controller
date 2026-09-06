package serial_test

import (
	"testing"

	"github.com/andygeiss/esp32-controller/serial"
)

func TestAvailable(t *testing.T) {
	// Begin sets AvailableN, and -shuffle=on may have run its test first.
	serial.AvailableN = 0
	if got := serial.Available(); got != 0 {
		t.Errorf("Available() = %d, want 0", got)
	}
}

func TestBegin(t *testing.T) {
	serial.AvailableN = 0
	serial.Baud = 0
	serial.Begin(serial.BaudRate115200)
	if serial.Baud != serial.BaudRate115200 {
		t.Errorf("Baud = %d, want %d", serial.Baud, serial.BaudRate115200)
	}
	if got := serial.Available(); got != 1 {
		t.Errorf("Available() = %d, want 1", got)
	}
}
