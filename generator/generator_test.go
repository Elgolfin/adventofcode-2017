package generator

import (
	"testing"
)

func TestJudge(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"65,8921", 588},
	}
	for _, c := range cases {
		got := Judge(c.in)
		if got != c.want {
			t.Errorf("Expected Judge(%q) to return %d got %d", c.in, c.want, got)
		}
	}
}

func TestImpatientJudge(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"65,8921", 309},
	}
	for _, c := range cases {
		got := ImpatientJudge(c.in)
		if got != c.want {
			t.Errorf("Expected ImpatientJudge(%q) to return %d got %d", c.in, c.want, got)
		}
	}
}

func TestNextValueGenA(t *testing.T) {
	cases := []struct {
		want int
	}{
		{1092455},
		{1181022009},
		{245556042},
		{1744312007},
		{1352636452},
	}
	genA := Generator{65, 16807, 2147483647, -1}
	for _, c := range cases {
		got := genA.NextValue()
		if got != c.want {
			t.Errorf("Expected NextValue() to return %d got %d", c.want, got)
		}
	}
}

func TestNextValueGenB(t *testing.T) {
	cases := []struct {
		want int
	}{
		{430625591},
		{1233683848},
		{1431495498},
		{137874439},
		{285222916},
	}
	genA := Generator{8921, 48271, 2147483647, -1}
	for _, c := range cases {
		got := genA.NextValue()
		if got != c.want {
			t.Errorf("Expected NextValue() to return %d got %d", c.want, got)
		}
	}
}

func TestNextValuePickyGenA(t *testing.T) {
	cases := []struct {
		want int
	}{
		{1352636452},
		{1992081072},
		{530830436},
		{1980017072},
		{740335192},
	}
	genA := Generator{65, 16807, 2147483647, -1}
	for _, c := range cases {
		got := genA.NextValuePicky(4)
		if got != c.want {
			t.Errorf("Expected NextValuePicky() to return %d got %d", c.want, got)
		}
	}
}

func TestNextValuePickyGenB(t *testing.T) {
	cases := []struct {
		want int
	}{
		{1233683848},
		{862516352},
		{1159784568},
		{1616057672},
		{412269392},
	}
	genA := Generator{8921, 48271, 2147483647, -1}
	for _, c := range cases {
		got := genA.NextValuePicky(8)
		if got != c.want {
			t.Errorf("Expected NextValuePicky() to return %d got %d", c.want, got)
		}
	}
}
