package conflagration

import (
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`set a 1
set b 2
mul a 2
sub b 4
mul b 0
mul c 2
sub d 3`, 3},
	}
	for _, c := range cases {
		p := InitializeProgram(0)
		p.Load(c.in)
		got := p.Run()
		if got != c.want {
			t.Errorf("Expected Run(%q) to return %v got %v", c.in, c.want, got)
		}
	}
}
