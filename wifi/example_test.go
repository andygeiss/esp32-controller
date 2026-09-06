package wifi_test

import (
	"fmt"

	"github.com/andygeiss/esp32-controller/serial"
	"github.com/andygeiss/esp32-controller/wifi"
)

// Example joins a protected network, then reads back what the board is on.
func Example() {
	wifi.BeginEncrypted("my-network", "my-passphrase")
	if wifi.Status() != wifi.StatusConnected {
		return
	}

	fmt.Println(wifi.SSID())
	// Output: my-network
}

// ExampleLocalIP prints the board's own address the way the sketch does.
func ExampleLocalIP() {
	wifi.CurrentLocalIP = &wifi.IPAddress{A: 192, B: 168, C: 0, D: 10}

	serial.Println(wifi.LocalIP())
	// Output: 192.168.0.10
}

// ExampleClient sends one line to a server, and prints how many characters went
// out.
func ExampleClient() {
	var client wifi.Client
	if !client.Connect("example.com", 80) {
		return
	}
	n := client.Println("GET /")

	fmt.Println(n)
	// Output:
	// GET /
	// 6
}
