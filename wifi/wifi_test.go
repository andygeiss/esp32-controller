package wifi_test

import (
	"testing"

	"github.com/andygeiss/cloud-native-utils/assert"
	"github.com/andygeiss/esp32-controller/wifi"
)

func TestWifiBegin(t *testing.T) {
	ssid := "test"
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.Begin(ssid)
	assert.That(t, "current status is connected", wifi.CurrentStatus, wifi.StatusConnected)
}

func TestWifiBeginEncrypted(t *testing.T) {
	ssid := "test"
	passphrase := "passphrase"
	ipv4 := &wifi.IPAddress{127, 0, 0, 1}
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.BeginEncrypted(ssid, passphrase)
	assert.That(t, "current status is connected", wifi.CurrentStatus, wifi.StatusConnected)
	assert.That(t, "current local ip is ipv4", wifi.CurrentLocalIP, ipv4)
}
func TestWifiDisconnect(t *testing.T) {
	ssid := "test"
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.Begin(ssid)  // StatusConnected
	wifi.Disconnect() // back to idle?
	assert.That(t, "current status is idle", wifi.CurrentStatus, wifi.StatusIdle)
}

func TestWifiRSSIShouldBeNotMinusOne(t *testing.T) {
	ssid := "test"
	wifi.CurrentRSSI = -1
	wifi.Begin(ssid)
	assert.That(t, "RSSI is set to 0", wifi.RSSI(), 0)
}

func TestWifiSSIDShouldNotBeEmpty(t *testing.T) {
	ssid := "test"
	wifi.CurrentSSID = ""
	wifi.Begin(ssid)
	assert.That(t, "SSID is set to test", wifi.SSID(), "test")
}
