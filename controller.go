// Package esp32_controller declares the shape of a program that runs on an ESP32.
//
// A program implements [Controller]: Setup runs once when the board powers up,
// then Loop runs over and over. The other packages in this module mirror the
// Arduino calls such a program makes — digital, random, serial, timer and wifi —
// so esp32-transpiler can rewrite the Go source into an Arduino sketch.
//
// See https://github.com/andygeiss/esp32-transpiler
package esp32_controller

// Controller is a program the board runs. It becomes the setup() and loop()
// functions of the generated sketch.
type Controller interface {
	// Loop does the program's work. The board calls it again as soon as it returns.
	Loop() error
	// Setup prepares the board: pin modes, serial speed, WiFi. It runs once.
	Setup() error
}
