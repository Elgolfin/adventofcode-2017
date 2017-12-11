package hexed

import "strings"

// CountSteps returns the fewest number of steps
func CountSteps(content string) int {
	path := strings.Split(content, ",")
	countDir := map[string]int{"nw": 0, "n": 0, "ne": 0, "se": 0, "s": 0, "sw": 0}
	for _, direction := range path {
		// ne cancels sw; se cancels nw, n cancels s
		switch {
		case direction == "ne":
			if countDir["sw"] > 0 {
				countDir["sw"]--
			} else {
				countDir["ne"]++
			}
		case direction == "sw":
			if countDir["ne"] > 0 {
				countDir["ne"]--
			} else {
				countDir["sw"]++
			}
		case direction == "se":
			if countDir["nw"] > 0 {
				countDir["nw"]--
			} else {
				countDir["se"]++
			}
		case direction == "nw":
			if countDir["se"] > 0 {
				countDir["se"]--
			} else {
				countDir["nw"]++
			}
		case direction == "n":
			if countDir["s"] > 0 {
				countDir["s"]--
			} else {
				countDir["n"]++
			}
		case direction == "s":
			if countDir["n"] > 0 {
				countDir["n"]--
			} else {
				countDir["s"]++
			}
		}
	}

	// Change all the possible combinations until no more changes can be done
	for {
		change := false
		// combination of ne,s or s,ne can be replaced by se
		change = change || cleanDirections(countDir, "ne", "s", "se")

		// combination of ne,nw|nw,ne can be replaced by n
		change = change || cleanDirections(countDir, "ne", "nw", "n")

		// combination of se,sw|sw,se can be replaced by s
		change = change || cleanDirections(countDir, "se", "sw", "s")

		// combination of se,n|n,se can be replaced by ne
		change = change || cleanDirections(countDir, "se", "n", "ne")

		// combination of sw,n|n,sw can be replaced by nw
		change = change || cleanDirections(countDir, "sw", "n", "nw")

		// combination of s,nw|nw,s can be replaced by sw
		change = change || cleanDirections(countDir, "s", "nw", "sw")

		if !change {
			break
		}
	}

	// fmt.Printf("%v\n", countDir)
	return countSteps(countDir)
}

func countSteps(hexSteps map[string]int) int {
	steps := 0
	for _, v := range hexSteps {
		steps += v
	}
	return steps
}

func cleanDirections(dirMap map[string]int, dir1 string, dir2 string, transformDir string) bool {
	change := false
	switch {
	case dirMap[dir2] > 0 && dirMap[dir1] > dirMap[dir2]:
		transformCount := dirMap[dir1] - dirMap[dir2]
		dirMap[transformDir] += transformCount
		dirMap[dir2] = 0
		dirMap[dir1] -= transformCount
		change = true
	case dirMap[dir2] > 0 && dirMap[dir1] == dirMap[dir2]:
		dirMap[transformDir] += dirMap[dir1]
		dirMap[dir1] = 0
		dirMap[transformDir] = 0
		change = true
	case dirMap[dir1] > 0 && dirMap[dir1] < dirMap[dir2]:
		transformCount := dirMap[dir2] - dirMap[dir1]
		dirMap[transformDir] += transformCount
		dirMap[dir1] = 0
		dirMap[dir2] -= transformCount
		change = true
	}
	return change
}
