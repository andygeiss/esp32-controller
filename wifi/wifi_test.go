package wifi_test

import (
	"slices"
	"testing"

	"github.com/andygeiss/esp32-controller/wifi"
)

func TestBegin(t *testing.T) {
	wifi.CurrentRSSI = -1
	wifi.CurrentSSID = ""
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.Begin("test")
	if got := wifi.Status(); got != wifi.StatusConnected {
		t.Errorf("Status() = %d, want %d", got, wifi.StatusConnected)
	}
	if got := wifi.SSID(); got != "test" {
		t.Errorf("SSID() = %q, want %q", got, "test")
	}
	if got := wifi.RSSI(); got != 0 {
		t.Errorf("RSSI() = %d, want 0", got)
	}
}

func TestBeginEncrypted(t *testing.T) {
	wifi.CurrentSSID = ""
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.BeginEncrypted("test", "passphrase")
	if got := wifi.Status(); got != wifi.StatusConnected {
		t.Errorf("Status() = %d, want %d", got, wifi.StatusConnected)
	}
	if got := wifi.SSID(); got != "test" {
		t.Errorf("SSID() = %q, want %q", got, "test")
	}
}

func TestDisconnect(t *testing.T) {
	wifi.Begin("test")
	wifi.Disconnect()
	if got := wifi.Status(); got != wifi.StatusIdle {
		t.Errorf("Status() = %d, want %d", got, wifi.StatusIdle)
	}
}

func TestBSSID(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}
	wifi.CurrentBSSID = want
	if got := wifi.BSSID(); !slices.Equal(got, want) {
		t.Errorf("BSSID() = %v, want %v", got, want)
	}
}

func TestEncryptionType(t *testing.T) {
	wifi.CurrentEncryptionType = wifi.EncryptionTypeCCMP
	if got := wifi.EncryptionType(); got != wifi.EncryptionTypeCCMP {
		t.Errorf("EncryptionType() = %d, want %d", got, wifi.EncryptionTypeCCMP)
	}
}

func TestLocalIP(t *testing.T) {
	want := wifi.IPAddress{A: 192, B: 168, C: 0, D: 10}
	wifi.CurrentLocalIP = &want
	if got := wifi.LocalIP(); *got != want {
		t.Errorf("LocalIP() = %v, want %v", *got, want)
	}
}

func TestScanNetworks(t *testing.T) {
	wifi.CurrentNetworks = 3
	if got := wifi.ScanNetworks(); got != 3 {
		t.Errorf("ScanNetworks() = %d, want 3", got)
	}
}

func TestSetDNS(t *testing.T) {
	want := []int{8, 8, 8, 8}
	wifi.CurrentDNS = nil
	wifi.SetDNS(want)
	if !slices.Equal(wifi.CurrentDNS, want) {
		t.Errorf("CurrentDNS = %v, want %v", wifi.CurrentDNS, want)
	}
}

func TestHostByNameAlwaysFails(t *testing.T) {
	// The real lookup only happens in the generated sketch, so the Go side
	// reports failure rather than pretending it resolved.
	if got := wifi.HostByName("example.com", ""); got != 0 {
		t.Errorf("HostByName() = %d, want 0", got)
	}
}
