package fractart

import "testing"
import "github.com/elgolfin/adventofcode-2017/sliceutil"

func TestDraw(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#`, 12,
		}}
	for _, c := range cases {
		got := Draw(c.in, 2)
		if got != c.want {
			t.Errorf("Expected Draw(%v) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestInitializeSquare(t *testing.T) {
	cases := []struct {
		in   string
		want Square
	}{
		{"../.#", Square{
			[][]string{
				[]string{".", "."},
				[]string{".", "#"},
			},
		}},
	}
	for _, c := range cases {
		got := InitializeSquare(c.in)
		if !sliceutil.Equal2DString(got.grid, c.want.grid) {
			t.Errorf("Expected InitializeSquare(%v) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestSplit(t *testing.T) {
	cases := []struct {
		in   string
		want SquareCollection
	}{
		{".#./..#/###", SquareCollection{[][]Square{
			[]Square{
				Square{[][]string{
					[]string{".", "#", "."},
					[]string{".", ".", "#"},
					[]string{"#", "#", "#"},
				}}},
		}}},
		{"#..#/..../..../#..#", SquareCollection{[][]Square{
			[]Square{
				Square{[][]string{
					[]string{"#", "."},
					[]string{".", "."},
				}},
				Square{[][]string{
					[]string{".", "#"},
					[]string{".", "."},
				}}},
			[]Square{
				Square{[][]string{
					[]string{".", "."},
					[]string{"#", "."},
				}},
				Square{[][]string{
					[]string{".", "."},
					[]string{".", "#"},
				}}},
		}}},
	}
	for _, c := range cases {
		s := InitializeSquare(c.in)
		got := s.Split()
		if !sliceutil.Equal2DString(got.items[0][0].grid, c.want.items[0][0].grid) || len(got.items) != len(c.want.items) {
			t.Errorf("Expected %v.Split() to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestMerge(t *testing.T) {
	cases := []struct {
		in   SquareCollection
		want Square
	}{
		{SquareCollection{[][]Square{
			[]Square{
				Square{[][]string{
					[]string{".", "#", "."},
					[]string{".", ".", "#"},
					[]string{"#", "#", "#"},
				}}},
		}},
			Square{[][]string{
				[]string{".", "#", "."},
				[]string{".", ".", "#"},
				[]string{"#", "#", "#"},
			}},
		},
		{SquareCollection{[][]Square{
			[]Square{
				Square{[][]string{
					[]string{"#", ".", ".", "#"},
					[]string{".", ".", ".", "."},
					[]string{".", ".", ".", "."},
					[]string{"#", ".", ".", "#"},
				}}},
		}},
			Square{[][]string{
				[]string{"#", ".", ".", "#"},
				[]string{".", ".", ".", "."},
				[]string{".", ".", ".", "."},
				[]string{"#", ".", ".", "#"},
			}},
		},
		{SquareCollection{[][]Square{
			[]Square{
				Square{[][]string{
					[]string{"#", "."},
					[]string{".", "."},
				}},
				Square{[][]string{
					[]string{".", "#"},
					[]string{".", "."},
				}}},
			[]Square{
				Square{[][]string{
					[]string{".", "."},
					[]string{"#", "."},
				}},
				Square{[][]string{
					[]string{".", "."},
					[]string{".", "#"},
				}}},
		}},
			Square{[][]string{
				[]string{"#", ".", ".", "#"},
				[]string{".", ".", ".", "."},
				[]string{".", ".", ".", "."},
				[]string{"#", ".", ".", "#"},
			}}},
	}
	for _, c := range cases {
		got := c.in.Merge()
		if !sliceutil.Equal2DString(got.grid, c.want.grid) {
			t.Errorf("Expected %v.Merge() to return %v, got %v", c.in, c.want.grid, got.grid)
		}
	}
}

func TestDoesMatchRule(t *testing.T) {
	cases := []struct {
		in   string
		rule string
		want bool
	}{
		{"#./..", "../.#", true},
		{".#/..", "../.#", true},
		{"../.#", "../.#", true},
		{"../#.", "../.#", true},
		{"../##", "../.#", false},
		{".#./..#/###", ".#./#../###", true},
		{".#./..#/###", ".../..#/.#.", false},
	}
	for _, c := range cases {
		s := InitializeSquare(c.in)
		rule := InitializeSquare(c.rule)
		got := s.DoesMatchRule(rule)
		if got != c.want {
			t.Errorf("Expected %v.DoesMatchRule(%v) to return %v, got %v", c.in, c.rule, c.want, got)
		}
	}
}

func TestCountPixelOn(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"../.#", 1},
		{".../.#./###", 4},
		{".../.##/..#", 3},
	}
	for _, c := range cases {
		s := InitializeSquare(c.in)
		got := s.CountPixelOn()
		if got != c.want {
			t.Errorf("Expected CountPixelOn(%v) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestFlipSquare(t *testing.T) {
	cases := []struct {
		in   string
		want Square
	}{
		{"../.#", Square{
			[][]string{
				[]string{".", "."},
				[]string{"#", "."},
			},
		}},
		{".../.#./###", Square{
			[][]string{
				[]string{".", ".", "."},
				[]string{".", "#", "."},
				[]string{"#", "#", "#"},
			},
		}},
		{".../.##/..#", Square{
			[][]string{
				[]string{".", ".", "."},
				[]string{"#", "#", "."},
				[]string{"#", ".", "."},
			},
		}},
	}
	for _, c := range cases {
		got := InitializeSquare(c.in)
		got.Flip()
		if !sliceutil.Equal2DString(got.grid, c.want.grid) {
			t.Errorf("Expected Flip(%v) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestRotateSquare(t *testing.T) {
	cases := []struct {
		in   string
		want Square
	}{
		{"../#.", Square{
			[][]string{
				[]string{"#", "."},
				[]string{".", "."},
			},
		}},
		{"##./.../...", Square{
			[][]string{
				[]string{".", ".", "#"},
				[]string{".", ".", "#"},
				[]string{".", ".", "."},
			},
		}},
		{"###/#.#/###", Square{
			[][]string{
				[]string{"#", "#", "#"},
				[]string{"#", ".", "#"},
				[]string{"#", "#", "#"},
			},
		}},
	}
	for _, c := range cases {
		got := InitializeSquare(c.in)
		got.Rotate()
		if !sliceutil.Equal2DString(got.grid, c.want.grid) {
			t.Errorf("Expected Rotate(%v) to return %v, got %v", c.in, c.want, got)
		}
	}
}
