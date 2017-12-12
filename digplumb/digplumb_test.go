package digplumb

import (
	"testing"

	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

func TestGetProgGroupCount(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`, 6},
	}
	for _, c := range cases {
		got := GetProgGroupCount(c.in, 0)
		if got != c.want {
			t.Errorf("Expected GetProgGroupCount(%q) to return %v got %v", c.in, c.want, got)
		}
	}
}

func TestParseLine(t *testing.T) {
	cases := []struct {
		in    string
		want1 int
		want2 []int
	}{
		{`0 <-> 2`, 0, []int{2}},
		{`2 <-> 0, 3, 4`, 2, []int{0, 3, 4}},
	}
	for _, c := range cases {
		got1, got2 := parseLine(c.in)
		if got1 != c.want1 || !sliceutil.EqualInt(c.want2, got2) {
			t.Errorf("Expected parseLine(%q) to return %v, %v got %v, %v", c.in, c.want1, c.want2, got1, got2)
		}
	}
}
