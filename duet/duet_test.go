package duet

import (
	"testing"
)

func TestPlayMusicTilNonZeroRcv(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2`, 4},
	}
	for _, c := range cases {
		got := PlayMusicTilNonZeroRcv(c.in)
		if got != c.want {
			t.Errorf("Expected PlayMusicTilNonZeroRcv(%q) to return %v got %v", c.in, c.want, got)
		}
	}
}
