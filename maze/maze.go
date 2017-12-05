package maze

import (
	"strconv"
	"strings"
)

// FindExit returns the number of steps to escape the maze of intructions
func FindExit(content string) int {
	instructions := initializeArray(content)
	steps := 0
	offset := 0
	previousOffset := 0
	for true {
		previousOffset = offset
		offset += instructions[offset]
		instructions[previousOffset]++
		steps++
		// fmt.Printf("%d. %v\n", steps, instructions)
		if offset >= len(instructions) {
			break
		}
	}
	return steps
}

func initializeArray(content string) []int {
	lines := strings.Split(content, "\n")
	arrayOfInts := make([]int, len(lines))
	for i, instruction := range lines {
		arrayOfInts[i], _ = strconv.Atoi(instruction)
	}
	return arrayOfInts
}
