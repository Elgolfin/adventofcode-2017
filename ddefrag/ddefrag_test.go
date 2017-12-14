package ddefrag

import (
	"testing"
)

// func TestCountSquares(t *testing.T) {
// 	cases := []struct {
// 		in   string
// 		want int
// 	}{
// 		{"flqrgnkx", 8108},
// 	}
// 	for _, c := range cases {
// 		got, _ := CountSquares(c.in)
// 		if got != c.want {
// 			t.Errorf("Expected CountSquares(%q) to return %v, got %v", c.in, c.want, got)
// 		}
// 	}
// }

func TestConvertHexaToBits(t *testing.T) {
	cases := []struct {
		in    string
		want1 int
		want2 string
	}{
		{"012", 2, "000000010010"},
	}
	for _, c := range cases {
		got1, got2 := OnesCount(c.in)
		if got1 != c.want1 || got2 != c.want2 {
			t.Errorf("Expected OnesCount(%s) to return %d, %s got %d, %s", c.in, c.want1, c.want2, got1, got2)
		}
	}
}

// func TestAreasCount(t *testing.T) {
// 	cases := []struct {
// 		in   string
// 		want int
// 	}{
// 		{"flqrgnkx", 8108},
// 	}
// 	for _, c := range cases {
// 		_, grid := CountSquares(c.in)
// 		got := AreasCount(grid)
// 		if got != c.want {
// 			t.Errorf("Expected CountSquares(%q) to return %v, got %v", c.in, c.want, got)
// 		}
// 	}
// }

func TestAreasCount(t *testing.T) {
	grid := [][]int{
		[]int{1, 1, 0, 1, 0, 1, 0, 0},
		[]int{0, 1, 0, 1, 0, 1, 0, 1},
		[]int{0, 0, 0, 0, 1, 0, 1, 0},
		[]int{1, 0, 1, 0, 1, 1, 0, 1},
		[]int{0, 1, 1, 0, 1, 0, 0, 0},
		[]int{1, 1, 0, 0, 1, 0, 0, 1},
		[]int{0, 1, 0, 0, 0, 1, 0, 0},
		[]int{1, 1, 0, 0, 0, 1, 1, 0},
	}
	// gridResult := [][]int{
	// 	[]int{2, 2, 0, 3, 0, 4, 0, 0},
	// 	[]int{0, 2, 0, 3, 0, 4, 0, 5},
	// 	[]int{0, 0, 0, 0, 6, 0, 7, 0},
	// 	[]int{8, 0, 9, 0, 6, 6, 0, 10},
	// 	[]int{0, 9, 9, 0, 6, 0, 0, 0},
	// 	[]int{9, 9, 0, 0, 6, 0, 0, 11},
	// 	[]int{0, 9, 0, 0, 0, 12, 0, 0},
	// 	[]int{9, 9, 0, 0, 0, 12, 12, 0},
	// }
	regions, grid := AreasCount(grid)
	if regions != 11 {
		t.Errorf("Expected AreasCount to return %v, got %v", 11, regions)
	}
}
