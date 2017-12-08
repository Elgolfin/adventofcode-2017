package recurscircus

import (
	"testing"

	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

func TestGetRootProgramName(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{`pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`, "tknk"},
	}
	for _, c := range cases {
		got := GetRootProgramName(c.in)
		if got != c.want {
			t.Errorf("Expected GetRootProgramName(%v) to return %s, got %s", c.in, c.want, got)
		}
	}
}

func TestGetRootProgram(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{`pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`, "tknk"},
	}
	for _, c := range cases {
		got := GetRootProgram(c.in)
		if got.Name != c.want {
			t.Errorf("Expected GetRootProgram(%v) to return %s, got %s", c.in, c.want, got.Name)
		}
	}
}

func TestBalanceWeight(t *testing.T) {
	cases := []struct {
		in    string
		want1 int
		want2 string
	}{
		{`pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`, 60, "ugml"},
	}
	for _, c := range cases {
		got1, got2 := BalanceWeight(c.in)
		if got1 != c.want1 || got2.Name != c.want2 {
			t.Errorf("Expected BalanceWeight(%v) to return %d, %s got %d, %s", c.in, c.want1, c.want2, got1, got2.Name)
		}
	}
}
func TestParseProgramLine(t *testing.T) {
	cases := []struct {
		in   string
		want Program
	}{
		{"pbga (66)", Program{"pbga", 66, nil, nil}},
		{"qmnsyex (427) -> ucytsw, fnnxu, ntiulub, dbpvy", Program{"qmnsyex", 427, []string{"ucytsw", "fnnxu", "ntiulub", "dbpvy"}, nil}},
	}
	for _, c := range cases {
		got := ParseProgramLine(c.in)
		if got.Name != c.want.Name || got.Weight != c.want.Weight || !sliceutil.EqualString(got.abovePrograms, c.want.abovePrograms) {
			t.Errorf("Expected ParseProgramLine(\"%s\") to return %s, %d, %v, got %s, %d, %v", c.in, c.want.Name, c.want.Weight, c.want.abovePrograms, got.Name, got.Weight, got.abovePrograms)
		}
	}
}
