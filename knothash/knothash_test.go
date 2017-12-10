package knothash

import (
	"testing"

	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

func TestHash(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`3`, 2},
		{`3,4`, 12},
		{`3,4,1,5`, 12},
	}
	for _, c := range cases {
		got := Hash(c.in, 5)
		if got != c.want {
			t.Errorf("Expected Hash(%q) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestGetStringCircle(t *testing.T) {
	cases := []struct {
		in   int
		want StringCircle
	}{
		{10, StringCircle{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 0}},
	}
	for _, c := range cases {
		got := GetStringCircle(c.in)
		if got.currentPosition != c.want.currentPosition || got.skipSize != c.want.skipSize || !sliceutil.EqualInt(got.list, c.want.list) {
			t.Errorf("Expected GetStringCircle(%d) to return %v, got %v", c.in, c.want, got)
		}
	}
}
