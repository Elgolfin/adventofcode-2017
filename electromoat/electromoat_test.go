package electromoat

import (
	"testing"
)

func TestFindHeaviestBridges(t *testing.T) {
	cases := []struct {
		in    string
		want1 int
		want2 int
	}{
		{`0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`, 31, 19},
	}
	for _, c := range cases {
		got1, got2 := FindHeaviestBridges(c.in)
		if got1 != c.want1 || got2 != c.want2 {
			t.Errorf("Expected FindHeaviestBridges(%q) to return %d, %d got %d, %d", c.in, c.want1, c.want2, got1, got2)
		}
	}
}

func TestBuildChains(t *testing.T) {
	cases := []struct {
		in   string
		want map[string]int
	}{
		//{``, []Component{}},
		{`0/1
0/2`, map[string]int{
			"--0,1": 1,
			"--0,2": 2,
		}},
		{`0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`, map[string]int{
			"--0,1":                1,
			"--0,1--10,1":          12,
			"--0,1--10,1--9,10":    31,
			"--0,2":                2,
			"--0,2--2,3":           7,
			"--0,2--2,3--3,4":      14,
			"--0,2--2,3--3,5":      15,
			"--0,2--2,2":           6,
			"--0,2--2,2--2,3":      11,
			"--0,2--2,2--2,3--3,4": 18,
			"--0,2--2,2--2,3--3,5": 19,
		}},
	}
	for _, c := range cases {
		got := BuildBridges("", make(map[string]int), 0, LoadComponents(c.in))
		equal := true
		// fmt.Printf("%v\n", got)
		for k, v := range got {
			if c.want[k] != v {
				equal = false
				break
			}
		}
		if !equal {
			t.Errorf("Expected BuildChains(%q) to return %v got %v", c.in, c.want, got)
		}
	}
}

func TestLoadComponents(t *testing.T) {
	cases := []struct {
		in string

		want map[int][]Component
	}{
		//{``, []Component{}},
		{`1/1
2/2`, map[int][]Component{
			1: []Component{Component{1, 1}},
			2: []Component{Component{2, 2}},
		}},
		{`0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`, map[int][]Component{
			0:  []Component{Component{0, 2}, Component{0, 1}},
			1:  []Component{Component{0, 1}, Component{10, 1}},
			2:  []Component{Component{0, 2}, Component{2, 2}, Component{2, 3}},
			3:  []Component{Component{2, 3}, Component{3, 4}, Component{3, 5}},
			4:  []Component{Component{3, 4}},
			5:  []Component{Component{3, 5}},
			9:  []Component{Component{9, 10}},
			10: []Component{Component{10, 1}, Component{9, 10}},
		}},
	}
	for _, c := range cases {
		got := LoadComponents(c.in)
		equal := true
		for key, val := range got {
			if len(val) != len(c.want[key]) {
				equal = false
				break
			}
			for i, component := range val {
				if component.port1 != c.want[key][i].port1 || component.port2 != c.want[key][i].port2 {
					equal = false
					break
				}
			}
		}
		if !equal {
			t.Errorf("Expected LoadComponents(%q) to return %v got %v", c.in, c.want, got)
		}
	}
}
