package spinlock

import (
	"testing"
)

func TestProcess(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{3, 638},
	}
	for _, c := range cases {
		got := Process(c.in)
		if got != c.want {
			t.Errorf("Expected Process(%d) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestAngryProcess(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{304, 1930815},
	}
	for _, c := range cases {
		got := AngryProcess(c.in)
		if got != c.want {
			t.Errorf("Expected AngryProcess(%d) to return %v, got %v", c.in, c.want, got)
		}
	}
}
