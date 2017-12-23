package sporifica

import (
	"testing"

	"github.com/elgolfin/adventofcode-2017/maputil"
)

func TestBurst(t *testing.T) {
	cases := []struct {
		in1  string
		in2  int
		want int
	}{
		{`..#
#..
...`, 7, 5},
		{`..#
#..
...`, 70, 41},
		{`..#
#..
...`, 10000, 5587},
	}
	for _, c := range cases {
		got := Burst(c.in1, c.in2)
		if got != c.want {
			t.Errorf("Expected Burst(%s, %d) to return %v got %v", c.in1, c.in2, c.want, got)
		}
	}
}

func TestLoadGrid(t *testing.T) {
	cases := []struct {
		in   string
		want map[string]string
	}{
		{`..#
#..
...`, map[string]string{
			"-1,1":  ".",
			"0,1":   ".",
			"1,1":   "#",
			"-1,0":  "#",
			"0,0":   ".",
			"1,0":   ".",
			"-1,-1": ".",
			"0,-1":  ".",
			"1,-1":  ".",
		}},
	}
	for _, c := range cases {
		got := loadGrid(c.in)
		if !maputil.EqualStringString(got, c.want) {
			t.Errorf("Expected loadGrid(%v) to return %v got %v", c.in, c.want, got)
		}
	}
}
