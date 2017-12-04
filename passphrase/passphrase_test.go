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
			t.Errorf("Expected GetValidPassphrases(%q) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestIsValidNew(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
	}
	for _, c := range cases {
		got := IsValidNew(c.in)
		if got != c.want {
			t.Errorf("Expected IsValidNew(%q) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestGetValidPassphrasesNew(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`abcde fghij
abcde xyz ecdab
a ab abc abd abf abj
iiii oiii ooii oooi oooo
oiii ioii iioi iiio
bqs bqs dbutvgf mmzb izpyud rap izpyud xlzeb mnj hjncs`, 3},
	}
	for _, c := range cases {
		got := GetValidPassphrasesNew(c.in)
		if got != c.want {
			t.Errorf("Expected GetValidPassphrasesNew(%q) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		in1  string
		in2  string
		want bool
	}{
		{"iouty", "tyoui", true},
		{"abcd", "abce", false},
		{"bqs", "bqs", true},
	}
	for _, c := range cases {
		got := IsAnagram(c.in1, c.in2)
		if got != c.want {
			t.Errorf("Expected IsAnagram(%q,%q) to return %v, got %v", c.in1, c.in2, c.want, got)
		}
	}
}
