package digplumb

import (
	"testing"

	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

func TestGetProgGroupCount(t *testing.T) {
	cases := []struct {
		in    string
		want1 int // number of program groups
		want2 int // number og programs in the group of progID 0
	}{
		{`0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`, 2, 6},
	}
	for _, c := range cases {
		n, progGroups := GetAllProgGroups(c.in)
		if n != c.want1 || progGroups[0] != c.want2 {
			t.Errorf("Expected GetAllProgGroups(%q) to return %d, %d got %d, %d", c.in, c.want1, c.want2, n, progGroups[0])
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
