package sliceutil

import (
	"testing"
)

func TestAtoi(t *testing.T) {
	cases := []struct {
		in   string
		sep  string
		want []int
	}{
		{"5	1	9	5", "	", []int{5, 1, 9, 5}},
		{"7 5 3 7", " ", []int{7, 5, 3, 7}},
		{`2
4
6
8`, "\n", []int{2, 4, 6, 8}},
	}
	for _, c := range cases {
		got := Atoi(c.in, c.sep)
		if !EqualInt(got, c.want) {
			t.Errorf("Expected Atoi(%q, %q) to return %v, got %v", c.in, c.sep, c.want, got)
		}
	}
}

func TestEqualInt(t *testing.T) {
	cases := []struct {
		s1   []int
		s2   []int
		want bool
	}{
		{[]int{5, 1, 9, 5}, []int{5, 1, 9, 5}, true},
		{[]int{7, 5, 3, 7}, []int{2, 1, 9, 5}, false},
		{[]int{2, 1, 9, 7}, []int{2, 1, 9}, false},
		{nil, nil, true},
		{[]int{2, 1, 9, 7}, nil, false},
	}
	for _, c := range cases {
		got := EqualInt(c.s1, c.s2)
		if got != c.want {
			t.Errorf("Expected EqualInt(%v, %v) to return %v, got %v", c.s1, c.s2, c.want, got)
		}
	}
}

func TestEqualString(t *testing.T) {
	cases := []struct {
		s1   []string
		s2   []string
		want bool
	}{
		{[]string{"5", "1", "9", "5"}, []string{"5", "1", "9", "5"}, true},
		{[]string{"7", "5", "3", "7"}, []string{"2", "1", "9", "5"}, false},
		{[]string{"2", "1", "9", "7"}, []string{"2", "1", "9"}, false},
		{nil, nil, true},
		{[]string{"2", "1", "9", "7"}, nil, false},
		{make([]string, 0), make([]string, 0), true},
	}
	for _, c := range cases {
		got := EqualString(c.s1, c.s2)
		if got != c.want {
			t.Errorf("Expected EqualString(%v, %v) to return %v, got %v", c.s1, c.s2, c.want, got)
		}
	}
}

func TestExtendString(t *testing.T) {
	cases := []struct {
		s1   []string
		s2   string
		want []string
	}{
		{[]string{"5", "1", "9"}, "5", []string{"5", "1", "9", "5"}},
		{[]string{"7", "5", "3"}, "2", []string{"7", "5", "3", "2"}},
		{[]string{"2", "1"}, "9", []string{"2", "1", "9"}},
	}
	for _, c := range cases {
		got := ExtendString(c.s1, c.s2)
		if !EqualString(got, c.want) {
			t.Errorf("Expected ExtendString(%v, %v) to return %v, got %v", c.s1, c.s2, c.want, got)
		}
	}
}

func TestGetLargest(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{5, 1, 9, 5}, 2},
		{[]int{7, 5, 3, 7}, 0},
		{[]int{2, 1, 9, 17}, 3},
		{nil, -1},
		{[]int{2, 1, 9, 7, 56, 21}, 4},
	}
	for _, c := range cases {
		got := GetLargest(c.in)
		if got != c.want {
			t.Errorf("Expected GetLargest(%v) to return %v, got %v", c.in, c.want, got)
		}
	}
}
