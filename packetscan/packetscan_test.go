package packetscan

import (
	"testing"

	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

func TestGoThroughTheFirewall(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`0: 3
1: 2
4: 4
6: 4`, 24},
	}
	for _, c := range cases {
		got := GoThroughTheFirewall(c.in)
		if got != c.want {
			t.Errorf("Expected GoThroughTheFirewall(%q) to return %d, got %d", c.in, c.want, got)
		}
	}
}

func TestParseLine(t *testing.T) {
	cases := []struct {
		in    string
		want1 int
		want2 int
	}{
		{`0: 2`, 0, 2},
		{`88: 43`, 88, 43},
	}
	for _, c := range cases {
		got1, got2 := parseLine(c.in)
		if got1 != c.want1 || got2 != c.want2 {
			t.Errorf("Expected parseLine(%q) to return %v, %v got %v, %v", c.in, c.want1, c.want2, got1, got2)
		}
	}
}

func TestInitializeFirewall(t *testing.T) {
	cases := []struct {
		in    string
		want1 int
		want2 []int
		want3 int
	}{
		{`0: 3
1: 2
4: 4
6: 4`, 4, []int{0, 0, -1, -1, 0, -1, 0}, 0},
	}
	for _, c := range cases {
		got := InitializeFirewall(c.in)
		if len(got.layers) != c.want1 || !sliceutil.EqualInt(got.scanners, c.want2) || got.picosecond != c.want3 {
			t.Errorf("Expected InitializeFirewall(%q) to return %d, %v got %d %v, %v", c.in, c.want1, c.want2, len(got.layers), got.layers, got.scanners)
		}
	}
}

func TestNextPicosecond(t *testing.T) {
	in := `0: 3
1: 2
4: 4
6: 4`
	cases := []struct {
		in   int
		want []int
	}{
		// {0, []int{0, 0, -1, -1, 0, -1, 0}},
		// {1, []int{1, 1, -1, -1, 1, -1, 1}},
		// {2, []int{2, 0, -1, -1, 2, -1, 2}},
		// {3, []int{1, 1, -1, -1, 3, -1, 3}},
		{4, []int{0, 0, -1, -1, 2, -1, 2}},
	}
	for _, c := range cases {
		got := InitializeFirewall(in)
		for i := 1; i <= c.in; i++ {
			got.NextPicosecond()
		}
		if !sliceutil.EqualInt(got.scanners, c.want) || got.picosecond != c.in {
			t.Errorf("Expected NextPicosecond() times %d to return %v, %d got %v, %d, %v", c.in, c.want, c.in, got.scanners, got.picosecond, got.layers)
		}
	}
}
