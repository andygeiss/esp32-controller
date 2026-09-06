package timer_test

import (
	"testing"
	"testing/synctest"
	"time"

	"github.com/andygeiss/esp32-controller/timer"
)

func TestDelay(t *testing.T) {
	tests := []struct {
		name string
		ms   int
		want time.Duration
	}{
		{"a delay of a second and a half", 1500, 1500 * time.Millisecond},
		{"a delay of nothing returns at once", 0, 0},
		{"a negative delay returns at once", -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// The bubble's clock is fake, so this asserts the exact pause
			// without waiting for it.
			synctest.Test(t, func(t *testing.T) {
				start := time.Now()
				timer.Delay(tt.ms)
				if got := time.Since(start); got != tt.want {
					t.Errorf("Delay(%d) paused %v, want %v", tt.ms, got, tt.want)
				}
			})
		})
	}
}
