// Package wifi joins the ESP32 to a WiFi network and opens connections over it.
//
// [Begin] joins an open network and [BeginEncrypted] joins a protected one; both
// leave the result in [Status]. [Client] then opens a connection to a server.
//
// Network state lives in the Current* variables, the way an Arduino sketch keeps
// it in globals. Nothing here is safe to use from more than one goroutine.
//
// See https://www.arduino.cc/en/Reference/WiFi
package wifi

// IPAddress is an IPv4 address, one field per dotted part: 127.0.0.1 is
// {127, 0, 0, 1}.
type IPAddress struct {
	A int
	B int
	C int
	D int
}

// How a network protects itself, as [EncryptionType] reports it.
const (
	EncryptionTypeAuto = 8
	EncryptionTypeCCMP = 4
	EncryptionTypeNone = 7
	EncryptionTypeTKIP = 2
	EncryptionTypeWEP  = 5
)

// Where the board stands with the network, as [Status] reports it.
const (
	StatusConnected       = 3
	StatusConnectionLost  = 4
	StatusConnectFailed   = 5
	StatusDisconnected    = 6
	StatusIdle            = 0
	StatusNoShield        = 255
	StatusNoSSIDAvailable = 1
	StatusScanCompleted   = 2
)

// MaxSocketNum is how many connections the board tracks at once.
const MaxSocketNum = 4096

var (
	// CurrentBSSID is the MAC address [BSSID] reports.
	CurrentBSSID = []int{0, 0, 0, 0, 0, 0, 0, 0}
	// CurrentDNS is the name server [SetDNS] last set.
	CurrentDNS = []int{0, 0, 0, 0}
	// CurrentEncryptionType is the protection [EncryptionType] reports.
	CurrentEncryptionType = EncryptionTypeNone
	// CurrentGateway is the router the board sends other traffic through.
	CurrentGateway = &IPAddress{127, 0, 0, 255}
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
	// SocketPort holds the port each open connection uses.
	SocketPort = make(map[int]int, MaxSocketNum)
	// SocketState holds how far along each open connection is.
	SocketState = make(map[int]int, MaxSocketNum)
)

// BSSID returns the MAC address of the router the board is joined to.
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

// HostByName looks hostname up and returns 1 when it resolved. The board writes
// the address into addr; this Go version cannot, because addr is a copy.
//
// See https://www.arduino.cc/en/Reference/WiFiHostByName
func HostByName(hostname string, addr string) int {
	return 0
}

// LocalIP returns the board's own address on the network.
//
// See https://www.arduino.cc/en/Reference/WiFiLocalIP
func LocalIP() *IPAddress {
	return CurrentLocalIP
}

// RSSI returns the signal strength to the router in decibels. It is negative,
// and closer to zero is stronger.
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
