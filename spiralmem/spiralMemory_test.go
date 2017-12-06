package spiralmem

import (
	"testing"
)

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

func TestGetFirstValueLargerThan(t *testing.T) {
	cases := []struct {
		in   float64
		want float64
	}{
		{750, 806},
		{1, 2},
		{145, 147},
	}
	for _, c := range cases {
		got := GetFirstValueLargerThan(c.in)
		if got != c.want {
			t.Errorf("Expected GetFirstValueLargerThan(%f) to return %f, got %f", c.in, c.want, got)
		}
	}
}
