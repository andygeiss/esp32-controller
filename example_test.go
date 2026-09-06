package esp32_controller_test

import (
	"github.com/andygeiss/esp32-controller"
	"github.com/andygeiss/esp32-controller/digital"
	"github.com/andygeiss/esp32-controller/serial"
	"github.com/andygeiss/esp32-controller/timer"
)

// blinker flashes the LED on one pin, once per call to Loop.
type blinker struct {
	pin int
}

func (b *blinker) Setup() error {
	serial.Begin(serial.BaudRate115200)
	digital.PinMode(b.pin, digital.ModeOutput)
	serial.Println("ready")
	return nil
}

func (b *blinker) Loop() error {
	digital.Write(b.pin, digital.High)
	timer.Delay(500)
	digital.Write(b.pin, digital.Low)
	timer.Delay(500)
	return nil
}

// Example writes the program esp32-transpiler turns into an Arduino sketch:
// Setup becomes setup(), and Loop becomes loop(), which the board calls over and
// over.
func Example() {
	var c esp32_controller.Controller = &blinker{pin: 2}

	if err := c.Setup(); err != nil {
		return
	}
	for range 3 {
		if err := c.Loop(); err != nil {
			return
		}
	}
}
