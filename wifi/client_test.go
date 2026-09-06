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

func TestClientCountsWhatItSends(t *testing.T) {
	tests := []struct {
		name string
		send func(*wifi.Client) int
		want int
	}{
		{"Println counts the newline too", func(c *wifi.Client) int { return c.Println("GET /") }, 6},
		{"Write counts the characters", func(c *wifi.Client) int { return c.Write("GET /") }, 5},
		{"a number goes out as digits", func(c *wifi.Client) int { return c.Write(123) }, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c wifi.Client
			if got := tt.send(&c); got != tt.want {
				t.Errorf("sent %d characters, want %d", got, tt.want)
			}
		})
	}
}
