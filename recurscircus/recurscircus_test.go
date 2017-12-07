package recurscircus

import (
	"testing"

	"github.com/elgolfin/adventofcode-2017/sliceutil"
)

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
		if got != c.want {
			t.Errorf("Expected GetRootProgram(%v) to return %s", c.in, c.want)
		}
	}
}
func TestParseProgramLine(t *testing.T) {
	cases := []struct {
		in    string
		want1 string
		want2 int
		want3 []string
	}{
		{"pbga (66)", "pbga", 66, []string{}},
		{"qmnsyex (427) -> ucytsw, fnnxu, ntiulub, dbpvy", "qmnsyex", 427, []string{"ucytsw", "fnnxu", "ntiulub", "dbpvy"}},
	}
	for _, c := range cases {
		got1, got2, got3 := ParseProgramLine(c.in)
		if got1 != c.want1 || got2 != c.want2 || sliceutil.EqualString(got3, c.want3) {
			t.Errorf("Expected ParseProgramLine(\"%s\") to return %s, %d, %v, got %s, %d, %v", c.in, c.want1, c.want2, c.want3, got1, got2, got3)
		}
	}
}
