package hexed

import (
	"testing"
)

func TestCountSteps(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`ne,ne,ne`, 3},
		{`ne,ne,sw,sw`, 0},
		{`ne,ne,s,s`, 2},
		{`se,sw,se,sw,sw`, 3},
	}
	for _, c := range cases {
		got := CountSteps(c.in)
		if got != c.want {
			t.Errorf("Expected CountSteps(%q) to return %v, got %v", c.in, c.want, got)
		}
	}
}
