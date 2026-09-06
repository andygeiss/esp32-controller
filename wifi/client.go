package wifi

import "fmt"

// Client is a connection to a server, opened with [Client.Connect]. Its zero
// value is ready to use.
type Client struct{}

// Connect opens a connection to host on port, and reports whether it worked.
// host may be a name, which the board looks up first. Only the sketch opens a
// real connection, so the Go side always reports success.
//
// See https://www.arduino.cc/en/Reference/WiFiClientConnect
func (c *Client) Connect(host string, port int) bool {
	return true
}

// Println sends data to the server, followed by a newline, and returns the
// number of characters sent. Numbers go out as digits: 123 becomes '1', '2',
// '3'. The sketch's client.println sends a carriage return and a newline.
//
// See https://www.arduino.cc/en/Reference/WiFiClientPrintln
func (c *Client) Println(data any) int {
	n, _ := fmt.Println(data)
	return n
}

// Write sends data to the server and returns the number of characters sent.
// Most sketches ignore the count.
//
// See https://www.arduino.cc/en/Reference/WiFiClientWrite
func (c *Client) Write(data any) int {
	n, _ := fmt.Print(data)
	return n
}
