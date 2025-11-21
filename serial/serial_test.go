package serial_test

import (
	"testing"

	"github.com/andygeiss/cloud-native-utils/assert"
	"github.com/andygeiss/esp32-controller/serial"
)

func TestSerialAvailable(t *testing.T) {
	available := serial.Available()
	assert.That(t, "available must be 0", available, 0)
}

func TestSerialBegin(t *testing.T) {
	baud := serial.BaudRate115200
	serial.Baud = 0
	serial.Begin(baud)
	available := serial.Available()
	assert.That(t, "baud rate is 115200", serial.Baud, serial.BaudRate115200)
	assert.That(t, "available must be 1", available, 1)
}
