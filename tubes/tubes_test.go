package tubes

import "testing"

/*
    |
    |  +--+
    A  |  C
F---|----E|--+
    |  |  |  D
    +B-+  +--+`
*/
func TestWalkTheLine(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{`    |         
    |  +--+   
    A  |  C   
F---|----E|--+
    |  |  |  D
    +B-+  +--+`, "ABCDEF"},
	}
	for _, c := range cases {
		got := WalkTheLine(c.in)
		if got != c.want {
			t.Errorf("Expected WalkTheLine(%q) to return %v got %v", c.in, c.want, got)
		}
	}
}

func TestInitializeNetwork(t *testing.T) {
	cases := []struct {
		in    string
		want1 int
		want2 int
	}{
		{`    |         
    |  +--+   
    A  |  C   
F---|----E|--+
    |  |  |  D
    +B-+  +--+`, 14, 6},
	}
	for _, c := range cases {
		got := initializeNetwork(c.in)
		if got.width != c.want1 || got.height != c.want2 {
			t.Errorf("Expected initializeNetwork() to return %v, %v got %v, %v", c.want1, c.want2, got.width, got.height)
		}
	}
}

func TestIsEnd(t *testing.T) {
	cases := []struct {
		in1  point
		in2  point
		want bool
	}{
		{point{4, 0}, point{4, 1}, false},
		{point{1, 3}, point{0, 3}, true},
		{point{4, 5}, point{5, 5}, false},
	}
	network := initializeNetwork(`    |         
    |  +--+   
    A  |  C   
F---|----E|--+
    |  |  |  D
    +B-+  +--+`)
	for _, c := range cases {
		got := network.isEnd(c.in1, c.in2)
		if got != c.want {
			t.Errorf("Expected %v.isEnd(%v, %v) to return %v got %v", network, c.in1, c.in2, c.want, got)
		}
	}
}

func TestLeft(t *testing.T) {
	cases := []struct {
		in    point
		want1 point
		want2 bool
	}{
		{point{0, 0}, point{-1, -1}, true},
		{point{1, 0}, point{0, 0}, false},
		{point{2, 3}, point{1, 3}, false},
	}
	for _, c := range cases {
		got1, got2 := c.in.left()
		if got1 != c.want1 || got2 != c.want2 {
			t.Errorf("Expected %v.left() to return %v, %v got %v, %v", c.in, c.want1, c.want2, got1, got2)
		}
	}
}

func TestTop(t *testing.T) {
	cases := []struct {
		in    point
		want1 point
		want2 bool
	}{
		{point{0, 0}, point{-1, -1}, true},
		{point{0, 1}, point{0, 0}, false},
		{point{3, 2}, point{3, 1}, false},
	}
	for _, c := range cases {
		got1, got2 := c.in.top()
		if got1 != c.want1 || got2 != c.want2 {
			t.Errorf("Expected %v.top() to return %v, %v got %v, %v", c.in, c.want1, c.want2, got1, got2)
		}
	}
}

func TestRight(t *testing.T) {
	cases := []struct {
		in    point
		want1 point
		want2 bool
	}{
		{point{0, 0}, point{1, 0}, false},
		{point{2, 0}, point{-1, -1}, true},
		{point{1, 3}, point{2, 3}, false},
	}
	for _, c := range cases {
		got1, got2 := c.in.right(3)
		if got1 != c.want1 || got2 != c.want2 {
			t.Errorf("Expected %v.right(3) to return %v, %v got %v, %v", c.in, c.want1, c.want2, got1, got2)
		}
	}
}

func TestBottom(t *testing.T) {
	cases := []struct {
		in    point
		want1 point
		want2 bool
	}{
		{point{0, 0}, point{0, 1}, false},
		{point{2, 1}, point{2, 2}, false},
		{point{0, 2}, point{-1, -1}, true},
		{point{1, 3}, point{-1, -1}, true},
	}
	for _, c := range cases {
		got1, got2 := c.in.bottom(3)
		if got1 != c.want1 || got2 != c.want2 {
			t.Errorf("Expected %v.bottom(3) to return %v, %v got %v, %v", c.in, c.want1, c.want2, got1, got2)
		}
	}
}
