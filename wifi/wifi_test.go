package wifi_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/esp32-controller/wifi"
	"testing"
)

func TestWifiBegin(t *testing.T) {
	ssid := "test"
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.Begin(ssid)
	assert.That("current status is connected", t, wifi.CurrentStatus, wifi.StatusConnected)
}

func TestWifiBeginEncrypted(t *testing.T) {
	ssid := "test"
	passphrase := "passphrase"
	ipv4 := &wifi.IPAddress{127, 0, 0, 1}
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.BeginEncrypted(ssid, passphrase)
	assert.That("current status is connected", t, wifi.CurrentStatus, wifi.StatusConnected)
	assert.That("current local ip is ipv4", t, wifi.CurrentLocalIP, ipv4)
}
func TestWifiDisBegin(t *testing.T) {
	ssid := "test"
	wifi.CurrentStatus = wifi.StatusIdle
	wifi.Begin(ssid)  // StatusConnected
	wifi.Disconnect() // back to idle?
	assert.That("current status is idle", t, wifi.CurrentStatus, wifi.StatusIdle)
}

func TestWifiRSSIShouldBeNotMinusOne(t *testing.T) {
	ssid := "test"
	wifi.CurrentRSSI = -1
	wifi.Begin(ssid)
	assert.That("RSSI is set to 0", t, wifi.RSSI(), 0)
}
func TestWifiSSIDShouldNotBeEmpty(t *testing.T) {
	ssid := "test"
	wifi.CurrentSSID = ""
	wifi.Begin(ssid)
	assert.That("SSID is set to test", t, wifi.SSID(), "test")
}
