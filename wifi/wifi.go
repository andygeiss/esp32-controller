// Package wifi joins the ESP32 to a WiFi network and opens connections over it.
//
// [Begin] joins an open network and [BeginEncrypted] joins a protected one; both
// leave the result in [Status]. [Client.Connect] then opens a connection to a
// server.
//
// Network state lives in the Current* variables, the way an Arduino sketch keeps
// it in globals. Nothing here is safe to use from more than one goroutine.
//
// See https://www.arduino.cc/en/Reference/WiFi
package wifi

import "fmt"

// IPAddress is an IPv4 address, one field per dotted part: 127.0.0.1 is
// {127, 0, 0, 1}.
type IPAddress struct {
	A int
	B int
	C int
	D int
}

// String returns the address in dotted form, so serial.Println(wifi.LocalIP())
// prints what the sketch's Serial.println(WiFi.localIP()) prints.
func (a IPAddress) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a.A, a.B, a.C, a.D)
}

// How a network protects itself, as [EncryptionType] reports it. The numbers are
// Arduino's wl_enc_type.
const (
	EncryptionTypeAuto = 8
	EncryptionTypeCCMP = 4
	EncryptionTypeNone = 7
	EncryptionTypeTKIP = 2
	EncryptionTypeWEP  = 5
)

// Where the board stands with the network, as [Status] reports it. The numbers
// are Arduino's wl_status_t, so a sketch and this package agree.
const (
	StatusConnected       = 3
	StatusConnectFailed   = 4
	StatusConnectionLost  = 5
	StatusDisconnected    = 6
	StatusIdle            = 0
	StatusNoShield        = 255
	StatusNoSSIDAvailable = 1
	StatusScanCompleted   = 2
)

var (
	// CurrentBSSID is the MAC address [BSSID] reports, one field per octet.
	CurrentBSSID = []int{0, 0, 0, 0, 0, 0}
	// CurrentDNS is the name server [SetDNS] last set.
	CurrentDNS = []int{0, 0, 0, 0}
	// CurrentEncryptionType is the protection [EncryptionType] reports.
	CurrentEncryptionType = EncryptionTypeNone
	// CurrentLocalIP is the board's own address, as [LocalIP] reports it.
	CurrentLocalIP = &IPAddress{127, 0, 0, 1}
	// CurrentNetworks is the count [ScanNetworks] reports.
	CurrentNetworks = 0
	// CurrentRSSI is the signal strength [RSSI] reports, in decibels.
	CurrentRSSI = -1
	// CurrentSSID is the network name [SSID] reports.
	CurrentSSID = ""
	// CurrentStatus is the value [Status] reports.
	CurrentStatus = StatusIdle
)

// BSSID returns the MAC address of the router the board is joined to, one
// field per octet.
//
// See https://www.arduino.cc/en/Reference/WiFiBSSID
func BSSID() []int {
	return CurrentBSSID
}

// Begin joins the open network named ssid. Read [Status] for the result.
//
// See https://www.arduino.cc/en/Reference/WiFiBegin
func Begin(ssid string) {
	CurrentRSSI = 0
	CurrentSSID = ssid
	CurrentStatus = StatusConnected
}

// BeginEncrypted joins the protected network named ssid using passphrase.
// Read [Status] for the result.
//
// See https://www.arduino.cc/en/Reference/WiFiBegin
func BeginEncrypted(ssid, passphrase string) {
	CurrentRSSI = 0
	CurrentSSID = ssid
	CurrentStatus = StatusConnected
}

// Disconnect leaves the current network, putting [Status] back to StatusIdle.
//
// See https://www.arduino.cc/en/Reference/WiFiDisconnect
func Disconnect() {
	CurrentStatus = StatusIdle
}

// EncryptionType returns how the current network protects itself, as one of the
// EncryptionType constants.
//
// See https://www.arduino.cc/en/Reference/WiFiEncryptionType
func EncryptionType() int {
	return CurrentEncryptionType
}

// LocalIP returns the board's own address on the network.
//
// See https://www.arduino.cc/en/Reference/WiFiLocalIP
func LocalIP() *IPAddress {
	return CurrentLocalIP
}

// RSSI returns the signal strength to the router in decibels. Closer to zero
// is stronger.
//
// See https://www.arduino.cc/en/Reference/WiFiRSSI
func RSSI() int {
	return CurrentRSSI
}

// ScanNetworks looks for networks in range and returns how many it found.
//
// See https://www.arduino.cc/en/Reference/WiFiScanNetworks
func ScanNetworks() int {
	return CurrentNetworks
}

// SetDNS points the board at the name server dns, one field per dotted part.
//
// See https://www.arduino.cc/en/Reference/WiFiSetDns
func SetDNS(dns []int) {
	CurrentDNS = dns
}

// Status returns where the board stands with the network, as one of the Status
// constants.
//
// See https://www.arduino.cc/en/Reference/WiFiStatus
func Status() int {
	return CurrentStatus
}

// SSID returns the name of the network the board is joined to.
//
// See https://www.arduino.cc/en/Reference/WiFiSSID
func SSID() string {
	return CurrentSSID
}
