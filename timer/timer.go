// Package timer pauses a program.
//
// [Delay] is the whole package: it stops for a number of milliseconds, the way
// Arduino's delay() does.
package timer

import "time"

// Delay stops the program for ms milliseconds. It returns at once when ms is
// zero or less.
//
// See https://www.arduino.cc/reference/en/language/functions/time/delay/
func Delay(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
