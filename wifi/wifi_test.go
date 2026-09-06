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
	want := []int{1, 2, 3, 4, 5, 6}
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

func TestIPAddressString(t *testing.T) {
	addr := wifi.IPAddress{A: 192, B: 168, C: 0, D: 10}
	if got := addr.String(); got != "192.168.0.10" {
		t.Errorf("String() = %q, want %q", got, "192.168.0.10")
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

// TestConstantsMatchArduino pins each constant to the number the board reports.
// The transpiler emits the C++ name, so only this test catches a value that
// drifts from Arduino's wl_status_t and wl_enc_type — as StatusConnectFailed
// and StatusConnectionLost had, swapped with each other.
func TestConstantsMatchArduino(t *testing.T) {
	tests := []struct {
		name string
		got  int
		want int
	}{
		{"WL_IDLE_STATUS", wifi.StatusIdle, 0},
		{"WL_NO_SSID_AVAIL", wifi.StatusNoSSIDAvailable, 1},
		{"WL_SCAN_COMPLETED", wifi.StatusScanCompleted, 2},
		{"WL_CONNECTED", wifi.StatusConnected, 3},
		{"WL_CONNECT_FAILED", wifi.StatusConnectFailed, 4},
		{"WL_CONNECTION_LOST", wifi.StatusConnectionLost, 5},
		{"WL_DISCONNECTED", wifi.StatusDisconnected, 6},
		{"WL_NO_SHIELD", wifi.StatusNoShield, 255},
		{"ENC_TYPE_TKIP", wifi.EncryptionTypeTKIP, 2},
		{"ENC_TYPE_CCMP", wifi.EncryptionTypeCCMP, 4},
		{"ENC_TYPE_WEP", wifi.EncryptionTypeWEP, 5},
		{"ENC_TYPE_NONE", wifi.EncryptionTypeNone, 7},
		{"ENC_TYPE_AUTO", wifi.EncryptionTypeAuto, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.got != tt.want {
				t.Errorf("%s = %d, want %d", tt.name, tt.got, tt.want)
			}
		})
	}
}
