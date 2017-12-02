package captcha

import (
	"testing"
)

// Get solves the puzzle of the Day 1
func TestGet(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}
	for _, c := range cases {
		got := Get(c.in)
		if got != c.want {
			t.Errorf("Expected Get(%q) to return %d, got %d", c.in, c.want, got)
		}
	}
}
