package electromoat

import (
	"testing"
)

func TestFindHeaviestBridge(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{`0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`, 31},
	}
	for _, c := range cases {
		got := FindHeaviestBridge(c.in)
		if got != c.want {
			t.Errorf("Expected FindHeaviestBridge(%q) to return %d got %d", c.in, c.want, got)
		}
	}
}

func TestBuildChains(t *testing.T) {
	cases := []struct {
		in   string
		want map[string][]Component
	}{
		//{``, []Component{}},
		{`0/1
0/2`, map[string][]Component{
			"--0,1": []Component{Component{0, 1}},
			"--0,2": []Component{Component{0, 2}},
		}},
		{`0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10`, map[string][]Component{
			"--0,1":                []Component{Component{0, 1}},
			"--0,1--10,1":          []Component{Component{0, 1}, Component{10, 1}},
			"--0,1--10,1--9,10":    []Component{Component{0, 1}, Component{10, 1}, Component{9, 10}},
			"--0,2":                []Component{Component{0, 2}},
			"--0,2--2,3":           []Component{Component{0, 2}, Component{2, 3}},
			"--0,2--2,3--3,4":      []Component{Component{0, 2}, Component{2, 3}, Component{3, 4}},
			"--0,2--2,3--3,5":      []Component{Component{0, 2}, Component{2, 3}, Component{3, 5}},
			"--0,2--2,2":           []Component{Component{0, 2}, Component{2, 2}},
			"--0,2--2,2--2,3":      []Component{Component{0, 2}, Component{2, 2}, Component{2, 3}},
			"--0,2--2,2--2,3--3,4": []Component{Component{0, 2}, Component{2, 2}, Component{2, 3}, Component{3, 4}},
			"--0,2--2,2--2,3--3,5": []Component{Component{0, 2}, Component{2, 2}, Component{2, 3}, Component{3, 5}},
		}},
	}
	for _, c := range cases {
		got := BuildChains("", make(map[string][]Component), 0, LoadComponents(c.in))
		equal := true
		// fmt.Printf("%v\n", got)
		for k, v := range got {
			if len(c.want[k]) != len(v) {
				equal = false
				break
			}
		}
		if !equal {
			t.Errorf("Expected LoadComponents(%q) to return %v got %v", c.in, c.want, got)
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
