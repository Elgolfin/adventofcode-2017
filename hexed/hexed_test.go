package hexed

import (
	"testing"
)

func TestGetFewerSteps(t *testing.T) {
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
		got := GetFewerSteps(c.in)
		if got != c.want {
			t.Errorf("Expected GetFewerSteps(%q) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestGetFurthestSteps(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`ne,ne,ne`, 3},
		{`ne,ne,sw,sw`, 2},
		{`ne,ne,s,s`, 2},
		{`se,sw,se,sw,sw`, 3},
	}
	for _, c := range cases {
		got := GetFurthestSteps(c.in)
		if got != c.want {
			t.Errorf("Expected GetFurthestSteps(%q) to return %v, got %v", c.in, c.want, got)
		}
	}
}
