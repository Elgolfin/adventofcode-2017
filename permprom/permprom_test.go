package permprom

import (
	"testing"
)

func TestDance(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{`s1,x3/4,pe/b`, "baedc"},
	}
	for _, c := range cases {
		got := Dance(c.in, 5)
		if got != c.want {
			t.Errorf("Expected Dance(%s, %d) to return %v, got %v", c.in, 5, c.want, got)
		}
	}
}

func TestSpin(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{3, "cdeab"},
		{1, "eabcd"},
		{0, "abcde"},
	}
	for _, c := range cases {
		p := InitializePrograms(5)
		p.Spin(c.in)
		got := string(p.progs)
		if got != c.want {
			t.Errorf("Expected Spin(%d) to return %v, got %v", c.in, c.want, got)
		}
	}
}

func TestExchange(t *testing.T) {
	cases := []struct {
		in1  int
		in2  int
		want string
	}{
		{3, 4, "abced"},
	}
	for _, c := range cases {
		p := InitializePrograms(5)
		p.Exchange(c.in1, c.in2)
		got := string(p.progs)
		if got != c.want {
			t.Errorf("Expected Exchange(%d, %d) to return %v, got %v", c.in1, c.in2, c.want, got)
		}
	}
}

func TestPartner(t *testing.T) {
	cases := []struct {
		in1  rune
		in2  rune
		want string
	}{
		{'a', 'd', "dbcae"},
	}
	for _, c := range cases {
		p := InitializePrograms(5)
		p.Partner(c.in1, c.in2)
		got := string(p.progs)
		if got != c.want {
			t.Errorf("Expected Partner(%v, %v) to return %v, got %v", c.in1, c.in2, c.want, got)
		}
	}
}
