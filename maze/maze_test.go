package maze

import (
	"testing"
)

func TestFindExit(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`0
3
0
1
-3`, 5},
	}
	for _, c := range cases {
		got := FindExit(c.in)
		if got != c.want {
			t.Errorf("Expected FindExit(%q) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestFindExitSranger(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`0
3
0
1
-3`, 10},
	}
	for _, c := range cases {
		got := FindExitSranger(c.in)
		if got != c.want {
			t.Errorf("Expected FindExitSranger(%q) to return %v, got %v", c.in, c.want, got)
		}
	}
}
