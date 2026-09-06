package serial_test

import (
	"testing"

	"github.com/andygeiss/esp32-controller/serial"
)

func TestAvailable(t *testing.T) {
	serial.AvailableN = 2
	if got := serial.Available(); got != 2 {
		t.Errorf("Available() = %d, want 2", got)
	}
}

func TestBegin(t *testing.T) {
	// Opening a port receives nothing: a sketch that waits on Available()
	// would spin here if Begin invented a byte.
	serial.AvailableN = 0
	serial.Baud = 0
	serial.Begin(serial.BaudRate115200)
	if serial.Baud != serial.BaudRate115200 {
		t.Errorf("Baud = %d, want %d", serial.Baud, serial.BaudRate115200)
	}
	if got := serial.Available(); got != 0 {
		t.Errorf("Available() = %d, want 0", got)
	}
}
