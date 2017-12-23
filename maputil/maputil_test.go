package maputil

import (
	"testing"
)

func TestEqualStringString(t *testing.T) {
	cases := []struct {
		s1   map[string]string
		s2   map[string]string
		want bool
	}{
		{nil, nil, true},
		{map[string]string{}, nil, false},
		{map[string]string{}, map[string]string{}, true},
		{map[string]string{"k": "v"}, map[string]string{}, false},
		{map[string]string{"k": "v"}, map[string]string{"k": ""}, false},
		{map[string]string{"k": "v"}, map[string]string{"k": "v"}, true},
		{map[string]string{"k": "v"}, map[string]string{"c": "v"}, false},
	}
	for _, c := range cases {
		got := EqualStringString(c.s1, c.s2)
		if got != c.want {
			t.Errorf("Expected EqualStringString(%v, %v) to return %v, got %v", c.s1, c.s2, c.want, got)
		}
	}
}
