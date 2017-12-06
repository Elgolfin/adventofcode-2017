package captcha

import (
	"testing"
)

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

func TestGetNew(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}
	for _, c := range cases {
		got := GetNew(c.in)
		if got != c.want {
			t.Errorf("Expected GetNew(%q) to return %d, got %d\n", c.in, c.want, got)
		}
	}
}
