package passphrase

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}
	for _, c := range cases {
		got := IsValid(c.in)
		if got != c.want {
			t.Errorf("Expected IsValid(%q) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestGetValidPassphrases(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`aa bb cc dd ee
aa bb cc dd aa
aa bb cc dd aaa`, 2},
	}
	for _, c := range cases {
		got := GetValidPassphrases(c.in)
		if got != c.want {
			t.Errorf("Expected IsValid(%q) to return %v, got %v", c.in, c.want, got)
		}
	}
}
