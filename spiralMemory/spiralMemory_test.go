package spiralMemory

import (
	"testing"
)

// Get solves the puzzle of the Day 1
func TestGetSteps(t *testing.T) {
	cases := []struct {
		in   float64
		want float64
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}
	for _, c := range cases {
		got := GetSteps(c.in)
		if got != c.want {
			t.Errorf("Expected GetSteps(%f) to return %f, got %f", c.in, c.want, got)
		}
	}
}
