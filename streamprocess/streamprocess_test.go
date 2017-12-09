package streamprocess

import (
	"testing"
)

func TestScore(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}},", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}
	for _, c := range cases {
		got := Score(c.in)
		if got != c.want {
			t.Errorf("Expected Score(%s) to return %d, got %d", c.in, c.want, got)
		}
	}
}
