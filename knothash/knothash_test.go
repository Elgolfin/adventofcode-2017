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
			t.Errorf("Expected Hash(%q, 5) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestFullHash(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{`1,2,3`, "3efbe78a8d82f29979031a4aa0b16a9d"},
		{`AoC 2017`, "33efeb34ea91902bb2f59c9920caa6cd"},
		{`1,2,4`, "63960835bcdc130f0b66d7ff4f6a5a8e"},
		{``, "a2582a3a0e66e6e86e3812dcb672a272"},
	}
	for _, c := range cases {
		got := FullHash(c.in, 256)
		if got != c.want {
			t.Errorf("Expected FullHash(%q, 256) to return %s, got %s", c.in, c.want, got)
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
