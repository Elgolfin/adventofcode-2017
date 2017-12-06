package memrealloc

import (
	"testing"
)

func TestRedistribute(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"0	2	7	0", 5},
	}
	for _, c := range cases {
		got, _, _ := Redistribute(c.in)
		if got != c.want {
			t.Errorf("Expected Redistribute(%q) to return %d, got %d", c.in, c.want, got)
		}
	}
}

func TestGetLoopSize(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"0	2	7	0", 4},
	}
	for _, c := range cases {
		got := GetLoopSize(c.in)
		if got != c.want {
			t.Errorf("Expected GetLoopSize(%q) to return %d, got %d", c.in, c.want, got)
		}
	}
}
