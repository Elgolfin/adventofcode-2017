package helper

import (
	"testing"

	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

func TestPermut(t *testing.T) {
	cases := []struct {
		in   []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{
			[]int{1, 2, 3},
			[]int{2, 1, 3},
			[]int{3, 2, 1},
			[]int{2, 3, 1},
			[]int{3, 1, 2},
			[]int{1, 3, 2},
		}},
		{[]int{}, [][]int{}},
	}
	for _, c := range cases {
		got := Permut(c.in)
		if !sliceutil.Equal2DInt(got, c.want) {
			t.Errorf("Expected Permut(%v) to return %v got %v", c.in, c.want, got)
		}
	}
}
