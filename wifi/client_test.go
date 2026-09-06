package wifi_test

import (
	"testing"

	"github.com/andygeiss/esp32-controller/wifi"
)

func TestClientConnect(t *testing.T) {
	var c wifi.Client
	if !c.Connect("example.com", 80) {
		t.Error("Connect() = false, want true")
	}
}
