package ddefrag

import (
	"testing"
)

func TestCountSquares(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"flqrgnkx", 8108},
	}
	for _, c := range cases {
		got := CountSquares(c.in)
		if got != c.want {
			t.Errorf("Expected CountSquares(%q) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestConvertHexaToBits(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"012", 2},
	}
	for _, c := range cases {
		got := OnesCount(c.in)
		if got != c.want {
			t.Errorf("Expected OnesCount(%s) to return %d, got %d", c.in, c.want, got)
		}
	}
}
