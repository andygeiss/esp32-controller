package random_test

import (
	"testing"

	"github.com/andygeiss/esp32-controller/random"
)

func TestNumStaysUnderMax(t *testing.T) {
	for range 1000 {
		if got := random.Num(10); got < 0 || got >= 10 {
			t.Fatalf("Num(10) = %d, want 0 to 9", got)
		}
	}
}

func TestNumBetweenStaysInRange(t *testing.T) {
	tests := []struct {
		name     string
		min, max int
	}{
		// (10, 100) is the case the old implementation got wrong: it drew from
		// 90 to 99 instead of 10 to 99.
		{"a range wider than its start", 10, 100},
		{"a range as wide as its start", 100, 200},
		{"a range of one", 5, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			low, high := tt.max, tt.min
			for range 1000 {
				got := random.NumBetween(tt.min, tt.max)
				if got < tt.min || got >= tt.max {
					t.Fatalf("NumBetween(%d, %d) = %d, want %d to %d", tt.min, tt.max, got, tt.min, tt.max-1)
				}
				low, high = min(low, got), max(high, got)
			}
			// 1000 draws over a range this small should reach both ends.
			if low != tt.min || high != tt.max-1 {
				t.Errorf("1000 draws covered %d to %d, want %d to %d", low, high, tt.min, tt.max-1)
			}
		})
	}
}

func TestSeedRepeatsTheSequence(t *testing.T) {
	draw := func() []int {
		random.Seed(42)
		return []int{random.Num(1000), random.Num(1000), random.Num(1000)}
	}
	first, second := draw(), draw()
	for i := range first {
		if first[i] != second[i] {
			t.Fatalf("after the same seed, draw %d was %d then %d", i, first[i], second[i])
		}
	}
}

func TestSeedChangesTheSequence(t *testing.T) {
	random.Seed(1)
	first := []int{random.Num(1000), random.Num(1000), random.Num(1000)}
	random.Seed(2)
	second := []int{random.Num(1000), random.Num(1000), random.Num(1000)}
	same := true
	for i := range first {
		if first[i] != second[i] {
			same = false
		}
	}
	if same {
		t.Errorf("seeds 1 and 2 drew the same sequence %v", first)
	}
}
