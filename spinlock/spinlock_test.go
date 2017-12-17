package spinlock

import (
	"testing"
)

func TestDance(t *testing.T) {
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
