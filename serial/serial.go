// Package serial sends text over the ESP32's serial port.
//
// [Begin] picks the speed, then [Print] and [Println] send. In the generated
// sketch these become Serial.begin, Serial.print and Serial.println; while the
// Go program runs on a computer they write to standard output.
//
// Port state lives in [AvailableN] and [Baud], the way a sketch keeps it in
// globals. Nothing here is safe to use from more than one goroutine.
package serial

import "fmt"

// Speeds in bits per second that both ends of the port can agree on. Pass one to
// [Begin]; BaudRate9600 and BaudRate115200 are the usual choices.
const (
	BaudRate300    = 300
	BaudRate600    = 600
	BaudRate1200   = 1200
	BaudRate2400   = 2400
	BaudRate4800   = 4800
	BaudRate9600   = 9600
	BaudRate14400  = 14400
	BaudRate19200  = 19200
	BaudRate28800  = 28800
	BaudRate38400  = 38400
	BaudRate57600  = 57600
	BaudRate115200 = 115200
)

var (
	// AvailableN is the number of bytes [Available] reports as waiting. Nothing
	// in this package puts bytes there; set it to stand in for input the board
	// would have received.
	AvailableN = 0
	// Baud is the speed [Begin] last set, or 0 before the first call.
	Baud = 0
)

// Available returns the number of bytes waiting to be read from the port.
//
// See https://www.arduino.cc/reference/en/language/functions/communication/serial/available/
func Available() int {
	return AvailableN
}

// Begin opens the port at baud bits per second. Pass one of the BaudRate
// constants. Opening the port receives nothing, so it leaves [AvailableN] alone.
//
// See https://www.arduino.cc/reference/en/language/functions/communication/serial/begin/
func Begin(baud int) {
	Baud = baud
}

// Print sends val as readable text, with nothing after it.
//
// See https://www.arduino.cc/reference/en/language/functions/communication/serial/print/
func Print(val any) {
	fmt.Print(val)
}

// Println sends val as readable text, followed by a newline. The sketch's
// Serial.println sends a carriage return and a newline.
//
// See https://www.arduino.cc/reference/en/language/functions/communication/serial/println/
func Println(val any) {
	fmt.Println(val)
}
