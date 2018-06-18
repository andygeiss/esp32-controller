package wifi

// Client ...
type Client struct {}

// Connect to the IP address and port specified in the constructor.
// The return value indicates success or failure.
// connect() also supports DNS lookups when using a domain name (ex:google.com).
// @see: https://www.arduino.cc/en/Reference/WiFiClientConnect
func (c *Client) Connect(host string, port int) int {
	return StatusConnected
}
