package wifi

import "fmt"

// Client ...
type Client struct{}

// Connect to the IP address and port specified in the constructor.
// The return value indicates success or failure.
// connect() also supports DNS lookups when using a domain name (ex:google.com).
// @see: https://www.arduino.cc/en/Reference/WiFiClientConnect
func (c *Client) Connect(host string, port int) bool {
	return true
}

// Print data, followed by a carriage return and newline, to the server a client is connected to.
// Prints numbers as a sequence of digits, each an ASCII character (e.g. the number 123 is sent as the three characters '1', '2', '3').
// @see: https://www.arduino.cc/en/Reference/WiFiClientPrintln
func (c *Client) Println(data interface{}) int {
	fmt.Println(data)
	return 0
}

// Write data to the server the client is connected to.
// Return the number of characters written.
// It is not necessary to read this value.
// @see: https://www.arduino.cc/en/Reference/WiFiClientWrite
func (c *Client) Write(data interface{}) int {
	fmt.Print(data)
	return 0
}
