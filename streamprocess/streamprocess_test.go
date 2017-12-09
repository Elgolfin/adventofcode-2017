package streamprocess

import (
	"testing"
)

func TestScore(t *testing.T) {
	cases := []struct {
		in    string
		want1 int
		want2 int
	}{
		{"{}", 1, 0},
		{"{{{}}}", 6, 0},
		{"{{},{}}", 5, 0},
		{"{{{},{},{{}}}},", 16, 0},
		{"{<a>,<a>,<a>,<a>}", 1, 4},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9, 8},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9, 0},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3, 17},
	}
	for _, c := range cases {
		got1, got2 := Score(c.in)
		if got1 != c.want1 || got2 != c.want2 {
			t.Errorf("Expected Score(%s) to return %d, %d got %d, %d", c.in, c.want1, c.want2, got1, got2)
		}
	}
}
