package checksum

import (
	"testing"
)

func TestGetSmallest(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"5	1	9	5", 1},
		{"7	5	3", 3},
		{"2	4	6	8", 2},
		{"2	4	0	8", 0},
		{"179	64	150	88", 64},
	}
	for _, c := range cases {
		got := GetSmallest(c.in)
		if got != c.want {
			t.Errorf("Expected GetSmallest(%q) to return %d, got %d", c.in, c.want, got)
		}
	}
}
func TestGetLargest(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"5	1	9	5", 9},
		{"7	5	3", 7},
		{"2	4	6	8", 8},
		{"2	4	6	8	5	9	7", 9},
		{"179	64	150	88", 179},
	}
	for _, c := range cases {
		got := GetLargest(c.in)
		if got != c.want {
			t.Errorf("Expected GetLargest(%q) to return %d, got %d", c.in, c.want, got)
		}
	}
}
func TestGetDivisible(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"5	9	2	8", 4},
		{"9	4	7	3", 3},
		{"3	8	6	5", 2},
	}
	for _, c := range cases {
		got := GetDivisible(c.in)
		if got != c.want {
			t.Errorf("Expected GetDivisible(%q) to return %d, got %d", c.in, c.want, got)
		}
	}
}
func TestGenerate(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`5	1	9	5
7	5	3
2	4	6	8`, 18},
	}
	for _, c := range cases {
		got := Generate(c.in)
		if got != c.want {
			t.Errorf("Expected Generate(%q) to return %d, got %d", c.in, c.want, got)
		}
	}
}
func TestGenerateNew(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`5	9	2	8
9	4	7	3
3	8	6	5`, 9},
	}
	for _, c := range cases {
		got := GenerateNew(c.in)
		if got != c.want {
			t.Errorf("Expected GenerateNew(%q) to return %d, got %d", c.in, c.want, got)
		}
	}
}
